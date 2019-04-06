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

/*VIDeoStreamStatus Information about the status of a video stream. */
type VIDeoStreamStatus struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Framerate Frame rate */
	Framerate float32
	/*Bitrate Bit rate */
	Bitrate uint32
	/*Flags Bitmap of stream status flags */
	Flags uint16
	/*ResolutionH Horizontal resolution */
	ResolutionH uint16
	/*ResolutionV Vertical resolution */
	ResolutionV uint16
	/*Rotation Video image rotation clockwise */
	Rotation uint16
	/*Hfov Horizontal Field of view */
	Hfov uint16
	/*StreamID Video Stream ID (1 for first, 2 for second, etc.) */
	StreamID uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *VIDeoStreamStatus) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *VIDeoStreamStatus) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *VIDeoStreamStatus) GetName() string {
	return "VIDeoStreamStatus"
}

// GetID gets the ID of the Message
func (m *VIDeoStreamStatus) GetID() uint32 {
	return 270
}

// Read sets the field values of the message from the raw message payload
func (m *VIDeoStreamStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Framerate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Bitrate)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ResolutionH)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ResolutionV)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Rotation)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Hfov)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.StreamID)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *VIDeoStreamStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Framerate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Bitrate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ResolutionH)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ResolutionV)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Rotation)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Hfov)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.StreamID)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
