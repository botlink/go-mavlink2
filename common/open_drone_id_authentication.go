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

/*OpenDroneIDAuthentication Data for filling the OpenDroneID Authentication message. The Authentication Message defines a field that can provide a means of authenticity for the identity of the UAS (Unmanned Aircraft System). The Authentication message can have two different formats. Five data pages are supported. For data page 0, the fields PageCount, Length and TimeStamp are present and AuthData is only 17 bytes. For data page 1 through 4, PageCount,Length and TimeStamp are not present and the size of AuthData is 23 bytes. */
type OpenDroneIDAuthentication struct {
	/*Timestamp This field is only present for page 0. 32 bit Unix Timestamp in seconds since 00:00:00 01/01/2019. */
	Timestamp uint32
	/*AuthenticationType Indicates the type of authentication. */
	AuthenticationType uint8
	/*DataPage Allowed range is 0 - 4. */
	DataPage uint8
	/*PageCount This field is only present for page 0. Allowed range is 0 - 5. */
	PageCount uint8
	/*Length This field is only present for page 0. Total bytes of authentication_data from all data pages. Allowed range is 0 - 109 (17 + 23*4). */
	Length uint8
	/*AuthenticationData Opaque authentication data. For page 0, the size is only 17 bytes. For other pages, the size is 23 bytes. Shall be filled with nulls in the unused portion of the field. */
	AuthenticationData [23]uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *OpenDroneIDAuthentication) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [s]\n"
	format += "AuthenticationType:\t%v \n"
	format += "DataPage:\t%v \n"
	format += "PageCount:\t%v \n"
	format += "Length:\t%v [bytes]\n"
	format += "AuthenticationData:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.AuthenticationType,
		m.DataPage,
		m.PageCount,
		m.Length,
		m.AuthenticationData,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *OpenDroneIDAuthentication) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *OpenDroneIDAuthentication) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *OpenDroneIDAuthentication) GetMessageName() string {
	return "OpenDroneIDAuthentication"
}

// GetID gets the ID of the Message
func (m *OpenDroneIDAuthentication) GetID() uint32 {
	return 12902
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *OpenDroneIDAuthentication) HasExtensionFields() bool {
	return false
}

func (m *OpenDroneIDAuthentication) getV1Length() int {
	return 31
}

func (m *OpenDroneIDAuthentication) getIOSlice() []byte {
	return make([]byte, 31+1)
}

// Read sets the field values of the message from the raw message payload
func (m *OpenDroneIDAuthentication) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the OpenDroneIDAuthentication fields
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
func (m *OpenDroneIDAuthentication) Write(version int) (output []byte, err error) {
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
