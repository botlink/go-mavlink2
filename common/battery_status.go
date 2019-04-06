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
)

/*BatteryStatus Battery information. Updates GCS with flight controller battery status. Use SMART_BATTERY_* messages instead for smart batteries. */
type BatteryStatus struct {
	/*FrameVersion indicates the wire format of the frame this message was read from */
	FrameVersion int
	/*CurrentConsumed Consumed charge, -1: autopilot does not provide consumption estimate */
	CurrentConsumed int32
	/*EnergyConsumed Consumed energy, -1: autopilot does not provide energy consumption estimate */
	EnergyConsumed int32
	/*Temperature Temperature of the battery. INT16_MAX for unknown temperature. */
	Temperature int16
	/*Voltages Battery voltage of cells. Cells above the valid cell count for this battery should have the UINT16_MAX value. */
	Voltages []uint16
	/*CurrentBattery Battery current, -1: autopilot does not measure the current */
	CurrentBattery int16
	/*ID Battery ID */
	ID uint8
	/*BatteryFunction Function of the battery */
	BatteryFunction uint8
	/*Type Type (chemistry) of the battery */
	Type uint8
	/*BatteryRemaining Remaining battery energy. Values: [0-100], -1: autopilot does not estimate the remaining battery. */
	BatteryRemaining int8
	/*TimeRemaining Remaining battery time, 0: autopilot does not provide remaining battery time estimate */
	TimeRemaining int32
	/*ChargeState State for extent of discharge, provided by autopilot for warning or external reactions */
	ChargeState uint8
}

// GetVersion gets the MAVLink version of the Message contents
func (m *BatteryStatus) GetVersion() int {
	return m.FrameVersion
}

// GetDialect gets the name of the dialect that defines the Message
func (m *BatteryStatus) GetDialect() string {
	return "common"
}

// GetName gets the name of the Message
func (m *BatteryStatus) GetName() string {
	return "BatteryStatus"
}

// GetID gets the ID of the Message
func (m *BatteryStatus) GetID() uint32 {
	return 147
}

// Read sets the field values of the message from the raw message payload
func (m *BatteryStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.FrameVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.CurrentConsumed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.EnergyConsumed)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Voltages)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CurrentBattery)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BatteryFunction)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Type)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.BatteryRemaining)
	if err != nil {
		return
	}

	if version == 2 {
		err = binary.Read(reader, binary.LittleEndian, &m.TimeRemaining)
		if err != nil {
			return
		}

		err = binary.Read(reader, binary.LittleEndian, &m.ChargeState)
		if err != nil {
			return
		}

	}
	return
}

// Write encodes the field values of the message to a byte array
func (m *BatteryStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.CurrentConsumed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.EnergyConsumed)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Voltages)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CurrentBattery)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BatteryFunction)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Type)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.BatteryRemaining)
	if err != nil {
		return nil, err
	}

	if version == 2 {
		err = binary.Write(&buffer, binary.LittleEndian, m.BatteryRemaining)
		if err != nil {
			return nil, err
		}

		err = binary.Write(&buffer, binary.LittleEndian, m.BatteryRemaining)
		if err != nil {
			return nil, err
		}

	}

	return buffer.Bytes(), nil
}
