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

/*GlobalPositionIntCov The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It  is designed as scaled integer message since the resolution of float is not sufficient. NOTE: This message is intended for onboard networks / companion computers and higher-bandwidth links and optimized for accuracy and completeness. Please use the GLOBAL_POSITION_INT message for a minimal subset. */
type GlobalPositionIntCov struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Lat Latitude */
	Lat int32
	/*Lon Longitude */
	Lon int32
	/*Alt Altitude in meters above MSL */
	Alt int32
	/*RelativeAlt Altitude above ground */
	RelativeAlt int32
	/*Vx Ground X Speed (Latitude) */
	Vx float32
	/*Vy Ground Y Speed (Longitude) */
	Vy float32
	/*Vz Ground Z Speed (Altitude) */
	Vz float32
	/*Covariance Row-major representation of a 6x6 position and velocity 6x6 cross-covariance matrix (states: lat, lon, alt, vx, vy, vz; first six entries are the first ROW, next six entries are the second row, etc.). If unknown, assign NaN value to first element in the array. */
	Covariance []float32
	/*EstimatorType Class id of the estimator this estimate originated from. */
	EstimatorType uint8
}

// Read sets the field values of the message from the raw message payload
func (m *GlobalPositionIntCov) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Vx)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vy)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vz)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Covariance)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.EstimatorType)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *GlobalPositionIntCov) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Vx)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vy)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vz)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Covariance)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.EstimatorType)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
