package asluav

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
	"text/tabwriter"

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*CommandIntStamped Message encoding a command with parameters as scaled integers and additional metadata. Scaling depends on the actual command value. */
type CommandIntStamped struct {
	/*VehicleTimestamp Microseconds elapsed since vehicle boot */
	VehicleTimestamp uint64
	/*UtcTime UTC time, seconds elapsed since 01.01.1970 */
	UtcTime uint32
	/*Param1 PARAM1, see MAV_CMD enum */
	Param1 float32
	/*Param2 PARAM2, see MAV_CMD enum */
	Param2 float32
	/*Param3 PARAM3, see MAV_CMD enum */
	Param3 float32
	/*Param4 PARAM4, see MAV_CMD enum */
	Param4 float32
	/*X PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7 */
	X int32
	/*Y PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7 */
	Y int32
	/*Z PARAM7 / z position: global: altitude in meters (MSL, WGS84, AGL or relative to home - depending on frame). */
	Z float32
	/*Command The scheduled action for the mission item, as defined by MAV_CMD enum */
	Command uint16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*Frame The coordinate system of the COMMAND, as defined by MAV_FRAME enum */
	Frame uint8
	/*Current false:0, true:1 */
	Current uint8
	/*Autocontinue autocontinue to next wp */
	Autocontinue uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *CommandIntStamped) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "VehicleTimestamp:\t%v \n"
	format += "UtcTime:\t%v \n"
	format += "Param1:\t%v \n"
	format += "Param2:\t%v \n"
	format += "Param3:\t%v \n"
	format += "Param4:\t%v \n"
	format += "X:\t%v \n"
	format += "Y:\t%v \n"
	format += "Z:\t%v \n"
	format += "Command:\t%v \n"
	format += "TargetSystem:\t%v \n"
	format += "TargetComponent:\t%v \n"
	format += "Frame:\t%v \n"
	format += "Current:\t%v \n"
	format += "Autocontinue:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.VehicleTimestamp,
		m.UtcTime,
		m.Param1,
		m.Param2,
		m.Param3,
		m.Param4,
		m.X,
		m.Y,
		m.Z,
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
func (m *CommandIntStamped) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CommandIntStamped) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *CommandIntStamped) GetMessageName() string {
	return "CommandIntStamped"
}

// GetID gets the ID of the Message
func (m *CommandIntStamped) GetID() uint32 {
	return 78
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *CommandIntStamped) HasExtensionFields() bool {
	return false
}

func (m *CommandIntStamped) getV1Length() int {
	return 47
}

func (m *CommandIntStamped) getIOSlice() []byte {
	return make([]byte, 47+1)
}

// Read sets the field values of the message from the raw message payload
func (m *CommandIntStamped) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the CommandIntStamped fields
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
func (m *CommandIntStamped) Write(version int) (output []byte, err error) {
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
