package matrixpilot

import mavlink2 "github.com/queue-b/go-mavlink2"

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
	return "matrixpilot"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = matrixpilotMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := matrixpilotParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var matrixpilotMessages = map[uint32]mavlink2.MessageMeta{
	150: mavlink2.MessageMeta{CRCExtra: 181, MinimumLength: 2, MaximumLength: 2},
	151: mavlink2.MessageMeta{CRCExtra: 26, MinimumLength: 6, MaximumLength: 6},
	152: mavlink2.MessageMeta{CRCExtra: 101, MinimumLength: 58, MaximumLength: 58},
	153: mavlink2.MessageMeta{CRCExtra: 109, MinimumLength: 6, MaximumLength: 6},
	155: mavlink2.MessageMeta{CRCExtra: 12, MinimumLength: 53, MaximumLength: 53},
	156: mavlink2.MessageMeta{CRCExtra: 218, MinimumLength: 7, MaximumLength: 7},
	157: mavlink2.MessageMeta{CRCExtra: 133, MinimumLength: 3, MaximumLength: 3},
	158: mavlink2.MessageMeta{CRCExtra: 208, MinimumLength: 4, MaximumLength: 4},
	170: mavlink2.MessageMeta{CRCExtra: 103, MinimumLength: 61, MaximumLength: 61},
	171: mavlink2.MessageMeta{CRCExtra: 245, MinimumLength: 108, MaximumLength: 108},
	172: mavlink2.MessageMeta{CRCExtra: 191, MinimumLength: 10, MaximumLength: 10},
	173: mavlink2.MessageMeta{CRCExtra: 54, MinimumLength: 16, MaximumLength: 16},
	174: mavlink2.MessageMeta{CRCExtra: 54, MinimumLength: 20, MaximumLength: 20},
	175: mavlink2.MessageMeta{CRCExtra: 171, MinimumLength: 24, MaximumLength: 24},
	176: mavlink2.MessageMeta{CRCExtra: 142, MinimumLength: 28, MaximumLength: 28},
	177: mavlink2.MessageMeta{CRCExtra: 249, MinimumLength: 14, MaximumLength: 14},
	178: mavlink2.MessageMeta{CRCExtra: 123, MinimumLength: 17, MaximumLength: 17},
	179: mavlink2.MessageMeta{CRCExtra: 7, MinimumLength: 60, MaximumLength: 60},
	180: mavlink2.MessageMeta{CRCExtra: 222, MinimumLength: 110, MaximumLength: 110},
	181: mavlink2.MessageMeta{CRCExtra: 55, MinimumLength: 28, MaximumLength: 28},
	182: mavlink2.MessageMeta{CRCExtra: 154, MinimumLength: 16, MaximumLength: 16},
	183: mavlink2.MessageMeta{CRCExtra: 175, MinimumLength: 12, MaximumLength: 12},
	184: mavlink2.MessageMeta{CRCExtra: 41, MinimumLength: 20, MaximumLength: 20},
	185: mavlink2.MessageMeta{CRCExtra: 87, MinimumLength: 8, MaximumLength: 8},
	186: mavlink2.MessageMeta{CRCExtra: 144, MinimumLength: 25, MaximumLength: 25},
	187: mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 12, MaximumLength: 12},
	188: mavlink2.MessageMeta{CRCExtra: 91, MinimumLength: 12, MaximumLength: 12},
}

var matrixpilotParsers = map[uint32]mavlink2.FrameParser{
	150: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionSet{}

		err = message.Read(frame)

		return
	}),
	151: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionReadReq{}

		err = message.Read(frame)

		return
	}),
	152: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionBufferFunction{}

		err = message.Read(frame)

		return
	}),
	153: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionBufferFunctionAck{}

		err = message.Read(frame)

		return
	}),
	155: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionDirectory{}

		err = message.Read(frame)

		return
	}),
	156: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionDirectoryAck{}

		err = message.Read(frame)

		return
	}),
	157: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionCommand{}

		err = message.Read(frame)

		return
	}),
	158: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlexifunctionCommandAck{}

		err = message.Read(frame)

		return
	}),
	170: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF2A{}

		err = message.Read(frame)

		return
	}),
	171: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF2B{}

		err = message.Read(frame)

		return
	}),
	172: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF4{}

		err = message.Read(frame)

		return
	}),
	173: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF5{}

		err = message.Read(frame)

		return
	}),
	174: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF6{}

		err = message.Read(frame)

		return
	}),
	175: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF7{}

		err = message.Read(frame)

		return
	}),
	176: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF8{}

		err = message.Read(frame)

		return
	}),
	177: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF13{}

		err = message.Read(frame)

		return
	}),
	178: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF14{}

		err = message.Read(frame)

		return
	}),
	179: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF15{}

		err = message.Read(frame)

		return
	}),
	180: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF16{}

		err = message.Read(frame)

		return
	}),
	181: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Altitudes{}

		err = message.Read(frame)

		return
	}),
	182: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Airspeeds{}

		err = message.Read(frame)

		return
	}),
	183: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF17{}

		err = message.Read(frame)

		return
	}),
	184: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF18{}

		err = message.Read(frame)

		return
	}),
	185: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF19{}

		err = message.Read(frame)

		return
	}),
	186: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF20{}

		err = message.Read(frame)

		return
	}),
	187: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF21{}

		err = message.Read(frame)

		return
	}),
	188: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialUdbExtraF22{}

		err = message.Read(frame)

		return
	}),
}
