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
	"text/tabwriter"

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/util"
)

/*OnboardComputerStatus Hardware status sent by an onboard computer. */
type OnboardComputerStatus struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Uptime Time since system boot. */
	Uptime uint32
	/*RamUsage Amount of used RAM on the component system. A value of UINT32_MAX implies the field is unused. */
	RamUsage uint32
	/*RamTotal Total amount of RAM on the component system. A value of UINT32_MAX implies the field is unused. */
	RamTotal uint32
	/*StorageType Storage type: 0: HDD, 1: SSD, 2: EMMC, 3: SD card (non-removable), 4: SD card (removable). A value of UINT32_MAX implies the field is unused. */
	StorageType [4]uint32
	/*StorageUsage Amount of used storage space on the component system. A value of UINT32_MAX implies the field is unused. */
	StorageUsage [4]uint32
	/*StorageTotal Total amount of storage space on the component system. A value of UINT32_MAX implies the field is unused. */
	StorageTotal [4]uint32
	/*LinkType Link type: 0-9: UART, 10-19: Wired network, 20-29: Wifi, 30-39: Point-to-point proprietary, 40-49: Mesh proprietary */
	LinkType [6]uint32
	/*LinkTxRate Network traffic from the component system. A value of UINT32_MAX implies the field is unused. */
	LinkTxRate [6]uint32
	/*LinkRxRate Network traffic to the component system. A value of UINT32_MAX implies the field is unused. */
	LinkRxRate [6]uint32
	/*LinkTxMax Network capacity from the component system. A value of UINT32_MAX implies the field is unused. */
	LinkTxMax [6]uint32
	/*LinkRxMax Network capacity to the component system. A value of UINT32_MAX implies the field is unused. */
	LinkRxMax [6]uint32
	/*FanSpeed Fan speeds. A value of INT16_MAX implies the field is unused. */
	FanSpeed [4]int16
	/*Type Type of the onboard computer: 0: Mission computer primary, 1: Mission computer backup 1, 2: Mission computer backup 2, 3: Compute node, 4-5: Compute spares, 6-9: Payload computers. */
	Type uint8
	/*CpuCores CPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused. */
	CpuCores [8]uint8
	/*CpuCombined Combined CPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused. */
	CpuCombined [10]uint8
	/*GpuCores GPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused. */
	GpuCores [4]uint8
	/*GpuCombined Combined GPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused. */
	GpuCombined [10]uint8
	/*TemperatureBoard Temperature of the board. A value of INT8_MAX implies the field is unused. */
	TemperatureBoard int8
	/*TemperatureCore Temperature of the CPU core. A value of INT8_MAX implies the field is unused. */
	TemperatureCore [8]int8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *OnboardComputerStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "Uptime:\t%v [ms]\n"
	format += "RamUsage:\t%v [MiB]\n"
	format += "RamTotal:\t%v [MiB]\n"
	format += "StorageType:\t%v \n"
	format += "StorageUsage:\t%v [MiB]\n"
	format += "StorageTotal:\t%v [MiB]\n"
	format += "LinkType:\t%v \n"
	format += "LinkTxRate:\t%v [KiB/s]\n"
	format += "LinkRxRate:\t%v [KiB/s]\n"
	format += "LinkTxMax:\t%v [KiB/s]\n"
	format += "LinkRxMax:\t%v [KiB/s]\n"
	format += "FanSpeed:\t%v [rpm]\n"
	format += "Type:\t%v \n"
	format += "CpuCores:\t%v \n"
	format += "CpuCombined:\t%v \n"
	format += "GpuCores:\t%v \n"
	format += "GpuCombined:\t%v \n"
	format += "TemperatureBoard:\t%v [degC]\n"
	format += "TemperatureCore:\t%v [degC]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.Uptime,
		m.RamUsage,
		m.RamTotal,
		m.StorageType,
		m.StorageUsage,
		m.StorageTotal,
		m.LinkType,
		m.LinkTxRate,
		m.LinkRxRate,
		m.LinkTxMax,
		m.LinkRxMax,
		m.FanSpeed,
		m.Type,
		m.CpuCores,
		m.CpuCombined,
		m.GpuCores,
		m.GpuCombined,
		m.TemperatureBoard,
		m.TemperatureCore,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *OnboardComputerStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *OnboardComputerStatus) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *OnboardComputerStatus) GetMessageName() string {
	return "OnboardComputerStatus"
}

// GetID gets the ID of the Message
func (m *OnboardComputerStatus) GetID() uint32 {
	return 390
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *OnboardComputerStatus) HasExtensionFields() bool {
	return false
}

func (m *OnboardComputerStatus) getV1Length() int {
	return 238
}

func (m *OnboardComputerStatus) getIOSlice() []byte {
	return make([]byte, 238+1)
}

// Read sets the field values of the message from the raw message payload
func (m *OnboardComputerStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the OnboardComputerStatus fields
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
func (m *OnboardComputerStatus) Write(version int) (output []byte, err error) {
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
