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

/*SensorpodStatus Monitoring of sensorpod status */
type SensorpodStatus struct {
	/*Timestamp Timestamp in linuxtime (since 1.1.1970) */
	Timestamp uint64
	/*FreeSpace Free space available in recordings directory in [Gb] * 1e2 */
	FreeSpace uint16
	/*VisensorRate1 Rate of ROS topic 1 */
	VisensorRate1 uint8
	/*VisensorRate2 Rate of ROS topic 2 */
	VisensorRate2 uint8
	/*VisensorRate3 Rate of ROS topic 3 */
	VisensorRate3 uint8
	/*VisensorRate4 Rate of ROS topic 4 */
	VisensorRate4 uint8
	/*RecordingNodesCount Number of recording nodes */
	RecordingNodesCount uint8
	/*CpuTemp Temperature of sensorpod CPU in */
	CpuTemp uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SensorpodStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [ms]\n"
	format += "FreeSpace:\t%v \n"
	format += "VisensorRate1:\t%v \n"
	format += "VisensorRate2:\t%v \n"
	format += "VisensorRate3:\t%v \n"
	format += "VisensorRate4:\t%v \n"
	format += "RecordingNodesCount:\t%v \n"
	format += "CpuTemp:\t%v [degC]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.FreeSpace,
		m.VisensorRate1,
		m.VisensorRate2,
		m.VisensorRate3,
		m.VisensorRate4,
		m.RecordingNodesCount,
		m.CpuTemp,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SensorpodStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SensorpodStatus) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *SensorpodStatus) GetMessageName() string {
	return "SensorpodStatus"
}

// GetID gets the ID of the Message
func (m *SensorpodStatus) GetID() uint32 {
	return 211
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SensorpodStatus) HasExtensionFields() bool {
	return false
}

func (m *SensorpodStatus) getV1Length() int {
	return 16
}

func (m *SensorpodStatus) getIOSlice() []byte {
	return make([]byte, 16+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SensorpodStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SensorpodStatus fields
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
func (m *SensorpodStatus) Write(version int) (output []byte, err error) {
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
