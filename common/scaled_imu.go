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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*ScaledIMU The RAW IMU readings for the usual 9DOF sensor setup. This message should contain the scaled values to the described units */
type ScaledIMU struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Xacc X acceleration */
	Xacc int16
	/*Yacc Y acceleration */
	Yacc int16
	/*Zacc Z acceleration */
	Zacc int16
	/*Xgyro Angular speed around X axis */
	Xgyro int16
	/*Ygyro Angular speed around Y axis */
	Ygyro int16
	/*Zgyro Angular speed around Z axis */
	Zgyro int16
	/*Xmag X Magnetic field */
	Xmag int16
	/*Ymag Y Magnetic field */
	Ymag int16
	/*Zmag Z Magnetic field */
	Zmag int16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *ScaledIMU) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v [ms]\n"
	format += "Xacc:\t%v [mG]\n"
	format += "Yacc:\t%v [mG]\n"
	format += "Zacc:\t%v [mG]\n"
	format += "Xgyro:\t%v [mrad/s]\n"
	format += "Ygyro:\t%v [mrad/s]\n"
	format += "Zgyro:\t%v [mrad/s]\n"
	format += "Xmag:\t%v [mT]\n"
	format += "Ymag:\t%v [mT]\n"
	format += "Zmag:\t%v [mT]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *ScaledIMU) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *ScaledIMU) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *ScaledIMU) GetMessageName() string {
	return "ScaledIMU"
}

// GetID gets the ID of the Message
func (m *ScaledIMU) GetID() uint32 {
	return 26
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *ScaledIMU) HasExtensionFields() bool {
	return false
}

func (m *ScaledIMU) getV1Length() int {
	return 22
}

func (m *ScaledIMU) getIOSlice() []byte {
	return make([]byte, 22+1)
}

// Read sets the field values of the message from the raw message payload
func (m *ScaledIMU) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the ScaledIMU fields
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
func (m *ScaledIMU) Write(version int) (output []byte, err error) {
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
