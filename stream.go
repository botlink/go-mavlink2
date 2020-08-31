package mavlink2

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2020 queue-b <https://github.com/queue-b>

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
	"errors"
	"fmt"
	"io"
	"sync"
)

// V2StartByte is the magic frame start indicator for a V2 MAVLink Frame
const V2StartByte = 0xfd

// V1StartByte is the magic frame start indicator for a V1 MAVLink Frame
const V1StartByte = 0xfe

// CompatibilityFlagSignature is the flag indicating that a V2 Frame is signed
const CompatibilityFlagSignature = 0x01

// FrameParser represents a function that takes a Frame as input and returns a Message
type FrameParser func(Frame) (Message, error)

var ErrUserClosed = errors.New("User closed")
var ErrWriteChannelClosed = errors.New("Write channel closed")

type frameLoggingMetaData struct {
	version           int
	messageID         uint32
	senderSystemID    uint8
	senderComponentID uint8
}

// FrameStream represents a stream of MAVLink Frames
type FrameStream struct {
	sync.RWMutex
	inputFrames         chan Frame
	outputFrames        chan Frame
	rwc                 io.ReadWriteCloser
	reader              *bufio.Reader
	writer              io.Writer
	closeErr            error
	closeOnce           sync.Once
	closed              chan struct{}
	buffer              []byte
	bufferIndex         int
	dialects            Dialects
	returnInvalidFrames bool
	invalidLogged       map[frameLoggingMetaData]bool
}

// See https://mavlink.io/en/guide/serialization.html
// mavlink v1 has max length 263, v2 has max length 279
const maxFrameLength = 279

func NewFrameStream(rwc io.ReadWriteCloser, inputFrames chan Frame, dialects Dialects, returnInvalidFrames bool) *FrameStream {
	f := &FrameStream{
		inputFrames:         inputFrames,
		outputFrames:        make(chan Frame),
		reader:              bufio.NewReader(rwc),
		rwc:                 rwc,
		closeOnce:           sync.Once{},
		closed:              make(chan struct{}),
		buffer:              make([]byte, maxFrameLength*2, maxFrameLength*2),
		bufferIndex:         0,
		dialects:            dialects,
		returnInvalidFrames: returnInvalidFrames,
		invalidLogged:       make(map[frameLoggingMetaData]bool),
	}

	return f
}

// Write returns a
func (s *FrameStream) Write() chan<- Frame {
	return s.inputFrames
}

func (s *FrameStream) Read() <-chan Frame {
	return s.outputFrames
}

// Close closes the FrameStream
func (s *FrameStream) Close() {
	s.close(ErrUserClosed)
}

// Closed returns a channel that is closed when the FrameStream is closed
func (s *FrameStream) Closed() <-chan struct{} {
	return s.closed
}

func (s *FrameStream) close(err error) {
	s.closeOnce.Do(func() {
		s.Lock()
		defer s.Unlock()

		s.closeErr = err

		s.rwc.Close()
		close(s.outputFrames)
		close(s.closed)
	})
}

func (s *FrameStream) WriteContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.close(ctx.Err())
			return
		case frame, ok := <-s.inputFrames:
			if !ok {
				s.close(ErrWriteChannelClosed)
				return
			}

			_, err := s.rwc.Write(frame.Bytes())
			if err != nil {
				s.close(err)
				return
			}
		}
	}
}

func (s *FrameStream) ReadContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.close(ctx.Err())
			return
		default:
			frame, err := s.readFrame(s.reader)

			if err != nil {
				s.close(err)
				return
			}

			s.outputFrames <- frame
		}
	}
}

func peekHeader(buffer []byte) (frameLength uint16, startFound bool) {
	messageLength := buffer[1]
	startByte := buffer[0]

	if startByte == V2StartByte {
		frameLength = uint16(messageLength) + 12

		if buffer[2]&CompatibilityFlagSignature == CompatibilityFlagSignature {
			frameLength += 13
		}

		return frameLength, true
	} else if startByte == V1StartByte {
		frameLength = uint16(messageLength) + 8

		return frameLength, true
	} else {
		return 0, false
	}
}

func (s *FrameStream) readFrame(reader *bufio.Reader) (frame Frame, err error) {
	for {
		if s.bufferIndex < 3 {
			bytesNeeded := 3 - s.bufferIndex
			n, err := io.ReadAtLeast(reader, s.buffer[s.bufferIndex:], bytesNeeded)
			if err != nil {
				return frame, err
			}
			s.bufferIndex += n
		}

		frameLength, startFound := peekHeader(s.buffer)
		if startFound {
			if s.bufferIndex < int(frameLength) {
				bytesNeeded := int(frameLength) - s.bufferIndex
				//bytes := make([]byte, bytesNeeded)
				n, err := io.ReadAtLeast(reader, s.buffer[s.bufferIndex:], bytesNeeded)
				if err != nil {
					return frame, err
				}
				s.bufferIndex += n
			}
			// validate
			valid := false
			frame, err = FrameFromBytes(s.buffer, false)
			if err != nil {
				return frame, err
			}

			if err = s.dialects.Validate(frame); err == nil {
				valid = true
			} else {
				// Check if already logged so we don't spam the logs
				metaData := frameLoggingMetaData{frame.GetVersion(),
					frame.GetMessageID(),
					frame.GetSenderSystemID(),
					frame.GetSenderComponentID()}
				if _, ok := s.invalidLogged[metaData]; !ok {
					fmt.Printf("Got invalid frame, version %d, message id %d, system id %d, component id %d\n",
						metaData.version, metaData.messageID,
						metaData.senderSystemID, metaData.senderComponentID)
					s.invalidLogged[metaData] = true
				}
			}

			if valid {
				// need to copy frame as it's constructed using s.buffer without copying from s.buffer
				frame, err = FrameFromBytes(s.buffer, true)
				copy(s.buffer, s.buffer[frameLength:])
				s.bufferIndex -= int(frameLength)
				return frame, err
			} else {
				if s.returnInvalidFrames {
					frame, err = FrameFromBytes(s.buffer, true)
				}

				copy(s.buffer, s.buffer[1:])
				s.bufferIndex -= 1

				if s.returnInvalidFrames {
					return frame, err
				}
			}
		} else {
			copy(s.buffer, s.buffer[1:])
			s.bufferIndex -= 1
		}
	}
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
