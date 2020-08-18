package common

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2020 queue-b <https://github.com/queue-b>

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

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*AisVessel The location and information of an AIS vessel */
type AisVessel struct {
	/*Mmsi Mobile Marine Service Identifier, 9 decimal digits */
	Mmsi uint32
	/*Lat Latitude */
	Lat int32
	/*Lon Longitude */
	Lon int32
	/*Cog Course over ground */
	Cog uint16
	/*Heading True heading */
	Heading uint16
	/*Velocity Speed over ground */
	Velocity uint16
	/*DimensionBow Distance from lat/lon location to bow */
	DimensionBow uint16
	/*DimensionStern Distance from lat/lon location to stern */
	DimensionStern uint16
	/*Tslc Time since last communication in seconds */
	Tslc uint16
	/*Flags Bitmask to indicate various statuses including valid data fields */
	Flags uint16
	/*TurnRate Turn rate */
	TurnRate int8
	/*NAVigationalStatus Navigational status */
	NAVigationalStatus uint8
	/*Type Type of vessels */
	Type uint8
	/*DimensionPort Distance from lat/lon location to port side */
	DimensionPort uint8
	/*DimensionStarboard Distance from lat/lon location to starboard side */
	DimensionStarboard uint8
	/*Callsign The vessel callsign */
	Callsign [7]byte
	/*Name The vessel name */
	Name [20]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AisVessel) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Mmsi:\t%v \n"
	format += "Lat:\t%v [degE7]\n"
	format += "Lon:\t%v [degE7]\n"
	format += "Cog:\t%v [cdeg]\n"
	format += "Heading:\t%v [cdeg]\n"
	format += "Velocity:\t%v [cm/s]\n"
	format += "DimensionBow:\t%v [m]\n"
	format += "DimensionStern:\t%v [m]\n"
	format += "Tslc:\t%v [s]\n"
	format += "Flags:\t%v \n"
	format += "TurnRate:\t%v [cdeg/s]\n"
	format += "NAVigationalStatus:\t%v \n"
	format += "Type:\t%v \n"
	format += "DimensionPort:\t%v [m]\n"
	format += "DimensionStarboard:\t%v [m]\n"
	format += "Callsign:\t%v \n"
	format += "Name:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Mmsi,
		m.Lat,
		m.Lon,
		m.Cog,
		m.Heading,
		m.Velocity,
		m.DimensionBow,
		m.DimensionStern,
		m.Tslc,
		m.Flags,
		m.TurnRate,
		m.NAVigationalStatus,
		m.Type,
		m.DimensionPort,
		m.DimensionStarboard,
		m.Callsign,
		m.Name,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetCallsign encodes the input string to the Callsign array
func (m *AisVessel) SetCallsign(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(7)))
	copy(m.Callsign[:], []byte(input)[:clen])

	if len(input) > 7 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetCallsign decodes the null-terminated string in the Callsign
func (m *AisVessel) GetCallsign() string {
	clen := util.CStrLen(m.Callsign[:])

	return string(m.Callsign[:clen])
}

// SetName encodes the input string to the Name array
func (m *AisVessel) SetName(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(20)))
	copy(m.Name[:], []byte(input)[:clen])

	if len(input) > 20 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetName decodes the null-terminated string in the Name
func (m *AisVessel) GetName() string {
	clen := util.CStrLen(m.Name[:])

	return string(m.Name[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AisVessel) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AisVessel) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *AisVessel) GetMessageName() string {
	return "AisVessel"
}

// GetID gets the ID of the Message
func (m *AisVessel) GetID() uint32 {
	return 301
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AisVessel) HasExtensionFields() bool {
	return false
}

func (m *AisVessel) getV1Length() int {
	return 58
}

func (m *AisVessel) getIOSlice() []byte {
	return make([]byte, 58+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AisVessel) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the AisVessel fields
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
func (m *AisVessel) Write(version int) (output []byte, err error) {
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
