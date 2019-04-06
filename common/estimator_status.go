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

/*EstimatorStatus Estimator status message including flags, innovation test ratios and estimated accuracies. The flags message is an integer bitmask containing information on which EKF outputs are valid. See the ESTIMATOR_STATUS_FLAGS enum definition for further information. The innovation test ratios show the magnitude of the sensor innovation divided by the innovation check threshold. Under normal operation the innovation test ratios should be below 0.5 with occasional values up to 1.0. Values greater than 1.0 should be rare under normal operation and indicate that a measurement has been rejected by the filter. The user should be notified if an innovation test ratio greater than 1.0 is recorded. Notifications for values in the range between 0.5 and 1.0 should be optional and controllable by the user. */
type EstimatorStatus struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*VelRatio Velocity innovation test ratio */
	VelRatio float32
	/*PosHorizRatio Horizontal position innovation test ratio */
	PosHorizRatio float32
	/*PosVertRatio Vertical position innovation test ratio */
	PosVertRatio float32
	/*MagRatio Magnetometer innovation test ratio */
	MagRatio float32
	/*HaglRatio Height above terrain innovation test ratio */
	HaglRatio float32
	/*TasRatio True airspeed innovation test ratio */
	TasRatio float32
	/*PosHorizAccuracy Horizontal position 1-STD accuracy relative to the EKF local origin */
	PosHorizAccuracy float32
	/*PosVertAccuracy Vertical position 1-STD accuracy relative to the EKF local origin */
	PosVertAccuracy float32
	/*Flags Bitmap indicating which EKF outputs are valid. */
	Flags uint16
}

// Read sets the field values of the message from the raw message payload
func (m *EstimatorStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosHorizRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosVertRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MagRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HaglRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TasRatio)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosHorizAccuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosVertAccuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *EstimatorStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosHorizRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosVertRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MagRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HaglRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TasRatio)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosHorizAccuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosVertAccuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
