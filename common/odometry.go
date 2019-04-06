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

/*Odometry Odometry message to communicate odometry information with an external interface. Fits ROS REP 147 standard for aerial vehicles (http://www.ros.org/reps/rep-0147.html). */
type Odometry struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*X X Position */
	X float32
	/*Y Y Position */
	Y float32
	/*Z Z Position */
	Z float32
	/*Q Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation) */
	Q []float32
	/*Vx X linear speed */
	Vx float32
	/*Vy Y linear speed */
	Vy float32
	/*Vz Z linear speed */
	Vz float32
	/*Rollspeed Roll angular speed */
	Rollspeed float32
	/*Pitchspeed Pitch angular speed */
	Pitchspeed float32
	/*Yawspeed Yaw angular speed */
	Yawspeed float32
	/*PoseCovariance Row-major representation of a 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array. */
	PoseCovariance []float32
	/*VelocityCovariance Row-major representation of a 6x6 velocity cross-covariance matrix upper right triangle (states: vx, vy, vz, rollspeed, pitchspeed, yawspeed; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array. */
	VelocityCovariance []float32
	/*FrameID Coordinate frame of reference for the pose data. */
	FrameID uint8
	/*ChildFrameID Coordinate frame of reference for the velocity in free space (twist) data. */
	ChildFrameID uint8
}

// Read sets the field values of the message from the raw message payload
func (m *Odometry) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Q)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Rollspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Pitchspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yawspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PoseCovariance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelocityCovariance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FrameID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ChildFrameID)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *Odometry) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Q)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Rollspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Pitchspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yawspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PoseCovariance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelocityCovariance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FrameID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ChildFrameID)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
