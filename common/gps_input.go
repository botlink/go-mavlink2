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

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*GPSInput GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system. */
type GPSInput struct {
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
	/*Vn GPS velocity in north direction in earth-fixed NED frame */
	Vn float32
	/*Ve GPS velocity in east direction in earth-fixed NED frame */
	Ve float32
	/*Vd GPS velocity in down direction in earth-fixed NED frame */
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
	/*GPSID ID of the GPS for multiple GPS inputs */
	GPSID uint8
	/*FixType 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK */
	FixType uint8
	/*SatellitesVisible Number of satellites visible. */
	SatellitesVisible uint8
	/*Yaw Yaw of vehicle relative to Earth's North, zero means not available, use 36000 for north */
	Yaw uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GPSInput) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "TimeWeekMs:\t%v [ms]\n"
	format += "Lat:\t%v [degE7]\n"
	format += "Lon:\t%v [degE7]\n"
	format += "Alt:\t%v [m]\n"
	format += "Hdop:\t%v [m]\n"
	format += "Vdop:\t%v [m]\n"
	format += "Vn:\t%v [m/s]\n"
	format += "Ve:\t%v [m/s]\n"
	format += "Vd:\t%v [m/s]\n"
	format += "SpeedAccuracy:\t%v [m/s]\n"
	format += "HorizAccuracy:\t%v [m]\n"
	format += "VertAccuracy:\t%v [m]\n"
	format += "IgnoreFlags:\t%v \n"
	format += "TimeWeek:\t%v \n"
	format += "GPSID:\t%v \n"
	format += "FixType:\t%v \n"
	format += "SatellitesVisible:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "Yaw:\t%v\n"
	}

	if m.HasExtensionFieldValues {
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
			m.GPSID,
			m.FixType,
			m.SatellitesVisible,
			m.Yaw,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

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
		m.GPSID,
		m.FixType,
		m.SatellitesVisible,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GPSInput) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GPSInput) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GPSInput) GetMessageName() string {
	return "GPSInput"
}

// GetID gets the ID of the Message
func (m *GPSInput) GetID() uint32 {
	return 232
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GPSInput) HasExtensionFields() bool {
	return true
}

func (m *GPSInput) getV1Length() int {
	return 63
}

func (m *GPSInput) getIOSlice() []byte {
	return make([]byte, 65+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GPSInput) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GPSInput fields
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
func (m *GPSInput) Write(version int) (output []byte, err error) {
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
