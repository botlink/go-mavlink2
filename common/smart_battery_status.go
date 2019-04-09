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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SmartBatteryStatus Smart Battery information (dynamic). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries. */
type SmartBatteryStatus struct {
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
	Voltages [16]uint16
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SmartBatteryStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "FaultBitmask:\t%v \n"
	format += "TimeRemaining:\t%v [s]\n"
	format += "ID:\t%v \n"
	format += "CapacityRemaining:\t%v \n"
	format += "Current:\t%v [cA]\n"
	format += "Temperature:\t%v [cdegC]\n"
	format += "CellOffset:\t%v \n"
	format += "Voltages:\t%v [mV]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.FaultBitmask,
		m.TimeRemaining,
		m.ID,
		m.CapacityRemaining,
		m.Current,
		m.Temperature,
		m.CellOffset,
		m.Voltages,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SmartBatteryStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SmartBatteryStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *SmartBatteryStatus) GetMessageName() string {
	return "SmartBatteryStatus"
}

// GetID gets the ID of the Message
func (m *SmartBatteryStatus) GetID() uint32 {
	return 371
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SmartBatteryStatus) HasExtensionFields() bool {
	return false
}

func (m *SmartBatteryStatus) getV1Length() int {
	return 50
}

func (m *SmartBatteryStatus) getIOSlice() []byte {
	return make([]byte, 50+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SmartBatteryStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the SmartBatteryStatus fields
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
func (m *SmartBatteryStatus) Write(version int) (output []byte, err error) {
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
