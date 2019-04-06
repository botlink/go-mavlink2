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

/*Heartbeat The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). */
type Heartbeat struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*CustomMode A bitfield for use for autopilot-specific flags */
	CustomMode uint32
	/*Type Type of the system (quadrotor, helicopter, etc.). Components use the same type as their associated system. */
	Type uint8
	/*Autopilot Autopilot type / class. */
	Autopilot uint8
	/*BaseMode System mode bitmap. */
	BaseMode uint8
	/*SystemStatus System status flag. */
	SystemStatus uint8
	/*MavlinkVersion MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version */
	MavlinkVersion uint8
}

// Read sets the field values of the message from the raw message payload
func (m *Heartbeat) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.CustomMode)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Type)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Autopilot)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaseMode)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SystemStatus)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MavlinkVersion)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *Heartbeat) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.CustomMode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Type)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Autopilot)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaseMode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SystemStatus)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MavlinkVersion)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
