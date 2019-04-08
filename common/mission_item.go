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

/*MissionItem Message encoding a mission item. This message is emitted to announce
  the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). See also https://mavlink.io/en/services/mission.html. */
type MissionItem struct {
	/*Param1 PARAM1, see MAV_CMD enum */
	Param1 float32
	/*Param2 PARAM2, see MAV_CMD enum */
	Param2 float32
	/*Param3 PARAM3, see MAV_CMD enum */
	Param3 float32
	/*Param4 PARAM4, see MAV_CMD enum */
	Param4 float32
	/*X PARAM5 / local: X coordinate, global: latitude */
	X float32
	/*Y PARAM6 / local: Y coordinate, global: longitude */
	Y float32
	/*Z PARAM7 / local: Z coordinate, global: altitude (relative or absolute, depending on frame). */
	Z float32
	/*Seq Sequence */
	Seq uint16
	/*Command The scheduled action for the waypoint. */
	Command uint16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*Frame The coordinate system of the waypoint. */
	Frame uint8
	/*Current false:0, true:1 */
	Current uint8
	/*Autocontinue Autocontinue to next waypoint */
	Autocontinue uint8
	/*MissionType Mission type. */
	MissionType uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *MissionItem) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Param1:\t%v \n")
	builder.WriteString("Param2:\t%v \n")
	builder.WriteString("Param3:\t%v \n")
	builder.WriteString("Param4:\t%v \n")
	builder.WriteString("X:\t%v \n")
	builder.WriteString("Y:\t%v \n")
	builder.WriteString("Z:\t%v \n")
	builder.WriteString("Seq:\t%v \n")
	builder.WriteString("Command:\t%v \n")
	builder.WriteString("TargetSystem:\t%v \n")
	builder.WriteString("TargetComponent:\t%v \n")
	builder.WriteString("Frame:\t%v \n")
	builder.WriteString("Current:\t%v \n")
	builder.WriteString("Autocontinue:\t%v \n")
	if m.HasExtensionFieldValues {
		builder.WriteString("MissionType:\t%v\n")
	}
	format := builder.String()

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.Param1,
			m.Param2,
			m.Param3,
			m.Param4,
			m.X,
			m.Y,
			m.Z,
			m.Seq,
			m.Command,
			m.TargetSystem,
			m.TargetComponent,
			m.Frame,
			m.Current,
			m.Autocontinue,
			m.MissionType,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.X,
		m.Y,
		m.Z,
		m.Seq,
		m.Command,
		m.TargetSystem,
		m.TargetComponent,
		m.Frame,
		m.Current,
		m.Autocontinue,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *MissionItem) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *MissionItem) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *MissionItem) GetMessageName() string {
	return "MissionItem"
}

// GetID gets the ID of the Message
func (m *MissionItem) GetID() uint32 {
	return 39
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *MissionItem) HasExtensionFields() bool {
	return true
}

func (m *MissionItem) getV1Length() int {
	return 37
}

func (m *MissionItem) getIOSlice() []byte {
	return make([]byte, 38+1)
}

// Read sets the field values of the message from the raw message payload
func (m *MissionItem) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the MissionItem fields
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
func (m *MissionItem) Write(version int) (output []byte, err error) {
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
