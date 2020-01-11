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

/*HighLatency2 Message appropriate for high latency connections like Iridium (version 2) */
type HighLatency2 struct {
	/*Timestamp Timestamp (milliseconds since boot or Unix epoch) */
	Timestamp uint32
	/*Latitude Latitude */
	Latitude int32
	/*Longitude Longitude */
	Longitude int32
	/*CustomMode A bitfield for use for autopilot-specific flags (2 byte version). */
	CustomMode uint16
	/*Altitude Altitude above mean sea level */
	Altitude int16
	/*TargetAltitude Altitude setpoint */
	TargetAltitude int16
	/*TargetDistance Distance to target waypoint or position */
	TargetDistance uint16
	/*WpNum Current waypoint number */
	WpNum uint16
	/*FailureFlags Bitmap of failure flags. */
	FailureFlags uint16
	/*Type Type of the MAV (quadrotor, helicopter, etc.) */
	Type uint8
	/*Autopilot Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers. */
	Autopilot uint8
	/*Heading Heading */
	Heading uint8
	/*TargetHeading Heading setpoint */
	TargetHeading uint8
	/*Throttle Throttle */
	Throttle uint8
	/*Airspeed Airspeed */
	Airspeed uint8
	/*AirspeedSp Airspeed setpoint */
	AirspeedSp uint8
	/*Groundspeed Groundspeed */
	Groundspeed uint8
	/*Windspeed Windspeed */
	Windspeed uint8
	/*WindHeading Wind heading */
	WindHeading uint8
	/*Eph Maximum error horizontal position since last message */
	Eph uint8
	/*Epv Maximum error vertical position since last message */
	Epv uint8
	/*TemperatureAir Air temperature from airspeed sensor */
	TemperatureAir int8
	/*ClimbRate Maximum climb rate magnitude since last message */
	ClimbRate int8
	/*Battery Battery level (-1 if field not provided). */
	Battery int8
	/*Custom0 Field for custom payload. */
	Custom0 int8
	/*Custom1 Field for custom payload. */
	Custom1 int8
	/*Custom2 Field for custom payload. */
	Custom2 int8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *HighLatency2) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [ms]\n"
	format += "Latitude:\t%v [degE7]\n"
	format += "Longitude:\t%v [degE7]\n"
	format += "CustomMode:\t%v \n"
	format += "Altitude:\t%v [m]\n"
	format += "TargetAltitude:\t%v [m]\n"
	format += "TargetDistance:\t%v [dam]\n"
	format += "WpNum:\t%v \n"
	format += "FailureFlags:\t%v \n"
	format += "Type:\t%v \n"
	format += "Autopilot:\t%v \n"
	format += "Heading:\t%v [deg/2]\n"
	format += "TargetHeading:\t%v [deg/2]\n"
	format += "Throttle:\t%v \n"
	format += "Airspeed:\t%v [m/s*5]\n"
	format += "AirspeedSp:\t%v [m/s*5]\n"
	format += "Groundspeed:\t%v [m/s*5]\n"
	format += "Windspeed:\t%v [m/s*5]\n"
	format += "WindHeading:\t%v [deg/2]\n"
	format += "Eph:\t%v [dm]\n"
	format += "Epv:\t%v [dm]\n"
	format += "TemperatureAir:\t%v [degC]\n"
	format += "ClimbRate:\t%v [dm/s]\n"
	format += "Battery:\t%v \n"
	format += "Custom0:\t%v \n"
	format += "Custom1:\t%v \n"
	format += "Custom2:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.Latitude,
		m.Longitude,
		m.CustomMode,
		m.Altitude,
		m.TargetAltitude,
		m.TargetDistance,
		m.WpNum,
		m.FailureFlags,
		m.Type,
		m.Autopilot,
		m.Heading,
		m.TargetHeading,
		m.Throttle,
		m.Airspeed,
		m.AirspeedSp,
		m.Groundspeed,
		m.Windspeed,
		m.WindHeading,
		m.Eph,
		m.Epv,
		m.TemperatureAir,
		m.ClimbRate,
		m.Battery,
		m.Custom0,
		m.Custom1,
		m.Custom2,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HighLatency2) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HighLatency2) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *HighLatency2) GetMessageName() string {
	return "HighLatency2"
}

// GetID gets the ID of the Message
func (m *HighLatency2) GetID() uint32 {
	return 235
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *HighLatency2) HasExtensionFields() bool {
	return false
}

func (m *HighLatency2) getV1Length() int {
	return 42
}

func (m *HighLatency2) getIOSlice() []byte {
	return make([]byte, 42+1)
}

// Read sets the field values of the message from the raw message payload
func (m *HighLatency2) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the HighLatency2 fields
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
func (m *HighLatency2) Write(version int) (output []byte, err error) {
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
