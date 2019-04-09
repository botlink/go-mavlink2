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

/*HighresIMU The IMU readings in SI units in NED body frame */
type HighresIMU struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Xacc X acceleration */
	Xacc float32
	/*Yacc Y acceleration */
	Yacc float32
	/*Zacc Z acceleration */
	Zacc float32
	/*Xgyro Angular speed around X axis */
	Xgyro float32
	/*Ygyro Angular speed around Y axis */
	Ygyro float32
	/*Zgyro Angular speed around Z axis */
	Zgyro float32
	/*Xmag X Magnetic field */
	Xmag float32
	/*Ymag Y Magnetic field */
	Ymag float32
	/*Zmag Z Magnetic field */
	Zmag float32
	/*AbsPressure Absolute pressure */
	AbsPressure float32
	/*DiffPressure Differential pressure */
	DiffPressure float32
	/*PressureAlt Altitude calculated from pressure */
	PressureAlt float32
	/*Temperature Temperature */
	Temperature float32
	/*FieldsUpdated Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature */
	FieldsUpdated uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *HighresIMU) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "Xacc:\t%v [m/s/s]\n"
	format += "Yacc:\t%v [m/s/s]\n"
	format += "Zacc:\t%v [m/s/s]\n"
	format += "Xgyro:\t%v [rad/s]\n"
	format += "Ygyro:\t%v [rad/s]\n"
	format += "Zgyro:\t%v [rad/s]\n"
	format += "Xmag:\t%v [gauss]\n"
	format += "Ymag:\t%v [gauss]\n"
	format += "Zmag:\t%v [gauss]\n"
	format += "AbsPressure:\t%v [mbar]\n"
	format += "DiffPressure:\t%v [mbar]\n"
	format += "PressureAlt:\t%v \n"
	format += "Temperature:\t%v [degC]\n"
	format += "FieldsUpdated:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Xmag,
		m.Ymag,
		m.Zmag,
		m.AbsPressure,
		m.DiffPressure,
		m.PressureAlt,
		m.Temperature,
		m.FieldsUpdated,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HighresIMU) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HighresIMU) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *HighresIMU) GetMessageName() string {
	return "HighresIMU"
}

// GetID gets the ID of the Message
func (m *HighresIMU) GetID() uint32 {
	return 105
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *HighresIMU) HasExtensionFields() bool {
	return false
}

func (m *HighresIMU) getV1Length() int {
	return 62
}

func (m *HighresIMU) getIOSlice() []byte {
	return make([]byte, 62+1)
}

// Read sets the field values of the message from the raw message payload
func (m *HighresIMU) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the HighresIMU fields
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
func (m *HighresIMU) Write(version int) (output []byte, err error) {
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
