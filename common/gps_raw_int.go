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

/*GpsRawInt The global position, as returned by the Global Positioning System (GPS). This is
  NOT the global position estimate of the system, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate. */
type GpsRawInt struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Lat Latitude (WGS84, EGM96 ellipsoid) */
	Lat int32
	/*Lon Longitude (WGS84, EGM96 ellipsoid) */
	Lon int32
	/*Alt Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude. */
	Alt int32
	/*Eph GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX */
	Eph uint16
	/*Epv GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX */
	Epv uint16
	/*Vel GPS ground speed. If unknown, set to: UINT16_MAX */
	Vel uint16
	/*Cog Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX */
	Cog uint16
	/*FixType GPS fix type. */
	FixType uint8
	/*SatellitesVisible Number of satellites visible. If unknown, set to 255 */
	SatellitesVisible uint8
	/*AltEllipsoID Altitude (above WGS84, EGM96 ellipsoid). Positive for up. */
	AltEllipsoID int32
	/*HAcc Position uncertainty. Positive for up. */
	HAcc uint32
	/*VAcc Altitude uncertainty. Positive for up. */
	VAcc uint32
	/*VelAcc Speed uncertainty. Positive for up. */
	VelAcc uint32
	/*HdgAcc Heading / track uncertainty */
	HdgAcc uint32
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GpsRawInt) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GpsRawInt) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *GpsRawInt) GetName() string {
	return "GpsRawInt"
}

// GetID gets the ID of the Message
func (m *GpsRawInt) GetID() uint32 {
	return 24
}

// Read sets the field values of the message from the raw message payload
func (m *GpsRawInt) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
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

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.AltEllipsoID)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.HAcc)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.VAcc)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.VelAcc)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.HdgAcc)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *GpsRawInt) Write(version int) ([]byte, error) {
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

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.SatellitesVisible)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
