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

/*Vibration Vibration levels and accelerometer clipping */
type Vibration struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*VibrationX Vibration levels on X-axis */
	VibrationX float32
	/*VibrationY Vibration levels on Y-axis */
	VibrationY float32
	/*VibrationZ Vibration levels on Z-axis */
	VibrationZ float32
	/*Clipping0 first accelerometer clipping count */
	Clipping0 uint32
	/*Clipping1 second accelerometer clipping count */
	Clipping1 uint32
	/*Clipping2 third accelerometer clipping count */
	Clipping2 uint32
}

// Read sets the field values of the message from the raw message payload
func (m *Vibration) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VibrationX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VibrationY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VibrationZ)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Clipping0)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Clipping1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Clipping2)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *Vibration) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VibrationX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VibrationY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VibrationZ)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Clipping0)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Clipping1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Clipping2)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
