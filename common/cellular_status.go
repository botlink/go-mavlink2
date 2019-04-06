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

/*CellularStatus Report current used cellular network status */
type CellularStatus struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*CID Cell ID. If unknown, set to: UINT32_MAX */
	CID uint32
	/*Status Status bitmap */
	Status uint16
	/*Mcc Mobile country code. If unknown, set to: UINT16_MAX */
	Mcc uint16
	/*Mnc Mobile network code. If unknown, set to: UINT16_MAX */
	Mnc uint16
	/*Lac Location area code. If unknown, set to: 0 */
	Lac uint16
	/*Type Cellular network radio type: gsm, cdma, lte... */
	Type uint8
	/*Quality Cellular network RSSI/RSRP in dBm, absolute value */
	Quality uint8
}

// Read sets the field values of the message from the raw message payload
func (m *CellularStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.CID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Status)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Mcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Mnc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lac)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Type)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Quality)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CellularStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.CID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Status)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Mcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Mnc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lac)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Type)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Quality)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
