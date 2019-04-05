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

/*SimState Status of simulation environment, if used */
type SimState struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Q1 True attitude quaternion component 1, w (1 in null-rotation) */
	Q1 float32
	/*Q2 True attitude quaternion component 2, x (0 in null-rotation) */
	Q2 float32
	/*Q3 True attitude quaternion component 3, y (0 in null-rotation) */
	Q3 float32
	/*Q4 True attitude quaternion component 4, z (0 in null-rotation) */
	Q4 float32
	/*Roll Attitude roll expressed as Euler angles, not recommended except for human-readable outputs */
	Roll float32
	/*Pitch Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs */
	Pitch float32
	/*Yaw Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs */
	Yaw float32
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
	/*Lat Latitude */
	Lat float32
	/*Lon Longitude */
	Lon float32
	/*Alt Altitude */
	Alt float32
	/*StdDevHorz Horizontal position standard deviation */
	StdDevHorz float32
	/*StdDevVert Vertical position standard deviation */
	StdDevVert float32
	/*Vn True velocity in NORTH direction in earth-fixed NED frame */
	Vn float32
	/*Ve True velocity in EAST direction in earth-fixed NED frame */
	Ve float32
	/*Vd True velocity in DOWN direction in earth-fixed NED frame */
	Vd float32
}

// Read sets the field values of the message from the raw message payload
func (m *SimState) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Q1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q2)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q3)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q4)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Roll)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Pitch)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yaw)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Lat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Alt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.StdDevHorz)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.StdDevVert)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vn)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Ve)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vd)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *SimState) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Q1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q2)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q3)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q4)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Roll)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Pitch)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yaw)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Lat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Alt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.StdDevHorz)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.StdDevVert)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vn)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Ve)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vd)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
