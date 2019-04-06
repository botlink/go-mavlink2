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

/*SmartBatteryStatus Smart Battery information (dynamic). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries. */
type SmartBatteryStatus struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*FaultBitmask Fault/health indications. */
	FaultBitmask int32
	/*TimeRemaining Estimated remaining battery time. -1: field not provided. */
	TimeRemaining int32
	/*ID Battery ID */
	ID uint16
	/*CapacityRemaining Remaining battery energy. Values: [0-100], -1: field not provided. */
	CapacityRemaining int16
	/*Current Battery current (through all cells/loads). Positive if discharging, negative if charging. UINT16_MAX: field not provided. */
	Current int16
	/*Temperature Battery temperature. -1: field not provided. */
	Temperature int16
	/*CellOffset The cell number of the first index in the 'voltages' array field. Using this field allows you to specify cell voltages for batteries with more than 16 cells. */
	CellOffset uint16
	/*Voltages Individual cell voltages. Batteries with more 16 cells can use the cell_offset field to specify the cell offset for the array specified in the current message . Index values above the valid cell count for this battery should have the UINT16_MAX value. */
	Voltages []uint16
}

// Read sets the field values of the message from the raw message payload
func (m *SmartBatteryStatus) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.FaultBitmask)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.TimeRemaining)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CapacityRemaining)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Current)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Temperature)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CellOffset)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Voltages)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *SmartBatteryStatus) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.FaultBitmask)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.TimeRemaining)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CapacityRemaining)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Current)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Temperature)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CellOffset)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Voltages)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
