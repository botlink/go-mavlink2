package matrixpilot

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

/*SerialUdbExtraF4 Backwards compatible version of SERIAL_UDB_EXTRA F4: format */
type SerialUdbExtraF4 struct {
	/*SueRollStabilizationAilerons Serial UDB Extra Roll Stabilization with Ailerons Enabled */
	SueRollStabilizationAilerons uint8
	/*SueRollStabilizationRudder Serial UDB Extra Roll Stabilization with Rudder Enabled */
	SueRollStabilizationRudder uint8
	/*SuePitchStabilization Serial UDB Extra Pitch Stabilization Enabled */
	SuePitchStabilization uint8
	/*SueYawStabilizationRudder Serial UDB Extra Yaw Stabilization using Rudder Enabled */
	SueYawStabilizationRudder uint8
	/*SueYawStabilizationAileron Serial UDB Extra Yaw Stabilization using Ailerons Enabled */
	SueYawStabilizationAileron uint8
	/*SueAileronNAVigation Serial UDB Extra Navigation with Ailerons Enabled */
	SueAileronNAVigation uint8
	/*SueRudderNAVigation Serial UDB Extra Navigation with Rudder Enabled */
	SueRudderNAVigation uint8
	/*SueAltitudeholdStabilized Serial UDB Extra Type of Alitude Hold when in Stabilized Mode */
	SueAltitudeholdStabilized uint8
	/*SueAltitudeholdWaypoint Serial UDB Extra Type of Alitude Hold when in Waypoint Mode */
	SueAltitudeholdWaypoint uint8
	/*SueRacingMode Serial UDB Extra Firmware racing mode enabled */
	SueRacingMode uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF4) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "SueRollStabilizationAilerons:\t%v \n"
	format += "SueRollStabilizationRudder:\t%v \n"
	format += "SuePitchStabilization:\t%v \n"
	format += "SueYawStabilizationRudder:\t%v \n"
	format += "SueYawStabilizationAileron:\t%v \n"
	format += "SueAileronNAVigation:\t%v \n"
	format += "SueRudderNAVigation:\t%v \n"
	format += "SueAltitudeholdStabilized:\t%v \n"
	format += "SueAltitudeholdWaypoint:\t%v \n"
	format += "SueRacingMode:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueRollStabilizationAilerons,
		m.SueRollStabilizationRudder,
		m.SuePitchStabilization,
		m.SueYawStabilizationRudder,
		m.SueYawStabilizationAileron,
		m.SueAileronNAVigation,
		m.SueRudderNAVigation,
		m.SueAltitudeholdStabilized,
		m.SueAltitudeholdWaypoint,
		m.SueRacingMode,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF4) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF4) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF4) GetMessageName() string {
	return "SerialUdbExtraF4"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF4) GetID() uint32 {
	return 172
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF4) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF4) getV1Length() int {
	return 10
}

func (m *SerialUdbExtraF4) getIOSlice() []byte {
	return make([]byte, 10+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF4) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SerialUdbExtraF4 fields
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
func (m *SerialUdbExtraF4) Write(version int) (output []byte, err error) {
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
