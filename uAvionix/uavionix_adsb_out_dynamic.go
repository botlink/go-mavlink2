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
	"text/tabwriter"

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*UavionixAdsbOutDynamic Dynamic data used to generate ADS-B out transponder data (send at 5Hz) */
type UavionixAdsbOutDynamic struct {
	/*Utctime UTC time in seconds since GPS epoch (Jan 6, 1980). If unknown set to UINT32_MAX */
	Utctime uint32
	/*GPSlat Latitude WGS84 (deg * 1E7). If unknown set to INT32_MAX */
	GPSlat int32
	/*GPSlon Longitude WGS84 (deg * 1E7). If unknown set to INT32_MAX */
	GPSlon int32
	/*GPSalt Altitude (WGS84). UP +ve. If unknown set to INT32_MAX */
	GPSalt int32
	/*Baroaltmsl Barometric pressure altitude (MSL) relative to a standard atmosphere of 1013.2 mBar and NOT bar corrected altitude (m * 1E-3). (up +ve). If unknown set to INT32_MAX */
	Baroaltmsl int32
	/*Accuracyhor Horizontal accuracy in mm (m * 1E-3). If unknown set to UINT32_MAX */
	Accuracyhor uint32
	/*Accuracyvert Vertical accuracy in cm. If unknown set to UINT16_MAX */
	Accuracyvert uint16
	/*Accuracyvel Velocity accuracy in mm/s (m * 1E-3). If unknown set to UINT16_MAX */
	Accuracyvel uint16
	/*Velvert GPS vertical speed in cm/s. If unknown set to INT16_MAX */
	Velvert int16
	/*Velns North-South velocity over ground in cm/s North +ve. If unknown set to INT16_MAX */
	Velns int16
	/*Velew East-West velocity over ground in cm/s East +ve. If unknown set to INT16_MAX */
	Velew int16
	/*State ADS-B transponder dynamic input state flags */
	State uint16
	/*Squawk Mode A code (typically 1200 [0x04B0] for VFR) */
	Squawk uint16
	/*GPSfix 0-1: no fix, 2: 2D fix, 3: 3D fix, 4: DGPS, 5: RTK */
	GPSfix uint8
	/*Numsats Number of satellites visible. If unknown set to UINT8_MAX */
	Numsats uint8
	/*Emergencystatus Emergency status */
	Emergencystatus uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *UavionixAdsbOutDynamic) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Utctime:\t%v [s]\n"
	format += "GPSlat:\t%v [degE7]\n"
	format += "GPSlon:\t%v [degE7]\n"
	format += "GPSalt:\t%v [mm]\n"
	format += "Baroaltmsl:\t%v [mbar]\n"
	format += "Accuracyhor:\t%v [mm]\n"
	format += "Accuracyvert:\t%v [cm]\n"
	format += "Accuracyvel:\t%v [mm/s]\n"
	format += "Velvert:\t%v [cm/s]\n"
	format += "Velns:\t%v [cm/s]\n"
	format += "Velew:\t%v [cm/s]\n"
	format += "State:\t%v \n"
	format += "Squawk:\t%v \n"
	format += "GPSfix:\t%v \n"
	format += "Numsats:\t%v \n"
	format += "Emergencystatus:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Utctime,
		m.GPSlat,
		m.GPSlon,
		m.GPSalt,
		m.Baroaltmsl,
		m.Accuracyhor,
		m.Accuracyvert,
		m.Accuracyvel,
		m.Velvert,
		m.Velns,
		m.Velew,
		m.State,
		m.Squawk,
		m.GPSfix,
		m.Numsats,
		m.Emergencystatus,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *UavionixAdsbOutDynamic) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *UavionixAdsbOutDynamic) GetDialect() string {
	return "uavionix"
}

// GetMessageName gets the name of the Message
func (m *UavionixAdsbOutDynamic) GetMessageName() string {
	return "UavionixAdsbOutDynamic"
}

// GetID gets the ID of the Message
func (m *UavionixAdsbOutDynamic) GetID() uint32 {
	return 10002
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *UavionixAdsbOutDynamic) HasExtensionFields() bool {
	return false
}

func (m *UavionixAdsbOutDynamic) getV1Length() int {
	return 41
}

func (m *UavionixAdsbOutDynamic) getIOSlice() []byte {
	return make([]byte, 41+1)
}

// Read sets the field values of the message from the raw message payload
func (m *UavionixAdsbOutDynamic) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the UavionixAdsbOutDynamic fields
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
func (m *UavionixAdsbOutDynamic) Write(version int) (output []byte, err error) {
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
