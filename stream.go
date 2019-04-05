/*
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

package mavlink

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

// FrameStream represents a stream of MAVLink Frames
type FrameStream struct {
	Version int
	reader  bufio.ReadWriter
}

func (s *FrameStream) peekHeader() (frameLength int, err error) {
	for {
		var headerBytes []byte
		headerBytes, err = s.reader.Peek(3)

		if err != nil {
			return
		}

		messageLength := headerBytes[1]
		startByte := headerBytes[0]

		if startByte == V2StartByte {
			frameLength = int(messageLength + 12)

			if headerBytes[2]&CompatibilityFlagSignature == CompatibilityFlagSignature {
				frameLength += 13
			}

			return
		} else if startByte == V1StartByte {
			frameLength = int(messageLength + 8)

			return
		} else {
			_, err = s.reader.ReadByte()

			if err != nil {
				return
			}
		}
	}
}

func (s *FrameStream) readFrame() (frame Frame, err error) {
	frameLength, err := s.peekHeader()

	if err != nil {
		return
	}

	frameBytes := make([]byte, frameLength)

	_, err = io.ReadFull(s.reader, frameBytes)

	if err != nil {
		return
	}

	frame, err = FrameFromBytes(frameBytes)

	return
}

// FrameFromBytes returns a Frame containing the input bytes,
// or an error if the first byte is not a valid frame start
func FrameFromBytes(raw []byte) (frame Frame, err error) {
	if len(raw) == 0 {
		err = fmt.Errorf("Runt frame")
		return
	}

	startByte := raw[0]

	if startByte == V1StartByte {
		frame = FrameV1(raw)
	} else if startByte == V2StartByte {
		frame = FrameV2(raw)
	} else {
		err = fmt.Errorf("Unknown start byte")
	}

	return
}
