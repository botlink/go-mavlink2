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

/*ParamMapRc Bind a RC channel to a parameter. The parameter should change according to the RC channel value. */
type ParamMapRc struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*ParamValue0 Initial parameter value */
	ParamValue0 float32
	/*Scale Scale, maps the RC range [-1, 1] to a parameter value */
	Scale float32
	/*ParamValueMin Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation) */
	ParamValueMin float32
	/*ParamValueMax Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation) */
	ParamValueMax float32
	/*ParamIndex Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index. */
	ParamIndex int16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*ParamID Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string */
	ParamID []byte
	/*ParameterRcChannelIndex Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC. */
	ParameterRcChannelIndex uint8
}

// Read sets the field values of the message from the raw message payload
func (m *ParamMapRc) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.ParamValue0)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Scale)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamValueMin)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamValueMax)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParamIndex)
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

	err = binary.Read(reader, binary.LittleEndian, &m.ParamID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ParameterRcChannelIndex)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ParamMapRc) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.ParamValue0)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Scale)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamValueMin)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamValueMax)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamIndex)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.ParamID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ParameterRcChannelIndex)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
