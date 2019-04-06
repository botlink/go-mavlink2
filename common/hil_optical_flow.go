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

/*HilOpticalFlow Simulated optical flow from a flow sensor (e.g. PX4FLOW or optical mouse sensor) */
type HilOpticalFlow struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*IntegrationTimeUs Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the. */
	IntegrationTimeUs uint32
	/*IntegratedX Flow in radians around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.) */
	IntegratedX float32
	/*IntegratedY Flow in radians around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.) */
	IntegratedY float32
	/*IntegratedXgyro RH rotation around X axis */
	IntegratedXgyro float32
	/*IntegratedYgyro RH rotation around Y axis */
	IntegratedYgyro float32
	/*IntegratedZgyro RH rotation around Z axis */
	IntegratedZgyro float32
	/*TimeDeltaDistanceUs Time since the distance was sampled. */
	TimeDeltaDistanceUs uint32
	/*Distance Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance. */
	Distance float32
	/*Temperature Temperature */
	Temperature int16
	/*SensorID Sensor ID */
	SensorID uint8
	/*Quality Optical flow quality / confidence. 0: no valid flow, 255: maximum quality */
	Quality uint8
}

// Read sets the field values of the message from the raw message payload
func (m *HilOpticalFlow) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegrationTimeUs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegratedX)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegratedY)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegratedXgyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegratedYgyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IntegratedZgyro)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeDeltaDistanceUs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Distance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SensorID)
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
func (m *HilOpticalFlow) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegrationTimeUs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegratedX)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegratedY)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegratedXgyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegratedYgyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IntegratedZgyro)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeDeltaDistanceUs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Distance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SensorID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Quality)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
