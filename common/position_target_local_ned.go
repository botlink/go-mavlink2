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

/*PositionTargetLocalNed Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_LOCAL_NED if the vehicle is being controlled this way. */
type PositionTargetLocalNed struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*X X Position in NED frame */
	X float32
	/*Y Y Position in NED frame */
	Y float32
	/*Z Z Position in NED frame (note, altitude is negative in NED) */
	Z float32
	/*Vx X velocity in NED frame */
	Vx float32
	/*Vy Y velocity in NED frame */
	Vy float32
	/*Vz Z velocity in NED frame */
	Vz float32
	/*Afx X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afx float32
	/*Afy Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afy float32
	/*Afz Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N */
	Afz float32
	/*Yaw yaw setpoint */
	Yaw float32
	/*YawRate yaw rate setpoint */
	YawRate float32
	/*TypeMask Bitmap to indicate which dimensions should be ignored by the vehicle. */
	TypeMask uint16
	/*CoordinateFrame Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9 */
	CoordinateFrame uint8
}

// Read sets the field values of the message from the raw message payload
func (m *PositionTargetLocalNed) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Afx)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Afy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Afz)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yaw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.YawRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TypeMask)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CoordinateFrame)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *PositionTargetLocalNed) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Afx)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Afy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Afz)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yaw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.YawRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TypeMask)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CoordinateFrame)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
