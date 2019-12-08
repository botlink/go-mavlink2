package common

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
	return "common"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = commonMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := commonParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var commonMessages = map[uint32]mavlink2.MessageMeta{
	0:    mavlink2.MessageMeta{CRCExtra: 50, MinimumLength: 9, MaximumLength: 9},
	1:    mavlink2.MessageMeta{CRCExtra: 124, MinimumLength: 31, MaximumLength: 31},
	2:    mavlink2.MessageMeta{CRCExtra: 137, MinimumLength: 12, MaximumLength: 12},
	4:    mavlink2.MessageMeta{CRCExtra: 237, MinimumLength: 14, MaximumLength: 14},
	5:    mavlink2.MessageMeta{CRCExtra: 217, MinimumLength: 28, MaximumLength: 28},
	6:    mavlink2.MessageMeta{CRCExtra: 104, MinimumLength: 3, MaximumLength: 3},
	7:    mavlink2.MessageMeta{CRCExtra: 119, MinimumLength: 32, MaximumLength: 32},
	11:   mavlink2.MessageMeta{CRCExtra: 89, MinimumLength: 6, MaximumLength: 6},
	20:   mavlink2.MessageMeta{CRCExtra: 214, MinimumLength: 20, MaximumLength: 20},
	21:   mavlink2.MessageMeta{CRCExtra: 159, MinimumLength: 2, MaximumLength: 2},
	22:   mavlink2.MessageMeta{CRCExtra: 220, MinimumLength: 25, MaximumLength: 25},
	23:   mavlink2.MessageMeta{CRCExtra: 168, MinimumLength: 23, MaximumLength: 23},
	24:   mavlink2.MessageMeta{CRCExtra: 24, MinimumLength: 30, MaximumLength: 50},
	25:   mavlink2.MessageMeta{CRCExtra: 23, MinimumLength: 101, MaximumLength: 101},
	26:   mavlink2.MessageMeta{CRCExtra: 170, MinimumLength: 22, MaximumLength: 22},
	27:   mavlink2.MessageMeta{CRCExtra: 144, MinimumLength: 26, MaximumLength: 26},
	28:   mavlink2.MessageMeta{CRCExtra: 67, MinimumLength: 16, MaximumLength: 16},
	29:   mavlink2.MessageMeta{CRCExtra: 115, MinimumLength: 14, MaximumLength: 14},
	30:   mavlink2.MessageMeta{CRCExtra: 39, MinimumLength: 28, MaximumLength: 28},
	31:   mavlink2.MessageMeta{CRCExtra: 246, MinimumLength: 32, MaximumLength: 32},
	32:   mavlink2.MessageMeta{CRCExtra: 185, MinimumLength: 28, MaximumLength: 28},
	33:   mavlink2.MessageMeta{CRCExtra: 104, MinimumLength: 28, MaximumLength: 28},
	34:   mavlink2.MessageMeta{CRCExtra: 237, MinimumLength: 22, MaximumLength: 22},
	35:   mavlink2.MessageMeta{CRCExtra: 244, MinimumLength: 22, MaximumLength: 22},
	36:   mavlink2.MessageMeta{CRCExtra: 222, MinimumLength: 21, MaximumLength: 37},
	37:   mavlink2.MessageMeta{CRCExtra: 212, MinimumLength: 6, MaximumLength: 7},
	38:   mavlink2.MessageMeta{CRCExtra: 9, MinimumLength: 6, MaximumLength: 7},
	39:   mavlink2.MessageMeta{CRCExtra: 254, MinimumLength: 37, MaximumLength: 38},
	40:   mavlink2.MessageMeta{CRCExtra: 230, MinimumLength: 4, MaximumLength: 5},
	41:   mavlink2.MessageMeta{CRCExtra: 28, MinimumLength: 4, MaximumLength: 4},
	42:   mavlink2.MessageMeta{CRCExtra: 28, MinimumLength: 2, MaximumLength: 2},
	43:   mavlink2.MessageMeta{CRCExtra: 132, MinimumLength: 2, MaximumLength: 3},
	44:   mavlink2.MessageMeta{CRCExtra: 221, MinimumLength: 4, MaximumLength: 5},
	45:   mavlink2.MessageMeta{CRCExtra: 232, MinimumLength: 2, MaximumLength: 3},
	46:   mavlink2.MessageMeta{CRCExtra: 11, MinimumLength: 2, MaximumLength: 2},
	47:   mavlink2.MessageMeta{CRCExtra: 153, MinimumLength: 3, MaximumLength: 4},
	48:   mavlink2.MessageMeta{CRCExtra: 41, MinimumLength: 13, MaximumLength: 21},
	49:   mavlink2.MessageMeta{CRCExtra: 39, MinimumLength: 12, MaximumLength: 20},
	50:   mavlink2.MessageMeta{CRCExtra: 78, MinimumLength: 37, MaximumLength: 37},
	51:   mavlink2.MessageMeta{CRCExtra: 196, MinimumLength: 4, MaximumLength: 5},
	54:   mavlink2.MessageMeta{CRCExtra: 15, MinimumLength: 27, MaximumLength: 27},
	55:   mavlink2.MessageMeta{CRCExtra: 3, MinimumLength: 25, MaximumLength: 25},
	61:   mavlink2.MessageMeta{CRCExtra: 167, MinimumLength: 72, MaximumLength: 72},
	62:   mavlink2.MessageMeta{CRCExtra: 183, MinimumLength: 26, MaximumLength: 26},
	63:   mavlink2.MessageMeta{CRCExtra: 119, MinimumLength: 181, MaximumLength: 181},
	64:   mavlink2.MessageMeta{CRCExtra: 191, MinimumLength: 225, MaximumLength: 225},
	65:   mavlink2.MessageMeta{CRCExtra: 118, MinimumLength: 42, MaximumLength: 42},
	66:   mavlink2.MessageMeta{CRCExtra: 148, MinimumLength: 6, MaximumLength: 6},
	67:   mavlink2.MessageMeta{CRCExtra: 21, MinimumLength: 4, MaximumLength: 4},
	69:   mavlink2.MessageMeta{CRCExtra: 243, MinimumLength: 11, MaximumLength: 11},
	70:   mavlink2.MessageMeta{CRCExtra: 124, MinimumLength: 18, MaximumLength: 38},
	73:   mavlink2.MessageMeta{CRCExtra: 38, MinimumLength: 37, MaximumLength: 38},
	74:   mavlink2.MessageMeta{CRCExtra: 20, MinimumLength: 20, MaximumLength: 20},
	75:   mavlink2.MessageMeta{CRCExtra: 158, MinimumLength: 35, MaximumLength: 35},
	76:   mavlink2.MessageMeta{CRCExtra: 152, MinimumLength: 33, MaximumLength: 33},
	77:   mavlink2.MessageMeta{CRCExtra: 143, MinimumLength: 3, MaximumLength: 10},
	81:   mavlink2.MessageMeta{CRCExtra: 106, MinimumLength: 22, MaximumLength: 22},
	82:   mavlink2.MessageMeta{CRCExtra: 49, MinimumLength: 39, MaximumLength: 39},
	83:   mavlink2.MessageMeta{CRCExtra: 22, MinimumLength: 37, MaximumLength: 37},
	84:   mavlink2.MessageMeta{CRCExtra: 143, MinimumLength: 53, MaximumLength: 53},
	85:   mavlink2.MessageMeta{CRCExtra: 140, MinimumLength: 51, MaximumLength: 51},
	86:   mavlink2.MessageMeta{CRCExtra: 5, MinimumLength: 53, MaximumLength: 53},
	87:   mavlink2.MessageMeta{CRCExtra: 150, MinimumLength: 51, MaximumLength: 51},
	89:   mavlink2.MessageMeta{CRCExtra: 231, MinimumLength: 28, MaximumLength: 28},
	90:   mavlink2.MessageMeta{CRCExtra: 183, MinimumLength: 56, MaximumLength: 56},
	91:   mavlink2.MessageMeta{CRCExtra: 63, MinimumLength: 42, MaximumLength: 42},
	92:   mavlink2.MessageMeta{CRCExtra: 54, MinimumLength: 33, MaximumLength: 33},
	93:   mavlink2.MessageMeta{CRCExtra: 47, MinimumLength: 81, MaximumLength: 81},
	100:  mavlink2.MessageMeta{CRCExtra: 175, MinimumLength: 26, MaximumLength: 34},
	101:  mavlink2.MessageMeta{CRCExtra: 102, MinimumLength: 32, MaximumLength: 117},
	102:  mavlink2.MessageMeta{CRCExtra: 158, MinimumLength: 32, MaximumLength: 117},
	103:  mavlink2.MessageMeta{CRCExtra: 208, MinimumLength: 20, MaximumLength: 57},
	104:  mavlink2.MessageMeta{CRCExtra: 56, MinimumLength: 32, MaximumLength: 116},
	105:  mavlink2.MessageMeta{CRCExtra: 93, MinimumLength: 62, MaximumLength: 62},
	106:  mavlink2.MessageMeta{CRCExtra: 138, MinimumLength: 44, MaximumLength: 44},
	107:  mavlink2.MessageMeta{CRCExtra: 108, MinimumLength: 64, MaximumLength: 64},
	108:  mavlink2.MessageMeta{CRCExtra: 32, MinimumLength: 84, MaximumLength: 84},
	109:  mavlink2.MessageMeta{CRCExtra: 185, MinimumLength: 9, MaximumLength: 9},
	110:  mavlink2.MessageMeta{CRCExtra: 84, MinimumLength: 254, MaximumLength: 254},
	111:  mavlink2.MessageMeta{CRCExtra: 34, MinimumLength: 16, MaximumLength: 16},
	112:  mavlink2.MessageMeta{CRCExtra: 174, MinimumLength: 12, MaximumLength: 12},
	113:  mavlink2.MessageMeta{CRCExtra: 124, MinimumLength: 36, MaximumLength: 36},
	114:  mavlink2.MessageMeta{CRCExtra: 237, MinimumLength: 44, MaximumLength: 44},
	115:  mavlink2.MessageMeta{CRCExtra: 4, MinimumLength: 64, MaximumLength: 64},
	116:  mavlink2.MessageMeta{CRCExtra: 76, MinimumLength: 22, MaximumLength: 22},
	117:  mavlink2.MessageMeta{CRCExtra: 128, MinimumLength: 6, MaximumLength: 6},
	118:  mavlink2.MessageMeta{CRCExtra: 56, MinimumLength: 14, MaximumLength: 14},
	119:  mavlink2.MessageMeta{CRCExtra: 116, MinimumLength: 12, MaximumLength: 12},
	120:  mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 97, MaximumLength: 97},
	121:  mavlink2.MessageMeta{CRCExtra: 237, MinimumLength: 2, MaximumLength: 2},
	122:  mavlink2.MessageMeta{CRCExtra: 203, MinimumLength: 2, MaximumLength: 2},
	123:  mavlink2.MessageMeta{CRCExtra: 250, MinimumLength: 113, MaximumLength: 113},
	124:  mavlink2.MessageMeta{CRCExtra: 87, MinimumLength: 35, MaximumLength: 35},
	125:  mavlink2.MessageMeta{CRCExtra: 203, MinimumLength: 6, MaximumLength: 6},
	126:  mavlink2.MessageMeta{CRCExtra: 220, MinimumLength: 79, MaximumLength: 79},
	127:  mavlink2.MessageMeta{CRCExtra: 25, MinimumLength: 35, MaximumLength: 35},
	128:  mavlink2.MessageMeta{CRCExtra: 226, MinimumLength: 35, MaximumLength: 35},
	129:  mavlink2.MessageMeta{CRCExtra: 46, MinimumLength: 22, MaximumLength: 22},
	130:  mavlink2.MessageMeta{CRCExtra: 29, MinimumLength: 13, MaximumLength: 13},
	131:  mavlink2.MessageMeta{CRCExtra: 223, MinimumLength: 255, MaximumLength: 255},
	132:  mavlink2.MessageMeta{CRCExtra: 85, MinimumLength: 14, MaximumLength: 38},
	133:  mavlink2.MessageMeta{CRCExtra: 6, MinimumLength: 18, MaximumLength: 18},
	134:  mavlink2.MessageMeta{CRCExtra: 229, MinimumLength: 43, MaximumLength: 43},
	135:  mavlink2.MessageMeta{CRCExtra: 203, MinimumLength: 8, MaximumLength: 8},
	136:  mavlink2.MessageMeta{CRCExtra: 1, MinimumLength: 22, MaximumLength: 22},
	137:  mavlink2.MessageMeta{CRCExtra: 195, MinimumLength: 14, MaximumLength: 14},
	138:  mavlink2.MessageMeta{CRCExtra: 109, MinimumLength: 36, MaximumLength: 120},
	139:  mavlink2.MessageMeta{CRCExtra: 168, MinimumLength: 43, MaximumLength: 43},
	140:  mavlink2.MessageMeta{CRCExtra: 181, MinimumLength: 41, MaximumLength: 41},
	141:  mavlink2.MessageMeta{CRCExtra: 47, MinimumLength: 32, MaximumLength: 32},
	142:  mavlink2.MessageMeta{CRCExtra: 72, MinimumLength: 243, MaximumLength: 243},
	143:  mavlink2.MessageMeta{CRCExtra: 131, MinimumLength: 14, MaximumLength: 14},
	144:  mavlink2.MessageMeta{CRCExtra: 127, MinimumLength: 93, MaximumLength: 93},
	146:  mavlink2.MessageMeta{CRCExtra: 103, MinimumLength: 100, MaximumLength: 100},
	147:  mavlink2.MessageMeta{CRCExtra: 154, MinimumLength: 36, MaximumLength: 41},
	148:  mavlink2.MessageMeta{CRCExtra: 178, MinimumLength: 60, MaximumLength: 78},
	149:  mavlink2.MessageMeta{CRCExtra: 200, MinimumLength: 30, MaximumLength: 60},
	230:  mavlink2.MessageMeta{CRCExtra: 163, MinimumLength: 42, MaximumLength: 42},
	231:  mavlink2.MessageMeta{CRCExtra: 105, MinimumLength: 40, MaximumLength: 40},
	232:  mavlink2.MessageMeta{CRCExtra: 151, MinimumLength: 63, MaximumLength: 63},
	233:  mavlink2.MessageMeta{CRCExtra: 35, MinimumLength: 182, MaximumLength: 182},
	234:  mavlink2.MessageMeta{CRCExtra: 150, MinimumLength: 40, MaximumLength: 40},
	235:  mavlink2.MessageMeta{CRCExtra: 179, MinimumLength: 42, MaximumLength: 42},
	241:  mavlink2.MessageMeta{CRCExtra: 90, MinimumLength: 32, MaximumLength: 32},
	242:  mavlink2.MessageMeta{CRCExtra: 104, MinimumLength: 52, MaximumLength: 60},
	243:  mavlink2.MessageMeta{CRCExtra: 85, MinimumLength: 53, MaximumLength: 61},
	244:  mavlink2.MessageMeta{CRCExtra: 95, MinimumLength: 6, MaximumLength: 6},
	245:  mavlink2.MessageMeta{CRCExtra: 130, MinimumLength: 2, MaximumLength: 2},
	246:  mavlink2.MessageMeta{CRCExtra: 184, MinimumLength: 38, MaximumLength: 38},
	247:  mavlink2.MessageMeta{CRCExtra: 81, MinimumLength: 19, MaximumLength: 19},
	248:  mavlink2.MessageMeta{CRCExtra: 8, MinimumLength: 254, MaximumLength: 254},
	249:  mavlink2.MessageMeta{CRCExtra: 204, MinimumLength: 36, MaximumLength: 36},
	250:  mavlink2.MessageMeta{CRCExtra: 49, MinimumLength: 30, MaximumLength: 30},
	251:  mavlink2.MessageMeta{CRCExtra: 170, MinimumLength: 18, MaximumLength: 18},
	252:  mavlink2.MessageMeta{CRCExtra: 44, MinimumLength: 18, MaximumLength: 18},
	253:  mavlink2.MessageMeta{CRCExtra: 83, MinimumLength: 51, MaximumLength: 51},
	254:  mavlink2.MessageMeta{CRCExtra: 46, MinimumLength: 9, MaximumLength: 9},
	256:  mavlink2.MessageMeta{CRCExtra: 71, MinimumLength: 42, MaximumLength: 42},
	257:  mavlink2.MessageMeta{CRCExtra: 131, MinimumLength: 9, MaximumLength: 9},
	258:  mavlink2.MessageMeta{CRCExtra: 187, MinimumLength: 32, MaximumLength: 232},
	259:  mavlink2.MessageMeta{CRCExtra: 92, MinimumLength: 235, MaximumLength: 235},
	260:  mavlink2.MessageMeta{CRCExtra: 146, MinimumLength: 5, MaximumLength: 13},
	261:  mavlink2.MessageMeta{CRCExtra: 179, MinimumLength: 27, MaximumLength: 27},
	262:  mavlink2.MessageMeta{CRCExtra: 12, MinimumLength: 18, MaximumLength: 18},
	263:  mavlink2.MessageMeta{CRCExtra: 133, MinimumLength: 255, MaximumLength: 255},
	264:  mavlink2.MessageMeta{CRCExtra: 49, MinimumLength: 28, MaximumLength: 28},
	265:  mavlink2.MessageMeta{CRCExtra: 26, MinimumLength: 16, MaximumLength: 20},
	266:  mavlink2.MessageMeta{CRCExtra: 193, MinimumLength: 255, MaximumLength: 255},
	267:  mavlink2.MessageMeta{CRCExtra: 35, MinimumLength: 255, MaximumLength: 255},
	268:  mavlink2.MessageMeta{CRCExtra: 14, MinimumLength: 4, MaximumLength: 4},
	269:  mavlink2.MessageMeta{CRCExtra: 109, MinimumLength: 213, MaximumLength: 213},
	270:  mavlink2.MessageMeta{CRCExtra: 59, MinimumLength: 19, MaximumLength: 19},
	299:  mavlink2.MessageMeta{CRCExtra: 19, MinimumLength: 96, MaximumLength: 96},
	300:  mavlink2.MessageMeta{CRCExtra: 217, MinimumLength: 22, MaximumLength: 22},
	310:  mavlink2.MessageMeta{CRCExtra: 28, MinimumLength: 17, MaximumLength: 17},
	311:  mavlink2.MessageMeta{CRCExtra: 95, MinimumLength: 116, MaximumLength: 116},
	320:  mavlink2.MessageMeta{CRCExtra: 243, MinimumLength: 20, MaximumLength: 20},
	321:  mavlink2.MessageMeta{CRCExtra: 88, MinimumLength: 2, MaximumLength: 2},
	322:  mavlink2.MessageMeta{CRCExtra: 243, MinimumLength: 149, MaximumLength: 149},
	323:  mavlink2.MessageMeta{CRCExtra: 78, MinimumLength: 147, MaximumLength: 147},
	324:  mavlink2.MessageMeta{CRCExtra: 132, MinimumLength: 146, MaximumLength: 146},
	330:  mavlink2.MessageMeta{CRCExtra: 23, MinimumLength: 158, MaximumLength: 158},
	331:  mavlink2.MessageMeta{CRCExtra: 91, MinimumLength: 230, MaximumLength: 231},
	332:  mavlink2.MessageMeta{CRCExtra: 91, MinimumLength: 229, MaximumLength: 229},
	333:  mavlink2.MessageMeta{CRCExtra: 231, MinimumLength: 109, MaximumLength: 109},
	334:  mavlink2.MessageMeta{CRCExtra: 135, MinimumLength: 14, MaximumLength: 14},
	340:  mavlink2.MessageMeta{CRCExtra: 99, MinimumLength: 70, MaximumLength: 70},
	350:  mavlink2.MessageMeta{CRCExtra: 232, MinimumLength: 20, MaximumLength: 252},
	360:  mavlink2.MessageMeta{CRCExtra: 11, MinimumLength: 25, MaximumLength: 25},
	365:  mavlink2.MessageMeta{CRCExtra: 36, MinimumLength: 255, MaximumLength: 255},
	370:  mavlink2.MessageMeta{CRCExtra: 98, MinimumLength: 73, MaximumLength: 73},
	371:  mavlink2.MessageMeta{CRCExtra: 161, MinimumLength: 50, MaximumLength: 50},
	380:  mavlink2.MessageMeta{CRCExtra: 232, MinimumLength: 20, MaximumLength: 20},
	9000: mavlink2.MessageMeta{CRCExtra: 113, MinimumLength: 137, MaximumLength: 137},
}

var commonParsers = map[uint32]mavlink2.FrameParser{
	0: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Heartbeat{}

		err = message.Read(frame)

		return
	}),
	1: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SysStatus{}

		err = message.Read(frame)

		return
	}),
	2: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SystemTime{}

		err = message.Read(frame)

		return
	}),
	4: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Ping{}

		err = message.Read(frame)

		return
	}),
	5: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ChangeOperatorControl{}

		err = message.Read(frame)

		return
	}),
	6: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ChangeOperatorControlAck{}

		err = message.Read(frame)

		return
	}),
	7: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AuthKey{}

		err = message.Read(frame)

		return
	}),
	11: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetMode{}

		err = message.Read(frame)

		return
	}),
	20: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamRequestRead{}

		err = message.Read(frame)

		return
	}),
	21: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamRequestList{}

		err = message.Read(frame)

		return
	}),
	22: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamValue{}

		err = message.Read(frame)

		return
	}),
	23: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamSet{}

		err = message.Read(frame)

		return
	}),
	24: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSRawInt{}

		err = message.Read(frame)

		return
	}),
	25: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSStatus{}

		err = message.Read(frame)

		return
	}),
	26: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledIMU{}

		err = message.Read(frame)

		return
	}),
	27: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RawIMU{}

		err = message.Read(frame)

		return
	}),
	28: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RawPressure{}

		err = message.Read(frame)

		return
	}),
	29: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledPressure{}

		err = message.Read(frame)

		return
	}),
	30: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Attitude{}

		err = message.Read(frame)

		return
	}),
	31: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AttitudeQuaternion{}

		err = message.Read(frame)

		return
	}),
	32: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LocalPositionNed{}

		err = message.Read(frame)

		return
	}),
	33: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GlobalPositionInt{}

		err = message.Read(frame)

		return
	}),
	34: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RCChannelsScaled{}

		err = message.Read(frame)

		return
	}),
	35: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RCChannelsRaw{}

		err = message.Read(frame)

		return
	}),
	36: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ServoOutputRaw{}

		err = message.Read(frame)

		return
	}),
	37: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionRequestPartialList{}

		err = message.Read(frame)

		return
	}),
	38: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionWritePartialList{}

		err = message.Read(frame)

		return
	}),
	39: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionItem{}

		err = message.Read(frame)

		return
	}),
	40: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionRequest{}

		err = message.Read(frame)

		return
	}),
	41: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionSetCurrent{}

		err = message.Read(frame)

		return
	}),
	42: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionCurrent{}

		err = message.Read(frame)

		return
	}),
	43: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionRequestList{}

		err = message.Read(frame)

		return
	}),
	44: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionCount{}

		err = message.Read(frame)

		return
	}),
	45: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionClearAll{}

		err = message.Read(frame)

		return
	}),
	46: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionItemReached{}

		err = message.Read(frame)

		return
	}),
	47: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionAck{}

		err = message.Read(frame)

		return
	}),
	48: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetGPSGlobalOrigin{}

		err = message.Read(frame)

		return
	}),
	49: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSGlobalOrigin{}

		err = message.Read(frame)

		return
	}),
	50: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamMapRC{}

		err = message.Read(frame)

		return
	}),
	51: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionRequestInt{}

		err = message.Read(frame)

		return
	}),
	54: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SafetySetAllowedArea{}

		err = message.Read(frame)

		return
	}),
	55: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SafetyAllowedArea{}

		err = message.Read(frame)

		return
	}),
	61: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AttitudeQuaternionCov{}

		err = message.Read(frame)

		return
	}),
	62: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &NAVControllerOutput{}

		err = message.Read(frame)

		return
	}),
	63: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GlobalPositionIntCov{}

		err = message.Read(frame)

		return
	}),
	64: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LocalPositionNedCov{}

		err = message.Read(frame)

		return
	}),
	65: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RCChannels{}

		err = message.Read(frame)

		return
	}),
	66: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RequestDataStream{}

		err = message.Read(frame)

		return
	}),
	67: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DataStream{}

		err = message.Read(frame)

		return
	}),
	69: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ManualControl{}

		err = message.Read(frame)

		return
	}),
	70: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RCChannelsOverrIDe{}

		err = message.Read(frame)

		return
	}),
	73: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MissionItemInt{}

		err = message.Read(frame)

		return
	}),
	74: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VFRHUD{}

		err = message.Read(frame)

		return
	}),
	75: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CommandInt{}

		err = message.Read(frame)

		return
	}),
	76: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CommandLong{}

		err = message.Read(frame)

		return
	}),
	77: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CommandAck{}

		err = message.Read(frame)

		return
	}),
	81: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ManualSetpoint{}

		err = message.Read(frame)

		return
	}),
	82: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetAttitudeTarget{}

		err = message.Read(frame)

		return
	}),
	83: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AttitudeTarget{}

		err = message.Read(frame)

		return
	}),
	84: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetPositionTargetLocalNed{}

		err = message.Read(frame)

		return
	}),
	85: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PositionTargetLocalNed{}

		err = message.Read(frame)

		return
	}),
	86: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetPositionTargetGlobalInt{}

		err = message.Read(frame)

		return
	}),
	87: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PositionTargetGlobalInt{}

		err = message.Read(frame)

		return
	}),
	89: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LocalPositionNedSystemGlobalOffset{}

		err = message.Read(frame)

		return
	}),
	90: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilState{}

		err = message.Read(frame)

		return
	}),
	91: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilControls{}

		err = message.Read(frame)

		return
	}),
	92: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilRCInputsRaw{}

		err = message.Read(frame)

		return
	}),
	93: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilActuatorControls{}

		err = message.Read(frame)

		return
	}),
	100: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &OpticalFlow{}

		err = message.Read(frame)

		return
	}),
	101: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GlobalVisionPositionEstimate{}

		err = message.Read(frame)

		return
	}),
	102: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VisionPositionEstimate{}

		err = message.Read(frame)

		return
	}),
	103: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VisionSpeedEstimate{}

		err = message.Read(frame)

		return
	}),
	104: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ViconPositionEstimate{}

		err = message.Read(frame)

		return
	}),
	105: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HighresIMU{}

		err = message.Read(frame)

		return
	}),
	106: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &OpticalFlowRad{}

		err = message.Read(frame)

		return
	}),
	107: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilSensor{}

		err = message.Read(frame)

		return
	}),
	108: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SimState{}

		err = message.Read(frame)

		return
	}),
	109: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RadioStatus{}

		err = message.Read(frame)

		return
	}),
	110: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FileTransferProtocol{}

		err = message.Read(frame)

		return
	}),
	111: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Timesync{}

		err = message.Read(frame)

		return
	}),
	112: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraTrigger{}

		err = message.Read(frame)

		return
	}),
	113: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilGPS{}

		err = message.Read(frame)

		return
	}),
	114: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilOpticalFlow{}

		err = message.Read(frame)

		return
	}),
	115: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HilStateQuaternion{}

		err = message.Read(frame)

		return
	}),
	116: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledIMU2{}

		err = message.Read(frame)

		return
	}),
	117: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogRequestList{}

		err = message.Read(frame)

		return
	}),
	118: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogEntry{}

		err = message.Read(frame)

		return
	}),
	119: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogRequestData{}

		err = message.Read(frame)

		return
	}),
	120: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogData{}

		err = message.Read(frame)

		return
	}),
	121: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogErase{}

		err = message.Read(frame)

		return
	}),
	122: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LogRequestEnd{}

		err = message.Read(frame)

		return
	}),
	123: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSInjectData{}

		err = message.Read(frame)

		return
	}),
	124: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPS2Raw{}

		err = message.Read(frame)

		return
	}),
	125: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PowerStatus{}

		err = message.Read(frame)

		return
	}),
	126: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SerialControl{}

		err = message.Read(frame)

		return
	}),
	127: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSRtk{}

		err = message.Read(frame)

		return
	}),
	128: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPS2Rtk{}

		err = message.Read(frame)

		return
	}),
	129: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledIMU3{}

		err = message.Read(frame)

		return
	}),
	130: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DataTransmissionHandshake{}

		err = message.Read(frame)

		return
	}),
	131: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EncapsulatedData{}

		err = message.Read(frame)

		return
	}),
	132: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DistanceSensor{}

		err = message.Read(frame)

		return
	}),
	133: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TerrainRequest{}

		err = message.Read(frame)

		return
	}),
	134: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TerrainData{}

		err = message.Read(frame)

		return
	}),
	135: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TerrainCheck{}

		err = message.Read(frame)

		return
	}),
	136: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TerrainReport{}

		err = message.Read(frame)

		return
	}),
	137: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledPressure2{}

		err = message.Read(frame)

		return
	}),
	138: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AttPosMocap{}

		err = message.Read(frame)

		return
	}),
	139: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetActuatorControlTarget{}

		err = message.Read(frame)

		return
	}),
	140: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ActuatorControlTarget{}

		err = message.Read(frame)

		return
	}),
	141: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Altitude{}

		err = message.Read(frame)

		return
	}),
	142: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ResouRCeRequest{}

		err = message.Read(frame)

		return
	}),
	143: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ScaledPressure3{}

		err = message.Read(frame)

		return
	}),
	144: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FollowTarget{}

		err = message.Read(frame)

		return
	}),
	146: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ControlSystemState{}

		err = message.Read(frame)

		return
	}),
	147: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &BatteryStatus{}

		err = message.Read(frame)

		return
	}),
	148: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AutopilotVersion{}

		err = message.Read(frame)

		return
	}),
	149: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LandingTarget{}

		err = message.Read(frame)

		return
	}),
	230: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EstimatorStatus{}

		err = message.Read(frame)

		return
	}),
	231: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &WindCov{}

		err = message.Read(frame)

		return
	}),
	232: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSInput{}

		err = message.Read(frame)

		return
	}),
	233: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GPSRtcmData{}

		err = message.Read(frame)

		return
	}),
	234: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HighLatency{}

		err = message.Read(frame)

		return
	}),
	235: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HighLatency2{}

		err = message.Read(frame)

		return
	}),
	241: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Vibration{}

		err = message.Read(frame)

		return
	}),
	242: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HomePosition{}

		err = message.Read(frame)

		return
	}),
	243: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetHomePosition{}

		err = message.Read(frame)

		return
	}),
	244: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MessageInterval{}

		err = message.Read(frame)

		return
	}),
	245: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ExtendedSysState{}

		err = message.Read(frame)

		return
	}),
	246: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AdsbVehicle{}

		err = message.Read(frame)

		return
	}),
	247: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Collision{}

		err = message.Read(frame)

		return
	}),
	248: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &V2Extension{}

		err = message.Read(frame)

		return
	}),
	249: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MemoryVect{}

		err = message.Read(frame)

		return
	}),
	250: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DebugVect{}

		err = message.Read(frame)

		return
	}),
	251: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &NamedValueFloat{}

		err = message.Read(frame)

		return
	}),
	252: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &NamedValueInt{}

		err = message.Read(frame)

		return
	}),
	253: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Statustext{}

		err = message.Read(frame)

		return
	}),
	254: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Debug{}

		err = message.Read(frame)

		return
	}),
	256: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetupSigning{}

		err = message.Read(frame)

		return
	}),
	257: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ButtonChange{}

		err = message.Read(frame)

		return
	}),
	258: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PlayTune{}

		err = message.Read(frame)

		return
	}),
	259: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraInformation{}

		err = message.Read(frame)

		return
	}),
	260: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraSettings{}

		err = message.Read(frame)

		return
	}),
	261: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &StorageInformation{}

		err = message.Read(frame)

		return
	}),
	262: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraCaptureStatus{}

		err = message.Read(frame)

		return
	}),
	263: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraImageCaptured{}

		err = message.Read(frame)

		return
	}),
	264: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FlightInformation{}

		err = message.Read(frame)

		return
	}),
	265: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MountOrientation{}

		err = message.Read(frame)

		return
	}),
	266: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LoggingData{}

		err = message.Read(frame)

		return
	}),
	267: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LoggingDataAcked{}

		err = message.Read(frame)

		return
	}),
	268: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LoggingAck{}

		err = message.Read(frame)

		return
	}),
	269: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VIDeoStreamInformation{}

		err = message.Read(frame)

		return
	}),
	270: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VIDeoStreamStatus{}

		err = message.Read(frame)

		return
	}),
	299: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &WifiConfigAp{}

		err = message.Read(frame)

		return
	}),
	300: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ProtocolVersion{}

		err = message.Read(frame)

		return
	}),
	310: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &UavcanNodeStatus{}

		err = message.Read(frame)

		return
	}),
	311: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &UavcanNodeInfo{}

		err = message.Read(frame)

		return
	}),
	320: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamExtRequestRead{}

		err = message.Read(frame)

		return
	}),
	321: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamExtRequestList{}

		err = message.Read(frame)

		return
	}),
	322: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamExtValue{}

		err = message.Read(frame)

		return
	}),
	323: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamExtSet{}

		err = message.Read(frame)

		return
	}),
	324: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ParamExtAck{}

		err = message.Read(frame)

		return
	}),
	330: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ObstacleDistance{}

		err = message.Read(frame)

		return
	}),
	331: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Odometry{}

		err = message.Read(frame)

		return
	}),
	332: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TrajectoryRepresentationWaypoints{}

		err = message.Read(frame)

		return
	}),
	333: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TrajectoryRepresentationBezier{}

		err = message.Read(frame)

		return
	}),
	334: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CellularStatus{}

		err = message.Read(frame)

		return
	}),
	340: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &UtmGlobalPosition{}

		err = message.Read(frame)

		return
	}),
	350: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DebugFloatArray{}

		err = message.Read(frame)

		return
	}),
	360: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &OrbitExecutionStatus{}

		err = message.Read(frame)

		return
	}),
	365: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &StatustextLong{}

		err = message.Read(frame)

		return
	}),
	370: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SmartBatteryInfo{}

		err = message.Read(frame)

		return
	}),
	371: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SmartBatteryStatus{}

		err = message.Read(frame)

		return
	}),
	380: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &TimeEstimateToTarget{}

		err = message.Read(frame)

		return
	}),
	9000: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &WheelDistance{}

		err = message.Read(frame)

		return
	}),
}
