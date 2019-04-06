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

/*AutopilotVersion Version and capability of autopilot software */
type AutopilotVersion struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Capabilities Bitmap of capabilities */
	Capabilities uint64
	/*UID UID if provided by hardware (see uid2) */
	UID uint64
	/*FlightSwVersion Firmware version number */
	FlightSwVersion uint32
	/*MIDdlewareSwVersion Middleware version number */
	MIDdlewareSwVersion uint32
	/*OsSwVersion Operating system version number */
	OsSwVersion uint32
	/*BoardVersion HW / board version (last 8 bytes should be silicon ID, if any) */
	BoardVersion uint32
	/*VendorID ID of the board vendor */
	VendorID uint16
	/*ProductID ID of the product */
	ProductID uint16
	/*FlightCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	FlightCustomVersion []uint8
	/*MIDdlewareCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	MIDdlewareCustomVersion []uint8
	/*OsCustomVersion Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases. */
	OsCustomVersion []uint8
	/*UID2 UID if provided by hardware (supersedes the uid field. If this is non-zero, use this field, otherwise use uid) */
	UID2 []uint8
}

// Read sets the field values of the message from the raw message payload
func (m *AutopilotVersion) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Capabilities)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.UID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FlightSwVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MIDdlewareSwVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.OsSwVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BoardVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VendorID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ProductID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FlightCustomVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MIDdlewareCustomVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.OsCustomVersion)
	if err != nil {
		return
	}

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.UID2)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *AutopilotVersion) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Capabilities)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.UID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FlightSwVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MIDdlewareSwVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.OsSwVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BoardVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VendorID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ProductID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FlightCustomVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MIDdlewareCustomVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.OsCustomVersion)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.OsCustomVersion)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
