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

/*ManualSetpoint Setpoint in roll, pitch, yaw and thrust from the operator */
type ManualSetpoint struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Roll Desired roll rate */
	Roll float32
	/*Pitch Desired pitch rate */
	Pitch float32
	/*Yaw Desired yaw rate */
	Yaw float32
	/*Thrust Collective thrust, normalized to 0 .. 1 */
	Thrust float32
	/*ModeSwitch Flight mode switch position, 0.. 255 */
	ModeSwitch uint8
	/*ManualOverrIDeSwitch Override mode switch position, 0.. 255 */
	ManualOverrIDeSwitch uint8
}

// Read sets the field values of the message from the raw message payload
func (m *ManualSetpoint) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Roll)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Pitch)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yaw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Thrust)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ModeSwitch)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ManualOverrIDeSwitch)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ManualSetpoint) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Roll)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Pitch)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yaw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Thrust)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ModeSwitch)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ManualOverrIDeSwitch)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
