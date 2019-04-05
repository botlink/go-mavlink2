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
package common

import (
	"bytes"
	"encoding/binary"
)

/*ScaledPressure3 Barometer readings for 3rd barometer */
type ScaledPressure3 struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*PressAbs Absolute pressure */
	PressAbs float32
	/*PressDiff Differential pressure */
	PressDiff float32
	/*Temperature Temperature measurement */
	Temperature int16
}

// Read sets the field values of the message from the raw message payload
func (m *ScaledPressure3) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressAbs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressDiff)
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
func (m *ScaledPressure3) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressAbs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressDiff)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
