package common

/*
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
)

/*UavcanNodeInfo General information describing a particular UAVCAN node. Please refer to the definition of the UAVCAN service "uavcan.protocol.GetNodeInfo" for the background information. This message should be emitted by the system whenever a new node appears online, or an existing node reboots. Additionally, it can be emitted upon request from the other end of the MAVLink channel (see MAV_CMD_UAVCAN_GET_NODE_INFO). It is also not prohibited to emit this message unconditionally at a low frequency. The UAVCAN specification is available at http://uavcan.org. */
type UavcanNodeInfo struct {
	/*ReadVersion indicates the wire format the packet was read from */
	ReadVersion int
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*UptimeSec Time since the start-up of the node. */
	UptimeSec uint32
	/*SwVcsCommit Version control system (VCS) revision identifier (e.g. git short commit hash). Zero if unknown. */
	SwVcsCommit uint32
	/*Name Node name string. For example, "sapog.px4.io". */
	Name []byte
	/*HwVersionMajor Hardware major version number. */
	HwVersionMajor uint8
	/*HwVersionMinor Hardware minor version number. */
	HwVersionMinor uint8
	/*HwUniqueID Hardware unique 128-bit ID. */
	HwUniqueID []uint8
	/*SwVersionMajor Software major version number. */
	SwVersionMajor uint8
	/*SwVersionMinor Software minor version number. */
	SwVersionMinor uint8
}

// Read sets the field values of the message from the raw message payload
func (m *UavcanNodeInfo) Read(version int, payload []byte) (err error) {
	reader := bytes.NewReader(payload)

	m.ReadVersion = version
	err = binary.Read(reader, binary.LittleEndian, &m.TimeUsec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.UptimeSec)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SwVcsCommit)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.Name)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HwVersionMajor)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HwVersionMinor)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.HwUniqueID)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SwVersionMajor)
	if err != nil {
		return
	}

	err = binary.Read(reader, binary.LittleEndian, &m.SwVersionMinor)
	if err != nil {
		return
	}

	return
}

// Write encodes the field values of the message to a byte array
func (m *UavcanNodeInfo) Write(version int) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, m.TimeUsec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.UptimeSec)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SwVcsCommit)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.Name)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HwVersionMajor)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HwVersionMinor)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.HwUniqueID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SwVersionMajor)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, m.SwVersionMinor)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
