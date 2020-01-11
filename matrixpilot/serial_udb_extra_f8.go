package matrixpilot

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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SerialUdbExtraF8 Backwards compatible version of SERIAL_UDB_EXTRA F8: format */
type SerialUdbExtraF8 struct {
	/*SueHeightTargetMax Serial UDB Extra HEIGHT_TARGET_MAX */
	SueHeightTargetMax float32
	/*SueHeightTargetMin Serial UDB Extra HEIGHT_TARGET_MIN */
	SueHeightTargetMin float32
	/*SueAltHoldThrottleMin Serial UDB Extra ALT_HOLD_THROTTLE_MIN */
	SueAltHoldThrottleMin float32
	/*SueAltHoldThrottleMax Serial UDB Extra ALT_HOLD_THROTTLE_MAX */
	SueAltHoldThrottleMax float32
	/*SueAltHoldPitchMin Serial UDB Extra ALT_HOLD_PITCH_MIN */
	SueAltHoldPitchMin float32
	/*SueAltHoldPitchMax Serial UDB Extra ALT_HOLD_PITCH_MAX */
	SueAltHoldPitchMax float32
	/*SueAltHoldPitchHigh Serial UDB Extra ALT_HOLD_PITCH_HIGH */
	SueAltHoldPitchHigh float32
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF8) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "SueHeightTargetMax:\t%v \n"
	format += "SueHeightTargetMin:\t%v \n"
	format += "SueAltHoldThrottleMin:\t%v \n"
	format += "SueAltHoldThrottleMax:\t%v \n"
	format += "SueAltHoldPitchMin:\t%v \n"
	format += "SueAltHoldPitchMax:\t%v \n"
	format += "SueAltHoldPitchHigh:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueHeightTargetMax,
		m.SueHeightTargetMin,
		m.SueAltHoldThrottleMin,
		m.SueAltHoldThrottleMax,
		m.SueAltHoldPitchMin,
		m.SueAltHoldPitchMax,
		m.SueAltHoldPitchHigh,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF8) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF8) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF8) GetMessageName() string {
	return "SerialUdbExtraF8"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF8) GetID() uint32 {
	return 176
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF8) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF8) getV1Length() int {
	return 28
}

func (m *SerialUdbExtraF8) getIOSlice() []byte {
	return make([]byte, 28+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF8) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SerialUdbExtraF8 fields
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
func (m *SerialUdbExtraF8) Write(version int) (output []byte, err error) {
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
