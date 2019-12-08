package asluav

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

/*SensBatmon Battery pack monitoring data for Li-Ion batteries */
type SensBatmon struct {
	/*BatmonTimestamp Time since system start */
	BatmonTimestamp uint64
	/*Temperature Battery pack temperature */
	Temperature float32
	/*Safetystatus Battery monitor safetystatus report bits in Hex */
	Safetystatus uint32
	/*Operationstatus Battery monitor operation status report bits in Hex */
	Operationstatus uint32
	/*Voltage Battery pack voltage */
	Voltage uint16
	/*Current Battery pack current */
	Current int16
	/*Batterystatus Battery monitor status report bits in Hex */
	Batterystatus uint16
	/*Serialnumber Battery monitor serial number in Hex */
	Serialnumber uint16
	/*Cellvoltage1 Battery pack cell 1 voltage */
	Cellvoltage1 uint16
	/*Cellvoltage2 Battery pack cell 2 voltage */
	Cellvoltage2 uint16
	/*Cellvoltage3 Battery pack cell 3 voltage */
	Cellvoltage3 uint16
	/*Cellvoltage4 Battery pack cell 4 voltage */
	Cellvoltage4 uint16
	/*Cellvoltage5 Battery pack cell 5 voltage */
	Cellvoltage5 uint16
	/*Cellvoltage6 Battery pack cell 6 voltage */
	Cellvoltage6 uint16
	/*Soc Battery pack state-of-charge */
	Soc uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SensBatmon) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "BatmonTimestamp:\t%v [us]\n"
	format += "Temperature:\t%v [degC]\n"
	format += "Safetystatus:\t%v \n"
	format += "Operationstatus:\t%v \n"
	format += "Voltage:\t%v [mV]\n"
	format += "Current:\t%v [mA]\n"
	format += "Batterystatus:\t%v \n"
	format += "Serialnumber:\t%v \n"
	format += "Cellvoltage1:\t%v [mV]\n"
	format += "Cellvoltage2:\t%v [mV]\n"
	format += "Cellvoltage3:\t%v [mV]\n"
	format += "Cellvoltage4:\t%v [mV]\n"
	format += "Cellvoltage5:\t%v [mV]\n"
	format += "Cellvoltage6:\t%v [mV]\n"
	format += "Soc:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.BatmonTimestamp,
		m.Temperature,
		m.Safetystatus,
		m.Operationstatus,
		m.Voltage,
		m.Current,
		m.Batterystatus,
		m.Serialnumber,
		m.Cellvoltage1,
		m.Cellvoltage2,
		m.Cellvoltage3,
		m.Cellvoltage4,
		m.Cellvoltage5,
		m.Cellvoltage6,
		m.Soc,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SensBatmon) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SensBatmon) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *SensBatmon) GetMessageName() string {
	return "SensBatmon"
}

// GetID gets the ID of the Message
func (m *SensBatmon) GetID() uint32 {
	return 209
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SensBatmon) HasExtensionFields() bool {
	return false
}

func (m *SensBatmon) getV1Length() int {
	return 41
}

func (m *SensBatmon) getIOSlice() []byte {
	return make([]byte, 41+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SensBatmon) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SensBatmon fields
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
func (m *SensBatmon) Write(version int) (output []byte, err error) {
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
