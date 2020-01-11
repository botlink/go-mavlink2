package common

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2020 queue-b <https://github.com/queue-b>

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
	"text/tabwriter"

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*ComponentInformation Information about a component. For camera components instead use CAMERA_INFORMATION, and for autopilots use AUTOPILOT_VERSION. Components including GCSes should consider supporting requests of this message via MAV_CMD_REQUEST_MESSAGE. */
type ComponentInformation struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*FirmwareVersion Version of the component firmware (v << 24 & 0xff = Dev, v << 16 & 0xff = Patch, v << 8 & 0xff = Minor, v & 0xff = Major) */
	FirmwareVersion uint32
	/*HardwareVersion Version of the component hardware (v << 24 & 0xff = Dev, v << 16 & 0xff = Patch, v << 8 & 0xff = Minor, v & 0xff = Major) */
	HardwareVersion uint32
	/*CapabilityFlags Bitmap of component capability flags. */
	CapabilityFlags uint32
	/*ComponentDefinitionVersion Component definition version (iteration) */
	ComponentDefinitionVersion uint16
	/*VendorName Name of the component vendor */
	VendorName [32]uint8
	/*ModelName Name of the component model */
	ModelName [32]uint8
	/*ComponentDefinitionURI Component definition URI (if any, otherwise only basic functions will be available). The XML format is not yet specified and work in progress.  */
	ComponentDefinitionURI [140]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *ComponentInformation) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v [ms]\n"
	format += "FirmwareVersion:\t%v \n"
	format += "HardwareVersion:\t%v \n"
	format += "CapabilityFlags:\t%v \n"
	format += "ComponentDefinitionVersion:\t%v \n"
	format += "VendorName:\t%v \n"
	format += "ModelName:\t%v \n"
	format += "ComponentDefinitionURI:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.FirmwareVersion,
		m.HardwareVersion,
		m.CapabilityFlags,
		m.ComponentDefinitionVersion,
		m.VendorName,
		m.ModelName,
		m.ComponentDefinitionURI,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetComponentDefinitionURI encodes the input string to the ComponentDefinitionURI array
func (m *ComponentInformation) SetComponentDefinitionURI(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(140)))
	copy(m.ComponentDefinitionURI[:], []byte(input)[:clen])

	if len(input) > 140 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetComponentDefinitionURI decodes the null-terminated string in the ComponentDefinitionURI
func (m *ComponentInformation) GetComponentDefinitionURI() string {
	clen := util.CStrLen(m.ComponentDefinitionURI[:])

	return string(m.ComponentDefinitionURI[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *ComponentInformation) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *ComponentInformation) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *ComponentInformation) GetMessageName() string {
	return "ComponentInformation"
}

// GetID gets the ID of the Message
func (m *ComponentInformation) GetID() uint32 {
	return 395
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *ComponentInformation) HasExtensionFields() bool {
	return false
}

func (m *ComponentInformation) getV1Length() int {
	return 222
}

func (m *ComponentInformation) getIOSlice() []byte {
	return make([]byte, 222+1)
}

// Read sets the field values of the message from the raw message payload
func (m *ComponentInformation) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the ComponentInformation fields
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
func (m *ComponentInformation) Write(version int) (output []byte, err error) {
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

	output = buffer.Bytes()

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	// This also removes the trailing extra byte written for HasExtensionFieldValues
	if version == 1 {
		output = output[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		// Set HasExtensionFieldValues to zero so that it doesn't interfere with V2 truncation
		output[len(output)-1] = 0
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}
