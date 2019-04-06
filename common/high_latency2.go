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
)

/*HighLatency2 Message appropriate for high latency connections like Iridium (version 2) */
type HighLatency2 struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
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
	/*Autopilot Autopilot type / class. */
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
	/*Battery Battery (percentage, -1 for DNU) */
	Battery int8
	/*Custom0 Field for custom payload. */
	Custom0 int8
	/*Custom1 Field for custom payload. */
	Custom1 int8
	/*Custom2 Field for custom payload. */
	Custom2 int8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HighLatency2) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HighLatency2) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *HighLatency2) GetName() string {
	return "HighLatency2"
}

// GetID gets the ID of the Message
func (m *HighLatency2) GetID() uint32 {
	return 235
}

// Read sets the field values of the message from the raw message payload
func (m *HighLatency2) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Timestamp)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Latitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Longitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CustomMode)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Altitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetAltitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetDistance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WpNum)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FailureFlags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Type)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Autopilot)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Heading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetHeading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Throttle)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Airspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AirspeedSp)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Groundspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Windspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WindHeading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Eph)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Epv)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TemperatureAir)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ClimbRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Battery)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Custom0)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Custom1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Custom2)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HighLatency2) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Timestamp)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Latitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Longitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CustomMode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Altitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetAltitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetDistance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WpNum)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FailureFlags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Type)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Autopilot)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Heading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetHeading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Throttle)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Airspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AirspeedSp)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Groundspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Windspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WindHeading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Eph)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Epv)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TemperatureAir)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ClimbRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Battery)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Custom0)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Custom1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Custom2)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
