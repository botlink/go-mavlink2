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
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*LimitsStatus Status of AP_Limits. Sent in extended status stream when AP_Limits is enabled. */
type LimitsStatus struct {
	/*LastTrigger Time (since boot) of last breach. */
	LastTrigger uint32
	/*LastAction Time (since boot) of last recovery action. */
	LastAction uint32
	/*LastRecovery Time (since boot) of last successful recovery. */
	LastRecovery uint32
	/*LastClear Time (since boot) of last all-clear. */
	LastClear uint32
	/*BreachCount Number of fence breaches. */
	BreachCount uint16
	/*LimitsState State of AP_Limits. */
	LimitsState uint8
	/*ModsEnabled AP_Limit_Module bitfield of enabled modules. */
	ModsEnabled uint8
	/*ModsRequired AP_Limit_Module bitfield of required modules. */
	ModsRequired uint8
	/*ModsTriggered AP_Limit_Module bitfield of triggered modules. */
	ModsTriggered uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *LimitsStatus) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("LastTrigger:\t%v [ms]\n")
	builder.WriteString("LastAction:\t%v [ms]\n")
	builder.WriteString("LastRecovery:\t%v [ms]\n")
	builder.WriteString("LastClear:\t%v [ms]\n")
	builder.WriteString("BreachCount:\t%v \n")
	builder.WriteString("LimitsState:\t%v \n")
	builder.WriteString("ModsEnabled:\t%v \n")
	builder.WriteString("ModsRequired:\t%v \n")
	builder.WriteString("ModsTriggered:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.LastTrigger,
		m.LastAction,
		m.LastRecovery,
		m.LastClear,
		m.BreachCount,
		m.LimitsState,
		m.ModsEnabled,
		m.ModsRequired,
		m.ModsTriggered,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *LimitsStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *LimitsStatus) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *LimitsStatus) GetMessageName() string {
	return "LimitsStatus"
}

// GetID gets the ID of the Message
func (m *LimitsStatus) GetID() uint32 {
	return 167
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *LimitsStatus) HasExtensionFields() bool {
	return false
}

func (m *LimitsStatus) getV1Length() int {
	return 22
}

func (m *LimitsStatus) getIOSlice() []byte {
	return make([]byte, 22+1)
}

// Read sets the field values of the message from the raw message payload
func (m *LimitsStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the LimitsStatus fields
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
func (m *LimitsStatus) Write(version int) (output []byte, err error) {
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
