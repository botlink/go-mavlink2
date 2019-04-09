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
	"fmt"
	"io"
)

// V2StartByte is the magic frame start indicator for a V2 MAVLink Frame
const V2StartByte = 0xfd

// V1StartByte is the magic frame start indicator for a V1 MAVLink Frame
const V1StartByte = 0xfe

// CompatibilityFlagSignature is the flag indicating that a V2 Frame is signed
const CompatibilityFlagSignature = 0x01

// MAVLinkStream represents a stream of MAVLink Frames
type MAVLinkStream struct {
	Version  int
	Dialects Dialects
	Frames   chan Frame
	Messages chan Message
}

// Run consumes the MAVLink stream from the underlying Reader
func (s *MAVLinkStream) Run(stream io.Reader) {
	// Read Frames & Messages
	go func() {
		reader := bufio.NewReader(stream)
		for {
			frame, err := readFrame(reader)

			if err != nil {
				fmt.Println(err)
				continue
			}

			select {
			case s.Frames <- frame:
			default:
			}

			err = s.Dialects.Validate(frame)

			if err != nil && err != ErrUnknownMessage {
				fmt.Println(err)
				continue
			}

			message, err := s.Dialects.GetMessage(frame)

			if err != nil && err != ErrUnknownMessage {
				fmt.Println(err)
				continue
			}

			select {
			case s.Messages <- message:
			default:
			}
		}
	}()
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

// FrameParser represents a function that takes a Frame as input and returns a Message
type FrameParser func(Frame) (Message, error)
