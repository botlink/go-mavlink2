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

/*CameraInformation Information about a camera */
type CameraInformation struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*FirmwareVersion Version of the camera firmware (v << 24 & 0xff = Dev, v << 16 & 0xff = Patch, v << 8 & 0xff = Minor, v & 0xff = Major) */
	FirmwareVersion uint32
	/*FocalLength Focal length */
	FocalLength float32
	/*SensorSizeH Image sensor size horizontal */
	SensorSizeH float32
	/*SensorSizeV Image sensor size vertical */
	SensorSizeV float32
	/*Flags Bitmap of camera capability flags. */
	Flags uint32
	/*ResolutionH Horizontal image resolution */
	ResolutionH uint16
	/*ResolutionV Vertical image resolution */
	ResolutionV uint16
	/*CamDefinitionVersion Camera definition version (iteration) */
	CamDefinitionVersion uint16
	/*VendorName Name of the camera vendor */
	VendorName []uint8
	/*ModelName Name of the camera model */
	ModelName []uint8
	/*LensID Reserved for a lens ID */
	LensID uint8
	/*CamDefinitionURI Camera definition URI (if any, otherwise only basic functions will be available). */
	CamDefinitionURI []byte
}

// GetVersion gets the MAVLink version of the Message contents
func (m *CameraInformation) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *CameraInformation) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *CameraInformation) GetName() string {
	return "CameraInformation"
}

// GetID gets the ID of the Message
func (m *CameraInformation) GetID() uint32 {
	return 259
}

// Read sets the field values of the message from the raw message payload
func (m *CameraInformation) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FirmwareVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FocalLength)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SensorSizeH)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SensorSizeV)
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

	err = binary.Read(reader, binary.LittleEndian, &m.CamDefinitionVersion)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VendorName)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ModelName)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.LensID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CamDefinitionURI)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CameraInformation) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FirmwareVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FocalLength)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SensorSizeH)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SensorSizeV)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.CamDefinitionVersion)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VendorName)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ModelName)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.LensID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CamDefinitionURI)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
