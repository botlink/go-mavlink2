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

/*RawPressure The RAW pressure readings for the typical setup of one absolute pressure and one differential pressure sensor. The sensor values should be the raw, UNSCALED ADC values. */
type RawPressure struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*PressAbs Absolute pressure (raw) */
	PressAbs int16
	/*PressDiff1 Differential pressure 1 (raw, 0 if nonexistent) */
	PressDiff1 int16
	/*PressDiff2 Differential pressure 2 (raw, 0 if nonexistent) */
	PressDiff2 int16
	/*Temperature Raw Temperature measurement (raw) */
	Temperature int16
}

// GetVersion gets the MAVLink version of the Message contents
func (m *RawPressure) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *RawPressure) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *RawPressure) GetName() string {
	return "RawPressure"
}

// GetID gets the ID of the Message
func (m *RawPressure) GetID() uint32 {
	return 28
}

// Read sets the field values of the message from the raw message payload
func (m *RawPressure) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressAbs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressDiff1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressDiff2)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *RawPressure) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressAbs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressDiff1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressDiff2)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
