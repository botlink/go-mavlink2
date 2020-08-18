package asluav

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
	"text/tabwriter"

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*FwSoaringData Fixed-wing soaring (i.e. thermal seeking) data */
type FwSoaringData struct {
	/*Timestamp Timestamp */
	Timestamp uint64
	/*Timestampmodechanged Timestamp since last mode change */
	Timestampmodechanged uint64
	/*Xw Thermal core updraft strength */
	Xw float32
	/*Xr Thermal radius */
	Xr float32
	/*Xlat Thermal center latitude */
	Xlat float32
	/*Xlon Thermal center longitude */
	Xlon float32
	/*Varw Variance W */
	Varw float32
	/*Varr Variance R */
	Varr float32
	/*Varlat Variance Lat */
	Varlat float32
	/*Varlon Variance Lon  */
	Varlon float32
	/*Loiterradius Suggested loiter radius */
	Loiterradius float32
	/*Loiterdirection Suggested loiter direction */
	Loiterdirection float32
	/*Disttosoarpoint Distance to soar point */
	Disttosoarpoint float32
	/*Vsinkexp Expected sink rate at current airspeed, roll and throttle */
	Vsinkexp float32
	/*Z1Localupdraftspeed Measurement / updraft speed at current/local airplane position */
	Z1Localupdraftspeed float32
	/*Z2Deltaroll Measurement / roll angle tracking error */
	Z2Deltaroll float32
	/*Z1Exp Expected measurement 1 */
	Z1Exp float32
	/*Z2Exp Expected measurement 2 */
	Z2Exp float32
	/*Thermalgsnorth Thermal drift (from estimator prediction step only) */
	Thermalgsnorth float32
	/*Thermalgseast Thermal drift (from estimator prediction step only) */
	Thermalgseast float32
	/*TseDot  Total specific energy change (filtered) */
	TseDot float32
	/*Debugvar1  Debug variable 1 */
	Debugvar1 float32
	/*Debugvar2  Debug variable 2 */
	Debugvar2 float32
	/*Controlmode Control Mode [-] */
	Controlmode uint8
	/*ValID Data valid [-] */
	ValID uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *FwSoaringData) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [ms]\n"
	format += "Timestampmodechanged:\t%v [ms]\n"
	format += "Xw:\t%v [m/s]\n"
	format += "Xr:\t%v [m]\n"
	format += "Xlat:\t%v [deg]\n"
	format += "Xlon:\t%v [deg]\n"
	format += "Varw:\t%v \n"
	format += "Varr:\t%v \n"
	format += "Varlat:\t%v \n"
	format += "Varlon:\t%v \n"
	format += "Loiterradius:\t%v [m]\n"
	format += "Loiterdirection:\t%v \n"
	format += "Disttosoarpoint:\t%v [m]\n"
	format += "Vsinkexp:\t%v [m/s]\n"
	format += "Z1Localupdraftspeed:\t%v [m/s]\n"
	format += "Z2Deltaroll:\t%v [deg]\n"
	format += "Z1Exp:\t%v \n"
	format += "Z2Exp:\t%v \n"
	format += "Thermalgsnorth:\t%v [m/s]\n"
	format += "Thermalgseast:\t%v [m/s]\n"
	format += "TseDot:\t%v [m/s]\n"
	format += "Debugvar1:\t%v \n"
	format += "Debugvar2:\t%v \n"
	format += "Controlmode:\t%v \n"
	format += "ValID:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.Timestampmodechanged,
		m.Xw,
		m.Xr,
		m.Xlat,
		m.Xlon,
		m.Varw,
		m.Varr,
		m.Varlat,
		m.Varlon,
		m.Loiterradius,
		m.Loiterdirection,
		m.Disttosoarpoint,
		m.Vsinkexp,
		m.Z1Localupdraftspeed,
		m.Z2Deltaroll,
		m.Z1Exp,
		m.Z2Exp,
		m.Thermalgsnorth,
		m.Thermalgseast,
		m.TseDot,
		m.Debugvar1,
		m.Debugvar2,
		m.Controlmode,
		m.ValID,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *FwSoaringData) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *FwSoaringData) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *FwSoaringData) GetMessageName() string {
	return "FwSoaringData"
}

// GetID gets the ID of the Message
func (m *FwSoaringData) GetID() uint32 {
	return 210
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *FwSoaringData) HasExtensionFields() bool {
	return false
}

func (m *FwSoaringData) getV1Length() int {
	return 102
}

func (m *FwSoaringData) getIOSlice() []byte {
	return make([]byte, 102+1)
}

// Read sets the field values of the message from the raw message payload
func (m *FwSoaringData) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the FwSoaringData fields
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
func (m *FwSoaringData) Write(version int) (output []byte, err error) {
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
