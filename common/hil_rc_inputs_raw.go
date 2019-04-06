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

/*HilRcInputsRaw Sent from simulation to autopilot. The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification. */
type HilRcInputsRaw struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Chan1Raw RC channel 1 value */
	Chan1Raw uint16
	/*Chan2Raw RC channel 2 value */
	Chan2Raw uint16
	/*Chan3Raw RC channel 3 value */
	Chan3Raw uint16
	/*Chan4Raw RC channel 4 value */
	Chan4Raw uint16
	/*Chan5Raw RC channel 5 value */
	Chan5Raw uint16
	/*Chan6Raw RC channel 6 value */
	Chan6Raw uint16
	/*Chan7Raw RC channel 7 value */
	Chan7Raw uint16
	/*Chan8Raw RC channel 8 value */
	Chan8Raw uint16
	/*Chan9Raw RC channel 9 value */
	Chan9Raw uint16
	/*Chan10Raw RC channel 10 value */
	Chan10Raw uint16
	/*Chan11Raw RC channel 11 value */
	Chan11Raw uint16
	/*Chan12Raw RC channel 12 value */
	Chan12Raw uint16
	/*Rssi Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown. */
	Rssi uint8
}

// Read sets the field values of the message from the raw message payload
func (m *HilRcInputsRaw) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan1Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan2Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan3Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan4Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan5Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan6Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan7Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan8Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan9Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan10Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan11Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan12Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Rssi)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HilRcInputsRaw) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan1Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan2Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan3Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan4Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan5Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan6Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan7Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan8Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan9Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan10Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan11Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan12Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Rssi)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
