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

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*RCChannelsScaled The scaled values of the RC channels received: (-100%) -10000, (0%) 0, (100%) 10000. Channels that are inactive should be set to UINT16_MAX. */
type RCChannelsScaled struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*Chan1Scaled RC channel 1 value scaled. */
	Chan1Scaled int16
	/*Chan2Scaled RC channel 2 value scaled. */
	Chan2Scaled int16
	/*Chan3Scaled RC channel 3 value scaled. */
	Chan3Scaled int16
	/*Chan4Scaled RC channel 4 value scaled. */
	Chan4Scaled int16
	/*Chan5Scaled RC channel 5 value scaled. */
	Chan5Scaled int16
	/*Chan6Scaled RC channel 6 value scaled. */
	Chan6Scaled int16
	/*Chan7Scaled RC channel 7 value scaled. */
	Chan7Scaled int16
	/*Chan8Scaled RC channel 8 value scaled. */
	Chan8Scaled int16
	/*Port Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX. */
	Port uint8
	/*RSSI Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown. */
	RSSI uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *RCChannelsScaled) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v [ms]\n"
	format += "Chan1Scaled:\t%v \n"
	format += "Chan2Scaled:\t%v \n"
	format += "Chan3Scaled:\t%v \n"
	format += "Chan4Scaled:\t%v \n"
	format += "Chan5Scaled:\t%v \n"
	format += "Chan6Scaled:\t%v \n"
	format += "Chan7Scaled:\t%v \n"
	format += "Chan8Scaled:\t%v \n"
	format += "Port:\t%v \n"
	format += "RSSI:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.Chan1Scaled,
		m.Chan2Scaled,
		m.Chan3Scaled,
		m.Chan4Scaled,
		m.Chan5Scaled,
		m.Chan6Scaled,
		m.Chan7Scaled,
		m.Chan8Scaled,
		m.Port,
		m.RSSI,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *RCChannelsScaled) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *RCChannelsScaled) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *RCChannelsScaled) GetMessageName() string {
	return "RCChannelsScaled"
}

// GetID gets the ID of the Message
func (m *RCChannelsScaled) GetID() uint32 {
	return 34
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *RCChannelsScaled) HasExtensionFields() bool {
	return false
}

func (m *RCChannelsScaled) getV1Length() int {
	return 22
}

func (m *RCChannelsScaled) getIOSlice() []byte {
	return make([]byte, 22+1)
}

// Read sets the field values of the message from the raw message payload
func (m *RCChannelsScaled) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the RCChannelsScaled fields
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
func (m *RCChannelsScaled) Write(version int) (output []byte, err error) {
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
