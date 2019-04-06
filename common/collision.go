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

/*Collision Information about a potential collision */
type Collision struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*ID Unique identifier, domain based on src field */
	ID uint32
	/*TimeToMinimumDelta Estimated time until collision occurs */
	TimeToMinimumDelta float32
	/*AltitudeMinimumDelta Closest vertical distance between vehicle and object */
	AltitudeMinimumDelta float32
	/*HorizontalMinimumDelta Closest horizontal distance between vehicle and object */
	HorizontalMinimumDelta float32
	/*Src Collision data source */
	Src uint8
	/*Action Action that is being taken to avoid this collision */
	Action uint8
	/*ThreatLevel How concerned the aircraft is about this collision */
	ThreatLevel uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *Collision) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *Collision) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *Collision) GetName() string {
	return "Collision"
}

// GetID gets the ID of the Message
func (m *Collision) GetID() uint32 {
	return 247
}

// Read sets the field values of the message from the raw message payload
func (m *Collision) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.ID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeToMinimumDelta)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AltitudeMinimumDelta)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HorizontalMinimumDelta)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Src)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Action)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ThreatLevel)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *Collision) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeToMinimumDelta)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AltitudeMinimumDelta)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HorizontalMinimumDelta)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Src)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Action)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ThreatLevel)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
