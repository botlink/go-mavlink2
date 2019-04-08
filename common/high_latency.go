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

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*HighLatency Message appropriate for high latency connections like Iridium */
type HighLatency struct {
	/*CustomMode A bitfield for use for autopilot-specific flags. */
	CustomMode uint32
	/*Latitude Latitude */
	Latitude int32
	/*Longitude Longitude */
	Longitude int32
	/*Roll roll */
	Roll int16
	/*Pitch pitch */
	Pitch int16
	/*Heading heading */
	Heading uint16
	/*HeadingSp heading setpoint */
	HeadingSp int16
	/*AltitudeAmsl Altitude above mean sea level */
	AltitudeAmsl int16
	/*AltitudeSp Altitude setpoint relative to the home position */
	AltitudeSp int16
	/*WpDistance distance to target */
	WpDistance uint16
	/*BaseMode Bitmap of enabled system modes. */
	BaseMode uint8
	/*LandedState The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown. */
	LandedState uint8
	/*Throttle throttle (percentage) */
	Throttle int8
	/*Airspeed airspeed */
	Airspeed uint8
	/*AirspeedSp airspeed setpoint */
	AirspeedSp uint8
	/*Groundspeed groundspeed */
	Groundspeed uint8
	/*ClimbRate climb rate */
	ClimbRate int8
	/*GpsNsat Number of satellites visible. If unknown, set to 255 */
	GpsNsat uint8
	/*GpsFixType GPS Fix type. */
	GpsFixType uint8
	/*BatteryRemaining Remaining battery (percentage) */
	BatteryRemaining uint8
	/*Temperature Autopilot temperature (degrees C) */
	Temperature int8
	/*TemperatureAir Air temperature (degrees C) from airspeed sensor */
	TemperatureAir int8
	/*Failsafe failsafe (each bit represents a failsafe where 0=ok, 1=failsafe active (bit0:RC, bit1:batt, bit2:GPS, bit3:GCS, bit4:fence) */
	Failsafe uint8
	/*WpNum current waypoint number */
	WpNum uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *HighLatency) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("CustomMode:\t%v \n")
	builder.WriteString("Latitude:\t%v [degE7]\n")
	builder.WriteString("Longitude:\t%v [degE7]\n")
	builder.WriteString("Roll:\t%v [cdeg]\n")
	builder.WriteString("Pitch:\t%v [cdeg]\n")
	builder.WriteString("Heading:\t%v [cdeg]\n")
	builder.WriteString("HeadingSp:\t%v [cdeg]\n")
	builder.WriteString("AltitudeAmsl:\t%v [m]\n")
	builder.WriteString("AltitudeSp:\t%v [m]\n")
	builder.WriteString("WpDistance:\t%v [m]\n")
	builder.WriteString("BaseMode:\t%v \n")
	builder.WriteString("LandedState:\t%v \n")
	builder.WriteString("Throttle:\t%v [%]\n")
	builder.WriteString("Airspeed:\t%v [m/s]\n")
	builder.WriteString("AirspeedSp:\t%v [m/s]\n")
	builder.WriteString("Groundspeed:\t%v [m/s]\n")
	builder.WriteString("ClimbRate:\t%v [m/s]\n")
	builder.WriteString("GpsNsat:\t%v \n")
	builder.WriteString("GpsFixType:\t%v \n")
	builder.WriteString("BatteryRemaining:\t%v [%]\n")
	builder.WriteString("Temperature:\t%v [degC]\n")
	builder.WriteString("TemperatureAir:\t%v [degC]\n")
	builder.WriteString("Failsafe:\t%v \n")
	builder.WriteString("WpNum:\t%v \n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.CustomMode,
		m.Latitude,
		m.Longitude,
		m.Roll,
		m.Pitch,
		m.Heading,
		m.HeadingSp,
		m.AltitudeAmsl,
		m.AltitudeSp,
		m.WpDistance,
		m.BaseMode,
		m.LandedState,
		m.Throttle,
		m.Airspeed,
		m.AirspeedSp,
		m.Groundspeed,
		m.ClimbRate,
		m.GpsNsat,
		m.GpsFixType,
		m.BatteryRemaining,
		m.Temperature,
		m.TemperatureAir,
		m.Failsafe,
		m.WpNum,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HighLatency) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HighLatency) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *HighLatency) GetMessageName() string {
	return "HighLatency"
}

// GetID gets the ID of the Message
func (m *HighLatency) GetID() uint32 {
	return 234
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *HighLatency) HasExtensionFields() bool {
	return false
}

func (m *HighLatency) getV1Length() int {
	return 40
}

func (m *HighLatency) getIOSlice() []byte {
	return make([]byte, 40+1)
}

// Read sets the field values of the message from the raw message payload
func (m *HighLatency) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the HighLatency fields
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
func (m *HighLatency) Write(version int) (output []byte, err error) {
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
