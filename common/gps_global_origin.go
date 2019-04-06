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

/*GpsGlobalOrigin Once the MAV sets a new GPS-Local correspondence, this message announces the origin (0,0,0) position */
type GpsGlobalOrigin struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Latitude Latitude (WGS84) */
	Latitude int32
	/*Longitude Longitude (WGS84) */
	Longitude int32
	/*Altitude Altitude (MSL). Positive for up. */
	Altitude int32
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
}

// Read sets the field values of the message from the raw message payload
func (m *GpsGlobalOrigin) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Latitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Longitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Altitude)
	if err != nil {
		return
	}

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsGlobalOrigin) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Latitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Longitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Altitude)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.Altitude)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
