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

/*AutopilotVersion Version and capability of autopilot software. This should be emitted in response to a MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES command. */
type AutopilotVersion struct {
	/*Capabilities Bitmap of capabilities */
	Capabilities uint64
	/*UID UID if provided by hardware (see uid2) */
	UID uint64
	/*FlightSwVersion Firmware version number */
	FlightSwVersion uint32
	/*MIDdlewareSwVersion Middleware version number */
	MIDdlewareSwVersion uint32
	/*OsSwVersion Operating system version number */
	OsSwVersion uint32
	/*BoardVersion HW / board version (last 8 bytes should be silicon ID, if any) */
	BoardVersion uint32
	/*VendorID ID of the board vendor */
	VendorID uint16
	/*ProductID ID of the product */
	ProductID uint16
	/*FlightCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	FlightCustomVersion [8]uint8
	/*MIDdlewareCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	MIDdlewareCustomVersion [8]uint8
	/*OsCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	OsCustomVersion [8]uint8
	/*UID2 UID if provided by hardware (supersedes the uid field. If this is non-zero, use this field, otherwise use uid) */
	UID2 [18]uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AutopilotVersion) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Capabilities:\t%v \n"
	format += "UID:\t%v \n"
	format += "FlightSwVersion:\t%v \n"
	format += "MIDdlewareSwVersion:\t%v \n"
	format += "OsSwVersion:\t%v \n"
	format += "BoardVersion:\t%v \n"
	format += "VendorID:\t%v \n"
	format += "ProductID:\t%v \n"
	format += "FlightCustomVersion:\t%v \n"
	format += "MIDdlewareCustomVersion:\t%v \n"
	format += "OsCustomVersion:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "UID2:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.Capabilities,
			m.UID,
			m.FlightSwVersion,
			m.MIDdlewareSwVersion,
			m.OsSwVersion,
			m.BoardVersion,
			m.VendorID,
			m.ProductID,
			m.FlightCustomVersion,
			m.MIDdlewareCustomVersion,
			m.OsCustomVersion,
			m.UID2,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Capabilities,
		m.UID,
		m.FlightSwVersion,
		m.MIDdlewareSwVersion,
		m.OsSwVersion,
		m.BoardVersion,
		m.VendorID,
		m.ProductID,
		m.FlightCustomVersion,
		m.MIDdlewareCustomVersion,
		m.OsCustomVersion,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AutopilotVersion) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AutopilotVersion) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *AutopilotVersion) GetMessageName() string {
	return "AutopilotVersion"
}

// GetID gets the ID of the Message
func (m *AutopilotVersion) GetID() uint32 {
	return 148
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AutopilotVersion) HasExtensionFields() bool {
	return true
}

func (m *AutopilotVersion) getV1Length() int {
	return 60
}

func (m *AutopilotVersion) getIOSlice() []byte {
	return make([]byte, 78+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AutopilotVersion) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the AutopilotVersion fields
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
func (m *AutopilotVersion) Write(version int) (output []byte, err error) {
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
