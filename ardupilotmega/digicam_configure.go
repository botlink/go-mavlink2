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

/*DigicamConfigure Configure on-board Camera Control System. */
type DigicamConfigure struct {
	/*ExtraValue Correspondent value to given extra_param. */
	ExtraValue float32
	/*ShutterSpeed Divisor number //e.g. 1000 means 1/1000 (0 means ignore). */
	ShutterSpeed uint16
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*TargetComponent Component ID. */
	TargetComponent uint8
	/*Mode Mode enumeration from 1 to N //P, TV, AV, M, etc. (0 means ignore). */
	Mode uint8
	/*Aperture F stop number x 10 //e.g. 28 means 2.8 (0 means ignore). */
	Aperture uint8
	/*Iso ISO enumeration from 1 to N //e.g. 80, 100, 200, Etc (0 means ignore). */
	Iso uint8
	/*ExposureType Exposure type enumeration from 1 to N (0 means ignore). */
	ExposureType uint8
	/*CommandID Command Identity (incremental loop: 0 to 255). //A command sent multiple times will be executed or pooled just once. */
	CommandID uint8
	/*EngineCutOff Main engine cut-off time before camera trigger (0 means no cut-off). */
	EngineCutOff uint8
	/*ExtraParam Extra parameters enumeration (0 means ignore). */
	ExtraParam uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *DigicamConfigure) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "ExtraValue:\t%v \n"
	format += "ShutterSpeed:\t%v \n"
	format += "TargetSystem:\t%v \n"
	format += "TargetComponent:\t%v \n"
	format += "Mode:\t%v \n"
	format += "Aperture:\t%v \n"
	format += "Iso:\t%v \n"
	format += "ExposureType:\t%v \n"
	format += "CommandID:\t%v \n"
	format += "EngineCutOff:\t%v [ds]\n"
	format += "ExtraParam:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.ExtraValue,
		m.ShutterSpeed,
		m.TargetSystem,
		m.TargetComponent,
		m.Mode,
		m.Aperture,
		m.Iso,
		m.ExposureType,
		m.CommandID,
		m.EngineCutOff,
		m.ExtraParam,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *DigicamConfigure) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *DigicamConfigure) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *DigicamConfigure) GetMessageName() string {
	return "DigicamConfigure"
}

// GetID gets the ID of the Message
func (m *DigicamConfigure) GetID() uint32 {
	return 154
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *DigicamConfigure) HasExtensionFields() bool {
	return false
}

func (m *DigicamConfigure) getV1Length() int {
	return 15
}

func (m *DigicamConfigure) getIOSlice() []byte {
	return make([]byte, 15+1)
}

// Read sets the field values of the message from the raw message payload
func (m *DigicamConfigure) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the DigicamConfigure fields
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
func (m *DigicamConfigure) Write(version int) (output []byte, err error) {
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
