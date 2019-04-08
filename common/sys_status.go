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

/*SysStatus The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout. */
type SysStatus struct {
	/*OnboardControlSensorsPresent Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present. */
	OnboardControlSensorsPresent uint32
	/*OnboardControlSensorsEnabled Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled. */
	OnboardControlSensorsEnabled uint32
	/*OnboardControlSensorsHealth Bitmap showing which onboard controllers and sensors are operational or have an error:  Value of 0: not enabled. Value of 1: enabled. */
	OnboardControlSensorsHealth uint32
	/*Load Maximum usage in percent of the mainloop time. Values: [0-1000] - should always be below 1000 */
	Load uint16
	/*VoltageBattery Battery voltage */
	VoltageBattery uint16
	/*CurrentBattery Battery current, -1: autopilot does not measure the current */
	CurrentBattery int16
	/*DropRateComm Communication drop rate, (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV) */
	DropRateComm uint16
	/*ErrorsComm Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV) */
	ErrorsComm uint16
	/*ErrorsCount1 Autopilot-specific errors */
	ErrorsCount1 uint16
	/*ErrorsCount2 Autopilot-specific errors */
	ErrorsCount2 uint16
	/*ErrorsCount3 Autopilot-specific errors */
	ErrorsCount3 uint16
	/*ErrorsCount4 Autopilot-specific errors */
	ErrorsCount4 uint16
	/*BatteryRemaining Remaining battery energy, -1: autopilot estimate the remaining battery */
	BatteryRemaining int8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SysStatus) String() string {
	var builder strings.Builder

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("OnboardControlSensorsPresent:\t%v \n")
	builder.WriteString("OnboardControlSensorsEnabled:\t%v \n")
	builder.WriteString("OnboardControlSensorsHealth:\t%v \n")
	builder.WriteString("Load:\t%v [d%]\n")
	builder.WriteString("VoltageBattery:\t%v [mV]\n")
	builder.WriteString("CurrentBattery:\t%v [cA]\n")
	builder.WriteString("DropRateComm:\t%v [c%]\n")
	builder.WriteString("ErrorsComm:\t%v \n")
	builder.WriteString("ErrorsCount1:\t%v \n")
	builder.WriteString("ErrorsCount2:\t%v \n")
	builder.WriteString("ErrorsCount3:\t%v \n")
	builder.WriteString("ErrorsCount4:\t%v \n")
	builder.WriteString("BatteryRemaining:\t%v [%]\n")
	format := builder.String()

	return fmt.Sprintf(
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.OnboardControlSensorsPresent,
		m.OnboardControlSensorsEnabled,
		m.OnboardControlSensorsHealth,
		m.Load,
		m.VoltageBattery,
		m.CurrentBattery,
		m.DropRateComm,
		m.ErrorsComm,
		m.ErrorsCount1,
		m.ErrorsCount2,
		m.ErrorsCount3,
		m.ErrorsCount4,
		m.BatteryRemaining,
	)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SysStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SysStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SysStatus) GetMessageName() string {
	return "SysStatus"
}

// GetID gets the ID of the Message
func (m *SysStatus) GetID() uint32 {
	return 1
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SysStatus) HasExtensionFields() bool {
	return false
}

func (m *SysStatus) getV1Length() int {
	return 31
}

func (m *SysStatus) getIOSlice() []byte {
	return make([]byte, 31+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SysStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SysStatus fields
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
func (m *SysStatus) Write(version int) (output []byte, err error) {
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
