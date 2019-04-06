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

/*ServoOutputRaw The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. */
type ServoOutputRaw struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint32
	/*Servo1Raw Servo output 1 value */
	Servo1Raw uint16
	/*Servo2Raw Servo output 2 value */
	Servo2Raw uint16
	/*Servo3Raw Servo output 3 value */
	Servo3Raw uint16
	/*Servo4Raw Servo output 4 value */
	Servo4Raw uint16
	/*Servo5Raw Servo output 5 value */
	Servo5Raw uint16
	/*Servo6Raw Servo output 6 value */
	Servo6Raw uint16
	/*Servo7Raw Servo output 7 value */
	Servo7Raw uint16
	/*Servo8Raw Servo output 8 value */
	Servo8Raw uint16
	/*Port Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX. */
	Port uint8
	/*Servo9Raw Servo output 9 value */
	Servo9Raw uint16
	/*Servo10Raw Servo output 10 value */
	Servo10Raw uint16
	/*Servo11Raw Servo output 11 value */
	Servo11Raw uint16
	/*Servo12Raw Servo output 12 value */
	Servo12Raw uint16
	/*Servo13Raw Servo output 13 value */
	Servo13Raw uint16
	/*Servo14Raw Servo output 14 value */
	Servo14Raw uint16
	/*Servo15Raw Servo output 15 value */
	Servo15Raw uint16
	/*Servo16Raw Servo output 16 value */
	Servo16Raw uint16
}

// Read sets the field values of the message from the raw message payload
func (m *ServoOutputRaw) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo1Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo2Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo3Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo4Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo5Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo6Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo7Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Servo8Raw)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Port)
	if err != nil {
		return
	}

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.Servo9Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo10Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo11Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo12Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo13Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo14Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo15Raw)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.Servo16Raw)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *ServoOutputRaw) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo1Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo2Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo3Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo4Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo5Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo6Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo7Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Servo8Raw)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Port)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.Port)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
