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
	"bytes"
	"encoding/binary"
	"fmt"
	"text/tabwriter"
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
	Bytes() []byte
}

// FrameV1 represents a MAVLink V1 Frame
type FrameV1 []byte

// Bytes returns the Frame as a byte array
func (frame FrameV1) Bytes() []byte {
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
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	fmt.Fprintf(writer, "Version:\t%v\nMessageLength:\t%v\nSenderSystemID:\t%v\nSenderComponentID:\t%v\nMessageID:\t%v\nChecksum:\t%v",
		frame.GetVersion(),
		frame.GetMessageLength(),
		frame.GetSenderSystemID(),
		frame.GetSenderComponentID(),
		frame.GetMessageID(),
		frame.GetChecksum(),
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// FrameV2 represents a MAVLink V1 Frame
type FrameV2 []byte

// GetBytes returns the Frame as a byte array
func (frame FrameV2) Bytes() []byte {
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

	copy(messageID[0:3], frame[7:10])

	return binary.LittleEndian.Uint32(messageID)
}

// GetMessageBytes returns the message contained in the Frame as a byte array
func (frame FrameV2) GetMessageBytes() []byte {
	start := 10
	end := start + int(frame.GetMessageLength())
	return frame[start:end]
}

// GetChecksum returns the checksum of the Frame contents
func (frame FrameV2) GetChecksum() uint16 {
	start := int(frame.GetMessageLength()) + 10
	end := start + 2
	return binary.LittleEndian.Uint16(frame[start:end])
}

// GetChecksumInput returns the contents of the Frame used to calculate the checksum
func (frame FrameV2) GetChecksumInput() []byte {
	start := 1
	end := int(frame.GetMessageLength()) + 10
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
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	fmt.Fprintf(writer, "Version:\t%v\nSigned:\t%v\nMessageLength:\t%v\nIncompatibilityFlags:\t%v\nCompatibilityFlags:\t%v\nSenderSystemID:\t%v\nSenderComponentID:\t%v\nMessageID:\t%v\nChecksum:\t0x%x",
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
		fmt.Fprintf(writer, "\nSignature:\t%x", frame.GetSignature())
	}

	writer.Flush()
	return string(buffer.Bytes())
}

// ErrFrameTooShort indicates the Frame did not contain enough bytes to complete the requested operation
var ErrFrameTooShort = fmt.Errorf("Not enough bytes in Frame")

// ErrUnknownStartByte indicates the Frame started with an unknown start byte (magic byte)
var ErrUnknownStartByte = fmt.Errorf("Start byte is not 0x%x (V1) or 0x%x (V2)", V1StartByte, V2StartByte)

// ErrMessageTooShort indicates that the Message length in the header is shorter than the
// minimum length allowed for the Message
var ErrMessageTooShort = fmt.Errorf("Message too short")

// ErrMessageTooLong indicates that the Message length in the header is longer than the
// maximum length allowed for the Message
var ErrMessageTooLong = fmt.Errorf("Message too long")

// ErrInvalidChecksum indicates that the Message checksum is invalid
var ErrInvalidChecksum = fmt.Errorf("Invalid checksum")

// ErrUnknownMessage indicates that the Message is not defined in one of the loaded Dialects
var ErrUnknownMessage = fmt.Errorf("Unknown message")

// ErrPrivateField indicates that a Message contains private fields and cannot be unmarshalled using binary.Read
var ErrPrivateField = fmt.Errorf("Struct contains one or more private fields that cannot be unmarshalled using binary.Read")

// ErrUnsupportedVersion indicates that the version requested is not supported by this library
var ErrUnsupportedVersion = fmt.Errorf("The version requested is not supported by this library")

// ErrDecodeV2MessageV1Frame indicates that a request was made to decode a V2 Message from a V1 Frame
var ErrDecodeV2MessageV1Frame = fmt.Errorf("Unable to decode a V2 message from a V1 Frame")

// ErrEncodeV2MessageV1Frame indicates that a request was made to encode a V2 Message for use with a V1 Frame
var ErrEncodeV2MessageV1Frame = fmt.Errorf("Unable to encode a V2 message to a V1 Frame")

// ErrValueTooLong indicates that a Message field is unable to hold the requested string
// and that the string was truncated to fit
var ErrValueTooLong = fmt.Errorf("The input string was too long and was truncated")

// FrameFromBytes returns a Frame containing the input bytes,
// or an error if the first byte is not a valid frame start
func FrameFromBytes(raw []byte) (frame Frame, err error) {
	if len(raw) == 0 {
		err = ErrFrameTooShort
		return
	}

	startByte := raw[0]

	if startByte == V1StartByte {
		frame = FrameV1(raw)
	} else if startByte == V2StartByte {
		frame = FrameV2(raw)
	} else {
		err = ErrUnknownStartByte
	}

	return
}
