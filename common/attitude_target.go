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

/*AttitudeTarget Reports the current commanded attitude of the vehicle as specified by the autopilot. This should match the commands sent in a SET_ATTITUDE_TARGET message if the vehicle is being controlled this way. */
type AttitudeTarget struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Q Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0) */
	Q []float32
	/*BodyRollRate Body roll rate */
	BodyRollRate float32
	/*BodyPitchRate Body pitch rate */
	BodyPitchRate float32
	/*BodyYawRate Body yaw rate */
	BodyYawRate float32
	/*Thrust Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust) */
	Thrust float32
	/*TypeMask Mappings: If any of these bits are set, the corresponding input should be ignored: bit 1: body roll rate, bit 2: body pitch rate, bit 3: body yaw rate. bit 4-bit 7: reserved, bit 8: attitude */
	TypeMask uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AttitudeTarget) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AttitudeTarget) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *AttitudeTarget) GetName() string {
	return "AttitudeTarget"
}

// GetID gets the ID of the Message
func (m *AttitudeTarget) GetID() uint32 {
	return 83
}

// Read sets the field values of the message from the raw message payload
func (m *AttitudeTarget) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BodyRollRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BodyPitchRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BodyYawRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Thrust)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TypeMask)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *AttitudeTarget) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BodyRollRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BodyPitchRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BodyYawRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Thrust)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TypeMask)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
