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
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*ParamMapRc Bind a RC channel to a parameter. The parameter should change according to the RC channel value. */
type ParamMapRc struct {
	/*ParamValue0 Initial parameter value */
	ParamValue0 float32
	/*Scale Scale, maps the RC range [-1, 1] to a parameter value */
	Scale float32
	/*ParamValueMin Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation) */
	ParamValueMin float32
	/*ParamValueMax Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation) */
	ParamValueMax float32
	/*ParamIndex Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index. */
	ParamIndex int16
	/*TargetSystem System ID */
	TargetSystem uint8
	/*TargetComponent Component ID */
	TargetComponent uint8
	/*ParamID Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string */
	ParamID [16]byte
	/*ParameterRcChannelIndex Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC. */
	ParameterRcChannelIndex uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *ParamMapRc) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("ParamValue0:\t%v \n")
	builder.WriteString("Scale:\t%v \n")
	builder.WriteString("ParamValueMin:\t%v \n")
	builder.WriteString("ParamValueMax:\t%v \n")
	builder.WriteString("ParamIndex:\t%v \n")
	builder.WriteString("TargetSystem:\t%v \n")
	builder.WriteString("TargetComponent:\t%v \n")
	builder.WriteString("ParamID:\t%v \n")
	builder.WriteString("ParameterRcChannelIndex:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.ParamValue0,
		m.Scale,
		m.ParamValueMin,
		m.ParamValueMax,
		m.ParamIndex,
		m.TargetSystem,
		m.TargetComponent,
		m.ParamID,
		m.ParameterRcChannelIndex,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetParamID encodes the input string to the ParamID array
func (m *ParamMapRc) SetParamID(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(16)))
	copy(m.ParamID[:], []byte(input)[:clen])

	if len(input) > 16 {
		err = mavlink2.ErrStringTooLong
	}

	return
}

// GetParamID decodes the null-terminated string in the ParamID
func (m *ParamMapRc) GetParamID() string {
	clen := util.CStrLen(m.ParamID[:])

	return string(m.ParamID[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *ParamMapRc) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *ParamMapRc) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *ParamMapRc) GetMessageName() string {
	return "ParamMapRc"
}

// GetID gets the ID of the Message
func (m *ParamMapRc) GetID() uint32 {
	return 50
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *ParamMapRc) HasExtensionFields() bool {
	return false
}

func (m *ParamMapRc) getV1Length() int {
	return 37
}

func (m *ParamMapRc) getIOSlice() []byte {
	return make([]byte, 37+1)
}

// Read sets the field values of the message from the raw message payload
func (m *ParamMapRc) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the ParamMapRc fields
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
func (m *ParamMapRc) Write(version int) (output []byte, err error) {
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
