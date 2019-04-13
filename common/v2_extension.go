package common

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
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*V2Extension Message implementing parts of the V2 payload specs in V1 frames for transitional support. */
type V2Extension struct {
	/*MessageType A code that identifies the software component that understands this message (analogous to USB device classes or mime type strings).  If this code is less than 32768, it is considered a 'registered' protocol extension and the corresponding entry should be added to https://github.com/mavlink/mavlink/extension-message-ids.xml.  Software creators can register blocks of message IDs as needed (useful for GCS specific metadata, etc...). Message_types greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase. */
	MessageType uint16
	/*TargetNetwork Network ID (0 for broadcast) */
	TargetNetwork uint8
	/*TargetSystem System ID (0 for broadcast) */
	TargetSystem uint8
	/*TargetComponent Component ID (0 for broadcast) */
	TargetComponent uint8
	/*Payload Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields.  The entire content of this block is opaque unless you understand any the encoding message_type.  The particular encoding used can be extension specific and might not always be documented as part of the mavlink specification. */
	Payload [249]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
	/*PayloadLength Length of the variable length Payload*/
	PayloadLength uint8
}

func (m *V2Extension) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	format += "MessageType:\t%v\n"
	format += "TargetNetwork:	%v\n"
	format += "TargetSystem:\t%v\n"
	format += "TargetComponent:\t%v\n"
	format += "Payload:\t%v\n"
	format += "PayloadLength:\t%v\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.MessageType,
		m.TargetNetwork,
		m.TargetSystem,
		m.TargetComponent,
		m.Payload,
		m.PayloadLength,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *V2Extension) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *V2Extension) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *V2Extension) GetMessageName() string {
	return "V2Extension"
}

// GetID gets the ID of the Message
func (m *V2Extension) GetID() uint32 {
	return 248
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *V2Extension) HasExtensionFields() bool {
	return false
}

func (m *V2Extension) getV1Length() int {
	return 5
}

func (m *V2Extension) getIOSlice() []byte {
	// 249 payload + 5 description + bool + uint8
	return make([]byte, 256)
}

// GetPayloadSlice gets the valid Payload bytes as a slice
func (m *V2Extension) GetPayload() []byte {
	return m.Payload[0:m.PayloadLength]
}

// SetPayloadSlice sets the Payload from the slice value provided
func (m *V2Extension) SetPayloadSlice(payload []byte) error {
	if len(payload) > len(m.Payload) {
		copy(m.Payload[:], payload[:len(m.Payload)])
		m.PayloadLength = uint8(len(m.Payload))
		return ErrValueTooLong
	}

	copy(m.Payload[:], payload)
	m.PayloadLength = uint8(len(payload))
}

// Read sets the field values of the message from the raw message payload
func (m *V2Extension) Read(frame mavlink2.Frame) (err error) {
	version := frame.GetVersion()

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// binary.Read can panic; swallow the panic and return a sane error
	defer func() {
		if r := recover(); r != nil {
			err = mavlink2.ErrPrivateField
		}
	}()

	// Get a slice with enough capacity for a full read
	ioBytes := m.getIOSlice()

	copy(ioBytes, frame.GetMessageBytes())

	// Add the length of the payload to the end of the array so that it is read in as PayloadLength
	ioBytes[len(ioBytes)-1] = uint8(math.Max(0, float64(frame.GetMessageLength()-5)))

	buffer := bytes.NewBuffer(ioBytes)

	err = binary.Read(buffer, binary.LittleEndian, m)

	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *V2Extension) Write(version int) (output []byte, err error) {
	var buffer bytes.Buffer

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	err = binary.Write(&buffer, binary.LittleEndian, m)
	if err != nil {
		return
	}

	output = buffer.Bytes()
	// V2Extension packets are always variable length,
	// regardless of whether they are in a V1 or V2 stream
	// Cut off the trailing bytes past the end of the payload
	output = output[:5+m.PayloadLength]

	if version == 2 {
		output = util.TruncateV2(output)
	}

	return
}
