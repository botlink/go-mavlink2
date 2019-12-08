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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*LandingTarget The location of a landing target. See: https://mavlink.io/en/services/landing_target.html */
type LandingTarget struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*AngleX X-axis angular offset of the target from the center of the image */
	AngleX float32
	/*AngleY Y-axis angular offset of the target from the center of the image */
	AngleY float32
	/*Distance Distance to the target from the vehicle */
	Distance float32
	/*SizeX Size of target along x-axis */
	SizeX float32
	/*SizeY Size of target along y-axis */
	SizeY float32
	/*TargetNum The ID of the target if multiple targets are present */
	TargetNum uint8
	/*Frame Coordinate frame used for following fields. */
	Frame uint8
	/*X X Position of the landing target in MAV_FRAME */
	X float32
	/*Y Y Position of the landing target in MAV_FRAME */
	Y float32
	/*Z Z Position of the landing target in MAV_FRAME */
	Z float32
	/*Q Quaternion of landing target orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0) */
	Q [4]float32
	/*Type Type of landing target */
	Type uint8
	/*PositionValID Boolean indicating whether the position fields (x, y, z, q, type) contain valid target position information (valid: 1, invalid: 0). Default is 0 (invalid). */
	PositionValID uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *LandingTarget) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "AngleX:\t%v [rad]\n"
	format += "AngleY:\t%v [rad]\n"
	format += "Distance:\t%v [m]\n"
	format += "SizeX:\t%v [rad]\n"
	format += "SizeY:\t%v [rad]\n"
	format += "TargetNum:\t%v \n"
	format += "Frame:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "X:\t%v\n"
		format += "Y:\t%v\n"
		format += "Z:\t%v\n"
		format += "Q:\t%v\n"
		format += "Type:\t%v\n"
		format += "PositionValID:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.TimeUsec,
			m.AngleX,
			m.AngleY,
			m.Distance,
			m.SizeX,
			m.SizeY,
			m.TargetNum,
			m.Frame,
			m.X,
			m.Y,
			m.Z,
			m.Q,
			m.Type,
			m.PositionValID,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.AngleX,
		m.AngleY,
		m.Distance,
		m.SizeX,
		m.SizeY,
		m.TargetNum,
		m.Frame,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *LandingTarget) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *LandingTarget) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *LandingTarget) GetMessageName() string {
	return "LandingTarget"
}

// GetID gets the ID of the Message
func (m *LandingTarget) GetID() uint32 {
	return 149
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *LandingTarget) HasExtensionFields() bool {
	return true
}

func (m *LandingTarget) getV1Length() int {
	return 30
}

func (m *LandingTarget) getIOSlice() []byte {
	return make([]byte, 60+1)
}

// Read sets the field values of the message from the raw message payload
func (m *LandingTarget) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the LandingTarget fields
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
func (m *LandingTarget) Write(version int) (output []byte, err error) {
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

	output = buffer.Bytes()

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	// This also removes the trailing extra byte written for HasExtensionFieldValues
	if version == 1 {
		output = output[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		// Set HasExtensionFieldValues to zero so that it doesn't interfere with V2 truncation
		output[len(output)-1] = 0
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}
