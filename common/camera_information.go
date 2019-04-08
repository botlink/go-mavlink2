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
	"math"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*CameraInformation Information about a camera */
type CameraInformation struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*FirmwareVersion Version of the camera firmware (v << 24 & 0xff = Dev, v << 16 & 0xff = Patch, v << 8 & 0xff = Minor, v & 0xff = Major) */
	FirmwareVersion uint32
	/*FocalLength Focal length */
	FocalLength float32
	/*SensorSizeH Image sensor size horizontal */
	SensorSizeH float32
	/*SensorSizeV Image sensor size vertical */
	SensorSizeV float32
	/*Flags Bitmap of camera capability flags. */
	Flags uint32
	/*ResolutionH Horizontal image resolution */
	ResolutionH uint16
	/*ResolutionV Vertical image resolution */
	ResolutionV uint16
	/*CamDefinitionVersion Camera definition version (iteration) */
	CamDefinitionVersion uint16
	/*VendorName Name of the camera vendor */
	VendorName [32]uint8
	/*ModelName Name of the camera model */
	ModelName [32]uint8
	/*LensID Reserved for a lens ID */
	LensID uint8
	/*CamDefinitionURI Camera definition URI (if any, otherwise only basic functions will be available). */
	CamDefinitionURI [140]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

// SetCamDefinitionURI encodes the input string to the CamDefinitionURI array
func (m *CameraInformation) SetCamDefinitionURI(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(140)))
	copy(m.CamDefinitionURI[:], []byte(input)[:clen])

	if len(input) > 140 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetCamDefinitionURI decodes the null-terminated string in the CamDefinitionURI
func (m *CameraInformation) GetCamDefinitionURI() string {
	clen := util.CStrLen(m.CamDefinitionURI[:])

	return string(m.CamDefinitionURI[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CameraInformation) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CameraInformation) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *CameraInformation) GetMessageName() string {
	return "CameraInformation"
}

// GetID gets the ID of the Message
func (m *CameraInformation) GetID() uint32 {
	return 259
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *CameraInformation) HasExtensionFields() bool {
	return false
}

func (m *CameraInformation) getV1Length() int {
	return 235
}

func (m *CameraInformation) getIOSlice() []byte {
	return make([]byte, 235+1)
}

// Read sets the field values of the message from the raw message payload
func (m *CameraInformation) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the CameraInformation fields
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
func (m *CameraInformation) Write(version int) (output []byte, err error) {
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
