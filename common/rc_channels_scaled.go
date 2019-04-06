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

/*RcChannelsScaled The scaled values of the RC channels received: (-100%) -10000, (0%) 0, (100%) 10000. Channels that are inactive should be set to UINT16_MAX. */
type RcChannelsScaled struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Chan1Scaled RC channel 1 value scaled. */
	Chan1Scaled int16
	/*Chan2Scaled RC channel 2 value scaled. */
	Chan2Scaled int16
	/*Chan3Scaled RC channel 3 value scaled. */
	Chan3Scaled int16
	/*Chan4Scaled RC channel 4 value scaled. */
	Chan4Scaled int16
	/*Chan5Scaled RC channel 5 value scaled. */
	Chan5Scaled int16
	/*Chan6Scaled RC channel 6 value scaled. */
	Chan6Scaled int16
	/*Chan7Scaled RC channel 7 value scaled. */
	Chan7Scaled int16
	/*Chan8Scaled RC channel 8 value scaled. */
	Chan8Scaled int16
	/*Port Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX. */
	Port uint8
	/*Rssi Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown. */
	Rssi uint8
}

// Read sets the field values of the message from the raw message payload
func (m *RcChannelsScaled) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan1Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan2Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan3Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan4Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan5Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan6Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan7Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Chan8Scaled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Port)
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
func (m *RcChannelsScaled) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan1Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan2Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan3Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan4Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan5Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan6Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan7Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Chan8Scaled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Port)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Rssi)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
