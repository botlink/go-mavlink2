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

/*HighresImu The IMU readings in SI units in NED body frame */
type HighresImu struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Xacc X acceleration */
	Xacc float32
	/*Yacc Y acceleration */
	Yacc float32
	/*Zacc Z acceleration */
	Zacc float32
	/*Xgyro Angular speed around X axis */
	Xgyro float32
	/*Ygyro Angular speed around Y axis */
	Ygyro float32
	/*Zgyro Angular speed around Z axis */
	Zgyro float32
	/*Xmag X Magnetic field */
	Xmag float32
	/*Ymag Y Magnetic field */
	Ymag float32
	/*Zmag Z Magnetic field */
	Zmag float32
	/*AbsPressure Absolute pressure */
	AbsPressure float32
	/*DiffPressure Differential pressure */
	DiffPressure float32
	/*PressureAlt Altitude calculated from pressure */
	PressureAlt float32
	/*Temperature Temperature */
	Temperature float32
	/*FieldsUpdated Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature */
	FieldsUpdated uint16
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HighresImu) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HighresImu) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *HighresImu) GetName() string {
	return "HighresImu"
}

// GetID gets the ID of the Message
func (m *HighresImu) GetID() uint32 {
	return 105
}

// Read sets the field values of the message from the raw message payload
func (m *HighresImu) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
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

	err = binary.Read(reader, binary.LittleEndian, &m.AbsPressure)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.DiffPressure)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PressureAlt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FieldsUpdated)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HighresImu) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.AbsPressure)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.DiffPressure)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PressureAlt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FieldsUpdated)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
