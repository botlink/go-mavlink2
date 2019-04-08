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
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*GPSStatus The positioning status, as reported by GPS. This message is intended to display status information about each satellite visible to the receiver. See message GLOBAL_POSITION for the global position estimate. This message can contain information for up to 20 satellites. */
type GPSStatus struct {
	/*SatellitesVisible Number of satellites visible */
	SatellitesVisible uint8
	/*SatellitePrn Global satellite ID */
	SatellitePrn [20]uint8
	/*SatelliteUsed 0: Satellite not used, 1: used for localization */
	SatelliteUsed [20]uint8
	/*SatelliteElevation Elevation (0: right on top of receiver, 90: on the horizon) of satellite */
	SatelliteElevation [20]uint8
	/*SatelliteAzIMUth Direction of satellite, 0: 0 deg, 255: 360 deg. */
	SatelliteAzIMUth [20]uint8
	/*SatelliteSnr Signal to noise ratio of satellite */
	SatelliteSnr [20]uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GPSStatus) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("SatellitesVisible:\t%v \n")
	builder.WriteString("SatellitePrn:\t%v \n")
	builder.WriteString("SatelliteUsed:\t%v \n")
	builder.WriteString("SatelliteElevation:\t%v [deg]\n")
	builder.WriteString("SatelliteAzIMUth:\t%v [deg]\n")
	builder.WriteString("SatelliteSnr:\t%v [dB]\n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SatellitesVisible,
		m.SatellitePrn,
		m.SatelliteUsed,
		m.SatelliteElevation,
		m.SatelliteAzIMUth,
		m.SatelliteSnr,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GPSStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GPSStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GPSStatus) GetMessageName() string {
	return "GPSStatus"
}

// GetID gets the ID of the Message
func (m *GPSStatus) GetID() uint32 {
	return 25
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GPSStatus) HasExtensionFields() bool {
	return false
}

func (m *GPSStatus) getV1Length() int {
	return 101
}

func (m *GPSStatus) getIOSlice() []byte {
	return make([]byte, 101+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GPSStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GPSStatus fields
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
func (m *GPSStatus) Write(version int) (output []byte, err error) {
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
