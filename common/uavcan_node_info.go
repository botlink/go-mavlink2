package common

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
	"math"
	"text/tabwriter"

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*UavcanNodeInfo General information describing a particular UAVCAN node. Please refer to the definition of the UAVCAN service "uavcan.protocol.GetNodeInfo" for the background information. This message should be emitted by the system whenever a new node appears online, or an existing node reboots. Additionally, it can be emitted upon request from the other end of the MAVLink channel (see MAV_CMD_UAVCAN_GET_NODE_INFO). It is also not prohibited to emit this message unconditionally at a low frequency. The UAVCAN specification is available at http://uavcan.org. */
type UavcanNodeInfo struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*UptimeSec Time since the start-up of the node. */
	UptimeSec uint32
	/*SwVcsCommit Version control system (VCS) revision identifier (e.g. git short commit hash). Zero if unknown. */
	SwVcsCommit uint32
	/*Name Node name string. For example, "sapog.px4.io". */
	Name [80]byte
	/*HWVersionMajor Hardware major version number. */
	HWVersionMajor uint8
	/*HWVersionMinor Hardware minor version number. */
	HWVersionMinor uint8
	/*HWUniqueID Hardware unique 128-bit ID. */
	HWUniqueID [16]uint8
	/*SwVersionMajor Software major version number. */
	SwVersionMajor uint8
	/*SwVersionMinor Software minor version number. */
	SwVersionMinor uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *UavcanNodeInfo) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "UptimeSec:\t%v [s]\n"
	format += "SwVcsCommit:\t%v \n"
	format += "Name:\t%v \n"
	format += "HWVersionMajor:\t%v \n"
	format += "HWVersionMinor:\t%v \n"
	format += "HWUniqueID:\t%v \n"
	format += "SwVersionMajor:\t%v \n"
	format += "SwVersionMinor:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.UptimeSec,
		m.SwVcsCommit,
		m.Name,
		m.HWVersionMajor,
		m.HWVersionMinor,
		m.HWUniqueID,
		m.SwVersionMajor,
		m.SwVersionMinor,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetName encodes the input string to the Name array
func (m *UavcanNodeInfo) SetName(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(80)))
	copy(m.Name[:], []byte(input)[:clen])

	if len(input) > 80 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetName decodes the null-terminated string in the Name
func (m *UavcanNodeInfo) GetName() string {
	clen := util.CStrLen(m.Name[:])

	return string(m.Name[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *UavcanNodeInfo) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *UavcanNodeInfo) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *UavcanNodeInfo) GetMessageName() string {
	return "UavcanNodeInfo"
}

// GetID gets the ID of the Message
func (m *UavcanNodeInfo) GetID() uint32 {
	return 311
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *UavcanNodeInfo) HasExtensionFields() bool {
	return false
}

func (m *UavcanNodeInfo) getV1Length() int {
	return 116
}

func (m *UavcanNodeInfo) getIOSlice() []byte {
	return make([]byte, 116+1)
}

// Read sets the field values of the message from the raw message payload
func (m *UavcanNodeInfo) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the UavcanNodeInfo fields
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
func (m *UavcanNodeInfo) Write(version int) (output []byte, err error) {
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
