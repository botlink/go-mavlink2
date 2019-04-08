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

/*SetHomePosition The position the system will return to and land on. The position is set automatically by the system during the takeoff in case it was not explicitly set by the operator before or after. The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface. Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach. The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector. */
type SetHomePosition struct {
	/*Latitude Latitude (WGS84) */
	Latitude int32
	/*Longitude Longitude (WGS84) */
	Longitude int32
	/*Altitude Altitude (MSL). Positive for up. */
	Altitude int32
	/*X Local X position of this position in the local coordinate frame */
	X float32
	/*Y Local Y position of this position in the local coordinate frame */
	Y float32
	/*Z Local Z position of this position in the local coordinate frame */
	Z float32
	/*Q World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground */
	Q [4]float32
	/*ApproachX Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone. */
	ApproachX float32
	/*ApproachY Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone. */
	ApproachY float32
	/*ApproachZ Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone. */
	ApproachZ float32
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SetHomePosition) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Latitude:\t%v [degE7]\n")
	builder.WriteString("Longitude:\t%v [degE7]\n")
	builder.WriteString("Altitude:\t%v [mm]\n")
	builder.WriteString("X:\t%v [m]\n")
	builder.WriteString("Y:\t%v [m]\n")
	builder.WriteString("Z:\t%v [m]\n")
	builder.WriteString("Q:\t%v \n")
	builder.WriteString("ApproachX:\t%v [m]\n")
	builder.WriteString("ApproachY:\t%v [m]\n")
	builder.WriteString("ApproachZ:\t%v [m]\n")
	builder.WriteString("TargetSystem:\t%v \n")
	if m.HasExtensionFieldValues {
		builder.WriteString("TimeUsec:\t%v\n")
	}
	format := builder.String()

	if m.HasExtensionFieldValues {
		return fmt.Sprintf(
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.Latitude,
			m.Longitude,
			m.Altitude,
			m.X,
			m.Y,
			m.Z,
			m.Q,
			m.ApproachX,
			m.ApproachY,
			m.ApproachZ,
			m.TargetSystem,
			m.TimeUsec,
		)
	}

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Latitude,
		m.Longitude,
		m.Altitude,
		m.X,
		m.Y,
		m.Z,
		m.Q,
		m.ApproachX,
		m.ApproachY,
		m.ApproachZ,
		m.TargetSystem,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SetHomePosition) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SetHomePosition) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SetHomePosition) GetMessageName() string {
	return "SetHomePosition"
}

// GetID gets the ID of the Message
func (m *SetHomePosition) GetID() uint32 {
	return 243
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SetHomePosition) HasExtensionFields() bool {
	return true
}

func (m *SetHomePosition) getV1Length() int {
	return 53
}

func (m *SetHomePosition) getIOSlice() []byte {
	return make([]byte, 61+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SetHomePosition) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SetHomePosition fields
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
func (m *SetHomePosition) Write(version int) (output []byte, err error) {
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
