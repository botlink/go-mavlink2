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

/*DistanceSensor Distance sensor information for an onboard rangefinder. */
type DistanceSensor struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*MinDistance Minimum distance the sensor can measure */
	MinDistance uint16
	/*MaxDistance Maximum distance the sensor can measure */
	MaxDistance uint16
	/*CurrentDistance Current distance reading */
	CurrentDistance uint16
	/*Type Type of distance sensor. */
	Type uint8
	/*ID Onboard ID of the sensor */
	ID uint8
	/*Orientation Direction the sensor faces. downward-facing: ROTATION_PITCH_270, upward-facing: ROTATION_PITCH_90, backward-facing: ROTATION_PITCH_180, forward-facing: ROTATION_NONE, left-facing: ROTATION_YAW_90, right-facing: ROTATION_YAW_270 */
	Orientation uint8
	/*Covariance Measurement variance. Max standard deviation is 6cm. 255 if unknown. */
	Covariance uint8
	/*HorizontalFov Horizontal Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0. */
	HorizontalFov float32
	/*VerticalFov Vertical Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0. */
	VerticalFov float32
	/*Quaternion Quaternion of the sensor orientation in vehicle body frame (w, x, y, z order, zero-rotation is 1, 0, 0, 0). Zero-rotation is along the vehicle body x-axis. This field is required if the orientation is set to MAV_SENSOR_ROTATION_CUSTOM. Set it to 0 if invalid." */
	Quaternion [4]float32
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *DistanceSensor) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v [ms]\n"
	format += "MinDistance:\t%v [cm]\n"
	format += "MaxDistance:\t%v [cm]\n"
	format += "CurrentDistance:\t%v [cm]\n"
	format += "Type:\t%v \n"
	format += "ID:\t%v \n"
	format += "Orientation:\t%v \n"
	format += "Covariance:\t%v [cm^2]\n"
	if m.HasExtensionFieldValues {
		format += "HorizontalFov:\t%v\n"
		format += "VerticalFov:\t%v\n"
		format += "Quaternion:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.TimeBootMs,
			m.MinDistance,
			m.MaxDistance,
			m.CurrentDistance,
			m.Type,
			m.ID,
			m.Orientation,
			m.Covariance,
			m.HorizontalFov,
			m.VerticalFov,
			m.Quaternion,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.MinDistance,
		m.MaxDistance,
		m.CurrentDistance,
		m.Type,
		m.ID,
		m.Orientation,
		m.Covariance,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *DistanceSensor) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *DistanceSensor) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *DistanceSensor) GetMessageName() string {
	return "DistanceSensor"
}

// GetID gets the ID of the Message
func (m *DistanceSensor) GetID() uint32 {
	return 132
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *DistanceSensor) HasExtensionFields() bool {
	return true
}

func (m *DistanceSensor) getV1Length() int {
	return 14
}

func (m *DistanceSensor) getIOSlice() []byte {
	return make([]byte, 38+1)
}

// Read sets the field values of the message from the raw message payload
func (m *DistanceSensor) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the DistanceSensor fields
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
func (m *DistanceSensor) Write(version int) (output []byte, err error) {
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
