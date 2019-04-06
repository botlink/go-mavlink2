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

/*GpsInjectData Data for injecting into the onboard GPS (used for DGPS) */
type GpsInjectData struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*Len Data length */
	Len uint8
	/*Data Raw data (110 is enough for 12 satellites of RTCMv2) */
	Data []uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GpsInjectData) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GpsInjectData) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *GpsInjectData) GetName() string {
	return "GpsInjectData"
}

// GetID gets the ID of the Message
func (m *GpsInjectData) GetID() uint32 {
	return 123
}

// Read sets the field values of the message from the raw message payload
func (m *GpsInjectData) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TargetSystem)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetComponent)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Len)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Data)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsInjectData) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TargetSystem)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Len)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
