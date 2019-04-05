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

/*ScaledImu3 The RAW IMU readings for 3rd 9DOF sensor setup. This message should contain the scaled values to the described units */
type ScaledImu3 struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Xacc X acceleration */
	Xacc int16
	/*Yacc Y acceleration */
	Yacc int16
	/*Zacc Z acceleration */
	Zacc int16
	/*Xgyro Angular speed around X axis */
	Xgyro int16
	/*Ygyro Angular speed around Y axis */
	Ygyro int16
	/*Zgyro Angular speed around Z axis */
	Zgyro int16
	/*Xmag X Magnetic field */
	Xmag int16
	/*Ymag Y Magnetic field */
	Ymag int16
	/*Zmag Z Magnetic field */
	Zmag int16
}

// Read sets the field values of the message from the raw message payload
func (m *ScaledImu3) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Xacc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yacc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Zacc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Xgyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Ygyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Zgyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Xmag)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Ymag)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Zmag)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ScaledImu3) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Xacc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yacc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Zacc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Xgyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Ygyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Zgyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Xmag)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Ymag)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Zmag)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
