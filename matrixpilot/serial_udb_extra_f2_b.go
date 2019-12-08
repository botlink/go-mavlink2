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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SerialUdbExtraF2B Backwards compatible version of SERIAL_UDB_EXTRA - F2: Part B */
type SerialUdbExtraF2B struct {
	/*SueTime Serial UDB Extra Time */
	SueTime uint32
	/*SueFlags Serial UDB Extra Status Flags */
	SueFlags uint32
	/*SueBaromPress SUE barometer pressure */
	SueBaromPress int32
	/*SueBaromAlt SUE barometer altitude */
	SueBaromAlt int32
	/*SuePwmInput1 Serial UDB Extra PWM Input Channel 1 */
	SuePwmInput1 int16
	/*SuePwmInput2 Serial UDB Extra PWM Input Channel 2 */
	SuePwmInput2 int16
	/*SuePwmInput3 Serial UDB Extra PWM Input Channel 3 */
	SuePwmInput3 int16
	/*SuePwmInput4 Serial UDB Extra PWM Input Channel 4 */
	SuePwmInput4 int16
	/*SuePwmInput5 Serial UDB Extra PWM Input Channel 5 */
	SuePwmInput5 int16
	/*SuePwmInput6 Serial UDB Extra PWM Input Channel 6 */
	SuePwmInput6 int16
	/*SuePwmInput7 Serial UDB Extra PWM Input Channel 7 */
	SuePwmInput7 int16
	/*SuePwmInput8 Serial UDB Extra PWM Input Channel 8 */
	SuePwmInput8 int16
	/*SuePwmInput9 Serial UDB Extra PWM Input Channel 9 */
	SuePwmInput9 int16
	/*SuePwmInput10 Serial UDB Extra PWM Input Channel 10 */
	SuePwmInput10 int16
	/*SuePwmInput11 Serial UDB Extra PWM Input Channel 11 */
	SuePwmInput11 int16
	/*SuePwmInput12 Serial UDB Extra PWM Input Channel 12 */
	SuePwmInput12 int16
	/*SuePwmOutput1 Serial UDB Extra PWM Output Channel 1 */
	SuePwmOutput1 int16
	/*SuePwmOutput2 Serial UDB Extra PWM Output Channel 2 */
	SuePwmOutput2 int16
	/*SuePwmOutput3 Serial UDB Extra PWM Output Channel 3 */
	SuePwmOutput3 int16
	/*SuePwmOutput4 Serial UDB Extra PWM Output Channel 4 */
	SuePwmOutput4 int16
	/*SuePwmOutput5 Serial UDB Extra PWM Output Channel 5 */
	SuePwmOutput5 int16
	/*SuePwmOutput6 Serial UDB Extra PWM Output Channel 6 */
	SuePwmOutput6 int16
	/*SuePwmOutput7 Serial UDB Extra PWM Output Channel 7 */
	SuePwmOutput7 int16
	/*SuePwmOutput8 Serial UDB Extra PWM Output Channel 8 */
	SuePwmOutput8 int16
	/*SuePwmOutput9 Serial UDB Extra PWM Output Channel 9 */
	SuePwmOutput9 int16
	/*SuePwmOutput10 Serial UDB Extra PWM Output Channel 10 */
	SuePwmOutput10 int16
	/*SuePwmOutput11 Serial UDB Extra PWM Output Channel 11 */
	SuePwmOutput11 int16
	/*SuePwmOutput12 Serial UDB Extra PWM Output Channel 12 */
	SuePwmOutput12 int16
	/*SueIMULocationX Serial UDB Extra IMU Location X */
	SueIMULocationX int16
	/*SueIMULocationY Serial UDB Extra IMU Location Y */
	SueIMULocationY int16
	/*SueIMULocationZ Serial UDB Extra IMU Location Z */
	SueIMULocationZ int16
	/*SueLocationErrorEarthX Serial UDB Location Error Earth X */
	SueLocationErrorEarthX int16
	/*SueLocationErrorEarthY Serial UDB Location Error Earth Y */
	SueLocationErrorEarthY int16
	/*SueLocationErrorEarthZ Serial UDB Location Error Earth Z */
	SueLocationErrorEarthZ int16
	/*SueOscFails Serial UDB Extra Oscillator Failure Count */
	SueOscFails int16
	/*SueIMUVelocityX Serial UDB Extra IMU Velocity X */
	SueIMUVelocityX int16
	/*SueIMUVelocityY Serial UDB Extra IMU Velocity Y */
	SueIMUVelocityY int16
	/*SueIMUVelocityZ Serial UDB Extra IMU Velocity Z */
	SueIMUVelocityZ int16
	/*SueWaypointGoalX Serial UDB Extra Current Waypoint Goal X */
	SueWaypointGoalX int16
	/*SueWaypointGoalY Serial UDB Extra Current Waypoint Goal Y */
	SueWaypointGoalY int16
	/*SueWaypointGoalZ Serial UDB Extra Current Waypoint Goal Z */
	SueWaypointGoalZ int16
	/*SueAeroX Aeroforce in UDB X Axis */
	SueAeroX int16
	/*SueAeroY Aeroforce in UDB Y Axis */
	SueAeroY int16
	/*SueAeroZ Aeroforce in UDB Z axis */
	SueAeroZ int16
	/*SueBaromTemp SUE barometer temperature */
	SueBaromTemp int16
	/*SueBatVolt SUE battery voltage */
	SueBatVolt int16
	/*SueBatAmp SUE battery current */
	SueBatAmp int16
	/*SueBatAmpHours SUE battery milli amp hours used */
	SueBatAmpHours int16
	/*SueDesiredHeight Sue autopilot desired height */
	SueDesiredHeight int16
	/*SueMemoryStackFree Serial UDB Extra Stack Memory Free */
	SueMemoryStackFree int16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF2B) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "SueTime:\t%v \n"
	format += "SueFlags:\t%v \n"
	format += "SueBaromPress:\t%v \n"
	format += "SueBaromAlt:\t%v \n"
	format += "SuePwmInput1:\t%v \n"
	format += "SuePwmInput2:\t%v \n"
	format += "SuePwmInput3:\t%v \n"
	format += "SuePwmInput4:\t%v \n"
	format += "SuePwmInput5:\t%v \n"
	format += "SuePwmInput6:\t%v \n"
	format += "SuePwmInput7:\t%v \n"
	format += "SuePwmInput8:\t%v \n"
	format += "SuePwmInput9:\t%v \n"
	format += "SuePwmInput10:\t%v \n"
	format += "SuePwmInput11:\t%v \n"
	format += "SuePwmInput12:\t%v \n"
	format += "SuePwmOutput1:\t%v \n"
	format += "SuePwmOutput2:\t%v \n"
	format += "SuePwmOutput3:\t%v \n"
	format += "SuePwmOutput4:\t%v \n"
	format += "SuePwmOutput5:\t%v \n"
	format += "SuePwmOutput6:\t%v \n"
	format += "SuePwmOutput7:\t%v \n"
	format += "SuePwmOutput8:\t%v \n"
	format += "SuePwmOutput9:\t%v \n"
	format += "SuePwmOutput10:\t%v \n"
	format += "SuePwmOutput11:\t%v \n"
	format += "SuePwmOutput12:\t%v \n"
	format += "SueIMULocationX:\t%v \n"
	format += "SueIMULocationY:\t%v \n"
	format += "SueIMULocationZ:\t%v \n"
	format += "SueLocationErrorEarthX:\t%v \n"
	format += "SueLocationErrorEarthY:\t%v \n"
	format += "SueLocationErrorEarthZ:\t%v \n"
	format += "SueOscFails:\t%v \n"
	format += "SueIMUVelocityX:\t%v \n"
	format += "SueIMUVelocityY:\t%v \n"
	format += "SueIMUVelocityZ:\t%v \n"
	format += "SueWaypointGoalX:\t%v \n"
	format += "SueWaypointGoalY:\t%v \n"
	format += "SueWaypointGoalZ:\t%v \n"
	format += "SueAeroX:\t%v \n"
	format += "SueAeroY:\t%v \n"
	format += "SueAeroZ:\t%v \n"
	format += "SueBaromTemp:\t%v \n"
	format += "SueBatVolt:\t%v \n"
	format += "SueBatAmp:\t%v \n"
	format += "SueBatAmpHours:\t%v \n"
	format += "SueDesiredHeight:\t%v \n"
	format += "SueMemoryStackFree:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueTime,
		m.SueFlags,
		m.SueBaromPress,
		m.SueBaromAlt,
		m.SuePwmInput1,
		m.SuePwmInput2,
		m.SuePwmInput3,
		m.SuePwmInput4,
		m.SuePwmInput5,
		m.SuePwmInput6,
		m.SuePwmInput7,
		m.SuePwmInput8,
		m.SuePwmInput9,
		m.SuePwmInput10,
		m.SuePwmInput11,
		m.SuePwmInput12,
		m.SuePwmOutput1,
		m.SuePwmOutput2,
		m.SuePwmOutput3,
		m.SuePwmOutput4,
		m.SuePwmOutput5,
		m.SuePwmOutput6,
		m.SuePwmOutput7,
		m.SuePwmOutput8,
		m.SuePwmOutput9,
		m.SuePwmOutput10,
		m.SuePwmOutput11,
		m.SuePwmOutput12,
		m.SueIMULocationX,
		m.SueIMULocationY,
		m.SueIMULocationZ,
		m.SueLocationErrorEarthX,
		m.SueLocationErrorEarthY,
		m.SueLocationErrorEarthZ,
		m.SueOscFails,
		m.SueIMUVelocityX,
		m.SueIMUVelocityY,
		m.SueIMUVelocityZ,
		m.SueWaypointGoalX,
		m.SueWaypointGoalY,
		m.SueWaypointGoalZ,
		m.SueAeroX,
		m.SueAeroY,
		m.SueAeroZ,
		m.SueBaromTemp,
		m.SueBatVolt,
		m.SueBatAmp,
		m.SueBatAmpHours,
		m.SueDesiredHeight,
		m.SueMemoryStackFree,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF2B) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF2B) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF2B) GetMessageName() string {
	return "SerialUdbExtraF2B"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF2B) GetID() uint32 {
	return 171
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF2B) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF2B) getV1Length() int {
	return 108
}

func (m *SerialUdbExtraF2B) getIOSlice() []byte {
	return make([]byte, 108+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF2B) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SerialUdbExtraF2B fields
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
func (m *SerialUdbExtraF2B) Write(version int) (output []byte, err error) {
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
