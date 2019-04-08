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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SimState Status of simulation environment, if used */
type SimState struct {
	/*Q1 True attitude quaternion component 1, w (1 in null-rotation) */
	Q1 float32
	/*Q2 True attitude quaternion component 2, x (0 in null-rotation) */
	Q2 float32
	/*Q3 True attitude quaternion component 3, y (0 in null-rotation) */
	Q3 float32
	/*Q4 True attitude quaternion component 4, z (0 in null-rotation) */
	Q4 float32
	/*Roll Attitude roll expressed as Euler angles, not recommended except for human-readable outputs */
	Roll float32
	/*Pitch Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs */
	Pitch float32
	/*Yaw Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs */
	Yaw float32
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
	/*Lat Latitude */
	Lat float32
	/*Lon Longitude */
	Lon float32
	/*Alt Altitude */
	Alt float32
	/*StdDevHorz Horizontal position standard deviation */
	StdDevHorz float32
	/*StdDevVert Vertical position standard deviation */
	StdDevVert float32
	/*Vn True velocity in NORTH direction in earth-fixed NED frame */
	Vn float32
	/*Ve True velocity in EAST direction in earth-fixed NED frame */
	Ve float32
	/*Vd True velocity in DOWN direction in earth-fixed NED frame */
	Vd float32
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SimState) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Q1:\t%v \n")
	builder.WriteString("Q2:\t%v \n")
	builder.WriteString("Q3:\t%v \n")
	builder.WriteString("Q4:\t%v \n")
	builder.WriteString("Roll:\t%v \n")
	builder.WriteString("Pitch:\t%v \n")
	builder.WriteString("Yaw:\t%v \n")
	builder.WriteString("Xacc:\t%v [m/s/s]\n")
	builder.WriteString("Yacc:\t%v [m/s/s]\n")
	builder.WriteString("Zacc:\t%v [m/s/s]\n")
	builder.WriteString("Xgyro:\t%v [rad/s]\n")
	builder.WriteString("Ygyro:\t%v [rad/s]\n")
	builder.WriteString("Zgyro:\t%v [rad/s]\n")
	builder.WriteString("Lat:\t%v [deg]\n")
	builder.WriteString("Lon:\t%v [deg]\n")
	builder.WriteString("Alt:\t%v [m]\n")
	builder.WriteString("StdDevHorz:\t%v \n")
	builder.WriteString("StdDevVert:\t%v \n")
	builder.WriteString("Vn:\t%v [m/s]\n")
	builder.WriteString("Ve:\t%v [m/s]\n")
	builder.WriteString("Vd:\t%v [m/s]\n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Q1,
		m.Q2,
		m.Q3,
		m.Q4,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.Xacc,
		m.Yacc,
		m.Zacc,
		m.Xgyro,
		m.Ygyro,
		m.Zgyro,
		m.Lat,
		m.Lon,
		m.Alt,
		m.StdDevHorz,
		m.StdDevVert,
		m.Vn,
		m.Ve,
		m.Vd,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SimState) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SimState) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SimState) GetMessageName() string {
	return "SimState"
}

// GetID gets the ID of the Message
func (m *SimState) GetID() uint32 {
	return 108
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SimState) HasExtensionFields() bool {
	return false
}

func (m *SimState) getV1Length() int {
	return 84
}

func (m *SimState) getIOSlice() []byte {
	return make([]byte, 84+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SimState) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SimState fields
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
func (m *SimState) Write(version int) (output []byte, err error) {
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
