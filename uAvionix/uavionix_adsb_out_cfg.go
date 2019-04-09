package uavionix

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
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*UavionixAdsbOutCfg Static data to configure the ADS-B transponder (send within 10 sec of a POR and every 10 sec thereafter) */
type UavionixAdsbOutCfg struct {
	/*Icao Vehicle address (24 bit) */
	Icao uint32
	/*Stallspeed Aircraft stall speed in cm/s */
	Stallspeed uint16
	/*Callsign Vehicle identifier (8 characters, null terminated, valid characters are A-Z, 0-9, " " only) */
	Callsign [9]byte
	/*Emittertype Transmitting vehicle type. See ADSB_EMITTER_TYPE enum */
	Emittertype uint8
	/*AiRCraftsize Aircraft length and width encoding (table 2-35 of DO-282B) */
	AiRCraftsize uint8
	/*GPSoffsetlat GPS antenna lateral offset (table 2-36 of DO-282B) */
	GPSoffsetlat uint8
	/*GPSoffsetlon GPS antenna longitudinal offset from nose [if non-zero, take position (in meters) divide by 2 and add one] (table 2-37 DO-282B) */
	GPSoffsetlon uint8
	/*Rfselect ADS-B transponder reciever and transmit enable flags */
	Rfselect uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *UavionixAdsbOutCfg) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("Icao:\t%v \n")
	builder.WriteString("Stallspeed:\t%v [cm/s]\n")
	builder.WriteString("Callsign:\t%v \n")
	builder.WriteString("Emittertype:\t%v \n")
	builder.WriteString("AiRCraftsize:\t%v \n")
	builder.WriteString("GPSoffsetlat:\t%v \n")
	builder.WriteString("GPSoffsetlon:\t%v \n")
	builder.WriteString("Rfselect:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Icao,
		m.Stallspeed,
		m.Callsign,
		m.Emittertype,
		m.AiRCraftsize,
		m.GPSoffsetlat,
		m.GPSoffsetlon,
		m.Rfselect,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetCallsign encodes the input string to the Callsign array
func (m *UavionixAdsbOutCfg) SetCallsign(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(9)))
	copy(m.Callsign[:], []byte(input)[:clen])

	if len(input) > 9 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetCallsign decodes the null-terminated string in the Callsign
func (m *UavionixAdsbOutCfg) GetCallsign() string {
	clen := util.CStrLen(m.Callsign[:])

	return string(m.Callsign[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *UavionixAdsbOutCfg) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *UavionixAdsbOutCfg) GetDialect() string {
	return "uavionix"
}

// GetMessageName gets the name of the Message
func (m *UavionixAdsbOutCfg) GetMessageName() string {
	return "UavionixAdsbOutCfg"
}

// GetID gets the ID of the Message
func (m *UavionixAdsbOutCfg) GetID() uint32 {
	return 10001
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *UavionixAdsbOutCfg) HasExtensionFields() bool {
	return false
}

func (m *UavionixAdsbOutCfg) getV1Length() int {
	return 20
}

func (m *UavionixAdsbOutCfg) getIOSlice() []byte {
	return make([]byte, 20+1)
}

// Read sets the field values of the message from the raw message payload
func (m *UavionixAdsbOutCfg) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the UavionixAdsbOutCfg fields
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
func (m *UavionixAdsbOutCfg) Write(version int) (output []byte, err error) {
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

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	if version == 1 {
		output = buffer.Bytes()[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}
