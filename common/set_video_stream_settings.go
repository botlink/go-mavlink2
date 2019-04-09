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
	"math"
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SetVIDeoStreamSettings Message that sets video stream settings */
type SetVIDeoStreamSettings struct {
	/*Framerate Frame rate (set to -1 for highest framerate possible) */
	Framerate float32
	/*Bitrate Bit rate (set to -1 for auto) */
	Bitrate uint32
	/*ResolutionH Horizontal resolution (set to -1 for highest resolution possible) */
	ResolutionH uint16
	/*ResolutionV Vertical resolution (set to -1 for highest resolution possible) */
	ResolutionV uint16
	/*Rotation Video image rotation clockwise (0-359 degrees) */
	Rotation uint16
	/*TargetSystem system ID of the target */
	TargetSystem uint8
	/*TargetComponent component ID of the target */
	TargetComponent uint8
	/*CameraID Camera ID (1 for first, 2 for second, etc.) */
	CameraID uint8
	/*URI Video stream URI */
	URI [230]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SetVIDeoStreamSettings) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Framerate:\t%v [Hz]\n")
	builder.WriteString("Bitrate:\t%v [bits/s]\n")
	builder.WriteString("ResolutionH:\t%v [pix]\n")
	builder.WriteString("ResolutionV:\t%v [pix]\n")
	builder.WriteString("Rotation:\t%v [deg]\n")
	builder.WriteString("TargetSystem:\t%v \n")
	builder.WriteString("TargetComponent:\t%v \n")
	builder.WriteString("CameraID:\t%v \n")
	builder.WriteString("URI:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Framerate,
		m.Bitrate,
		m.ResolutionH,
		m.ResolutionV,
		m.Rotation,
		m.TargetSystem,
		m.TargetComponent,
		m.CameraID,
		m.URI,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetURI encodes the input string to the URI array
func (m *SetVIDeoStreamSettings) SetURI(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(230)))
	copy(m.URI[:], []byte(input)[:clen])

	if len(input) > 230 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetURI decodes the null-terminated string in the URI
func (m *SetVIDeoStreamSettings) GetURI() string {
	clen := util.CStrLen(m.URI[:])

	return string(m.URI[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SetVIDeoStreamSettings) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SetVIDeoStreamSettings) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SetVIDeoStreamSettings) GetMessageName() string {
	return "SetVIDeoStreamSettings"
}

// GetID gets the ID of the Message
func (m *SetVIDeoStreamSettings) GetID() uint32 {
	return 270
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SetVIDeoStreamSettings) HasExtensionFields() bool {
	return false
}

func (m *SetVIDeoStreamSettings) getV1Length() int {
	return 247
}

func (m *SetVIDeoStreamSettings) getIOSlice() []byte {
	return make([]byte, 247+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SetVIDeoStreamSettings) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SetVIDeoStreamSettings fields
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
func (m *SetVIDeoStreamSettings) Write(version int) (output []byte, err error) {
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