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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*GPSRawInt The global position, as returned by the Global Positioning System (GPS). This is
  NOT the global position estimate of the system, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate. */
type GPSRawInt struct {
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
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GPSRawInt) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "Lat:\t%v [degE7]\n"
	format += "Lon:\t%v [degE7]\n"
	format += "Alt:\t%v [mm]\n"
	format += "Eph:\t%v \n"
	format += "Epv:\t%v \n"
	format += "Vel:\t%v [cm/s]\n"
	format += "Cog:\t%v [cdeg]\n"
	format += "FixType:\t%v \n"
	format += "SatellitesVisible:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "AltEllipsoID:\t%v\n"
		format += "HAcc:\t%v\n"
		format += "VAcc:\t%v\n"
		format += "VelAcc:\t%v\n"
		format += "HdgAcc:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.TimeUsec,
			m.Lat,
			m.Lon,
			m.Alt,
			m.Eph,
			m.Epv,
			m.Vel,
			m.Cog,
			m.FixType,
			m.SatellitesVisible,
			m.AltEllipsoID,
			m.HAcc,
			m.VAcc,
			m.VelAcc,
			m.HdgAcc,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.Eph,
		m.Epv,
		m.Vel,
		m.Cog,
		m.FixType,
		m.SatellitesVisible,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GPSRawInt) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GPSRawInt) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GPSRawInt) GetMessageName() string {
	return "GPSRawInt"
}

// GetID gets the ID of the Message
func (m *GPSRawInt) GetID() uint32 {
	return 24
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GPSRawInt) HasExtensionFields() bool {
	return true
}

func (m *GPSRawInt) getV1Length() int {
	return 30
}

func (m *GPSRawInt) getIOSlice() []byte {
	return make([]byte, 50+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GPSRawInt) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GPSRawInt fields
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
func (m *GPSRawInt) Write(version int) (output []byte, err error) {
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
