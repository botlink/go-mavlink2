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
	"strings"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*CellularStatus Report current used cellular network status */
type CellularStatus struct {
	/*CID Cell ID. If unknown, set to: UINT32_MAX */
	CID uint32
	/*Status Status bitmap */
	Status uint16
	/*Mcc Mobile country code. If unknown, set to: UINT16_MAX */
	Mcc uint16
	/*Mnc Mobile network code. If unknown, set to: UINT16_MAX */
	Mnc uint16
	/*Lac Location area code. If unknown, set to: 0 */
	Lac uint16
	/*Type Cellular network radio type: gsm, cdma, lte... */
	Type uint8
	/*Quality Cellular network RSSI/RSRP in dBm, absolute value */
	Quality uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *CellularStatus) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("CID:\t%v \n")
	builder.WriteString("Status:\t%v \n")
	builder.WriteString("Mcc:\t%v \n")
	builder.WriteString("Mnc:\t%v \n")
	builder.WriteString("Lac:\t%v \n")
	builder.WriteString("Type:\t%v \n")
	builder.WriteString("Quality:\t%v \n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.CID,
		m.Status,
		m.Mcc,
		m.Mnc,
		m.Lac,
		m.Type,
		m.Quality,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CellularStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CellularStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *CellularStatus) GetMessageName() string {
	return "CellularStatus"
}

// GetID gets the ID of the Message
func (m *CellularStatus) GetID() uint32 {
	return 334
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *CellularStatus) HasExtensionFields() bool {
	return false
}

func (m *CellularStatus) getV1Length() int {
	return 14
}

func (m *CellularStatus) getIOSlice() []byte {
	return make([]byte, 14+1)
}

// Read sets the field values of the message from the raw message payload
func (m *CellularStatus) Read(frame mavlink2.Frame) (err error) {
	version := frame.GetVersion()

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Read V2 messages from V1 frames
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrDecodeV2MessageV1Frame
		return
	}

	// binary.Read can panic; swallow the panic and return a sane error
	defer func() {
		if r := recover(); r != nil {
			err = mavlink2.ErrPrivateField
		}
	}()

	// Get a slice of bytes long enough for the all the CellularStatus fields
	// binary.Read requires enough bytes in the reader to read all fields, even if
	// the fields are just zero values. This also simplifies handling MAVLink2
	// extensions and trailing zero truncation.
	ioSlice := m.getIOSlice()

	copy(ioSlice, frame.GetMessageBytes())

	// Indicate if
	if version == 2 && m.HasExtensionFields() {
		ioSlice[len(ioSlice)-1] = 1
	}

	reader := bytes.NewReader(ioSlice)

	err = binary.Read(reader, binary.LittleEndian, m)

	return
}

// Write encodes the field values of the message to a byte array
func (m *CellularStatus) Write(version int) (output []byte, err error) {
	var buffer bytes.Buffer

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Write V2 messages to V1 bodies
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrEncodeV2MessageV1Frame
		return
	}

	err = binary.Write(&buffer, binary.LittleEndian, *m)
	if err != nil {
		return
	}

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	if version == 1 {
		output = buffer.Bytes()[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}
