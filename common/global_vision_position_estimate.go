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

/*GlobalVisionPositionEstimate Global position/attitude estimate from a vision source. */
type GlobalVisionPositionEstimate struct {
	/*Usec Timestamp (UNIX time or since system boot) */
	Usec uint64
	/*X Global X position */
	X float32
	/*Y Global Y position */
	Y float32
	/*Z Global Z position */
	Z float32
	/*Roll Roll angle */
	Roll float32
	/*Pitch Pitch angle */
	Pitch float32
	/*Yaw Yaw angle */
	Yaw float32
	/*Covariance Pose covariance matrix upper right triangular (first six entries are the first ROW, next five entries are the second ROW, etc.) */
	Covariance [21]float32
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GlobalVisionPositionEstimate) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Usec:\t%v [us]\n")
	builder.WriteString("X:\t%v [m]\n")
	builder.WriteString("Y:\t%v [m]\n")
	builder.WriteString("Z:\t%v [m]\n")
	builder.WriteString("Roll:\t%v [rad]\n")
	builder.WriteString("Pitch:\t%v [rad]\n")
	builder.WriteString("Yaw:\t%v [rad]\n")
	if m.HasExtensionFieldValues {
		builder.WriteString("Covariance:\t%v\n")
	}
	format := builder.String()

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.Usec,
			m.X,
			m.Y,
			m.Z,
			m.Roll,
			m.Pitch,
			m.Yaw,
			m.Covariance,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Usec,
		m.X,
		m.Y,
		m.Z,
		m.Roll,
		m.Pitch,
		m.Yaw,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GlobalVisionPositionEstimate) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GlobalVisionPositionEstimate) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GlobalVisionPositionEstimate) GetMessageName() string {
	return "GlobalVisionPositionEstimate"
}

// GetID gets the ID of the Message
func (m *GlobalVisionPositionEstimate) GetID() uint32 {
	return 101
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GlobalVisionPositionEstimate) HasExtensionFields() bool {
	return true
}

func (m *GlobalVisionPositionEstimate) getV1Length() int {
	return 32
}

func (m *GlobalVisionPositionEstimate) getIOSlice() []byte {
	return make([]byte, 116+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GlobalVisionPositionEstimate) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GlobalVisionPositionEstimate fields
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
func (m *GlobalVisionPositionEstimate) Write(version int) (output []byte, err error) {
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
