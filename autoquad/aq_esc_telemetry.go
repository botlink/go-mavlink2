package autoquad

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

/*AqEscTelemetry Sends ESC32 telemetry data for up to 4 motors. Multiple messages may be sent in sequence when system has > 4 motors. Data is described as follows:
	// unsigned int state :   3;
    // unsigned int vin :	  12;  // x 100
    // unsigned int amps :	  14;  // x 100
    // unsigned int rpm :	  15;
    // unsigned int duty :	  8;   // x (255/100)
    // - Data Version 2 -
    //     unsigned int errors :    9;   // Bad detects error count
    // - Data Version 3 -
    //     unsigned int temp   :    9;   // (Deg C + 32) * 4
    // unsigned int errCode : 3;
*/
type AqEscTelemetry struct {
	/*TimeBootMs Timestamp of the component clock since boot time in ms. */
	TimeBootMs uint32
	/*Data0 Data bits 1-32 for each ESC. */
	Data0 [4]uint32
	/*Data1 Data bits 33-64 for each ESC. */
	Data1 [4]uint32
	/*StatusAge Age of each ESC telemetry reading in ms compared to boot time. A value of 0xFFFF means timeout/no data. */
	StatusAge [4]uint16
	/*Seq Sequence number of message (first set of 4 motors is #1, next 4 is #2, etc). */
	Seq uint8
	/*NumMotors Total number of active ESCs/motors on the system. */
	NumMotors uint8
	/*NumInSeq Number of active ESCs in this sequence (1 through this many array members will be populated with data) */
	NumInSeq uint8
	/*EscID ESC/Motor ID */
	EscID [4]uint8
	/*DataVersion Version of data structure (determines contents). */
	DataVersion [4]uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AqEscTelemetry) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v \n"
	format += "Data0:\t%v \n"
	format += "Data1:\t%v \n"
	format += "StatusAge:\t%v \n"
	format += "Seq:\t%v \n"
	format += "NumMotors:\t%v \n"
	format += "NumInSeq:\t%v \n"
	format += "EscID:\t%v \n"
	format += "DataVersion:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.Data0,
		m.Data1,
		m.StatusAge,
		m.Seq,
		m.NumMotors,
		m.NumInSeq,
		m.EscID,
		m.DataVersion,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AqEscTelemetry) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AqEscTelemetry) GetDialect() string {
	return "autoquad"
}

// GetMessageName gets the name of the Message
func (m *AqEscTelemetry) GetMessageName() string {
	return "AqEscTelemetry"
}

// GetID gets the ID of the Message
func (m *AqEscTelemetry) GetID() uint32 {
	return 152
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AqEscTelemetry) HasExtensionFields() bool {
	return false
}

func (m *AqEscTelemetry) getV1Length() int {
	return 55
}

func (m *AqEscTelemetry) getIOSlice() []byte {
	return make([]byte, 55+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AqEscTelemetry) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the AqEscTelemetry fields
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
func (m *AqEscTelemetry) Write(version int) (output []byte, err error) {
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
