package slugs

import "github.com/queue-b/go-mavlink2"

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

// Dialect represents a collection of MAVLink messages
type Dialect struct{}

// GetName gets the name of the Dialect
func (d Dialect) GetName() string {
	return "slugs"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = slugsMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := slugsParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var slugsMessages = map[uint32]mavlink2.MessageMeta{
	170: mavlink2.MessageMeta{CRCExtra: 75, MinimumLength: 4, MaximumLength: 4},
	172: mavlink2.MessageMeta{CRCExtra: 168, MinimumLength: 24, MaximumLength: 24},
	173: mavlink2.MessageMeta{CRCExtra: 2, MinimumLength: 18, MaximumLength: 18},
	176: mavlink2.MessageMeta{CRCExtra: 228, MinimumLength: 32, MaximumLength: 32},
	177: mavlink2.MessageMeta{CRCExtra: 167, MinimumLength: 24, MaximumLength: 24},
	179: mavlink2.MessageMeta{CRCExtra: 132, MinimumLength: 12, MaximumLength: 12},
	180: mavlink2.MessageMeta{CRCExtra: 146, MinimumLength: 13, MaximumLength: 13},
	181: mavlink2.MessageMeta{CRCExtra: 104, MinimumLength: 3, MaximumLength: 3},
	184: mavlink2.MessageMeta{CRCExtra: 45, MinimumLength: 5, MaximumLength: 5},
	185: mavlink2.MessageMeta{CRCExtra: 113, MinimumLength: 10, MaximumLength: 10},
	186: mavlink2.MessageMeta{CRCExtra: 101, MinimumLength: 9, MaximumLength: 9},
	188: mavlink2.MessageMeta{CRCExtra: 5, MinimumLength: 3, MaximumLength: 3},
	189: mavlink2.MessageMeta{CRCExtra: 246, MinimumLength: 16, MaximumLength: 16},
	191: mavlink2.MessageMeta{CRCExtra: 17, MinimumLength: 5, MaximumLength: 5},
	192: mavlink2.MessageMeta{CRCExtra: 187, MinimumLength: 5, MaximumLength: 5},
	193: mavlink2.MessageMeta{CRCExtra: 160, MinimumLength: 21, MaximumLength: 21},
	194: mavlink2.MessageMeta{CRCExtra: 51, MinimumLength: 11, MaximumLength: 11},
	195: mavlink2.MessageMeta{CRCExtra: 59, MinimumLength: 14, MaximumLength: 14},
	196: mavlink2.MessageMeta{CRCExtra: 129, MinimumLength: 11, MaximumLength: 11},
	197: mavlink2.MessageMeta{CRCExtra: 39, MinimumLength: 4, MaximumLength: 4},
}

var slugsParsers = map[uint32]mavlink2.FrameParser{
	170: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CpuLoad{}

		err = message.Read(frame)

		return
	}),
	172: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensorBias{}

		err = message.Read(frame)

		return
	}),
	173: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Diagnostic{}

		err = message.Read(frame)

		return
	}),
	176: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SlugsNAVigation{}

		err = message.Read(frame)

		return
	}),
	177: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DataLog{}

		err = message.Read(frame)

		return
	}),
	179: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSDateTime{}

		err = message.Read(frame)

		return
	}),
	180: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MIDLvlCmds{}

		err = message.Read(frame)

		return
	}),
	181: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CtrlSrfcPt{}

		err = message.Read(frame)

		return
	}),
	184: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SlugsCameraOrder{}

		err = message.Read(frame)

		return
	}),
	185: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ControlSurface{}

		err = message.Read(frame)

		return
	}),
	186: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SlugsMobileLocation{}

		err = message.Read(frame)

		return
	}),
	188: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SlugsConfigurationCamera{}

		err = message.Read(frame)

		return
	}),
	189: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &IsrLocation{}

		err = message.Read(frame)

		return
	}),
	191: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VoltSensor{}

		err = message.Read(frame)

		return
	}),
	192: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PtzStatus{}

		err = message.Read(frame)

		return
	}),
	193: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &UavStatus{}

		err = message.Read(frame)

		return
	}),
	194: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &StatusGPS{}

		err = message.Read(frame)

		return
	}),
	195: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &NovatelDiag{}

		err = message.Read(frame)

		return
	}),
	196: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensorDiag{}

		err = message.Read(frame)

		return
	}),
	197: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Boot{}

		err = message.Read(frame)

		return
	}),
}
