package common

/*
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
)

/*SysStatus The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout. */
type SysStatus struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
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
}

// Read sets the field values of the message from the raw message payload
func (m *SysStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.OnboardControlSensorsPresent)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.OnboardControlSensorsEnabled)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.OnboardControlSensorsHealth)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Load)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.VoltageBattery)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CurrentBattery)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.DropRateComm)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ErrorsComm)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ErrorsCount1)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ErrorsCount2)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ErrorsCount3)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ErrorsCount4)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BatteryRemaining)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *SysStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.OnboardControlSensorsPresent)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.OnboardControlSensorsEnabled)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.OnboardControlSensorsHealth)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Load)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.VoltageBattery)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CurrentBattery)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.DropRateComm)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ErrorsComm)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ErrorsCount1)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ErrorsCount2)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ErrorsCount3)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ErrorsCount4)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BatteryRemaining)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
