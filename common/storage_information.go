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

/*StorageInformation Information about a storage medium. */
type StorageInformation struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*TotalCapacity Total capacity. */
	TotalCapacity float32
	/*UsedCapacity Used capacity. */
	UsedCapacity float32
	/*AvailableCapacity Available storage capacity. */
	AvailableCapacity float32
	/*ReadSpeed Read speed. */
	ReadSpeed float32
	/*WriteSpeed Write speed. */
	WriteSpeed float32
	/*StorageID Storage ID (1 for first, 2 for second, etc.) */
	StorageID uint8
	/*StorageCount Number of storage devices */
	StorageCount uint8
	/*Status Status of storage (0 not available, 1 unformatted, 2 formatted) */
	Status uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *StorageInformation) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *StorageInformation) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *StorageInformation) GetName() string {
	return "StorageInformation"
}

// GetID gets the ID of the Message
func (m *StorageInformation) GetID() uint32 {
	return 261
}

// Read sets the field values of the message from the raw message payload
func (m *StorageInformation) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TotalCapacity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.UsedCapacity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AvailableCapacity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ReadSpeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WriteSpeed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.StorageID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.StorageCount)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Status)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *StorageInformation) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TotalCapacity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.UsedCapacity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AvailableCapacity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ReadSpeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WriteSpeed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.StorageID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.StorageCount)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Status)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
