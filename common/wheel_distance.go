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

/*WheelDistance Cumulative distance traveled for each reported wheel. */
type WheelDistance struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (synced to UNIX time or since system boot). */
	TimeUsec uint64
	/*Distance Distance reported by individual wheel encoders. Forward rotations increase values, reverse rotations decrease them. Not all wheels will necessarily have wheel encoders; the mapping of encoders to wheel positions must be agreed/understood by the endpoints. */
	Distance []float64
	/*Count Number of wheels reported. */
	Count uint8
}

// Read sets the field values of the message from the raw message payload
func (m *WheelDistance) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Distance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Count)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *WheelDistance) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Distance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Count)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
