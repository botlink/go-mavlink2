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

/*LoggingData A message containing logged data (see also MAV_CMD_LOGGING_START) */
type LoggingData struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Sequence sequence number (can wrap) */
	Sequence uint16
	/*TargetSystem system ID of the target */
	TargetSystem uint8
	/*TargetComponent component ID of the target */
	TargetComponent uint8
	/*Length data length */
	Length uint8
	/*FirstMessageOffset offset into data where first message starts. This can be used for recovery, when a previous message got lost (set to 255 if no start exists). */
	FirstMessageOffset uint8
	/*Data logged data */
	Data []uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *LoggingData) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *LoggingData) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *LoggingData) GetName() string {
	return "LoggingData"
}

// GetID gets the ID of the Message
func (m *LoggingData) GetID() uint32 {
	return 266
}

// Read sets the field values of the message from the raw message payload
func (m *LoggingData) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Sequence)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetSystem)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetComponent)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Length)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FirstMessageOffset)
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
func (m *LoggingData) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Sequence)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetSystem)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FirstMessageOffset)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
