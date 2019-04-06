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

/*GpsRtcmData RTCM message for injecting into the onboard GPS (used for DGPS) */
type GpsRtcmData struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Flags LSB: 1 means message is fragmented, next 2 bits are the fragment ID, the remaining 5 bits are used for the sequence ID. Messages are only to be flushed to the GPS when the entire message has been reconstructed on the autopilot. The fragment ID specifies which order the fragments should be assembled into a buffer, while the sequence ID is used to detect a mismatch between different buffers. The buffer is considered fully reconstructed when either all 4 fragments are present, or all the fragments before the first fragment with a non full payload is received. This management is used to ensure that normal GPS operation doesn't corrupt RTCM data, and to recover from a unreliable transport delivery order. */
	Flags uint8
	/*Len data length */
	Len uint8
	/*Data RTCM message (may be fragmented) */
	Data []uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GpsRtcmData) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GpsRtcmData) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *GpsRtcmData) GetName() string {
	return "GpsRtcmData"
}

// GetID gets the ID of the Message
func (m *GpsRtcmData) GetID() uint32 {
	return 233
}

// Read sets the field values of the message from the raw message payload
func (m *GpsRtcmData) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Len)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Data)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsRtcmData) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Len)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
