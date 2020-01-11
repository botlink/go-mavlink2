package ardupilotmega

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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*EfiStatus EFI Status Output */
type EfiStatus struct {
	/*EcuIndex ECU Index */
	EcuIndex float32
	/*Rpm RPM */
	Rpm float32
	/*FuelConsumed Fuel Consumed (grams) */
	FuelConsumed float32
	/*FuelFlow Fuel Flow Rate (g/min) */
	FuelFlow float32
	/*EngineLoad Engine Load (%) */
	EngineLoad float32
	/*ThrottlePosition Throttle Position (%) */
	ThrottlePosition float32
	/*SparkDwellTime Spark Dwell Time (ms) */
	SparkDwellTime float32
	/*BarometricPressure Barometric Pressure (kPa) */
	BarometricPressure float32
	/*IntakeManifoldPressure Intake Manifold Pressure (kPa)( */
	IntakeManifoldPressure float32
	/*IntakeManifoldTemperature Intake Manifold Temperature (degC) */
	IntakeManifoldTemperature float32
	/*CylinderHeadTemperature cylinder_head_temperature (degC) */
	CylinderHeadTemperature float32
	/*IgnitionTiming Ignition timing for cylinder i (Crank Angle degrees) */
	IgnitionTiming float32
	/*InjectionTime Injection time for injector i (ms) */
	InjectionTime float32
	/*Health EFI Health status */
	Health uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *EfiStatus) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "EcuIndex:\t%v \n"
	format += "Rpm:\t%v \n"
	format += "FuelConsumed:\t%v [g]\n"
	format += "FuelFlow:\t%v [g/min]\n"
	format += "EngineLoad:\t%v \n"
	format += "ThrottlePosition:\t%v \n"
	format += "SparkDwellTime:\t%v [ms]\n"
	format += "BarometricPressure:\t%v [kPa]\n"
	format += "IntakeManifoldPressure:\t%v [kPa]\n"
	format += "IntakeManifoldTemperature:\t%v [degC]\n"
	format += "CylinderHeadTemperature:\t%v [degC]\n"
	format += "IgnitionTiming:\t%v [deg]\n"
	format += "InjectionTime:\t%v [ms]\n"
	format += "Health:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.EcuIndex,
		m.Rpm,
		m.FuelConsumed,
		m.FuelFlow,
		m.EngineLoad,
		m.ThrottlePosition,
		m.SparkDwellTime,
		m.BarometricPressure,
		m.IntakeManifoldPressure,
		m.IntakeManifoldTemperature,
		m.CylinderHeadTemperature,
		m.IgnitionTiming,
		m.InjectionTime,
		m.Health,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *EfiStatus) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *EfiStatus) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *EfiStatus) GetMessageName() string {
	return "EfiStatus"
}

// GetID gets the ID of the Message
func (m *EfiStatus) GetID() uint32 {
	return 225
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *EfiStatus) HasExtensionFields() bool {
	return false
}

func (m *EfiStatus) getV1Length() int {
	return 53
}

func (m *EfiStatus) getIOSlice() []byte {
	return make([]byte, 53+1)
}

// Read sets the field values of the message from the raw message payload
func (m *EfiStatus) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the EfiStatus fields
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
func (m *EfiStatus) Write(version int) (output []byte, err error) {
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
