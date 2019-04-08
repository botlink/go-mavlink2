package mavlink2

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

import "fmt"

// MessageMeta stores metadata information about a MAVLink message
type MessageMeta struct {
	CRCExtra      byte
	MinimumLength uint8
	MaximumLength uint8
}

// Message defines a common interface for MAVLink messages
type Message interface {
	GetVersion() int
	GetDialect() string
	GetMessageName() string
	GetID() uint32
	Read(Frame) error
	Write(int) ([]byte, error)
}

// UnknownMessage represents a message that is not part of any loaded Dialect
type UnknownMessage struct {
	ID      uint32
	Version int
	Raw     []byte
}

func (m *UnknownMessage) String() string {
	return fmt.Sprintf("ID:\t%v\nVersion:\t%v\nRaw:\t%x\n", m.ID, m.Version, m.Raw)
}

// GetVersion gets the MAVLink version of the Message contents
func (m *UnknownMessage) GetVersion() int {
	return m.Version
}

// GetDialect gets the name of the dialect that defines the Message
func (m *UnknownMessage) GetDialect() string {
	return "unknown"
}

// GetMessageName gets the name of the Message
func (m *UnknownMessage) GetMessageName() string {
	return "unknown"
}

// GetID gets the ID of the Message
func (m *UnknownMessage) GetID() uint32 {
	return m.ID
}

func (m *UnknownMessage) Read(frame Frame) (err error) {
	rawBytes := make([]byte, frame.GetMessageLength())

	copy(rawBytes, frame.GetMessageBytes())

	m.ID = frame.GetMessageID()
	m.Version = frame.GetVersion()
	m.Raw = rawBytes

	return
}

func (m *UnknownMessage) Write(version int) ([]byte, error) {
	return m.Raw, nil
}
