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

/*StorageInformation Information about a storage medium. This message is sent in response to a request and whenever the status of the storage changes (STORAGE_STATUS). */
type StorageInformation struct {
	/*TimeBootMs Timestamp (time since system boot). */
	TimeBootMs uint32
	/*TotalCapacity Total capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored. */
	TotalCapacity float32
	/*UsedCapacity Used capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored. */
	UsedCapacity float32
	/*AvailableCapacity Available storage capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored. */
	AvailableCapacity float32
	/*ReadSpeed Read speed. */
	ReadSpeed float32
	/*WriteSpeed Write speed. */
	WriteSpeed float32
	/*StorageID Storage ID (1 for first, 2 for second, etc.) */
	StorageID uint8
	/*StorageCount Number of storage devices */
	StorageCount uint8
	/*Status Status of storage */
	Status uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *StorageInformation) String() string {
	format := ""
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	format += "Name:\t%v/%v\n"
	// Output field values based on the decoded message type
	format += "TimeBootMs:\t%v [ms]\n"
	format += "TotalCapacity:\t%v [MiB]\n"
	format += "UsedCapacity:\t%v [MiB]\n"
	format += "AvailableCapacity:\t%v [MiB]\n"
	format += "ReadSpeed:\t%v [MiB/s]\n"
	format += "WriteSpeed:\t%v [MiB/s]\n"
	format += "StorageID:\t%v \n"
	format += "StorageCount:\t%v \n"
	format += "Status:\t%v \n"

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeBootMs,
		m.TotalCapacity,
		m.UsedCapacity,
		m.AvailableCapacity,
		m.ReadSpeed,
		m.WriteSpeed,
		m.StorageID,
		m.StorageCount,
		m.Status,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *StorageInformation) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *StorageInformation) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *StorageInformation) GetMessageName() string {
	return "StorageInformation"
}

// GetID gets the ID of the Message
func (m *StorageInformation) GetID() uint32 {
	return 261
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *StorageInformation) HasExtensionFields() bool {
	return false
}

func (m *StorageInformation) getV1Length() int {
	return 27
}

func (m *StorageInformation) getIOSlice() []byte {
	return make([]byte, 27+1)
}

// Read sets the field values of the message from the raw message payload
func (m *StorageInformation) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the StorageInformation fields
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
func (m *StorageInformation) Write(version int) (output []byte, err error) {
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
