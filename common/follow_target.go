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

/*FollowTarget Current motion information from a designated system */
type FollowTarget struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Timestamp Timestamp (time since system boot). */
	Timestamp uint64
	/*CustomState button states or switches of a tracker device */
	CustomState uint64
	/*Lat Latitude (WGS84) */
	Lat int32
	/*Lon Longitude (WGS84) */
	Lon int32
	/*Alt Altitude (MSL) */
	Alt float32
	/*Vel target velocity (0,0,0) for unknown */
	Vel []float32
	/*Acc linear target acceleration (0,0,0) for unknown */
	Acc []float32
	/*AttitudeQ (1 0 0 0 for unknown) */
	AttitudeQ []float32
	/*Rates (0 0 0 for unknown) */
	Rates []float32
	/*PositionCov eph epv */
	PositionCov []float32
	/*EstCapabilities bit positions for tracker reporting capabilities (POS = 0, VEL = 1, ACCEL = 2, ATT + RATES = 3) */
	EstCapabilities uint8
}

// Read sets the field values of the message from the raw message payload
func (m *FollowTarget) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Timestamp)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CustomState)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Vel)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Acc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AttitudeQ)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Rates)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PositionCov)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.EstCapabilities)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *FollowTarget) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Timestamp)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CustomState)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Vel)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Acc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AttitudeQ)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Rates)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PositionCov)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.EstCapabilities)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
