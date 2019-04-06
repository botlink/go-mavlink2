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

/*ControlSystemState The smoothed, monotonic system state used to feed the control loops of the system. */
type ControlSystemState struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*XAcc X acceleration in body frame */
	XAcc float32
	/*YAcc Y acceleration in body frame */
	YAcc float32
	/*ZAcc Z acceleration in body frame */
	ZAcc float32
	/*XVel X velocity in body frame */
	XVel float32
	/*YVel Y velocity in body frame */
	YVel float32
	/*ZVel Z velocity in body frame */
	ZVel float32
	/*XPos X position in local frame */
	XPos float32
	/*YPos Y position in local frame */
	YPos float32
	/*ZPos Z position in local frame */
	ZPos float32
	/*Airspeed Airspeed, set to -1 if unknown */
	Airspeed float32
	/*VelVariance Variance of body velocity estimate */
	VelVariance []float32
	/*PosVariance Variance in local position */
	PosVariance []float32
	/*Q The attitude, represented as Quaternion */
	Q []float32
	/*RollRate Angular rate in roll axis */
	RollRate float32
	/*PitchRate Angular rate in pitch axis */
	PitchRate float32
	/*YawRate Angular rate in yaw axis */
	YawRate float32
}

// Read sets the field values of the message from the raw message payload
func (m *ControlSystemState) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.XAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.YAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ZAcc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.XVel)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.YVel)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ZVel)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.XPos)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.YPos)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ZPos)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Airspeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VelVariance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PosVariance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RollRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.PitchRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.YawRate)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ControlSystemState) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.XAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.YAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ZAcc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.XVel)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.YVel)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ZVel)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.XPos)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.YPos)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ZPos)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Airspeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VelVariance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PosVariance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RollRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.PitchRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.YawRate)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
