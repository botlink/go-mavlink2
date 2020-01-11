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

/*IsbdLinkStatus Status of the Iridium SBD link. */
type IsbdLinkStatus struct {
	/*Timestamp Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	Timestamp uint64
	/*LastHeartbeat Timestamp of the last successful sbd session. The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	LastHeartbeat uint64
	/*FailedSessions Number of failed SBD sessions. */
	FailedSessions uint16
	/*SuccessfulSessions Number of successful SBD sessions. */
	SuccessfulSessions uint16
	/*SignalQuality Signal quality equal to the number of bars displayed on the ISU signal strength indicator. Range is 0 to 5, where 0 indicates no signal and 5 indicates maximum signal strength. */
	SignalQuality uint8
	/*RingPending 1: Ring call pending, 0: No call pending. */
	RingPending uint8
	/*TxSessionPending 1: Transmission session pending, 0: No transmission session pending. */
	TxSessionPending uint8
	/*RxSessionPending 1: Receiving session pending, 0: No receiving session pending. */
	RxSessionPending uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *IsbdLinkStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [us]\n"
	format += "LastHeartbeat:\t%v [us]\n"
	format += "FailedSessions:\t%v \n"
	format += "SuccessfulSessions:\t%v \n"
	format += "SignalQuality:\t%v \n"
	format += "RingPending:\t%v \n"
	format += "TxSessionPending:\t%v \n"
	format += "RxSessionPending:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.LastHeartbeat,
		m.FailedSessions,
		m.SuccessfulSessions,
		m.SignalQuality,
		m.RingPending,
		m.TxSessionPending,
		m.RxSessionPending,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *IsbdLinkStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *IsbdLinkStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *IsbdLinkStatus) GetMessageName() string {
	return "IsbdLinkStatus"
}

// GetID gets the ID of the Message
func (m *IsbdLinkStatus) GetID() uint32 {
	return 335
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *IsbdLinkStatus) HasExtensionFields() bool {
	return false
}

func (m *IsbdLinkStatus) getV1Length() int {
	return 24
}

func (m *IsbdLinkStatus) getIOSlice() []byte {
	return make([]byte, 24+1)
}

// Read sets the field values of the message from the raw message payload
func (m *IsbdLinkStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the IsbdLinkStatus fields
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
func (m *IsbdLinkStatus) Write(version int) (output []byte, err error) {
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
