package icarous

import "github.com/queue-b/go-mavlink2"

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

// Dialect represents a collection of MAVLink messages
type Dialect struct{}

// GetName gets the name of the Dialect
func (d Dialect) GetName() string {
	return "icarous"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = icarousMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := icarousParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var icarousMessages = map[uint32]mavlink2.MessageMeta{
	42000: mavlink2.MessageMeta{CRCExtra: 227, MinimumLength: 1, MaximumLength: 1},
	42001: mavlink2.MessageMeta{CRCExtra: 239, MinimumLength: 46, MaximumLength: 46},
}

var icarousParsers = map[uint32]mavlink2.FrameParser{
	42000: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &IcarousHeartbeat{}

		err = message.Read(frame)

		return
	}),
	42001: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &IcarousKinematicBands{}

		err = message.Read(frame)

		return
	}),
}
