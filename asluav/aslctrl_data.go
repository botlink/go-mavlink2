package asluav

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

/*AslctrlData ASL-fixed-wing controller data */
type AslctrlData struct {
	/*Timestamp  Timestamp */
	Timestamp uint64
	/*H  See sourcecode for a description of these values...  */
	H float32
	/*Href   */
	Href float32
	/*HrefT   */
	HrefT float32
	/*Pitchangle Pitch angle */
	Pitchangle float32
	/*Pitchangleref Pitch angle reference */
	Pitchangleref float32
	/*Q   */
	Q float32
	/*Qref   */
	Qref float32
	/*Uelev   */
	Uelev float32
	/*Uthrot   */
	Uthrot float32
	/*Uthrot2   */
	Uthrot2 float32
	/*Nz   */
	Nz float32
	/*Airspeedref Airspeed reference */
	Airspeedref float32
	/*Yawangle Yaw angle */
	Yawangle float32
	/*Yawangleref Yaw angle reference */
	Yawangleref float32
	/*Rollangle Roll angle */
	Rollangle float32
	/*Rollangleref Roll angle reference */
	Rollangleref float32
	/*P   */
	P float32
	/*Pref   */
	Pref float32
	/*R   */
	R float32
	/*Rref   */
	Rref float32
	/*Uail   */
	Uail float32
	/*Urud   */
	Urud float32
	/*AslctrlMode  ASLCTRL control-mode (manual, stabilized, auto, etc...) */
	AslctrlMode uint8
	/*Spoilersengaged   */
	Spoilersengaged uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *AslctrlData) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Timestamp:\t%v [us]\n"
	format += "H:\t%v \n"
	format += "Href:\t%v \n"
	format += "HrefT:\t%v \n"
	format += "Pitchangle:\t%v [deg]\n"
	format += "Pitchangleref:\t%v [deg]\n"
	format += "Q:\t%v \n"
	format += "Qref:\t%v \n"
	format += "Uelev:\t%v \n"
	format += "Uthrot:\t%v \n"
	format += "Uthrot2:\t%v \n"
	format += "Nz:\t%v \n"
	format += "Airspeedref:\t%v [m/s]\n"
	format += "Yawangle:\t%v [deg]\n"
	format += "Yawangleref:\t%v [deg]\n"
	format += "Rollangle:\t%v [deg]\n"
	format += "Rollangleref:\t%v [deg]\n"
	format += "P:\t%v \n"
	format += "Pref:\t%v \n"
	format += "R:\t%v \n"
	format += "Rref:\t%v \n"
	format += "Uail:\t%v \n"
	format += "Urud:\t%v \n"
	format += "AslctrlMode:\t%v \n"
	format += "Spoilersengaged:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Timestamp,
		m.H,
		m.Href,
		m.HrefT,
		m.Pitchangle,
		m.Pitchangleref,
		m.Q,
		m.Qref,
		m.Uelev,
		m.Uthrot,
		m.Uthrot2,
		m.Nz,
		m.Airspeedref,
		m.Yawangle,
		m.Yawangleref,
		m.Rollangle,
		m.Rollangleref,
		m.P,
		m.Pref,
		m.R,
		m.Rref,
		m.Uail,
		m.Urud,
		m.AslctrlMode,
		m.Spoilersengaged,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *AslctrlData) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *AslctrlData) GetDialect() string {
	return "asluav"
}

// GetMessageName gets the name of the Message
func (m *AslctrlData) GetMessageName() string {
	return "AslctrlData"
}

// GetID gets the ID of the Message
func (m *AslctrlData) GetID() uint32 {
	return 203
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *AslctrlData) HasExtensionFields() bool {
	return false
}

func (m *AslctrlData) getV1Length() int {
	return 98
}

func (m *AslctrlData) getIOSlice() []byte {
	return make([]byte, 98+1)
}

// Read sets the field values of the message from the raw message payload
func (m *AslctrlData) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the AslctrlData fields
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
func (m *AslctrlData) Write(version int) (output []byte, err error) {
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
