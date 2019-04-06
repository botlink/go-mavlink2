package common

/*
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

/*HighLatency Message appropriate for high latency connections like Iridium */
type HighLatency struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
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
}

// Read sets the field values of the message from the raw message payload
func (m *HighLatency) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.CustomMode)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Roll)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Pitch)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Heading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HeadingSp)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AltitudeAmsl)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AltitudeSp)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WpDistance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaseMode)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.LandedState)
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

	err = binary.Read(reader, binary.LittleEndian, &m.ClimbRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.GpsNsat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.GpsFixType)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BatteryRemaining)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TemperatureAir)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Failsafe)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WpNum)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HighLatency) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.CustomMode)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Roll)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Pitch)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Heading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HeadingSp)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AltitudeAmsl)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AltitudeSp)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WpDistance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaseMode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.LandedState)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.ClimbRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.GpsNsat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.GpsFixType)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BatteryRemaining)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TemperatureAir)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Failsafe)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WpNum)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
