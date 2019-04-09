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

/*SerialUdbExtraF2A Backwards compatible MAVLink version of SERIAL_UDB_EXTRA - F2: Format Part A */
type SerialUdbExtraF2A struct {
	/*SueTime Serial UDB Extra Time */
	SueTime uint32
	/*SueLatitude Serial UDB Extra Latitude */
	SueLatitude int32
	/*SueLongitude Serial UDB Extra Longitude */
	SueLongitude int32
	/*SueAltitude Serial UDB Extra Altitude */
	SueAltitude int32
	/*SueWaypointIndex Serial UDB Extra Waypoint Index */
	SueWaypointIndex uint16
	/*SueRmat0 Serial UDB Extra Rmat 0 */
	SueRmat0 int16
	/*SueRmat1 Serial UDB Extra Rmat 1 */
	SueRmat1 int16
	/*SueRmat2 Serial UDB Extra Rmat 2 */
	SueRmat2 int16
	/*SueRmat3 Serial UDB Extra Rmat 3 */
	SueRmat3 int16
	/*SueRmat4 Serial UDB Extra Rmat 4 */
	SueRmat4 int16
	/*SueRmat5 Serial UDB Extra Rmat 5 */
	SueRmat5 int16
	/*SueRmat6 Serial UDB Extra Rmat 6 */
	SueRmat6 int16
	/*SueRmat7 Serial UDB Extra Rmat 7 */
	SueRmat7 int16
	/*SueRmat8 Serial UDB Extra Rmat 8 */
	SueRmat8 int16
	/*SueCog Serial UDB Extra GPS Course Over Ground */
	SueCog uint16
	/*SueSog Serial UDB Extra Speed Over Ground */
	SueSog int16
	/*SueCpuLoad Serial UDB Extra CPU Load */
	SueCpuLoad uint16
	/*SueAirSpeed3DIMU Serial UDB Extra 3D IMU Air Speed */
	SueAirSpeed3DIMU uint16
	/*SueEstimatedWind0 Serial UDB Extra Estimated Wind 0 */
	SueEstimatedWind0 int16
	/*SueEstimatedWind1 Serial UDB Extra Estimated Wind 1 */
	SueEstimatedWind1 int16
	/*SueEstimatedWind2 Serial UDB Extra Estimated Wind 2 */
	SueEstimatedWind2 int16
	/*SueMagfieldearth0 Serial UDB Extra Magnetic Field Earth 0  */
	SueMagfieldearth0 int16
	/*SueMagfieldearth1 Serial UDB Extra Magnetic Field Earth 1  */
	SueMagfieldearth1 int16
	/*SueMagfieldearth2 Serial UDB Extra Magnetic Field Earth 2  */
	SueMagfieldearth2 int16
	/*SueSvs Serial UDB Extra Number of Sattelites in View */
	SueSvs int16
	/*SueHdop Serial UDB Extra GPS Horizontal Dilution of Precision */
	SueHdop int16
	/*SueStatus Serial UDB Extra Status */
	SueStatus uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF2A) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("SueTime:\t%v \n")
	builder.WriteString("SueLatitude:\t%v \n")
	builder.WriteString("SueLongitude:\t%v \n")
	builder.WriteString("SueAltitude:\t%v \n")
	builder.WriteString("SueWaypointIndex:\t%v \n")
	builder.WriteString("SueRmat0:\t%v \n")
	builder.WriteString("SueRmat1:\t%v \n")
	builder.WriteString("SueRmat2:\t%v \n")
	builder.WriteString("SueRmat3:\t%v \n")
	builder.WriteString("SueRmat4:\t%v \n")
	builder.WriteString("SueRmat5:\t%v \n")
	builder.WriteString("SueRmat6:\t%v \n")
	builder.WriteString("SueRmat7:\t%v \n")
	builder.WriteString("SueRmat8:\t%v \n")
	builder.WriteString("SueCog:\t%v \n")
	builder.WriteString("SueSog:\t%v \n")
	builder.WriteString("SueCpuLoad:\t%v \n")
	builder.WriteString("SueAirSpeed3DIMU:\t%v \n")
	builder.WriteString("SueEstimatedWind0:\t%v \n")
	builder.WriteString("SueEstimatedWind1:\t%v \n")
	builder.WriteString("SueEstimatedWind2:\t%v \n")
	builder.WriteString("SueMagfieldearth0:\t%v \n")
	builder.WriteString("SueMagfieldearth1:\t%v \n")
	builder.WriteString("SueMagfieldearth2:\t%v \n")
	builder.WriteString("SueSvs:\t%v \n")
	builder.WriteString("SueHdop:\t%v \n")
	builder.WriteString("SueStatus:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueTime,
		m.SueLatitude,
		m.SueLongitude,
		m.SueAltitude,
		m.SueWaypointIndex,
		m.SueRmat0,
		m.SueRmat1,
		m.SueRmat2,
		m.SueRmat3,
		m.SueRmat4,
		m.SueRmat5,
		m.SueRmat6,
		m.SueRmat7,
		m.SueRmat8,
		m.SueCog,
		m.SueSog,
		m.SueCpuLoad,
		m.SueAirSpeed3DIMU,
		m.SueEstimatedWind0,
		m.SueEstimatedWind1,
		m.SueEstimatedWind2,
		m.SueMagfieldearth0,
		m.SueMagfieldearth1,
		m.SueMagfieldearth2,
		m.SueSvs,
		m.SueHdop,
		m.SueStatus,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF2A) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF2A) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF2A) GetMessageName() string {
	return "SerialUdbExtraF2A"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF2A) GetID() uint32 {
	return 170
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF2A) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF2A) getV1Length() int {
	return 61
}

func (m *SerialUdbExtraF2A) getIOSlice() []byte {
	return make([]byte, 61+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF2A) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SerialUdbExtraF2A fields
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
func (m *SerialUdbExtraF2A) Write(version int) (output []byte, err error) {
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
