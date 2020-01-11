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

/*OpenDroneIDLocation Data for filling the OpenDroneID Location message. The float data types are 32-bit IEEE 754. The Location message provides the location, altitude, direction and speed of the aircraft. */
type OpenDroneIDLocation struct {
	/*Latitude Current latitude of the UA (Unmanned Aircraft). If unknown: 0 deg (both Lat/Lon). */
	Latitude int32
	/*Longitude Current longitude of the UA (Unmanned Aircraft). If unknown: 0 deg (both Lat/Lon). */
	Longitude int32
	/*AltitudeBarometric The altitude calculated from the barometric pressue. Reference is against 29.92inHg or 1013.2mb. If unknown: -1000 m. */
	AltitudeBarometric float32
	/*AltitudeGeodetic The geodetic altitude as defined by WGS84. If unknown: -1000 m. */
	AltitudeGeodetic float32
	/*Height The current height of the UA (Unmanned Aircraft) above the take-off location or the ground as indicated by height_reference. If unknown: -1000 m. */
	Height float32
	/*Timestamp Seconds after the full hour. Typically the GPS outputs a time of week value in milliseconds. That value can be easily converted for this field using ((float) (time_week_ms % (60*60*1000))) / 1000. */
	Timestamp float32
	/*Direction Direction over ground (not heading, but direction of movement) in degrees * 100: 0.0 - 359.99 degrees. If unknown: 361.00 degrees. */
	Direction uint16
	/*SpeedHorizontal Ground speed. Positive only. If unknown: 255.00 m/s. If speed is larger than 254.25 m/s, use 254.25 m/s. */
	SpeedHorizontal uint16
	/*SpeedVertical The vertical speed. Up is positive. If unknown: 63.00 m/s. If speed is larger than 62.00 m/s, use 62.00 m/s. */
	SpeedVertical int16
	/*Status Indicates whether the Unmanned Aircraft is on the ground or in the air. */
	Status uint8
	/*HeightReference Indicates the reference point for the height field. */
	HeightReference uint8
	/*HorizontalAccuracy The accuracy of the horizontal position. */
	HorizontalAccuracy uint8
	/*VerticalAccuracy The accuracy of the vertical position. */
	VerticalAccuracy uint8
	/*BarometerAccuracy The accuracy of the barometric altitude. */
	BarometerAccuracy uint8
	/*SpeedAccuracy The accuracy of the horizontal and vertical speed. */
	SpeedAccuracy uint8
	/*TimestampAccuracy The accuracy of the timestamps. */
	TimestampAccuracy uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *OpenDroneIDLocation) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Latitude:\t%v [degE7]\n"
	format += "Longitude:\t%v [degE7]\n"
	format += "AltitudeBarometric:\t%v [m]\n"
	format += "AltitudeGeodetic:\t%v [m]\n"
	format += "Height:\t%v [m]\n"
	format += "Timestamp:\t%v [s]\n"
	format += "Direction:\t%v [cdeg]\n"
	format += "SpeedHorizontal:\t%v [cm/s]\n"
	format += "SpeedVertical:\t%v [cm/s]\n"
	format += "Status:\t%v \n"
	format += "HeightReference:\t%v \n"
	format += "HorizontalAccuracy:\t%v \n"
	format += "VerticalAccuracy:\t%v \n"
	format += "BarometerAccuracy:\t%v \n"
	format += "SpeedAccuracy:\t%v \n"
	format += "TimestampAccuracy:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Latitude,
		m.Longitude,
		m.AltitudeBarometric,
		m.AltitudeGeodetic,
		m.Height,
		m.Timestamp,
		m.Direction,
		m.SpeedHorizontal,
		m.SpeedVertical,
		m.Status,
		m.HeightReference,
		m.HorizontalAccuracy,
		m.VerticalAccuracy,
		m.BarometerAccuracy,
		m.SpeedAccuracy,
		m.TimestampAccuracy,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *OpenDroneIDLocation) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *OpenDroneIDLocation) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *OpenDroneIDLocation) GetMessageName() string {
	return "OpenDroneIDLocation"
}

// GetID gets the ID of the Message
func (m *OpenDroneIDLocation) GetID() uint32 {
	return 12901
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *OpenDroneIDLocation) HasExtensionFields() bool {
	return false
}

func (m *OpenDroneIDLocation) getV1Length() int {
	return 37
}

func (m *OpenDroneIDLocation) getIOSlice() []byte {
	return make([]byte, 37+1)
}

// Read sets the field values of the message from the raw message payload
func (m *OpenDroneIDLocation) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the OpenDroneIDLocation fields
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
func (m *OpenDroneIDLocation) Write(version int) (output []byte, err error) {
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
