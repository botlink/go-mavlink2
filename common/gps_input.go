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

/*GpsInput GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system. */
type GpsInput struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*TimeWeekMs GPS time (from start of GPS week) */
	TimeWeekMs uint32
	/*Lat Latitude (WGS84) */
	Lat int32
	/*Lon Longitude (WGS84) */
	Lon int32
	/*Alt Altitude (MSL). Positive for up. */
	Alt float32
	/*Hdop GPS HDOP horizontal dilution of position */
	Hdop float32
	/*Vdop GPS VDOP vertical dilution of position */
	Vdop float32
	/*Vn GPS velocity in NORTH direction in earth-fixed NED frame */
	Vn float32
	/*Ve GPS velocity in EAST direction in earth-fixed NED frame */
	Ve float32
	/*Vd GPS velocity in DOWN direction in earth-fixed NED frame */
	Vd float32
	/*SpeedAccuracy GPS speed accuracy */
	SpeedAccuracy float32
	/*HorizAccuracy GPS horizontal accuracy */
	HorizAccuracy float32
	/*VertAccuracy GPS vertical accuracy */
	VertAccuracy float32
	/*IgnoreFlags Bitmap indicating which GPS input flags fields to ignore.  All other fields must be provided. */
	IgnoreFlags uint16
	/*TimeWeek GPS week number */
	TimeWeek uint16
	/*GpsID ID of the GPS for multiple GPS inputs */
	GpsID uint8
	/*FixType 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK */
	FixType uint8
	/*SatellitesVisible Number of satellites visible. */
	SatellitesVisible uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GpsInput) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("TimeWeekMs:\t%v [ms]\n")
	builder.WriteString("Lat:\t%v [degE7]\n")
	builder.WriteString("Lon:\t%v [degE7]\n")
	builder.WriteString("Alt:\t%v [m]\n")
	builder.WriteString("Hdop:\t%v [m]\n")
	builder.WriteString("Vdop:\t%v [m]\n")
	builder.WriteString("Vn:\t%v [m/s]\n")
	builder.WriteString("Ve:\t%v [m/s]\n")
	builder.WriteString("Vd:\t%v [m/s]\n")
	builder.WriteString("SpeedAccuracy:\t%v [m/s]\n")
	builder.WriteString("HorizAccuracy:\t%v [m]\n")
	builder.WriteString("VertAccuracy:\t%v [m]\n")
	builder.WriteString("IgnoreFlags:\t%v \n")
	builder.WriteString("TimeWeek:\t%v \n")
	builder.WriteString("GpsID:\t%v \n")
	builder.WriteString("FixType:\t%v \n")
	builder.WriteString("SatellitesVisible:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.TimeWeekMs,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Hdop,
		m.Vdop,
		m.Vn,
		m.Ve,
		m.Vd,
		m.SpeedAccuracy,
		m.HorizAccuracy,
		m.VertAccuracy,
		m.IgnoreFlags,
		m.TimeWeek,
		m.GpsID,
		m.FixType,
		m.SatellitesVisible,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GpsInput) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GpsInput) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GpsInput) GetMessageName() string {
	return "GpsInput"
}

// GetID gets the ID of the Message
func (m *GpsInput) GetID() uint32 {
	return 232
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GpsInput) HasExtensionFields() bool {
	return false
}

func (m *GpsInput) getV1Length() int {
	return 63
}

func (m *GpsInput) getIOSlice() []byte {
	return make([]byte, 63+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GpsInput) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GpsInput fields
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
func (m *GpsInput) Write(version int) (output []byte, err error) {
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
