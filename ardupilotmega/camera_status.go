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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*CameraStatus Camera Event. */
type CameraStatus struct {
	/*TimeUsec Image timestamp (since UNIX epoch, according to camera clock). */
	TimeUsec uint64
	/*P1 Parameter 1 (meaning depends on event_id, see CAMERA_STATUS_TYPES enum). */
	P1 float32
	/*P2 Parameter 2 (meaning depends on event_id, see CAMERA_STATUS_TYPES enum). */
	P2 float32
	/*P3 Parameter 3 (meaning depends on event_id, see CAMERA_STATUS_TYPES enum). */
	P3 float32
	/*P4 Parameter 4 (meaning depends on event_id, see CAMERA_STATUS_TYPES enum). */
	P4 float32
	/*ImgIDx Image index. */
	ImgIDx uint16
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*CamIDx Camera ID. */
	CamIDx uint8
	/*EventID Event type. */
	EventID uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *CameraStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "P1:\t%v \n"
	format += "P2:\t%v \n"
	format += "P3:\t%v \n"
	format += "P4:\t%v \n"
	format += "ImgIDx:\t%v \n"
	format += "TargetSystem:\t%v \n"
	format += "CamIDx:\t%v \n"
	format += "EventID:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.P1,
		m.P2,
		m.P3,
		m.P4,
		m.ImgIDx,
		m.TargetSystem,
		m.CamIDx,
		m.EventID,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CameraStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CameraStatus) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *CameraStatus) GetMessageName() string {
	return "CameraStatus"
}

// GetID gets the ID of the Message
func (m *CameraStatus) GetID() uint32 {
	return 179
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *CameraStatus) HasExtensionFields() bool {
	return false
}

func (m *CameraStatus) getV1Length() int {
	return 29
}

func (m *CameraStatus) getIOSlice() []byte {
	return make([]byte, 29+1)
}

// Read sets the field values of the message from the raw message payload
func (m *CameraStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the CameraStatus fields
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
func (m *CameraStatus) Write(version int) (output []byte, err error) {
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
