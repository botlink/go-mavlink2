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

/*VisionSpeedEstimate Speed estimate from a vision source. */
type VisionSpeedEstimate struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Usec Timestamp (UNIX time or time since system boot) */
	Usec uint64
	/*X Global X speed */
	X float32
	/*Y Global Y speed */
	Y float32
	/*Z Global Z speed */
	Z float32
	/*Covariance Row-major representation of 3x3 linear velocity covariance matrix (states: vx, vy, vz; 1st three entries - 1st row, etc.). If unknown, assign NaN value to first element in the array. */
	Covariance []float32
}

// GetVersion gets the MAVLink version of the Message contents
func (m *VisionSpeedEstimate) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *VisionSpeedEstimate) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *VisionSpeedEstimate) GetName() string {
	return "VisionSpeedEstimate"
}

// GetID gets the ID of the Message
func (m *VisionSpeedEstimate) GetID() uint32 {
	return 103
}

// Read sets the field values of the message from the raw message payload
func (m *VisionSpeedEstimate) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Usec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.X)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Y)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Z)
	if err != nil {
		return
	}

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.Covariance)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *VisionSpeedEstimate) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Usec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.X)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Y)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Z)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.Z)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
