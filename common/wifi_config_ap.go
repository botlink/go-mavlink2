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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*WifiConfigAp Configure AP SSID and Password. */
type WifiConfigAp struct {
	/*SsID Name of Wi-Fi network (SSID). Leave it blank to leave it unchanged. */
	SsID [32]byte
	/*Password Password. Leave it blank for an open AP. */
	Password [64]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *WifiConfigAp) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "SsID:\t%v \n"
	format += "Password:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SsID,
		m.Password,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetSsID encodes the input string to the SsID array
func (m *WifiConfigAp) SetSsID(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(32)))
	copy(m.SsID[:], []byte(input)[:clen])

	if len(input) > 32 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetSsID decodes the null-terminated string in the SsID
func (m *WifiConfigAp) GetSsID() string {
	clen := util.CStrLen(m.SsID[:])

	return string(m.SsID[:clen])
}

// SetPassword encodes the input string to the Password array
func (m *WifiConfigAp) SetPassword(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(64)))
	copy(m.Password[:], []byte(input)[:clen])

	if len(input) > 64 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetPassword decodes the null-terminated string in the Password
func (m *WifiConfigAp) GetPassword() string {
	clen := util.CStrLen(m.Password[:])

	return string(m.Password[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *WifiConfigAp) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *WifiConfigAp) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *WifiConfigAp) GetMessageName() string {
	return "WifiConfigAp"
}

// GetID gets the ID of the Message
func (m *WifiConfigAp) GetID() uint32 {
	return 299
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *WifiConfigAp) HasExtensionFields() bool {
	return false
}

func (m *WifiConfigAp) getV1Length() int {
	return 96
}

func (m *WifiConfigAp) getIOSlice() []byte {
	return make([]byte, 96+1)
}

// Read sets the field values of the message from the raw message payload
func (m *WifiConfigAp) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the WifiConfigAp fields
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
func (m *WifiConfigAp) Write(version int) (output []byte, err error) {
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
