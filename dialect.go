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

import "encoding/binary"

// Dialect represents a collection of MAVLink message definitions
type Dialect interface {
	GetMeta(uint32) (MessageMeta, error)
	GetMessage(Frame) (Message, error)
}

// Dialects represents a collection of Dialects
type Dialects []Dialect

// Validate checks the frame to see if there are any obvious errors
func (d Dialects) Validate(frame Frame) error {
	checksum := new(X25CRC)
	checksum.Reset()

	var meta MessageMeta
	var err error
	for _, dialect := range d {
		meta, err = dialect.GetMeta(frame.GetMessageID())

		// If the Dialect doesn't know about the message, just try the next one
		if err == ErrUnknownMessage {
			continue
		}

		// If there's any error beyond ErrUnknownMessage just give up on the process
		if err != nil {
			return err
		}
	}

	// If the error is still ErrUnknownMessage, there isn't a Dialect that understands the message
	if err == ErrUnknownMessage {
		return err
	}

	// Check if the message is too short
	if frame.GetMessageLength() < meta.MinimumLength {
		return ErrMessageTooShort
	}

	// Check if the message is too long
	if frame.GetMessageLength() > meta.MaximumLength {
		return ErrMessageTooLong
	}

	_, err = checksum.Write(append(frame.GetChecksumInput(), meta.CRCExtra))

	if err != nil {
		return err
	}

	// Caclulate the checksum
	sum16 := binary.LittleEndian.Uint16(checksum.Sum(nil))

	// Check if the checksum matches
	if sum16 != frame.GetChecksum() {
		return ErrInvalidChecksum
	}

	return nil
}
