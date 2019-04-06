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

/*TrajectoryRepresentationBezier Describe a trajectory using an array of up-to 5 bezier points in the local frame. */
type TrajectoryRepresentationBezier struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*PosX X-coordinate of starting bezier point, set to NaN if not being used */
	PosX []float32
	/*PosY Y-coordinate of starting bezier point, set to NaN if not being used */
	PosY []float32
	/*PosZ Z-coordinate of starting bezier point, set to NaN if not being used */
	PosZ []float32
	/*Delta Bezier time horizon, set to NaN if velocity/acceleration should not be incorporated */
	Delta []float32
	/*PosYaw Yaw, set to NaN for unchanged */
	PosYaw []float32
	/*ValIDPoints Number of valid points (up-to 5 waypoints are possible) */
	ValIDPoints uint8
}

// Read sets the field values of the message from the raw message payload
func (m *TrajectoryRepresentationBezier) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosZ)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Delta)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosYaw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ValIDPoints)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *TrajectoryRepresentationBezier) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosZ)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Delta)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosYaw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ValIDPoints)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
