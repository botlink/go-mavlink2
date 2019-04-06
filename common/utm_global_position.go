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

/*UtmGlobalPosition The global position resulting from GPS and sensor fusion. */
type UtmGlobalPosition struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
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
	UasID []uint8
	/*FlightState Flight state */
	FlightState uint8
	/*Flags Bitwise OR combination of the data available flags. */
	Flags uint8
}

// Read sets the field values of the message from the raw message payload
func (m *UtmGlobalPosition) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Time)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Alt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RelativeAlt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.NextLat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.NextLon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.NextAlt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vx)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vz)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.UpdateRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.UasID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FlightState)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *UtmGlobalPosition) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Time)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Alt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RelativeAlt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.NextLat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.NextLon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.NextAlt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vx)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vz)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.UpdateRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.UasID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FlightState)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
