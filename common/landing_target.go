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

/*LandingTarget The location of a landing target. See: https://mavlink.io/en/services/landing_target.html */
type LandingTarget struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*AngleX X-axis angular offset of the target from the center of the image */
	AngleX float32
	/*AngleY Y-axis angular offset of the target from the center of the image */
	AngleY float32
	/*Distance Distance to the target from the vehicle */
	Distance float32
	/*SizeX Size of target along x-axis */
	SizeX float32
	/*SizeY Size of target along y-axis */
	SizeY float32
	/*TargetNum The ID of the target if multiple targets are present */
	TargetNum uint8
	/*Frame Coordinate frame used for following fields. */
	Frame uint8
	/*X X Position of the landing target in MAV_FRAME */
	X float32
	/*Y Y Position of the landing target in MAV_FRAME */
	Y float32
	/*Z Z Position of the landing target in MAV_FRAME */
	Z float32
	/*Q Quaternion of landing target orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0) */
	Q []float32
	/*Type Type of landing target */
	Type uint8
	/*PositionValID Boolean indicating whether the position fields (x, y, z, q, type) contain valid target position information (valid: 1, invalid: 0). Default is 0 (invalid). */
	PositionValID uint8
}

// Read sets the field values of the message from the raw message payload
func (m *LandingTarget) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AngleX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AngleY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Distance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SizeX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SizeY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetNum)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Frame)
	if err != nil {
		return
	}

	if version == 2 {
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

		err = binary.Read(reader, binary.LittleEndian, &m.Q)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Type)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.PositionValID)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *LandingTarget) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AngleX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AngleY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Distance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SizeX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SizeY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetNum)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
