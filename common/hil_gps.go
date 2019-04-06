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

/*HilGps The global position, as returned by the Global Positioning System (GPS). This is
  NOT the global position estimate of the sytem, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate. */
type HilGps struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Lat Latitude (WGS84) */
	Lat int32
	/*Lon Longitude (WGS84) */
	Lon int32
	/*Alt Altitude (MSL). Positive for up. */
	Alt int32
	/*Eph GPS HDOP horizontal dilution of position. If unknown, set to: 65535 */
	Eph uint16
	/*Epv GPS VDOP vertical dilution of position. If unknown, set to: 65535 */
	Epv uint16
	/*Vel GPS ground speed. If unknown, set to: 65535 */
	Vel uint16
	/*Vn GPS velocity in NORTH direction in earth-fixed NED frame */
	Vn int16
	/*Ve GPS velocity in EAST direction in earth-fixed NED frame */
	Ve int16
	/*Vd GPS velocity in DOWN direction in earth-fixed NED frame */
	Vd int16
	/*Cog Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: 65535 */
	Cog uint16
	/*FixType 0-1: no fix, 2: 2D fix, 3: 3D fix. Some applications will not use the value of this field unless it is at least two, so always correctly fill in the fix. */
	FixType uint8
	/*SatellitesVisible Number of satellites visible. If unknown, set to 255 */
	SatellitesVisible uint8
}

// Read sets the field values of the message from the raw message payload
func (m *HilGps) Read(version int, payload []byte) (err error) {
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

	err = binary.Read(reader, binary.LittleEndian, &m.Eph)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Epv)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vel)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vn)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Ve)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Vd)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Cog)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.FixType)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SatellitesVisible)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *HilGps) Write(version int) ([]byte, error) {
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Eph)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Epv)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vel)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vn)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Ve)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Vd)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Cog)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.FixType)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
