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

/*MissionWritePartialList This message is sent to the MAV to write a partial list. If start index == end index, only one item will be transmitted / updated. If the start index is NOT 0 and above the current list size, this request should be REJECTED! */
type MissionWritePartialList struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*StartIndex Start index. Must be smaller / equal to the largest index of the current onboard list. */
	StartIndex int16
	/*EndIndex End index, equal or greater than start index. */
	EndIndex int16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*MissionType Mission type. */
	MissionType uint8
}

// Read sets the field values of the message from the raw message payload
func (m *MissionWritePartialList) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.StartIndex)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.EndIndex)
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

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.MissionType)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *MissionWritePartialList) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.StartIndex)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.EndIndex)
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

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
