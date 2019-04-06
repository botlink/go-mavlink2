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

/*HilActuatorControls Sent from autopilot to simulation. Hardware in the loop control outputs (replacement for HIL_CONTROLS) */
type HilActuatorControls struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Flags Flags as bitfield, reserved for future use. */
	Flags uint64
	/*Controls Control outputs -1 .. 1. Channel assignment depends on the simulated hardware. */
	Controls []float32
	/*Mode System mode. Includes arming state. */
	Mode uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HilActuatorControls) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HilActuatorControls) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *HilActuatorControls) GetName() string {
	return "HilActuatorControls"
}

// GetID gets the ID of the Message
func (m *HilActuatorControls) GetID() uint32 {
	return 93
}

// Read sets the field values of the message from the raw message payload
func (m *HilActuatorControls) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Controls)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Mode)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HilActuatorControls) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Controls)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Mode)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
