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

/*TrajectoryRepresentationWaypoints Describe a trajectory using an array of up-to 5 waypoints in the local frame. */
type TrajectoryRepresentationWaypoints struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*PosX X-coordinate of waypoint, set to NaN if not being used */
	PosX []float32
	/*PosY Y-coordinate of waypoint, set to NaN if not being used */
	PosY []float32
	/*PosZ Z-coordinate of waypoint, set to NaN if not being used */
	PosZ []float32
	/*VelX X-velocity of waypoint, set to NaN if not being used */
	VelX []float32
	/*VelY Y-velocity of waypoint, set to NaN if not being used */
	VelY []float32
	/*VelZ Z-velocity of waypoint, set to NaN if not being used */
	VelZ []float32
	/*AccX X-acceleration of waypoint, set to NaN if not being used */
	AccX []float32
	/*AccY Y-acceleration of waypoint, set to NaN if not being used */
	AccY []float32
	/*AccZ Z-acceleration of waypoint, set to NaN if not being used */
	AccZ []float32
	/*PosYaw Yaw angle, set to NaN if not being used */
	PosYaw []float32
	/*VelYaw Yaw rate, set to NaN if not being used */
	VelYaw []float32
	/*ValIDPoints Number of valid points (up-to 5 waypoints are possible) */
	ValIDPoints uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *TrajectoryRepresentationWaypoints) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *TrajectoryRepresentationWaypoints) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *TrajectoryRepresentationWaypoints) GetName() string {
	return "TrajectoryRepresentationWaypoints"
}

// GetID gets the ID of the Message
func (m *TrajectoryRepresentationWaypoints) GetID() uint32 {
	return 332
}

// Read sets the field values of the message from the raw message payload
func (m *TrajectoryRepresentationWaypoints) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
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

	err = binary.Read(reader, binary.LittleEndian, &m.VelX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelZ)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AccX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AccY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AccZ)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosYaw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelYaw)
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
func (m *TrajectoryRepresentationWaypoints) Write(version int) ([]byte, error) {
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

	err = binary.Write(&buffer, binary.LittleEndian, m.VelX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelZ)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AccX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AccY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AccZ)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosYaw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelYaw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ValIDPoints)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
