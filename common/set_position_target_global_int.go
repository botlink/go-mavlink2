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

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SetPositionTargetGlobalInt Sets a desired vehicle position, velocity, and/or acceleration in a global coordinate system (WGS84). Used by an external controller to command the vehicle (manual controller or other system). */
type SetPositionTargetGlobalInt struct {
	/*TimeBootMs Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency. */
	TimeBootMs uint32
	/*LatInt X Position in WGS84 frame */
	LatInt int32
	/*LonInt Y Position in WGS84 frame */
	LonInt int32
	/*Alt Altitude (MSL, Relative to home, or AGL - depending on frame) */
	Alt float32
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
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*CoordinateFrame Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11 */
	CoordinateFrame uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SetPositionTargetGlobalInt) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SetPositionTargetGlobalInt) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *SetPositionTargetGlobalInt) GetMessageName() string {
	return "SetPositionTargetGlobalInt"
}

// GetID gets the ID of the Message
func (m *SetPositionTargetGlobalInt) GetID() uint32 {
	return 86
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SetPositionTargetGlobalInt) HasExtensionFields() bool {
	return false
}

func (m *SetPositionTargetGlobalInt) getV1Length() int {
	return 53
}

func (m *SetPositionTargetGlobalInt) getIOSlice() []byte {
	return make([]byte, 53+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SetPositionTargetGlobalInt) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SetPositionTargetGlobalInt fields
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

	err = binary.Read(reader, binary.LittleEndian, *m)

	return
}

// Write encodes the field values of the message to a byte array
func (m *SetPositionTargetGlobalInt) Write(version int) (output []byte, err error) {
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
