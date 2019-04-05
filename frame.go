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
	"encoding/binary"
	"fmt"
)

// Frame represents a MAVLink frame containing a MAVLink message
type Frame interface {
	GetVersion() int
	GetMessageLength() uint8
	GetMessageSequence() uint8
	GetSenderSystemID() uint8
	GetSenderComponentID() uint8
	GetMessageBytes() []byte
	GetMessageID() uint32
	GetChecksum() uint16
	GetChecksumInput() []byte
	GetBytes() []byte
}

// FrameV1 represents a MAVLink V1 Frame
type FrameV1 []byte

// GetBytes returns the Frame as a byte array
func (frame FrameV1) GetBytes() []byte {
	return frame
}

// GetVersion returns the version of the Frame
func (frame FrameV1) GetVersion() int {
	return 1
}

// GetMessageLength returns the length of the message contained in the Frame
func (frame FrameV1) GetMessageLength() uint8 {
	return frame[1]
}

// GetMessageSequence the sequence number of the message contained in the Frame
func (frame FrameV1) GetMessageSequence() uint8 {
	return frame[2]
}

// GetMessageBytes returns the message contained in the Frame as a byte array
func (frame FrameV1) GetMessageBytes() []byte {
	return frame[6 : 6+frame.GetMessageLength()]
}

// GetSenderSystemID returns the ID of the system that sent the Frame
func (frame FrameV1) GetSenderSystemID() uint8 {
	return frame[3]
}

// GetSenderComponentID returns the ID of the component that sent the Frame
func (frame FrameV1) GetSenderComponentID() uint8 {
	return frame[4]
}

// GetMessageID returns the ID of the message contained in the Frame
func (frame FrameV1) GetMessageID() uint32 {
	return uint32(frame[5])
}

// GetChecksum returns the checksum of the Frame contents
func (frame FrameV1) GetChecksum() uint16 {
	return binary.LittleEndian.Uint16(frame[len(frame)-2:])
}

// GetChecksumInput returns the contents of the Frame used to calculate the checksum
func (frame FrameV1) GetChecksumInput() []byte {
	return frame[1 : len(frame)-2]
}

func (frame FrameV1) String() string {
	return fmt.Sprintf("Version: %v\nMessageLength: %v\nSenderSystemID: %v\nSenderComponentID: %v\nMessageID: %v\nChecksum:%v",
		frame.GetVersion(),
		frame.GetMessageLength(),
		frame.GetSenderSystemID(),
		frame.GetSenderComponentID(),
		frame.GetMessageID(),
		frame.GetChecksum(),
	)
}

// FrameV2 represents a MAVLink V1 Frame
type FrameV2 []byte

// GetBytes returns the Frame as a byte array
func (frame FrameV2) GetBytes() []byte {
	return frame
}

// GetVersion returns the version of the Frame
func (frame FrameV2) GetVersion() int {
	return 2
}

// GetMessageLength returns the length of the message contained in the Frame
func (frame FrameV2) GetMessageLength() uint8 {
	return frame[1]
}

// GetIncompatibilityFlags returns the incompatibility flags for the Frame
func (frame FrameV2) GetIncompatibilityFlags() uint8 {
	return frame[2]
}

// GetCompatibilityFlags returns the compatibility flags for the Frame
func (frame FrameV2) GetCompatibilityFlags() uint8 {
	return frame[3]
}

// GetMessageSequence the sequence number of the message contained in the Frame
func (frame FrameV2) GetMessageSequence() uint8 {
	return frame[4]
}

// GetSenderSystemID returns the ID of the system that sent the Frame
func (frame FrameV2) GetSenderSystemID() uint8 {
	return frame[5]
}

// GetSenderComponentID returns the ID of the component that sent the Frame
func (frame FrameV2) GetSenderComponentID() uint8 {
	return frame[6]
}

// GetMessageID returns the ID of the message contained in the Frame
func (frame FrameV2) GetMessageID() uint32 {
	messageID := make([]byte, 4)

	messageID[0] = frame[7]
	messageID[1] = frame[8]
	messageID[2] = frame[9]
	messageID[3] = 0

	return binary.LittleEndian.Uint32(messageID)
}

// GetMessageBytes returns the message contained in the Frame as a byte array
func (frame FrameV2) GetMessageBytes() []byte {
	start := 10
	end := frame.GetMessageLength() + 10
	return frame[start:end]
}

// GetChecksum returns the checksum of the Frame contents
func (frame FrameV2) GetChecksum() uint16 {
	start := frame.GetMessageLength() + 10
	end := start + 2
	return binary.LittleEndian.Uint16(frame[start:end])
}

// GetChecksumInput returns the contents of the Frame used to calculate the checksum
func (frame FrameV2) GetChecksumInput() []byte {
	start := 1
	end := frame.GetMessageLength() + 10
	return frame[start:end]
}

// GetSignature returns the V2 Signature of the frame, if it exists
func (frame FrameV2) GetSignature() []byte {
	if frame.GetCompatibilityFlags()&CompatibilityFlagSignature != CompatibilityFlagSignature {
		return nil
	}

	return frame[frame.GetMessageLength()+12:]
}

func (frame FrameV2) String() string {
	format := fmt.Sprintf("Version: %v\nSigned: %v\nMessageLength: %v\nIncompatibilityFlags: %v\nCompatibilityFlags: %v\nSenderSystemID: %v\nSenderComponentID: %v\nMessageID: %v\nChecksum: 0x%x",
		frame.GetVersion(),
		frame.GetSignature() != nil,
		frame.GetMessageLength(),
		frame.GetIncompatibilityFlags(),
		frame.GetCompatibilityFlags(),
		frame.GetSenderSystemID(),
		frame.GetSenderComponentID(),
		frame.GetMessageID(),
		frame.GetChecksum(),
	)

	if frame.GetSignature() != nil {
		format = fmt.Sprintf(format+"\nSignature: %x", frame.GetSignature())
	}

	return format
}
