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

/*VIDeoStreamInformation Information about video stream */
type VIDeoStreamInformation struct {
	/*Framerate Frame rate. */
	Framerate float32
	/*Bitrate Bit rate. */
	Bitrate uint32
	/*Flags Bitmap of stream status flags. */
	Flags uint16
	/*ResolutionH Horizontal resolution. */
	ResolutionH uint16
	/*ResolutionV Vertical resolution. */
	ResolutionV uint16
	/*Rotation Video image rotation clockwise. */
	Rotation uint16
	/*Hfov Horizontal Field of view. */
	Hfov uint16
	/*StreamID Video Stream ID (1 for first, 2 for second, etc.) */
	StreamID uint8
	/*Count Number of streams available. */
	Count uint8
	/*Type Type of stream. */
	Type uint8
	/*Name Stream name. */
	Name [32]byte
	/*URI Video stream URI (TCP or RTSP URI ground station should connect to) or port number (UDP port ground station should listen to). */
	URI [160]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

// SetName encodes the input string to the Name array
func (m *VIDeoStreamInformation) SetName(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(32)))
	copy(m.Name[:], []byte(input)[:clen])

	if len(input) > 32 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetName decodes the null-terminated string in the Name
func (m *VIDeoStreamInformation) GetName() string {
	clen := util.CStrLen(m.Name[:])

	return string(m.Name[:clen])
}

// SetURI encodes the input string to the URI array
func (m *VIDeoStreamInformation) SetURI(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(160)))
	copy(m.URI[:], []byte(input)[:clen])

	if len(input) > 160 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetURI decodes the null-terminated string in the URI
func (m *VIDeoStreamInformation) GetURI() string {
	clen := util.CStrLen(m.URI[:])

	return string(m.URI[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *VIDeoStreamInformation) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *VIDeoStreamInformation) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *VIDeoStreamInformation) GetMessageName() string {
	return "VIDeoStreamInformation"
}

// GetID gets the ID of the Message
func (m *VIDeoStreamInformation) GetID() uint32 {
	return 269
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *VIDeoStreamInformation) HasExtensionFields() bool {
	return false
}

func (m *VIDeoStreamInformation) getV1Length() int {
	return 213
}

func (m *VIDeoStreamInformation) getIOSlice() []byte {
	return make([]byte, 213+1)
}

// Read sets the field values of the message from the raw message payload
func (m *VIDeoStreamInformation) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the VIDeoStreamInformation fields
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
func (m *VIDeoStreamInformation) Write(version int) (output []byte, err error) {
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
