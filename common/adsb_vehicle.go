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

/*AdsbVehicle The location and information of an ADSB vehicle */
type AdsbVehicle struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*IcaoAddress ICAO address */
	IcaoAddress uint32
	/*Lat Latitude */
	Lat int32
	/*Lon Longitude */
	Lon int32
	/*Altitude Altitude(ASL) */
	Altitude int32
	/*Heading Course over ground */
	Heading uint16
	/*HorVelocity The horizontal velocity */
	HorVelocity uint16
	/*VerVelocity The vertical velocity. Positive is up */
	VerVelocity int16
	/*Flags Bitmap to indicate various statuses including valid data fields */
	Flags uint16
	/*Squawk Squawk code */
	Squawk uint16
	/*AltitudeType ADSB altitude type. */
	AltitudeType uint8
	/*Callsign The callsign, 8+null */
	Callsign []byte
	/*EmitterType ADSB emitter type. */
	EmitterType uint8
	/*Tslc Time since last communication in seconds */
	Tslc uint8
}

// Read sets the field values of the message from the raw message payload
func (m *AdsbVehicle) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.IcaoAddress)
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

	err = binary.Read(reader, binary.LittleEndian, &m.Altitude)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Heading)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HorVelocity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VerVelocity)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Flags)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Squawk)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.AltitudeType)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Callsign)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.EmitterType)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Tslc)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *AdsbVehicle) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.IcaoAddress)
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

	err = binary.Write(&buffer, binary.LittleEndian, m.Altitude)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Heading)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HorVelocity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VerVelocity)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Flags)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Squawk)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.AltitudeType)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Callsign)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.EmitterType)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Tslc)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
