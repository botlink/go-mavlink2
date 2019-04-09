package asluav

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
	return "asluav"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = asluavMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := asluavParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var asluavMessages = map[uint32]mavlink2.MessageMeta{
	78:  mavlink2.MessageMeta{CRCExtra: 119, MinimumLength: 47, MaximumLength: 47},
	79:  mavlink2.MessageMeta{CRCExtra: 102, MinimumLength: 45, MaximumLength: 45},
	201: mavlink2.MessageMeta{CRCExtra: 218, MinimumLength: 16, MaximumLength: 16},
	202: mavlink2.MessageMeta{CRCExtra: 231, MinimumLength: 41, MaximumLength: 41},
	203: mavlink2.MessageMeta{CRCExtra: 172, MinimumLength: 98, MaximumLength: 98},
	204: mavlink2.MessageMeta{CRCExtra: 251, MinimumLength: 38, MaximumLength: 38},
	205: mavlink2.MessageMeta{CRCExtra: 97, MinimumLength: 14, MaximumLength: 14},
	206: mavlink2.MessageMeta{CRCExtra: 64, MinimumLength: 32, MaximumLength: 32},
	207: mavlink2.MessageMeta{CRCExtra: 234, MinimumLength: 33, MaximumLength: 33},
	208: mavlink2.MessageMeta{CRCExtra: 144, MinimumLength: 16, MaximumLength: 16},
	209: mavlink2.MessageMeta{CRCExtra: 155, MinimumLength: 41, MaximumLength: 41},
	210: mavlink2.MessageMeta{CRCExtra: 20, MinimumLength: 102, MaximumLength: 102},
	211: mavlink2.MessageMeta{CRCExtra: 54, MinimumLength: 16, MaximumLength: 16},
	212: mavlink2.MessageMeta{CRCExtra: 222, MinimumLength: 46, MaximumLength: 46},
	213: mavlink2.MessageMeta{CRCExtra: 200, MinimumLength: 14, MaximumLength: 14},
	214: mavlink2.MessageMeta{CRCExtra: 23, MinimumLength: 24, MaximumLength: 24},
}

var ASLUAVParsers = map[uint32]mavlink2.FrameParser{
	78: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CommandIntStamped{}

		err = message.Read(frame)

		return
	}),
	79: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CommandLongStamped{}

		err = message.Read(frame)

		return
	}),
	201: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensPower{}

		err = message.Read(frame)

		return
	}),
	202: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensMppt{}

		err = message.Read(frame)

		return
	}),
	203: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AslctrlData{}

		err = message.Read(frame)

		return
	}),
	204: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AslctrlDebug{}

		err = message.Read(frame)

		return
	}),
	205: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AsluavStatus{}

		err = message.Read(frame)

		return
	}),
	206: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EKFExt{}

		err = message.Read(frame)

		return
	}),
	207: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AslObctrl{}

		err = message.Read(frame)

		return
	}),
	208: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensAtmos{}

		err = message.Read(frame)

		return
	}),
	209: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensBatmon{}

		err = message.Read(frame)

		return
	}),
	210: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FwSoaringData{}

		err = message.Read(frame)

		return
	}),
	211: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensorpodStatus{}

		err = message.Read(frame)

		return
	}),
	212: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensPowerBoard{}

		err = message.Read(frame)

		return
	}),
	213: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GsmLinkStatus{}

		err = message.Read(frame)

		return
	}),
	214: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SatcomLinkStatus{}

		err = message.Read(frame)

		return
	}),
}
