package ardupilotmega

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
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*DeviceOpWrite Write registers for a device. */
type DeviceOpWrite struct {
	/*RequestID Request ID - copied to reply. */
	RequestID uint32
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*TargetComponent Component ID. */
	TargetComponent uint8
	/*Bustype The bus type. */
	Bustype uint8
	/*Bus Bus number. */
	Bus uint8
	/*Address Bus address. */
	Address uint8
	/*Busname Name of device on bus (for SPI). */
	Busname [40]byte
	/*Regstart First register to write. */
	Regstart uint8
	/*Count Count of registers to write. */
	Count uint8
	/*Data Write data. */
	Data [128]uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *DeviceOpWrite) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("RequestID:\t%v \n")
	builder.WriteString("TargetSystem:\t%v \n")
	builder.WriteString("TargetComponent:\t%v \n")
	builder.WriteString("Bustype:\t%v \n")
	builder.WriteString("Bus:\t%v \n")
	builder.WriteString("Address:\t%v \n")
	builder.WriteString("Busname:\t%v \n")
	builder.WriteString("Regstart:\t%v \n")
	builder.WriteString("Count:\t%v \n")
	builder.WriteString("Data:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.RequestID,
		m.TargetSystem,
		m.TargetComponent,
		m.Bustype,
		m.Bus,
		m.Address,
		m.Busname,
		m.Regstart,
		m.Count,
		m.Data,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetBusname encodes the input string to the Busname array
func (m *DeviceOpWrite) SetBusname(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(40)))
	copy(m.Busname[:], []byte(input)[:clen])

	if len(input) > 40 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetBusname decodes the null-terminated string in the Busname
func (m *DeviceOpWrite) GetBusname() string {
	clen := util.CStrLen(m.Busname[:])

	return string(m.Busname[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *DeviceOpWrite) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *DeviceOpWrite) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *DeviceOpWrite) GetMessageName() string {
	return "DeviceOpWrite"
}

// GetID gets the ID of the Message
func (m *DeviceOpWrite) GetID() uint32 {
	return 11002
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *DeviceOpWrite) HasExtensionFields() bool {
	return false
}

func (m *DeviceOpWrite) getV1Length() int {
	return 179
}

func (m *DeviceOpWrite) getIOSlice() []byte {
	return make([]byte, 179+1)
}

// Read sets the field values of the message from the raw message payload
func (m *DeviceOpWrite) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the DeviceOpWrite fields
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
func (m *DeviceOpWrite) Write(version int) (output []byte, err error) {
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