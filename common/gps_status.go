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

/*GpsStatus The positioning status, as reported by GPS. This message is intended to display status information about each satellite visible to the receiver. See message GLOBAL_POSITION for the global position estimate. This message can contain information for up to 20 satellites. */
type GpsStatus struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*SatellitesVisible Number of satellites visible */
	SatellitesVisible uint8
	/*SatellitePrn Global satellite ID */
	SatellitePrn []uint8
	/*SatelliteUsed 0: Satellite not used, 1: used for localization */
	SatelliteUsed []uint8
	/*SatelliteElevation Elevation (0: right on top of receiver, 90: on the horizon) of satellite */
	SatelliteElevation []uint8
	/*SatelliteAzimuth Direction of satellite, 0: 0 deg, 255: 360 deg. */
	SatelliteAzimuth []uint8
	/*SatelliteSnr Signal to noise ratio of satellite */
	SatelliteSnr []uint8
}

// Read sets the field values of the message from the raw message payload
func (m *GpsStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.SatellitesVisible)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatellitePrn)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatelliteUsed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatelliteElevation)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatelliteAzimuth)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatelliteSnr)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatellitePrn)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatelliteUsed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatelliteElevation)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatelliteAzimuth)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatelliteSnr)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
