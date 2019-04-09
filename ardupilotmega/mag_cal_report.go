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

/*MagCalReport Reports results of completed compass calibration. Sent until MAG_CAL_ACK received. */
type MagCalReport struct {
	/*Fitness RMS milligauss residuals. */
	Fitness float32
	/*OfsX X offset. */
	OfsX float32
	/*OfsY Y offset. */
	OfsY float32
	/*OfsZ Z offset. */
	OfsZ float32
	/*DiagX X diagonal (matrix 11). */
	DiagX float32
	/*DiagY Y diagonal (matrix 22). */
	DiagY float32
	/*DiagZ Z diagonal (matrix 33). */
	DiagZ float32
	/*OffdiagX X off-diagonal (matrix 12 and 21). */
	OffdiagX float32
	/*OffdiagY Y off-diagonal (matrix 13 and 31). */
	OffdiagY float32
	/*OffdiagZ Z off-diagonal (matrix 32 and 23). */
	OffdiagZ float32
	/*CompassID Compass being calibrated. */
	CompassID uint8
	/*CalMask Bitmask of compasses being calibrated. */
	CalMask uint8
	/*CalStatus Calibration Status. */
	CalStatus uint8
	/*Autosaved 0=requires a MAV_CMD_DO_ACCEPT_MAG_CAL, 1=saved to parameters. */
	Autosaved uint8
	/*OrientationConfIDence Confidence in orientation (higher is better). */
	OrientationConfIDence float32
	/*OldOrientation orientation before calibration. */
	OldOrientation uint8
	/*NewOrientation orientation after calibration. */
	NewOrientation uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *MagCalReport) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "Fitness:\t%v [mgauss]\n"
	format += "OfsX:\t%v \n"
	format += "OfsY:\t%v \n"
	format += "OfsZ:\t%v \n"
	format += "DiagX:\t%v \n"
	format += "DiagY:\t%v \n"
	format += "DiagZ:\t%v \n"
	format += "OffdiagX:\t%v \n"
	format += "OffdiagY:\t%v \n"
	format += "OffdiagZ:\t%v \n"
	format += "CompassID:\t%v \n"
	format += "CalMask:\t%v \n"
	format += "CalStatus:\t%v \n"
	format += "Autosaved:\t%v \n"
	if m.HasExtensionFieldValues {
		format += "OrientationConfIDence:\t%v\n"
		format += "OldOrientation:\t%v\n"
		format += "NewOrientation:\t%v\n"
	}

	if m.HasExtensionFieldValues {
		fmt.Fprintf(
			writer,
			format,
			m.GetDialect(),
			m.GetMessageName(),
			m.Fitness,
			m.OfsX,
			m.OfsY,
			m.OfsZ,
			m.DiagX,
			m.DiagY,
			m.DiagZ,
			m.OffdiagX,
			m.OffdiagY,
			m.OffdiagZ,
			m.CompassID,
			m.CalMask,
			m.CalStatus,
			m.Autosaved,
			m.OrientationConfIDence,
			m.OldOrientation,
			m.NewOrientation,
		)

		writer.Flush()
		return string(buffer.Bytes())
	}

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.Fitness,
		m.OfsX,
		m.OfsY,
		m.OfsZ,
		m.DiagX,
		m.DiagY,
		m.DiagZ,
		m.OffdiagX,
		m.OffdiagY,
		m.OffdiagZ,
		m.CompassID,
		m.CalMask,
		m.CalStatus,
		m.Autosaved,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *MagCalReport) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *MagCalReport) GetDialect() string {
	return "ardupilotmega"
}

// GetMessageName gets the name of the Message
func (m *MagCalReport) GetMessageName() string {
	return "MagCalReport"
}

// GetID gets the ID of the Message
func (m *MagCalReport) GetID() uint32 {
	return 192
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *MagCalReport) HasExtensionFields() bool {
	return true
}

func (m *MagCalReport) getV1Length() int {
	return 44
}

func (m *MagCalReport) getIOSlice() []byte {
	return make([]byte, 50+1)
}

// Read sets the field values of the message from the raw message payload
func (m *MagCalReport) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the MagCalReport fields
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
func (m *MagCalReport) Write(version int) (output []byte, err error) {
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
