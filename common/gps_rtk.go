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

/*GpsRtk RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting */
type GpsRtk struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeLastBaselineMs Time since boot of last baseline message received. */
	TimeLastBaselineMs uint32
	/*Tow GPS Time of Week of last baseline */
	Tow uint32
	/*BaselineAMm Current baseline in ECEF x or NED north component. */
	BaselineAMm int32
	/*BaselineBMm Current baseline in ECEF y or NED east component. */
	BaselineBMm int32
	/*BaselineCMm Current baseline in ECEF z or NED down component. */
	BaselineCMm int32
	/*Accuracy Current estimate of baseline accuracy. */
	Accuracy uint32
	/*IarNumHypotheses Current number of integer ambiguity hypotheses. */
	IarNumHypotheses int32
	/*Wn GPS Week Number of last baseline */
	Wn uint16
	/*RtkReceiverID Identification of connected RTK receiver. */
	RtkReceiverID uint8
	/*RtkHealth GPS-specific health report for RTK data. */
	RtkHealth uint8
	/*RtkRate Rate of baseline messages being received by GPS */
	RtkRate uint8
	/*Nsats Current number of sats used for RTK calculation. */
	Nsats uint8
	/*BaselineCoordsType Coordinate system of baseline */
	BaselineCoordsType uint8
}

// Read sets the field values of the message from the raw message payload
func (m *GpsRtk) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeLastBaselineMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Tow)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaselineAMm)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaselineBMm)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaselineCMm)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Accuracy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.IarNumHypotheses)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Wn)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RtkReceiverID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RtkHealth)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RtkRate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Nsats)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BaselineCoordsType)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsRtk) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeLastBaselineMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Tow)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaselineAMm)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaselineBMm)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaselineCMm)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Accuracy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.IarNumHypotheses)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Wn)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RtkReceiverID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RtkHealth)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RtkRate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Nsats)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BaselineCoordsType)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
