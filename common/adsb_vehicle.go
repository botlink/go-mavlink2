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
	"fmt"
	"math"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*AdsbVehicle The location and information of an ADSB vehicle */
type AdsbVehicle struct {
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
	Callsign [9]byte
	/*EmitterType ADSB emitter type. */
	EmitterType uint8
	/*Tslc Time since last communication in seconds */
	Tslc uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AdsbVehicle) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "IcaoAddress:\t%v \n"
	format += "Lat:\t%v [degE7]\n"
	format += "Lon:\t%v [degE7]\n"
	format += "Altitude:\t%v [mm]\n"
	format += "Heading:\t%v [cdeg]\n"
	format += "HorVelocity:\t%v [cm/s]\n"
	format += "VerVelocity:\t%v [cm/s]\n"
	format += "Flags:\t%v \n"
	format += "Squawk:\t%v \n"
	format += "AltitudeType:\t%v \n"
	format += "Callsign:\t%v \n"
	format += "EmitterType:\t%v \n"
	format += "Tslc:\t%v [s]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.IcaoAddress,
		m.Lat,
		m.Lon,
		m.Altitude,
		m.Heading,
		m.HorVelocity,
		m.VerVelocity,
		m.Flags,
		m.Squawk,
		m.AltitudeType,
		m.Callsign,
		m.EmitterType,
		m.Tslc,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetCallsign encodes the input string to the Callsign array
func (m *AdsbVehicle) SetCallsign(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(9)))
	copy(m.Callsign[:], []byte(input)[:clen])

	if len(input) > 9 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetCallsign decodes the null-terminated string in the Callsign
func (m *AdsbVehicle) GetCallsign() string {
	clen := util.CStrLen(m.Callsign[:])

	return string(m.Callsign[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AdsbVehicle) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AdsbVehicle) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *AdsbVehicle) GetMessageName() string {
	return "AdsbVehicle"
}

// GetID gets the ID of the Message
func (m *AdsbVehicle) GetID() uint32 {
	return 246
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AdsbVehicle) HasExtensionFields() bool {
	return false
}

func (m *AdsbVehicle) getV1Length() int {
	return 38
}

func (m *AdsbVehicle) getIOSlice() []byte {
	return make([]byte, 38+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AdsbVehicle) Read(frame mavlink2.Frame) (err error) {
	version := frame.GetVersion()

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Read V2 messages from V1 frames
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrDecodeV2MessageV1Frame
		return
	}

	// binary.Read can panic; swallow the panic and return a sane error
	defer func() {
		if r := recover(); r != nil {
			err = mavlink2.ErrPrivateField
		}
	}()

	// Get a slice of bytes long enough for the all the AdsbVehicle fields
	// binary.Read requires enough bytes in the reader to read all fields, even if
	// the fields are just zero values. This also simplifies handling MAVLink2
	// extensions and trailing zero truncation.
	ioSlice := m.getIOSlice()

	copy(ioSlice, frame.GetMessageBytes())

	// Indicate if
	if version == 2 && m.HasExtensionFields() {
		ioSlice[len(ioSlice)-1] = 1
	}

	reader := bytes.NewReader(ioSlice)

	err = binary.Read(reader, binary.LittleEndian, m)

	return
}

// Write encodes the field values of the message to a byte array
func (m *AdsbVehicle) Write(version int) (output []byte, err error) {
	var buffer bytes.Buffer

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Write V2 messages to V1 bodies
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrEncodeV2MessageV1Frame
		return
	}

	err = binary.Write(&buffer, binary.LittleEndian, *m)
	if err != nil {
		return
	}

	output = buffer.Bytes()

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	// This also removes the trailing extra byte written for HasExtensionFieldValues
	if version == 1 {
		output = output[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		// Set HasExtensionFieldValues to zero so that it doesn't interfere with V2 truncation
		output[len(output)-1] = 0
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}
