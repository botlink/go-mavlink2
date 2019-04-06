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

/*CommandInt Message encoding a command with parameters as scaled integers. Scaling depends on the actual command value. */
type CommandInt struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Param1 PARAM1, see MAV_CMD enum */
	Param1 float32
	/*Param2 PARAM2, see MAV_CMD enum */
	Param2 float32
	/*Param3 PARAM3, see MAV_CMD enum */
	Param3 float32
	/*Param4 PARAM4, see MAV_CMD enum */
	Param4 float32
	/*X PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7 */
	X int32
	/*Y PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7 */
	Y int32
	/*Z PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame). */
	Z float32
	/*Command The scheduled action for the mission item. */
	Command uint16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*Frame The coordinate system of the COMMAND. */
	Frame uint8
	/*Current false:0, true:1 */
	Current uint8
	/*Autocontinue autocontinue to next wp */
	Autocontinue uint8
}

// Read sets the field values of the message from the raw message payload
func (m *CommandInt) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Param1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Param2)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Param3)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Param4)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.X)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Y)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Z)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Command)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Frame)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Current)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Autocontinue)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CommandInt) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Param1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Param2)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Param3)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Param4)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.X)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Y)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Z)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Command)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Current)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Autocontinue)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
