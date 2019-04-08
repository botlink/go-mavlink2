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
	"strings"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*GpsRtk RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting */
type GpsRtk struct {
	/*TimeLastBaselineMs Time since boot of last baseline message received. */
	TimeLastBaselineMs uint32
	/*Tow GPS Time of Week of last baseline */
	Tow uint32
	/*BaselineAMm Current baseline in ECEF x or NED north component. */
	BaselineAMm int32
	/*BaselineBMm Current baseline in ECEF y or NED east component. */
	BaselineBMm int32
	/*BaselineCMm Current baseline in ECEF z or NED down component. */
	BaselineCMm int32
	/*Accuracy Current estimate of baseline accuracy. */
	Accuracy uint32
	/*IarNumHypotheses Current number of integer ambiguity hypotheses. */
	IarNumHypotheses int32
	/*Wn GPS Week Number of last baseline */
	Wn uint16
	/*RtkReceiverID Identification of connected RTK receiver. */
	RtkReceiverID uint8
	/*RtkHealth GPS-specific health report for RTK data. */
	RtkHealth uint8
	/*RtkRate Rate of baseline messages being received by GPS */
	RtkRate uint8
	/*Nsats Current number of sats used for RTK calculation. */
	Nsats uint8
	/*BaselineCoordsType Coordinate system of baseline */
	BaselineCoordsType uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GpsRtk) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeLastBaselineMs:\t%v [ms]\n")
	builder.WriteString("Tow:\t%v [ms]\n")
	builder.WriteString("BaselineAMm:\t%v [mm]\n")
	builder.WriteString("BaselineBMm:\t%v [mm]\n")
	builder.WriteString("BaselineCMm:\t%v [mm]\n")
	builder.WriteString("Accuracy:\t%v \n")
	builder.WriteString("IarNumHypotheses:\t%v \n")
	builder.WriteString("Wn:\t%v \n")
	builder.WriteString("RtkReceiverID:\t%v \n")
	builder.WriteString("RtkHealth:\t%v \n")
	builder.WriteString("RtkRate:\t%v [Hz]\n")
	builder.WriteString("Nsats:\t%v \n")
	builder.WriteString("BaselineCoordsType:\t%v \n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeLastBaselineMs,
		m.Tow,
		m.BaselineAMm,
		m.BaselineBMm,
		m.BaselineCMm,
		m.Accuracy,
		m.IarNumHypotheses,
		m.Wn,
		m.RtkReceiverID,
		m.RtkHealth,
		m.RtkRate,
		m.Nsats,
		m.BaselineCoordsType,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GpsRtk) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GpsRtk) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GpsRtk) GetMessageName() string {
	return "GpsRtk"
}

// GetID gets the ID of the Message
func (m *GpsRtk) GetID() uint32 {
	return 127
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GpsRtk) HasExtensionFields() bool {
	return false
}

func (m *GpsRtk) getV1Length() int {
	return 35
}

func (m *GpsRtk) getIOSlice() []byte {
	return make([]byte, 35+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GpsRtk) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GpsRtk fields
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
func (m *GpsRtk) Write(version int) (output []byte, err error) {
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
