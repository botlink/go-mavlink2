package ardupilotmega

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

/*GimbalReport 3 axis gimbal measurements. */
type GimbalReport struct {
	/*DeltaTime Time since last update. */
	DeltaTime float32
	/*DeltaAngleX Delta angle X. */
	DeltaAngleX float32
	/*DeltaAngleY Delta angle Y. */
	DeltaAngleY float32
	/*DeltaAngleZ Delta angle X. */
	DeltaAngleZ float32
	/*DeltaVelocityX Delta velocity X. */
	DeltaVelocityX float32
	/*DeltaVelocityY Delta velocity Y. */
	DeltaVelocityY float32
	/*DeltaVelocityZ Delta velocity Z. */
	DeltaVelocityZ float32
	/*JointRoll Joint ROLL. */
	JointRoll float32
	/*JointEl Joint EL. */
	JointEl float32
	/*JointAz Joint AZ. */
	JointAz float32
	/*TargetSystem System ID. */
	TargetSystem uint8
	/*TargetComponent Component ID. */
	TargetComponent uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GimbalReport) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "DeltaTime:\t%v [s]\n"
	format += "DeltaAngleX:\t%v [rad]\n"
	format += "DeltaAngleY:\t%v [rad]\n"
	format += "DeltaAngleZ:\t%v [rad]\n"
	format += "DeltaVelocityX:\t%v [m/s]\n"
	format += "DeltaVelocityY:\t%v [m/s]\n"
	format += "DeltaVelocityZ:\t%v [m/s]\n"
	format += "JointRoll:\t%v [rad]\n"
	format += "JointEl:\t%v [rad]\n"
	format += "JointAz:\t%v [rad]\n"
	format += "TargetSystem:\t%v \n"
	format += "TargetComponent:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.DeltaTime,
		m.DeltaAngleX,
		m.DeltaAngleY,
		m.DeltaAngleZ,
		m.DeltaVelocityX,
		m.DeltaVelocityY,
		m.DeltaVelocityZ,
		m.JointRoll,
		m.JointEl,
		m.JointAz,
		m.TargetSystem,
		m.TargetComponent,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GimbalReport) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GimbalReport) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *GimbalReport) GetMessageName() string {
	return "GimbalReport"
}

// GetID gets the ID of the Message
func (m *GimbalReport) GetID() uint32 {
	return 200
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GimbalReport) HasExtensionFields() bool {
	return false
}

func (m *GimbalReport) getV1Length() int {
	return 42
}

func (m *GimbalReport) getIOSlice() []byte {
	return make([]byte, 42+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GimbalReport) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GimbalReport fields
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
func (m *GimbalReport) Write(version int) (output []byte, err error) {
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
