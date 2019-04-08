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

/*PositionTargetLocalNed Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_LOCAL_NED if the vehicle is being controlled this way. */
type PositionTargetLocalNed struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*X X Position in NED frame */
	X float32
	/*Y Y Position in NED frame */
	Y float32
	/*Z Z Position in NED frame (note, altitude is negative in NED) */
	Z float32
	/*Vx X velocity in NED frame */
	Vx float32
	/*Vy Y velocity in NED frame */
	Vy float32
	/*Vz Z velocity in NED frame */
	Vz float32
	/*Afx X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afx float32
	/*Afy Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afy float32
	/*Afz Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afz float32
	/*Yaw yaw setpoint */
	Yaw float32
	/*YawRate yaw rate setpoint */
	YawRate float32
	/*TypeMask Bitmap to indicate which dimensions should be ignored by the vehicle. */
	TypeMask uint16
	/*CoordinateFrame Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9 */
	CoordinateFrame uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *PositionTargetLocalNed) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeBootMs:\t%v [ms]\n")
	builder.WriteString("X:\t%v [m]\n")
	builder.WriteString("Y:\t%v [m]\n")
	builder.WriteString("Z:\t%v [m]\n")
	builder.WriteString("Vx:\t%v [m/s]\n")
	builder.WriteString("Vy:\t%v [m/s]\n")
	builder.WriteString("Vz:\t%v [m/s]\n")
	builder.WriteString("Afx:\t%v [m/s/s]\n")
	builder.WriteString("Afy:\t%v [m/s/s]\n")
	builder.WriteString("Afz:\t%v [m/s/s]\n")
	builder.WriteString("Yaw:\t%v [rad]\n")
	builder.WriteString("YawRate:\t%v [rad/s]\n")
	builder.WriteString("TypeMask:\t%v \n")
	builder.WriteString("CoordinateFrame:\t%v \n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.X,
		m.Y,
		m.Z,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Afx,
		m.Afy,
		m.Afz,
		m.Yaw,
		m.YawRate,
		m.TypeMask,
		m.CoordinateFrame,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *PositionTargetLocalNed) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *PositionTargetLocalNed) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *PositionTargetLocalNed) GetMessageName() string {
	return "PositionTargetLocalNed"
}

// GetID gets the ID of the Message
func (m *PositionTargetLocalNed) GetID() uint32 {
	return 85
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *PositionTargetLocalNed) HasExtensionFields() bool {
	return false
}

func (m *PositionTargetLocalNed) getV1Length() int {
	return 51
}

func (m *PositionTargetLocalNed) getIOSlice() []byte {
	return make([]byte, 51+1)
}

// Read sets the field values of the message from the raw message payload
func (m *PositionTargetLocalNed) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the PositionTargetLocalNed fields
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
func (m *PositionTargetLocalNed) Write(version int) (output []byte, err error) {
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
