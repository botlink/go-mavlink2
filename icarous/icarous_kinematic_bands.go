package icarous

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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*IcarousKinematicBands Kinematic multi bands (track) output from Daidalus */
type IcarousKinematicBands struct {
	/*Min1 min angle (degrees) */
	Min1 float32
	/*Max1 max angle (degrees) */
	Max1 float32
	/*Min2 min angle (degrees) */
	Min2 float32
	/*Max2 max angle (degrees) */
	Max2 float32
	/*Min3 min angle (degrees) */
	Min3 float32
	/*Max3 max angle (degrees) */
	Max3 float32
	/*Min4 min angle (degrees) */
	Min4 float32
	/*Max4 max angle (degrees) */
	Max4 float32
	/*Min5 min angle (degrees) */
	Min5 float32
	/*Max5 max angle (degrees) */
	Max5 float32
	/*Numbands Number of track bands */
	Numbands int8
	/*Type1 See the TRACK_BAND_TYPES enum. */
	Type1 uint8
	/*Type2 See the TRACK_BAND_TYPES enum. */
	Type2 uint8
	/*Type3 See the TRACK_BAND_TYPES enum. */
	Type3 uint8
	/*Type4 See the TRACK_BAND_TYPES enum. */
	Type4 uint8
	/*Type5 See the TRACK_BAND_TYPES enum. */
	Type5 uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *IcarousKinematicBands) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Min1:\t%v [deg]\n"
	format += "Max1:\t%v [deg]\n"
	format += "Min2:\t%v [deg]\n"
	format += "Max2:\t%v [deg]\n"
	format += "Min3:\t%v [deg]\n"
	format += "Max3:\t%v [deg]\n"
	format += "Min4:\t%v [deg]\n"
	format += "Max4:\t%v [deg]\n"
	format += "Min5:\t%v [deg]\n"
	format += "Max5:\t%v [deg]\n"
	format += "Numbands:\t%v \n"
	format += "Type1:\t%v \n"
	format += "Type2:\t%v \n"
	format += "Type3:\t%v \n"
	format += "Type4:\t%v \n"
	format += "Type5:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Min1,
		m.Max1,
		m.Min2,
		m.Max2,
		m.Min3,
		m.Max3,
		m.Min4,
		m.Max4,
		m.Min5,
		m.Max5,
		m.Numbands,
		m.Type1,
		m.Type2,
		m.Type3,
		m.Type4,
		m.Type5,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *IcarousKinematicBands) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *IcarousKinematicBands) GetDialect() string {
	return "icarous"
}

// GetMessageName gets the name of the Message
func (m *IcarousKinematicBands) GetMessageName() string {
	return "IcarousKinematicBands"
}

// GetID gets the ID of the Message
func (m *IcarousKinematicBands) GetID() uint32 {
	return 42001
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *IcarousKinematicBands) HasExtensionFields() bool {
	return false
}

func (m *IcarousKinematicBands) getV1Length() int {
	return 46
}

func (m *IcarousKinematicBands) getIOSlice() []byte {
	return make([]byte, 46+1)
}

// Read sets the field values of the message from the raw message payload
func (m *IcarousKinematicBands) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the IcarousKinematicBands fields
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
func (m *IcarousKinematicBands) Write(version int) (output []byte, err error) {
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
