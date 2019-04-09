package matrixpilot

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

/*SerialUdbExtraF20 Backwards compatible version of SERIAL_UDB_EXTRA F20 format */
type SerialUdbExtraF20 struct {
	/*SueTrimValueInput1 SUE UDB PWM Trim Value on Input 1 */
	SueTrimValueInput1 int16
	/*SueTrimValueInput2 SUE UDB PWM Trim Value on Input 2 */
	SueTrimValueInput2 int16
	/*SueTrimValueInput3 SUE UDB PWM Trim Value on Input 3 */
	SueTrimValueInput3 int16
	/*SueTrimValueInput4 SUE UDB PWM Trim Value on Input 4 */
	SueTrimValueInput4 int16
	/*SueTrimValueInput5 SUE UDB PWM Trim Value on Input 5 */
	SueTrimValueInput5 int16
	/*SueTrimValueInput6 SUE UDB PWM Trim Value on Input 6 */
	SueTrimValueInput6 int16
	/*SueTrimValueInput7 SUE UDB PWM Trim Value on Input 7 */
	SueTrimValueInput7 int16
	/*SueTrimValueInput8 SUE UDB PWM Trim Value on Input 8 */
	SueTrimValueInput8 int16
	/*SueTrimValueInput9 SUE UDB PWM Trim Value on Input 9 */
	SueTrimValueInput9 int16
	/*SueTrimValueInput10 SUE UDB PWM Trim Value on Input 10 */
	SueTrimValueInput10 int16
	/*SueTrimValueInput11 SUE UDB PWM Trim Value on Input 11 */
	SueTrimValueInput11 int16
	/*SueTrimValueInput12 SUE UDB PWM Trim Value on Input 12 */
	SueTrimValueInput12 int16
	/*SueNumberOfInputs SUE Number of Input Channels */
	SueNumberOfInputs uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF20) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("SueTrimValueInput1:\t%v \n")
	builder.WriteString("SueTrimValueInput2:\t%v \n")
	builder.WriteString("SueTrimValueInput3:\t%v \n")
	builder.WriteString("SueTrimValueInput4:\t%v \n")
	builder.WriteString("SueTrimValueInput5:\t%v \n")
	builder.WriteString("SueTrimValueInput6:\t%v \n")
	builder.WriteString("SueTrimValueInput7:\t%v \n")
	builder.WriteString("SueTrimValueInput8:\t%v \n")
	builder.WriteString("SueTrimValueInput9:\t%v \n")
	builder.WriteString("SueTrimValueInput10:\t%v \n")
	builder.WriteString("SueTrimValueInput11:\t%v \n")
	builder.WriteString("SueTrimValueInput12:\t%v \n")
	builder.WriteString("SueNumberOfInputs:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueTrimValueInput1,
		m.SueTrimValueInput2,
		m.SueTrimValueInput3,
		m.SueTrimValueInput4,
		m.SueTrimValueInput5,
		m.SueTrimValueInput6,
		m.SueTrimValueInput7,
		m.SueTrimValueInput8,
		m.SueTrimValueInput9,
		m.SueTrimValueInput10,
		m.SueTrimValueInput11,
		m.SueTrimValueInput12,
		m.SueNumberOfInputs,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF20) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF20) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF20) GetMessageName() string {
	return "SerialUdbExtraF20"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF20) GetID() uint32 {
	return 186
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF20) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF20) getV1Length() int {
	return 25
}

func (m *SerialUdbExtraF20) getIOSlice() []byte {
	return make([]byte, 25+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF20) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SerialUdbExtraF20 fields
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
func (m *SerialUdbExtraF20) Write(version int) (output []byte, err error) {
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
