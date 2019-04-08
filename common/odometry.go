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

/*Odometry Odometry message to communicate odometry information with an external interface. Fits ROS REP 147 standard for aerial vehicles (http://www.ros.org/reps/rep-0147.html). */
type Odometry struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*X X Position */
	X float32
	/*Y Y Position */
	Y float32
	/*Z Z Position */
	Z float32
	/*Q Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation) */
	Q [4]float32
	/*Vx X linear speed */
	Vx float32
	/*Vy Y linear speed */
	Vy float32
	/*Vz Z linear speed */
	Vz float32
	/*Rollspeed Roll angular speed */
	Rollspeed float32
	/*Pitchspeed Pitch angular speed */
	Pitchspeed float32
	/*Yawspeed Yaw angular speed */
	Yawspeed float32
	/*PoseCovariance Row-major representation of a 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array. */
	PoseCovariance [21]float32
	/*VelocityCovariance Row-major representation of a 6x6 velocity cross-covariance matrix upper right triangle (states: vx, vy, vz, rollspeed, pitchspeed, yawspeed; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array. */
	VelocityCovariance [21]float32
	/*FrameID Coordinate frame of reference for the pose data. */
	FrameID uint8
	/*ChildFrameID Coordinate frame of reference for the velocity in free space (twist) data. */
	ChildFrameID uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *Odometry) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("X:\t%v [m]\n")
	builder.WriteString("Y:\t%v [m]\n")
	builder.WriteString("Z:\t%v [m]\n")
	builder.WriteString("Q:\t%v \n")
	builder.WriteString("Vx:\t%v [m/s]\n")
	builder.WriteString("Vy:\t%v [m/s]\n")
	builder.WriteString("Vz:\t%v [m/s]\n")
	builder.WriteString("Rollspeed:\t%v [rad/s]\n")
	builder.WriteString("Pitchspeed:\t%v [rad/s]\n")
	builder.WriteString("Yawspeed:\t%v [rad/s]\n")
	builder.WriteString("PoseCovariance:\t%v \n")
	builder.WriteString("VelocityCovariance:\t%v \n")
	builder.WriteString("FrameID:\t%v \n")
	builder.WriteString("ChildFrameID:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.X,
		m.Y,
		m.Z,
		m.Q,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Rollspeed,
		m.Pitchspeed,
		m.Yawspeed,
		m.PoseCovariance,
		m.VelocityCovariance,
		m.FrameID,
		m.ChildFrameID,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *Odometry) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *Odometry) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *Odometry) GetMessageName() string {
	return "Odometry"
}

// GetID gets the ID of the Message
func (m *Odometry) GetID() uint32 {
	return 331
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *Odometry) HasExtensionFields() bool {
	return false
}

func (m *Odometry) getV1Length() int {
	return 230
}

func (m *Odometry) getIOSlice() []byte {
	return make([]byte, 230+1)
}

// Read sets the field values of the message from the raw message payload
func (m *Odometry) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the Odometry fields
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
func (m *Odometry) Write(version int) (output []byte, err error) {
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
