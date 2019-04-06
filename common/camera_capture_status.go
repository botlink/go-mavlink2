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

/*CameraCaptureStatus Information about the status of a capture. */
type CameraCaptureStatus struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*ImageInterval Image capture interval */
	ImageInterval float32
	/*RecordingTimeMs Time since recording started */
	RecordingTimeMs uint32
	/*AvailableCapacity Available storage capacity. */
	AvailableCapacity float32
	/*ImageStatus Current status of image capturing (0: idle, 1: capture in progress, 2: interval set but idle, 3: interval set and capture in progress) */
	ImageStatus uint8
	/*VIDeoStatus Current status of video capturing (0: idle, 1: capture in progress) */
	VIDeoStatus uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CameraCaptureStatus) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CameraCaptureStatus) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *CameraCaptureStatus) GetName() string {
	return "CameraCaptureStatus"
}

// GetID gets the ID of the Message
func (m *CameraCaptureStatus) GetID() uint32 {
	return 262
}

// Read sets the field values of the message from the raw message payload
func (m *CameraCaptureStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ImageInterval)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RecordingTimeMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AvailableCapacity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ImageStatus)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VIDeoStatus)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CameraCaptureStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ImageInterval)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RecordingTimeMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AvailableCapacity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ImageStatus)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VIDeoStatus)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
