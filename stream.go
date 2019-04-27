package mavlink2

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2019 queue-b <https://github.com/queue-b>

Permission is hereby granted, free of charge, to any person obtaining a copy
of the generated software (the "Generated Software"), to deal
in the Generated Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Generated Software, and to permit persons to whom the Generated
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Generated Software.

THE GENERATED SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE GENERATED SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE GENERATED SOFTWARE.
*/

import (
	"bufio"
	"context"
	"encoding/binary"
	"io"
)

// V2StartByte is the magic frame start indicator for a V2 MAVLink Frame
const V2StartByte = 0xfd

// V1StartByte is the magic frame start indicator for a V1 MAVLink Frame
const V1StartByte = 0xfe

// CompatibilityFlagSignature is the flag indicating that a V2 Frame is signed
const CompatibilityFlagSignature = 0x01

// FrameParser represents a function that takes a Frame as input and returns a Message
type FrameParser func(Frame) (Message, error)

// FrameStream represents a stream of MAVLink Frames
type FrameStream struct {
	InputFrames  chan Frame
	OutputFrames chan Frame
	Errors       chan error
}

// RunForever consumes the MAVLink stream from the underlying Reaader
func (s *FrameStream) RunForever(readStream io.Reader, writeStream *io.Writer) {
	s.Run(context.Background(), readStream, writeStream)
}

// Run consumes the MAVLink stream from the underlying Reader until the context is cancelled
func (s *FrameStream) Run(ctx context.Context, readStream io.Reader, writeStream *io.Writer) {
	// Read Frames
	go func() {
		reader := bufio.NewReader(readStream)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				frame, err := readFrame(reader)

				if err != nil {
					select {
					case s.Errors <- err:
					default:
					}
					continue
				}

				select {
				case s.OutputFrames <- frame:
				default:
				}
			}
		}
	}()

	if writeStream != nil {
		// Write Frames
		go func() {
			writer := *writeStream

			for {
				select {
				case <-ctx.Done():
					return
				default:
					frame := <-s.InputFrames
					_, err := writer.Write(frame.Bytes())
					if err != nil {
						select {
						case s.Errors <- err:
						default:
						}
					}
				}
			}
		}()
	}
}

func peekHeader(reader *bufio.Reader) (frameLength uint16, err error) {
	for {
		var headerBytes []byte
		headerBytes, err = reader.Peek(3)

		if err != nil {
			return
		}

		messageLength := headerBytes[1]
		startByte := headerBytes[0]

		if startByte == V2StartByte {
			frameLength = uint16(messageLength) + 12

			if headerBytes[2]&CompatibilityFlagSignature == CompatibilityFlagSignature {
				frameLength += 13
			}

			return
		} else if startByte == V1StartByte {
			frameLength = uint16(messageLength) + 8

			return
		} else {
			_, err = reader.ReadByte()

			if err != nil {
				return
			}
		}
	}
}

func readFrame(reader *bufio.Reader) (frame Frame, err error) {
	frameLength, err := peekHeader(reader)

	if err != nil {
		return
	}

	frameBytes := make([]byte, frameLength)

	_, err = io.ReadFull(reader, frameBytes)

	if err != nil {
		return
	}

	frame, err = FrameFromBytes(frameBytes)

	return
}

func packFrame(version int, senderSystemID, senderComponentID, messageSequence uint8, message Message, meta MessageMeta) (frame Frame, err error) {
	// Make sure the version matches correctly
	if version != 1 && version != 2 {
		err = ErrUnsupportedVersion
		return
	}

	var frameBytes []byte
	var messageBytes []byte

	messageBytes, err = message.Write(version)

	if err != nil {
		return
	}

	x25 := X25CRC{}
	x25.Reset()

	if version == 1 {
		frameBytes = make([]byte, len(messageBytes)+6)
		frameBytes[0] = V1StartByte
		frameBytes[1] = uint8(len(messageBytes))
		frameBytes[2] = senderSystemID
		frameBytes[3] = senderComponentID
		frameBytes[4] = messageSequence

		idBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(idBytes, message.GetID())

		frameBytes[5] = idBytes[0]
		copy(frameBytes[6:], messageBytes)

		_, err = x25.Write(append(frameBytes[1:], meta.CRCExtra))

		if err != nil {
			return
		}

		frameBytes := x25.Sum(frameBytes)

		frame = FrameV1(frameBytes)
	}

	if version == 2 {
		frameBytes = make([]byte, len(messageBytes)+10)
		frameBytes[0] = V2StartByte
		frameBytes[1] = uint8(len(messageBytes))
		// Compatibility and incompatibility flags (bytes 2 & 3) not currently supported
		frameBytes[4] = messageSequence
		frameBytes[5] = senderSystemID
		frameBytes[6] = senderComponentID

		idBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(idBytes, message.GetID())

		copy(frameBytes[7:10], idBytes[:3])
		copy(frameBytes[10:], messageBytes)

		_, err = x25.Write(append(frameBytes[1:], meta.CRCExtra))

		if err != nil {
			return
		}

		frameBytes = x25.Sum(frameBytes)
		frame = FrameV2(frameBytes)
	}

	return
}
