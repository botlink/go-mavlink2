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

/*DataTransmissionHandshake Handshake message to initiate, control and stop image streaming when using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html. */
type DataTransmissionHandshake struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Size total data size (set on ACK only). */
	Size uint32
	/*WIDth Width of a matrix or image. */
	WIDth uint16
	/*Height Height of a matrix or image. */
	Height uint16
	/*Packets Number of packets being sent (set on ACK only). */
	Packets uint16
	/*Type Type of requested/acknowledged data. */
	Type uint8
	/*Payload Payload size per packet (normally 253 byte, see DATA field size in message ENCAPSULATED_DATA) (set on ACK only). */
	Payload uint8
	/*JpgQuality JPEG quality. Values: [1-100]. */
	JpgQuality uint8
}

// Read sets the field values of the message from the raw message payload
func (m *DataTransmissionHandshake) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Size)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.WIDth)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Height)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Packets)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Type)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Payload)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.JpgQuality)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *DataTransmissionHandshake) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Size)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.WIDth)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Height)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Packets)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Type)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Payload)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.JpgQuality)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
