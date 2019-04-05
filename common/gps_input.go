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
package common

import (
	"bytes"
	"encoding/binary"
)

/*GpsInput GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system. */
type GpsInput struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
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
}

// Read sets the field values of the message from the raw message payload
func (m *GpsInput) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeWeekMs)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Hdop)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vdop)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vn)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Ve)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vd)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SpeedAccuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HorizAccuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VertAccuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IgnoreFlags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeWeek)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.GpsID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FixType)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatellitesVisible)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsInput) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeWeekMs)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Hdop)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vdop)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vn)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Ve)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vd)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SpeedAccuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HorizAccuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VertAccuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IgnoreFlags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeWeek)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.GpsID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FixType)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
