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

/*HilStateQuaternion Sent from simulation to autopilot, avoids in contrast to HIL_STATE singularities. This packet is useful for high throughput applications such as hardware in the loop simulations. */
type HilStateQuaternion struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*AttitudeQuaternion Vehicle attitude expressed as normalized quaternion in w, x, y, z order (with 1 0 0 0 being the null-rotation) */
	AttitudeQuaternion []float32
	/*Rollspeed Body frame roll / phi angular speed */
	Rollspeed float32
	/*Pitchspeed Body frame pitch / theta angular speed */
	Pitchspeed float32
	/*Yawspeed Body frame yaw / psi angular speed */
	Yawspeed float32
	/*Lat Latitude */
	Lat int32
	/*Lon Longitude */
	Lon int32
	/*Alt Altitude */
	Alt int32
	/*Vx Ground X Speed (Latitude) */
	Vx int16
	/*Vy Ground Y Speed (Longitude) */
	Vy int16
	/*Vz Ground Z Speed (Altitude) */
	Vz int16
	/*IndAirspeed Indicated airspeed */
	IndAirspeed uint16
	/*TrueAirspeed True airspeed */
	TrueAirspeed uint16
	/*Xacc X acceleration */
	Xacc int16
	/*Yacc Y acceleration */
	Yacc int16
	/*Zacc Z acceleration */
	Zacc int16
}

// GetVersion gets the MAVLink version of the Message contents
func (m *HilStateQuaternion) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *HilStateQuaternion) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *HilStateQuaternion) GetName() string {
	return "HilStateQuaternion"
}

// GetID gets the ID of the Message
func (m *HilStateQuaternion) GetID() uint32 {
	return 115
}

// Read sets the field values of the message from the raw message payload
func (m *HilStateQuaternion) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AttitudeQuaternion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Rollspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Pitchspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yawspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Alt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vx)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vz)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IndAirspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TrueAirspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Xacc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Yacc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Zacc)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HilStateQuaternion) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AttitudeQuaternion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Rollspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Pitchspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yawspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Alt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vx)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vz)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IndAirspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TrueAirspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Xacc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Yacc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Zacc)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
