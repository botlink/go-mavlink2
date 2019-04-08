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

/*TrajectoryRepresentationWaypoints Describe a trajectory using an array of up-to 5 waypoints in the local frame. */
type TrajectoryRepresentationWaypoints struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*PosX X-coordinate of waypoint, set to NaN if not being used */
	PosX [5]float32
	/*PosY Y-coordinate of waypoint, set to NaN if not being used */
	PosY [5]float32
	/*PosZ Z-coordinate of waypoint, set to NaN if not being used */
	PosZ [5]float32
	/*VelX X-velocity of waypoint, set to NaN if not being used */
	VelX [5]float32
	/*VelY Y-velocity of waypoint, set to NaN if not being used */
	VelY [5]float32
	/*VelZ Z-velocity of waypoint, set to NaN if not being used */
	VelZ [5]float32
	/*AccX X-acceleration of waypoint, set to NaN if not being used */
	AccX [5]float32
	/*AccY Y-acceleration of waypoint, set to NaN if not being used */
	AccY [5]float32
	/*AccZ Z-acceleration of waypoint, set to NaN if not being used */
	AccZ [5]float32
	/*PosYaw Yaw angle, set to NaN if not being used */
	PosYaw [5]float32
	/*VelYaw Yaw rate, set to NaN if not being used */
	VelYaw [5]float32
	/*ValIDPoints Number of valid points (up-to 5 waypoints are possible) */
	ValIDPoints uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *TrajectoryRepresentationWaypoints) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("PosX:\t%v [m]\n")
	builder.WriteString("PosY:\t%v [m]\n")
	builder.WriteString("PosZ:\t%v [m]\n")
	builder.WriteString("VelX:\t%v [m/s]\n")
	builder.WriteString("VelY:\t%v [m/s]\n")
	builder.WriteString("VelZ:\t%v [m/s]\n")
	builder.WriteString("AccX:\t%v [m/s/s]\n")
	builder.WriteString("AccY:\t%v [m/s/s]\n")
	builder.WriteString("AccZ:\t%v [m/s/s]\n")
	builder.WriteString("PosYaw:\t%v [rad]\n")
	builder.WriteString("VelYaw:\t%v [rad/s]\n")
	builder.WriteString("ValIDPoints:\t%v \n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.PosX,
		m.PosY,
		m.PosZ,
		m.VelX,
		m.VelY,
		m.VelZ,
		m.AccX,
		m.AccY,
		m.AccZ,
		m.PosYaw,
		m.VelYaw,
		m.ValIDPoints,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *TrajectoryRepresentationWaypoints) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *TrajectoryRepresentationWaypoints) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *TrajectoryRepresentationWaypoints) GetMessageName() string {
	return "TrajectoryRepresentationWaypoints"
}

// GetID gets the ID of the Message
func (m *TrajectoryRepresentationWaypoints) GetID() uint32 {
	return 332
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *TrajectoryRepresentationWaypoints) HasExtensionFields() bool {
	return false
}

func (m *TrajectoryRepresentationWaypoints) getV1Length() int {
	return 229
}

func (m *TrajectoryRepresentationWaypoints) getIOSlice() []byte {
	return make([]byte, 229+1)
}

// Read sets the field values of the message from the raw message payload
func (m *TrajectoryRepresentationWaypoints) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the TrajectoryRepresentationWaypoints fields
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
func (m *TrajectoryRepresentationWaypoints) Write(version int) (output []byte, err error) {
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
