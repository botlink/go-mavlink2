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

/*SafetyAllowedArea Read out the safety zone the MAV currently assumes. */
type SafetyAllowedArea struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*P1X x position 1 / Latitude 1 */
	P1X float32
	/*P1Y y position 1 / Longitude 1 */
	P1Y float32
	/*P1Z z position 1 / Altitude 1 */
	P1Z float32
	/*P2X x position 2 / Latitude 2 */
	P2X float32
	/*P2Y y position 2 / Longitude 2 */
	P2Y float32
	/*P2Z z position 2 / Altitude 2 */
	P2Z float32
	/*Frame Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down. */
	Frame uint8
}

// Read sets the field values of the message from the raw message payload
func (m *SafetyAllowedArea) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.P1X)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.P1Y)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.P1Z)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.P2X)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.P2Y)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.P2Z)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Frame)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *SafetyAllowedArea) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.P1X)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.P1Y)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.P1Z)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.P2X)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.P2Y)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.P2Z)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Frame)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
