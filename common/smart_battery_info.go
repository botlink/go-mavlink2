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
package common

import (
	"bytes"
	"encoding/binary"
)

/*SmartBatteryInfo Smart Battery information (static/infrequent update). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries. */
type SmartBatteryInfo struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*CapacityFullSpecification Capacity when full according to manufacturer, -1: field not provided. */
	CapacityFullSpecification int32
	/*CapacityFull Capacity when full (accounting for battery degradation), -1: field not provided. */
	CapacityFull int32
	/*SerialNumber Serial number. -1: field not provided. */
	SerialNumber int32
	/*CycleCount Charge/discharge cycle count. -1: field not provided. */
	CycleCount uint16
	/*Weight Battery weight. 0: field not provided. */
	Weight uint16
	/*DischargeMinimumVoltage Minimum per-cell voltage when discharging. If not supplied set to UINT16_MAX value. */
	DischargeMinimumVoltage uint16
	/*ChargingMinimumVoltage Minimum per-cell voltage when charging. If not supplied set to UINT16_MAX value. */
	ChargingMinimumVoltage uint16
	/*RestingMinimumVoltage Minimum per-cell voltage when resting. If not supplied set to UINT16_MAX value. */
	RestingMinimumVoltage uint16
	/*ID Battery ID */
	ID uint8
	/*DeviceName Static device name. Encode as manufacturer and product names separated using an underscore. */
	DeviceName []byte
}

// Read sets the field values of the message from the raw message payload
func (m *SmartBatteryInfo) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.CapacityFullSpecification)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CapacityFull)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SerialNumber)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.CycleCount)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Weight)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.DischargeMinimumVoltage)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ChargingMinimumVoltage)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.RestingMinimumVoltage)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.ID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.DeviceName)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *SmartBatteryInfo) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.CapacityFullSpecification)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CapacityFull)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SerialNumber)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.CycleCount)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Weight)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.DischargeMinimumVoltage)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ChargingMinimumVoltage)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.RestingMinimumVoltage)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.DeviceName)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
