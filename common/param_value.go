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

/*ParamValue Emit the value of a onboard parameter. The inclusion of param_count and param_index in the message allows the recipient to keep track of received parameters and allows him to re-request missing parameters after a loss or timeout. */
type ParamValue struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*ParamValue Onboard parameter value */
	ParamValue float32
	/*ParamCount Total number of onboard parameters */
	ParamCount uint16
	/*ParamIndex Index of this onboard parameter */
	ParamIndex uint16
	/*ParamID Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string */
	ParamID []byte
	/*ParamType Onboard parameter type. */
	ParamType uint8
}

// Read sets the field values of the message from the raw message payload
func (m *ParamValue) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.ParamValue)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamCount)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamIndex)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamType)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ParamValue) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.ParamValue)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamCount)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamIndex)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamType)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
