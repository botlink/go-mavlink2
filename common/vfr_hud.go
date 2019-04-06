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

/*VfrHud Metrics typically displayed on a HUD for fixed wing aircraft. */
type VfrHud struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Airspeed Current indicated airspeed (IAS). */
	Airspeed float32
	/*Groundspeed Current ground speed. */
	Groundspeed float32
	/*Alt Current altitude (MSL). */
	Alt float32
	/*Climb Current climb rate. */
	Climb float32
	/*Heading Current heading in compass units (0-360, 0=north). */
	Heading int16
	/*Throttle Current throttle setting (0 to 100). */
	Throttle uint16
}

// GetVersion gets the MAVLink version of the Message contents
func (m *VfrHud) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *VfrHud) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *VfrHud) GetName() string {
	return "VfrHud"
}

// GetID gets the ID of the Message
func (m *VfrHud) GetID() uint32 {
	return 74
}

// Read sets the field values of the message from the raw message payload
func (m *VfrHud) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Airspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Groundspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Alt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Climb)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Heading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Throttle)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *VfrHud) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Airspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Groundspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Alt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Climb)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Heading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Throttle)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
