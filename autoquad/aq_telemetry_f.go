package autoquad

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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*AqTelemetryF Sends up to 20 raw float values. */
type AqTelemetryF struct {
	/*Value1 value1 */
	Value1 float32
	/*Value2 value2 */
	Value2 float32
	/*Value3 value3 */
	Value3 float32
	/*Value4 value4 */
	Value4 float32
	/*Value5 value5 */
	Value5 float32
	/*Value6 value6 */
	Value6 float32
	/*Value7 value7 */
	Value7 float32
	/*Value8 value8 */
	Value8 float32
	/*Value9 value9 */
	Value9 float32
	/*Value10 value10 */
	Value10 float32
	/*Value11 value11 */
	Value11 float32
	/*Value12 value12 */
	Value12 float32
	/*Value13 value13 */
	Value13 float32
	/*Value14 value14 */
	Value14 float32
	/*Value15 value15 */
	Value15 float32
	/*Value16 value16 */
	Value16 float32
	/*Value17 value17 */
	Value17 float32
	/*Value18 value18 */
	Value18 float32
	/*Value19 value19 */
	Value19 float32
	/*Value20 value20 */
	Value20 float32
	/*Index Index of message */
	Index uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AqTelemetryF) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Value1:\t%v \n"
	format += "Value2:\t%v \n"
	format += "Value3:\t%v \n"
	format += "Value4:\t%v \n"
	format += "Value5:\t%v \n"
	format += "Value6:\t%v \n"
	format += "Value7:\t%v \n"
	format += "Value8:\t%v \n"
	format += "Value9:\t%v \n"
	format += "Value10:\t%v \n"
	format += "Value11:\t%v \n"
	format += "Value12:\t%v \n"
	format += "Value13:\t%v \n"
	format += "Value14:\t%v \n"
	format += "Value15:\t%v \n"
	format += "Value16:\t%v \n"
	format += "Value17:\t%v \n"
	format += "Value18:\t%v \n"
	format += "Value19:\t%v \n"
	format += "Value20:\t%v \n"
	format += "Index:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Value1,
		m.Value2,
		m.Value3,
		m.Value4,
		m.Value5,
		m.Value6,
		m.Value7,
		m.Value8,
		m.Value9,
		m.Value10,
		m.Value11,
		m.Value12,
		m.Value13,
		m.Value14,
		m.Value15,
		m.Value16,
		m.Value17,
		m.Value18,
		m.Value19,
		m.Value20,
		m.Index,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AqTelemetryF) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AqTelemetryF) GetDialect() string {
	return "autoquad"
}

// GetMessageName gets the name of the Message
func (m *AqTelemetryF) GetMessageName() string {
	return "AqTelemetryF"
}

// GetID gets the ID of the Message
func (m *AqTelemetryF) GetID() uint32 {
	return 150
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AqTelemetryF) HasExtensionFields() bool {
	return false
}

func (m *AqTelemetryF) getV1Length() int {
	return 82
}

func (m *AqTelemetryF) getIOSlice() []byte {
	return make([]byte, 82+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AqTelemetryF) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the AqTelemetryF fields
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
func (m *AqTelemetryF) Write(version int) (output []byte, err error) {
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
