package ardupilotmega

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

/*CameraFeedback Camera Capture Feedback. */
type CameraFeedback struct {
	/*TimeUsec Image timestamp (since UNIX epoch), as passed in by CAMERA_STATUS message (or autopilot if no CCB). */
	TimeUsec uint64
	/*Lat Latitude. */
	Lat int32
	/*Lng Longitude. */
	Lng int32
	/*AltMsl Altitude (MSL). */
	AltMsl float32
	/*AltRel Altitude (Relative to HOME location). */
	AltRel float32
	/*Roll Camera Roll angle (earth frame, +-180). */
	Roll float32
	/*Pitch Camera Pitch angle (earth frame, +-180). */
	Pitch float32
	/*Yaw Camera Yaw (earth frame, 0-360, true). */
	Yaw float32
	/*FocLen Focal Length. */
	FocLen float32
	/*ImgIDx Image index. */
	ImgIDx uint16
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*CamIDx Camera ID. */
	CamIDx uint8
	/*Flags Feedback flags. */
	Flags uint8
	/*CompletedCaptures Completed image captures. */
	CompletedCaptures uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *CameraFeedback) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("Lat:\t%v [degE7]\n")
	builder.WriteString("Lng:\t%v [degE7]\n")
	builder.WriteString("AltMsl:\t%v [m]\n")
	builder.WriteString("AltRel:\t%v [m]\n")
	builder.WriteString("Roll:\t%v [deg]\n")
	builder.WriteString("Pitch:\t%v [deg]\n")
	builder.WriteString("Yaw:\t%v [deg]\n")
	builder.WriteString("FocLen:\t%v [mm]\n")
	builder.WriteString("ImgIDx:\t%v \n")
	builder.WriteString("TargetSystem:\t%v \n")
	builder.WriteString("CamIDx:\t%v \n")
	builder.WriteString("Flags:\t%v \n")
	if m.HasExtensionFieldValues {
		builder.WriteString("CompletedCaptures:\t%v\n")
	}
	format := builder.String()

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.TimeUsec,
			m.Lat,
			m.Lng,
			m.AltMsl,
			m.AltRel,
			m.Roll,
			m.Pitch,
			m.Yaw,
			m.FocLen,
			m.ImgIDx,
			m.TargetSystem,
			m.CamIDx,
			m.Flags,
			m.CompletedCaptures,
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
		m.Lat,
		m.Lng,
		m.AltMsl,
		m.AltRel,
		m.Roll,
		m.Pitch,
		m.Yaw,
		m.FocLen,
		m.ImgIDx,
		m.TargetSystem,
		m.CamIDx,
		m.Flags,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CameraFeedback) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CameraFeedback) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *CameraFeedback) GetMessageName() string {
	return "CameraFeedback"
}

// GetID gets the ID of the Message
func (m *CameraFeedback) GetID() uint32 {
	return 180
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *CameraFeedback) HasExtensionFields() bool {
	return true
}

func (m *CameraFeedback) getV1Length() int {
	return 45
}

func (m *CameraFeedback) getIOSlice() []byte {
	return make([]byte, 47+1)
}

// Read sets the field values of the message from the raw message payload
func (m *CameraFeedback) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the CameraFeedback fields
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
func (m *CameraFeedback) Write(version int) (output []byte, err error) {
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
