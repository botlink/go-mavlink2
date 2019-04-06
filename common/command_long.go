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

/*CommandLong Send a command with up to seven parameters to the MAV */
type CommandLong struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Param1 Parameter 1 (for the specific command). */
	Param1 float32
	/*Param2 Parameter 2 (for the specific command). */
	Param2 float32
	/*Param3 Parameter 3 (for the specific command). */
	Param3 float32
	/*Param4 Parameter 4 (for the specific command). */
	Param4 float32
	/*Param5 Parameter 5 (for the specific command). */
	Param5 float32
	/*Param6 Parameter 6 (for the specific command). */
	Param6 float32
	/*Param7 Parameter 7 (for the specific command). */
	Param7 float32
	/*Command Command ID (of command to send). */
	Command uint16
	/*TargetSystem System which should execute the command */
	TargetSystem uint8
	/*TargetComponent Component which should execute the command, 0 for all components */
	TargetComponent uint8
	/*Confirmation 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command) */
	Confirmation uint8
}

// Read sets the field values of the message from the raw message payload
func (m *CommandLong) Read(version int, payload []byte) (err error) {
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

	err = binary.Read(reader, binary.LittleEndian, &m.Param5)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Param6)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Param7)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Confirmation)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CommandLong) Write(version int) ([]byte, error) {
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Param5)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Param6)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Param7)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Confirmation)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
