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

/*ControlSystemState The smoothed, monotonic system state used to feed the control loops of the system. */
type ControlSystemState struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*XAcc X acceleration in body frame */
	XAcc float32
	/*YAcc Y acceleration in body frame */
	YAcc float32
	/*ZAcc Z acceleration in body frame */
	ZAcc float32
	/*XVel X velocity in body frame */
	XVel float32
	/*YVel Y velocity in body frame */
	YVel float32
	/*ZVel Z velocity in body frame */
	ZVel float32
	/*XPos X position in local frame */
	XPos float32
	/*YPos Y position in local frame */
	YPos float32
	/*ZPos Z position in local frame */
	ZPos float32
	/*Airspeed Airspeed, set to -1 if unknown */
	Airspeed float32
	/*VelVariance Variance of body velocity estimate */
	VelVariance [3]float32
	/*PosVariance Variance in local position */
	PosVariance [3]float32
	/*Q The attitude, represented as Quaternion */
	Q [4]float32
	/*RollRate Angular rate in roll axis */
	RollRate float32
	/*PitchRate Angular rate in pitch axis */
	PitchRate float32
	/*YawRate Angular rate in yaw axis */
	YawRate float32
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *ControlSystemState) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeUsec:\t%v [us]\n"
	format += "XAcc:\t%v [m/s/s]\n"
	format += "YAcc:\t%v [m/s/s]\n"
	format += "ZAcc:\t%v [m/s/s]\n"
	format += "XVel:\t%v [m/s]\n"
	format += "YVel:\t%v [m/s]\n"
	format += "ZVel:\t%v [m/s]\n"
	format += "XPos:\t%v [m]\n"
	format += "YPos:\t%v [m]\n"
	format += "ZPos:\t%v [m]\n"
	format += "Airspeed:\t%v [m/s]\n"
	format += "VelVariance:\t%v \n"
	format += "PosVariance:\t%v \n"
	format += "Q:\t%v \n"
	format += "RollRate:\t%v [rad/s]\n"
	format += "PitchRate:\t%v [rad/s]\n"
	format += "YawRate:\t%v [rad/s]\n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.XAcc,
		m.YAcc,
		m.ZAcc,
		m.XVel,
		m.YVel,
		m.ZVel,
		m.XPos,
		m.YPos,
		m.ZPos,
		m.Airspeed,
		m.VelVariance,
		m.PosVariance,
		m.Q,
		m.RollRate,
		m.PitchRate,
		m.YawRate,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *ControlSystemState) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *ControlSystemState) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *ControlSystemState) GetMessageName() string {
	return "ControlSystemState"
}

// GetID gets the ID of the Message
func (m *ControlSystemState) GetID() uint32 {
	return 146
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *ControlSystemState) HasExtensionFields() bool {
	return false
}

func (m *ControlSystemState) getV1Length() int {
	return 100
}

func (m *ControlSystemState) getIOSlice() []byte {
	return make([]byte, 100+1)
}

// Read sets the field values of the message from the raw message payload
func (m *ControlSystemState) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the ControlSystemState fields
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
func (m *ControlSystemState) Write(version int) (output []byte, err error) {
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
