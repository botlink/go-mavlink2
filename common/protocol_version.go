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

/*ProtocolVersion Version and capability of protocol version. This message is the response to REQUEST_PROTOCOL_VERSION and is used as part of the handshaking to establish which MAVLink version should be used on the network. Every node should respond to REQUEST_PROTOCOL_VERSION to enable the handshaking. Library implementers should consider adding this into the default decoding state machine to allow the protocol core to respond directly. */
type ProtocolVersion struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*Version Currently active MAVLink version number * 100: v1.0 is 100, v2.0 is 200, etc. */
	Version uint16
	/*MinVersion Minimum MAVLink version supported */
	MinVersion uint16
	/*MaxVersion Maximum MAVLink version supported (set to the same value as version by default) */
	MaxVersion uint16
	/*SpecVersionHash The first 8 bytes (not characters printed in hex!) of the git hash. */
	SpecVersionHash []uint8
	/*LibraryVersionHash The first 8 bytes (not characters printed in hex!) of the git hash. */
	LibraryVersionHash []uint8
}

// Read sets the field values of the message from the raw message payload
func (m *ProtocolVersion) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Version)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MinVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.MaxVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SpecVersionHash)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.LibraryVersionHash)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *ProtocolVersion) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Version)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MinVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.MaxVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SpecVersionHash)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.LibraryVersionHash)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
