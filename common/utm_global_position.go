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

/*UtmGlobalPosition The global position resulting from GPS and sensor fusion. */
type UtmGlobalPosition struct {
	/*Time Time of applicability of position (microseconds since UNIX epoch). */
	Time uint64
	/*Lat Latitude (WGS84) */
	Lat int32
	/*Lon Longitude (WGS84) */
	Lon int32
	/*Alt Altitude (WGS84) */
	Alt int32
	/*RelativeAlt Altitude above ground */
	RelativeAlt int32
	/*NextLat Next waypoint, latitude (WGS84) */
	NextLat int32
	/*NextLon Next waypoint, longitude (WGS84) */
	NextLon int32
	/*NextAlt Next waypoint, altitude (WGS84) */
	NextAlt int32
	/*Vx Ground X speed (latitude, positive north) */
	Vx int16
	/*Vy Ground Y speed (longitude, positive east) */
	Vy int16
	/*Vz Ground Z speed (altitude, positive down) */
	Vz int16
	/*HAcc Horizontal position uncertainty (standard deviation) */
	HAcc uint16
	/*VAcc Altitude uncertainty (standard deviation) */
	VAcc uint16
	/*VelAcc Speed uncertainty (standard deviation) */
	VelAcc uint16
	/*UpdateRate Time until next update. Set to 0 if unknown or in data driven mode. */
	UpdateRate uint16
	/*UasID Unique UAS ID. */
	UasID [18]uint8
	/*FlightState Flight state */
	FlightState uint8
	/*Flags Bitwise OR combination of the data available flags. */
	Flags uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *UtmGlobalPosition) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Time:\t%v [us]\n")
	builder.WriteString("Lat:\t%v [degE7]\n")
	builder.WriteString("Lon:\t%v [degE7]\n")
	builder.WriteString("Alt:\t%v [mm]\n")
	builder.WriteString("RelativeAlt:\t%v [mm]\n")
	builder.WriteString("NextLat:\t%v [degE7]\n")
	builder.WriteString("NextLon:\t%v [degE7]\n")
	builder.WriteString("NextAlt:\t%v [mm]\n")
	builder.WriteString("Vx:\t%v [cm/s]\n")
	builder.WriteString("Vy:\t%v [cm/s]\n")
	builder.WriteString("Vz:\t%v [cm/s]\n")
	builder.WriteString("HAcc:\t%v [mm]\n")
	builder.WriteString("VAcc:\t%v [mm]\n")
	builder.WriteString("VelAcc:\t%v [cm/s]\n")
	builder.WriteString("UpdateRate:\t%v [cs]\n")
	builder.WriteString("UasID:\t%v \n")
	builder.WriteString("FlightState:\t%v \n")
	builder.WriteString("Flags:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Time,
		m.Lat,
		m.Lon,
		m.Alt,
		m.RelativeAlt,
		m.NextLat,
		m.NextLon,
		m.NextAlt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.HAcc,
		m.VAcc,
		m.VelAcc,
		m.UpdateRate,
		m.UasID,
		m.FlightState,
		m.Flags,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *UtmGlobalPosition) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *UtmGlobalPosition) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *UtmGlobalPosition) GetMessageName() string {
	return "UtmGlobalPosition"
}

// GetID gets the ID of the Message
func (m *UtmGlobalPosition) GetID() uint32 {
	return 340
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *UtmGlobalPosition) HasExtensionFields() bool {
	return false
}

func (m *UtmGlobalPosition) getV1Length() int {
	return 70
}

func (m *UtmGlobalPosition) getIOSlice() []byte {
	return make([]byte, 70+1)
}

// Read sets the field values of the message from the raw message payload
func (m *UtmGlobalPosition) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the UtmGlobalPosition fields
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
func (m *UtmGlobalPosition) Write(version int) (output []byte, err error) {
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
