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

/*ParamExtRequestRead Request to read the value of a parameter with the either the param_id string id or param_index. */
type ParamExtRequestRead struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*ParamIndex Parameter index. Set to -1 to use the Parameter ID field as identifier (else param_id will be ignored) */
	ParamIndex int16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*ParamID Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string */
	ParamID []byte
}

// Read sets the field values of the message from the raw message payload
func (m *ParamExtRequestRead) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.ParamIndex)
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

	err = binary.Read(reader, binary.LittleEndian, &m.ParamID)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ParamExtRequestRead) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.ParamIndex)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamID)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
