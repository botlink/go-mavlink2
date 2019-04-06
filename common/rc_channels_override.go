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

/*RcChannelsOverrIDe The RAW values of the RC channels sent to the MAV to override info received from the RC radio. A value of UINT16_MAX means no change to that channel. A value of 0 means control of that channel should be released back to the RC radio. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification. */
type RcChannelsOverrIDe struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Chan1Raw RC channel 1 value. A value of UINT16_MAX means to ignore this field. */
	Chan1Raw uint16
	/*Chan2Raw RC channel 2 value. A value of UINT16_MAX means to ignore this field. */
	Chan2Raw uint16
	/*Chan3Raw RC channel 3 value. A value of UINT16_MAX means to ignore this field. */
	Chan3Raw uint16
	/*Chan4Raw RC channel 4 value. A value of UINT16_MAX means to ignore this field. */
	Chan4Raw uint16
	/*Chan5Raw RC channel 5 value. A value of UINT16_MAX means to ignore this field. */
	Chan5Raw uint16
	/*Chan6Raw RC channel 6 value. A value of UINT16_MAX means to ignore this field. */
	Chan6Raw uint16
	/*Chan7Raw RC channel 7 value. A value of UINT16_MAX means to ignore this field. */
	Chan7Raw uint16
	/*Chan8Raw RC channel 8 value. A value of UINT16_MAX means to ignore this field. */
	Chan8Raw uint16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*Chan9Raw RC channel 9 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan9Raw uint16
	/*Chan10Raw RC channel 10 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan10Raw uint16
	/*Chan11Raw RC channel 11 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan11Raw uint16
	/*Chan12Raw RC channel 12 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan12Raw uint16
	/*Chan13Raw RC channel 13 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan13Raw uint16
	/*Chan14Raw RC channel 14 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan14Raw uint16
	/*Chan15Raw RC channel 15 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan15Raw uint16
	/*Chan16Raw RC channel 16 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan16Raw uint16
	/*Chan17Raw RC channel 17 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan17Raw uint16
	/*Chan18Raw RC channel 18 value. A value of 0 or UINT16_MAX means to ignore this field. */
	Chan18Raw uint16
}

// GetVersion gets the MAVLink version of the Message contents
func (m *RcChannelsOverrIDe) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *RcChannelsOverrIDe) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *RcChannelsOverrIDe) GetName() string {
	return "RcChannelsOverrIDe"
}

// GetID gets the ID of the Message
func (m *RcChannelsOverrIDe) GetID() uint32 {
	return 70
}

// Read sets the field values of the message from the raw message payload
func (m *RcChannelsOverrIDe) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
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

	err = binary.Read(reader, binary.LittleEndian, &m.TargetSystem)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TargetComponent)
	if err != nil {
		return
	}

	if version == 2 {
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

		err = binary.Read(reader, binary.LittleEndian, &m.Chan13Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Chan14Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Chan15Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Chan16Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Chan17Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Chan18Raw)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *RcChannelsOverrIDe) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
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

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.TargetComponent)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
