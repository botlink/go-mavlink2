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
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*OpticalFlowRad Optical flow from an angular rate flow sensor (e.g. PX4FLOW or mouse sensor) */
type OpticalFlowRad struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*IntegrationTimeUs Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the. */
	IntegrationTimeUs uint32
	/*IntegratedX Flow around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.) */
	IntegratedX float32
	/*IntegratedY Flow around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.) */
	IntegratedY float32
	/*IntegratedXgyro RH rotation around X axis */
	IntegratedXgyro float32
	/*IntegratedYgyro RH rotation around Y axis */
	IntegratedYgyro float32
	/*IntegratedZgyro RH rotation around Z axis */
	IntegratedZgyro float32
	/*TimeDeltaDistanceUs Time since the distance was sampled. */
	TimeDeltaDistanceUs uint32
	/*Distance Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance. */
	Distance float32
	/*Temperature Temperature */
	Temperature int16
	/*SensorID Sensor ID */
	SensorID uint8
	/*Quality Optical flow quality / confidence. 0: no valid flow, 255: maximum quality */
	Quality uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *OpticalFlowRad) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("IntegrationTimeUs:\t%v [us]\n")
	builder.WriteString("IntegratedX:\t%v [rad]\n")
	builder.WriteString("IntegratedY:\t%v [rad]\n")
	builder.WriteString("IntegratedXgyro:\t%v [rad]\n")
	builder.WriteString("IntegratedYgyro:\t%v [rad]\n")
	builder.WriteString("IntegratedZgyro:\t%v [rad]\n")
	builder.WriteString("TimeDeltaDistanceUs:\t%v [us]\n")
	builder.WriteString("Distance:\t%v [m]\n")
	builder.WriteString("Temperature:\t%v [cdegC]\n")
	builder.WriteString("SensorID:\t%v \n")
	builder.WriteString("Quality:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.IntegrationTimeUs,
		m.IntegratedX,
		m.IntegratedY,
		m.IntegratedXgyro,
		m.IntegratedYgyro,
		m.IntegratedZgyro,
		m.TimeDeltaDistanceUs,
		m.Distance,
		m.Temperature,
		m.SensorID,
		m.Quality,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *OpticalFlowRad) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *OpticalFlowRad) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *OpticalFlowRad) GetMessageName() string {
	return "OpticalFlowRad"
}

// GetID gets the ID of the Message
func (m *OpticalFlowRad) GetID() uint32 {
	return 106
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *OpticalFlowRad) HasExtensionFields() bool {
	return false
}

func (m *OpticalFlowRad) getV1Length() int {
	return 44
}

func (m *OpticalFlowRad) getIOSlice() []byte {
	return make([]byte, 44+1)
}

// Read sets the field values of the message from the raw message payload
func (m *OpticalFlowRad) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the OpticalFlowRad fields
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
func (m *OpticalFlowRad) Write(version int) (output []byte, err error) {
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
