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

/*TerrainData Terrain data sent from GCS. The lat/lon and grid_spacing must be the same as a lat/lon from a TERRAIN_REQUEST */
type TerrainData struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*Lat Latitude of SW corner of first grid */
	Lat int32
	/*Lon Longitude of SW corner of first grid */
	Lon int32
	/*GrIDSpacing Grid spacing */
	GrIDSpacing uint16
	/*Data Terrain data MSL */
	Data []int16
	/*GrIDbit bit within the terrain request mask */
	GrIDbit uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *TerrainData) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *TerrainData) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *TerrainData) GetName() string {
	return "TerrainData"
}

// GetID gets the ID of the Message
func (m *TerrainData) GetID() uint32 {
	return 134
}

// Read sets the field values of the message from the raw message payload
func (m *TerrainData) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.Lat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.GrIDSpacing)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Data)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.GrIDbit)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *TerrainData) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.Lat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.GrIDSpacing)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Data)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.GrIDbit)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
