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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SensMppt Maximum Power Point Tracker (MPPT) sensor data for solar module power performance tracking */
type SensMppt struct {
	/*MpptTimestamp  MPPT last timestamp  */
	MpptTimestamp uint64
	/*Mppt1Volt  MPPT1 voltage  */
	Mppt1Volt float32
	/*Mppt1Amp  MPPT1 current  */
	Mppt1Amp float32
	/*Mppt2Volt  MPPT2 voltage  */
	Mppt2Volt float32
	/*Mppt2Amp  MPPT2 current  */
	Mppt2Amp float32
	/*Mppt3Volt MPPT3 voltage  */
	Mppt3Volt float32
	/*Mppt3Amp  MPPT3 current  */
	Mppt3Amp float32
	/*Mppt1Pwm  MPPT1 pwm  */
	Mppt1Pwm uint16
	/*Mppt2Pwm  MPPT2 pwm  */
	Mppt2Pwm uint16
	/*Mppt3Pwm  MPPT3 pwm  */
	Mppt3Pwm uint16
	/*Mppt1Status  MPPT1 status  */
	Mppt1Status uint8
	/*Mppt2Status  MPPT2 status  */
	Mppt2Status uint8
	/*Mppt3Status  MPPT3 status  */
	Mppt3Status uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SensMppt) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "MpptTimestamp:\t%v [us]\n"
	format += "Mppt1Volt:\t%v [V]\n"
	format += "Mppt1Amp:\t%v [A]\n"
	format += "Mppt2Volt:\t%v [V]\n"
	format += "Mppt2Amp:\t%v [A]\n"
	format += "Mppt3Volt:\t%v [V]\n"
	format += "Mppt3Amp:\t%v [A]\n"
	format += "Mppt1Pwm:\t%v [us]\n"
	format += "Mppt2Pwm:\t%v [us]\n"
	format += "Mppt3Pwm:\t%v [us]\n"
	format += "Mppt1Status:\t%v \n"
	format += "Mppt2Status:\t%v \n"
	format += "Mppt3Status:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.MpptTimestamp,
		m.Mppt1Volt,
		m.Mppt1Amp,
		m.Mppt2Volt,
		m.Mppt2Amp,
		m.Mppt3Volt,
		m.Mppt3Amp,
		m.Mppt1Pwm,
		m.Mppt2Pwm,
		m.Mppt3Pwm,
		m.Mppt1Status,
		m.Mppt2Status,
		m.Mppt3Status,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SensMppt) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SensMppt) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *SensMppt) GetMessageName() string {
	return "SensMppt"
}

// GetID gets the ID of the Message
func (m *SensMppt) GetID() uint32 {
	return 202
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SensMppt) HasExtensionFields() bool {
	return false
}

func (m *SensMppt) getV1Length() int {
	return 41
}

func (m *SensMppt) getIOSlice() []byte {
	return make([]byte, 41+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SensMppt) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SensMppt fields
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
func (m *SensMppt) Write(version int) (output []byte, err error) {
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
