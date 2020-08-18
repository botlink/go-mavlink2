package common

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
	"math"
	"text/tabwriter"

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*SmartBatteryInfo Smart Battery information (static/infrequent update). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries. */
type SmartBatteryInfo struct {
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
	/*DischargeMinIMUmVoltage Minimum per-cell voltage when discharging. If not supplied set to UINT16_MAX value. */
	DischargeMinIMUmVoltage uint16
	/*ChargingMinIMUmVoltage Minimum per-cell voltage when charging. If not supplied set to UINT16_MAX value. */
	ChargingMinIMUmVoltage uint16
	/*RestingMinIMUmVoltage Minimum per-cell voltage when resting. If not supplied set to UINT16_MAX value. */
	RestingMinIMUmVoltage uint16
	/*ID Battery ID */
	ID uint8
	/*DeviceName Static device name. Encode as manufacturer and product names separated using an underscore. */
	DeviceName [50]byte
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SmartBatteryInfo) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "CapacityFullSpecification:\t%v [mAh]\n"
	format += "CapacityFull:\t%v [mAh]\n"
	format += "SerialNumber:\t%v \n"
	format += "CycleCount:\t%v \n"
	format += "Weight:\t%v [g]\n"
	format += "DischargeMinIMUmVoltage:\t%v [mV]\n"
	format += "ChargingMinIMUmVoltage:\t%v [mV]\n"
	format += "RestingMinIMUmVoltage:\t%v [mV]\n"
	format += "ID:\t%v \n"
	format += "DeviceName:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.CapacityFullSpecification,
		m.CapacityFull,
		m.SerialNumber,
		m.CycleCount,
		m.Weight,
		m.DischargeMinIMUmVoltage,
		m.ChargingMinIMUmVoltage,
		m.RestingMinIMUmVoltage,
		m.ID,
		m.DeviceName,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// SetDeviceName encodes the input string to the DeviceName array
func (m *SmartBatteryInfo) SetDeviceName(input string) (err error) {
	clen := int(math.Min(float64(len(input)), float64(50)))
	copy(m.DeviceName[:], []byte(input)[:clen])

	if len(input) > 50 {
		err = mavlink2.ErrValueTooLong
	}

	return
}

// GetDeviceName decodes the null-terminated string in the DeviceName
func (m *SmartBatteryInfo) GetDeviceName() string {
	clen := util.CStrLen(m.DeviceName[:])

	return string(m.DeviceName[:clen])
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SmartBatteryInfo) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SmartBatteryInfo) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SmartBatteryInfo) GetMessageName() string {
	return "SmartBatteryInfo"
}

// GetID gets the ID of the Message
func (m *SmartBatteryInfo) GetID() uint32 {
	return 370
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SmartBatteryInfo) HasExtensionFields() bool {
	return false
}

func (m *SmartBatteryInfo) getV1Length() int {
	return 73
}

func (m *SmartBatteryInfo) getIOSlice() []byte {
	return make([]byte, 73+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SmartBatteryInfo) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SmartBatteryInfo fields
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
func (m *SmartBatteryInfo) Write(version int) (output []byte, err error) {
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
