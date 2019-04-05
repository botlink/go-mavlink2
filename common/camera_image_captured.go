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
package common

import (
	"bytes"
	"encoding/binary"
)

/*CameraImageCaptured Information about a captured image */
type CameraImageCaptured struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUtc Timestamp (time since UNIX epoch) in UTC. 0 for unknown. */
	TimeUtc uint64
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Lat Latitude where image was taken */
	Lat int32
	/*Lon Longitude where capture was taken */
	Lon int32
	/*Alt Altitude (MSL) where image was taken */
	Alt int32
	/*RelativeAlt Altitude above ground */
	RelativeAlt int32
	/*Q Quaternion of camera orientation (w, x, y, z order, zero-rotation is 0, 0, 0, 0) */
	Q []float32
	/*ImageIndex Zero based index of this image (image count since armed -1) */
	ImageIndex int32
	/*CameraID Camera ID (1 for first, 2 for second, etc.) */
	CameraID uint8
	/*CaptureResult Boolean indicating success (1) or failure (0) while capturing this image. */
	CaptureResult int8
	/*FileURL URL of image taken. Either local storage or http://foo.jpg if camera provides an HTTP interface. */
	FileURL []byte
}

// Read sets the field values of the message from the raw message payload
func (m *CameraImageCaptured) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUtc)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeBootMs)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lat)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Lon)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Alt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RelativeAlt)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Q)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ImageIndex)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CameraID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CaptureResult)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FileURL)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *CameraImageCaptured) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUtc)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeBootMs)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lat)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Lon)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Alt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RelativeAlt)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Q)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ImageIndex)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CameraID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CaptureResult)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FileURL)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
