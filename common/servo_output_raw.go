package common

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

/*ServoOutputRaw Superseded by ACTUATOR_OUTPUT_STATUS. The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. */
type ServoOutputRaw struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint32
	/*Servo1Raw Servo output 1 value */
	Servo1Raw uint16
	/*Servo2Raw Servo output 2 value */
	Servo2Raw uint16
	/*Servo3Raw Servo output 3 value */
	Servo3Raw uint16
	/*Servo4Raw Servo output 4 value */
	Servo4Raw uint16
	/*Servo5Raw Servo output 5 value */
	Servo5Raw uint16
	/*Servo6Raw Servo output 6 value */
	Servo6Raw uint16
	/*Servo7Raw Servo output 7 value */
	Servo7Raw uint16
	/*Servo8Raw Servo output 8 value */
	Servo8Raw uint16
	/*Port Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX. */
	Port uint8
	/*Servo9Raw Servo output 9 value */
	Servo9Raw uint16
	/*Servo10Raw Servo output 10 value */
	Servo10Raw uint16
	/*Servo11Raw Servo output 11 value */
	Servo11Raw uint16
	/*Servo12Raw Servo output 12 value */
	Servo12Raw uint16
	/*Servo13Raw Servo output 13 value */
	Servo13Raw uint16
	/*Servo14Raw Servo output 14 value */
	Servo14Raw uint16
	/*Servo15Raw Servo output 15 value */
	Servo15Raw uint16
	/*Servo16Raw Servo output 16 value */
	Servo16Raw uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *ServoOutputRaw) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "Servo1Raw:\t%v [us]\n"
	format += "Servo2Raw:\t%v [us]\n"
	format += "Servo3Raw:\t%v [us]\n"
	format += "Servo4Raw:\t%v [us]\n"
	format += "Servo5Raw:\t%v [us]\n"
	format += "Servo6Raw:\t%v [us]\n"
	format += "Servo7Raw:\t%v [us]\n"
	format += "Servo8Raw:\t%v [us]\n"
	format += "Port:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "Servo9Raw:\t%v\n"
		format += "Servo10Raw:\t%v\n"
		format += "Servo11Raw:\t%v\n"
		format += "Servo12Raw:\t%v\n"
		format += "Servo13Raw:\t%v\n"
		format += "Servo14Raw:\t%v\n"
		format += "Servo15Raw:\t%v\n"
		format += "Servo16Raw:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.TimeUsec,
			m.Servo1Raw,
			m.Servo2Raw,
			m.Servo3Raw,
			m.Servo4Raw,
			m.Servo5Raw,
			m.Servo6Raw,
			m.Servo7Raw,
			m.Servo8Raw,
			m.Port,
			m.Servo9Raw,
			m.Servo10Raw,
			m.Servo11Raw,
			m.Servo12Raw,
			m.Servo13Raw,
			m.Servo14Raw,
			m.Servo15Raw,
			m.Servo16Raw,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.Servo1Raw,
		m.Servo2Raw,
		m.Servo3Raw,
		m.Servo4Raw,
		m.Servo5Raw,
		m.Servo6Raw,
		m.Servo7Raw,
		m.Servo8Raw,
		m.Port,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *ServoOutputRaw) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *ServoOutputRaw) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *ServoOutputRaw) GetMessageName() string {
	return "ServoOutputRaw"
}

// GetID gets the ID of the Message
func (m *ServoOutputRaw) GetID() uint32 {
	return 36
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *ServoOutputRaw) HasExtensionFields() bool {
	return true
}

func (m *ServoOutputRaw) getV1Length() int {
	return 21
}

func (m *ServoOutputRaw) getIOSlice() []byte {
	return make([]byte, 37+1)
}

// Read sets the field values of the message from the raw message payload
func (m *ServoOutputRaw) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the ServoOutputRaw fields
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
func (m *ServoOutputRaw) Write(version int) (output []byte, err error) {
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
