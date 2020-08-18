package ardupilotmega

import "github.com/botlink/go-mavlink2"

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
	return "ardupilotmega"
}

// GetMeta retrieves the metadata for the message. If no metadata is found, ErrUnknownMessage is returned
func (d Dialect) GetMeta(messageID uint32) (meta mavlink2.MessageMeta, err error) {
	var ok bool

	if meta, ok = ardupilotmegaMessages[messageID]; !ok {
		err = mavlink2.ErrUnknownMessage
	}

	return
}

// GetMessage extracts and parses the message contained in the Frame
func (d Dialect) GetMessage(frame mavlink2.Frame) (message mavlink2.Message, err error) {
	if parser, ok := ardupilotmegaParsers[frame.GetMessageID()]; ok {
		message, err = parser(frame)
		return
	}

	err = mavlink2.ErrUnknownMessage
	return
}

var ardupilotmegaMessages = map[uint32]mavlink2.MessageMeta{
	150:   mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 42, MaximumLength: 42},
	151:   mavlink2.MessageMeta{CRCExtra: 219, MinimumLength: 8, MaximumLength: 8},
	152:   mavlink2.MessageMeta{CRCExtra: 208, MinimumLength: 4, MaximumLength: 8},
	153:   mavlink2.MessageMeta{CRCExtra: 188, MinimumLength: 12, MaximumLength: 12},
	154:   mavlink2.MessageMeta{CRCExtra: 84, MinimumLength: 15, MaximumLength: 15},
	155:   mavlink2.MessageMeta{CRCExtra: 22, MinimumLength: 13, MaximumLength: 13},
	156:   mavlink2.MessageMeta{CRCExtra: 19, MinimumLength: 6, MaximumLength: 6},
	157:   mavlink2.MessageMeta{CRCExtra: 21, MinimumLength: 15, MaximumLength: 15},
	158:   mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 14, MaximumLength: 14},
	160:   mavlink2.MessageMeta{CRCExtra: 78, MinimumLength: 12, MaximumLength: 12},
	161:   mavlink2.MessageMeta{CRCExtra: 68, MinimumLength: 3, MaximumLength: 3},
	163:   mavlink2.MessageMeta{CRCExtra: 127, MinimumLength: 28, MaximumLength: 28},
	164:   mavlink2.MessageMeta{CRCExtra: 154, MinimumLength: 44, MaximumLength: 44},
	165:   mavlink2.MessageMeta{CRCExtra: 21, MinimumLength: 3, MaximumLength: 3},
	166:   mavlink2.MessageMeta{CRCExtra: 21, MinimumLength: 9, MaximumLength: 9},
	167:   mavlink2.MessageMeta{CRCExtra: 144, MinimumLength: 22, MaximumLength: 22},
	168:   mavlink2.MessageMeta{CRCExtra: 1, MinimumLength: 12, MaximumLength: 12},
	169:   mavlink2.MessageMeta{CRCExtra: 234, MinimumLength: 18, MaximumLength: 18},
	170:   mavlink2.MessageMeta{CRCExtra: 73, MinimumLength: 34, MaximumLength: 34},
	171:   mavlink2.MessageMeta{CRCExtra: 181, MinimumLength: 66, MaximumLength: 66},
	172:   mavlink2.MessageMeta{CRCExtra: 22, MinimumLength: 98, MaximumLength: 98},
	173:   mavlink2.MessageMeta{CRCExtra: 83, MinimumLength: 8, MaximumLength: 8},
	174:   mavlink2.MessageMeta{CRCExtra: 167, MinimumLength: 48, MaximumLength: 48},
	175:   mavlink2.MessageMeta{CRCExtra: 138, MinimumLength: 19, MaximumLength: 19},
	176:   mavlink2.MessageMeta{CRCExtra: 234, MinimumLength: 3, MaximumLength: 3},
	177:   mavlink2.MessageMeta{CRCExtra: 240, MinimumLength: 20, MaximumLength: 20},
	178:   mavlink2.MessageMeta{CRCExtra: 47, MinimumLength: 24, MaximumLength: 24},
	179:   mavlink2.MessageMeta{CRCExtra: 189, MinimumLength: 29, MaximumLength: 29},
	180:   mavlink2.MessageMeta{CRCExtra: 52, MinimumLength: 45, MaximumLength: 47},
	181:   mavlink2.MessageMeta{CRCExtra: 174, MinimumLength: 4, MaximumLength: 4},
	182:   mavlink2.MessageMeta{CRCExtra: 229, MinimumLength: 40, MaximumLength: 40},
	183:   mavlink2.MessageMeta{CRCExtra: 85, MinimumLength: 2, MaximumLength: 2},
	184:   mavlink2.MessageMeta{CRCExtra: 159, MinimumLength: 206, MaximumLength: 206},
	185:   mavlink2.MessageMeta{CRCExtra: 186, MinimumLength: 7, MaximumLength: 7},
	186:   mavlink2.MessageMeta{CRCExtra: 72, MinimumLength: 29, MaximumLength: 29},
	191:   mavlink2.MessageMeta{CRCExtra: 92, MinimumLength: 27, MaximumLength: 27},
	192:   mavlink2.MessageMeta{CRCExtra: 36, MinimumLength: 44, MaximumLength: 54},
	193:   mavlink2.MessageMeta{CRCExtra: 71, MinimumLength: 22, MaximumLength: 26},
	194:   mavlink2.MessageMeta{CRCExtra: 98, MinimumLength: 25, MaximumLength: 25},
	195:   mavlink2.MessageMeta{CRCExtra: 120, MinimumLength: 37, MaximumLength: 37},
	200:   mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 42, MaximumLength: 42},
	201:   mavlink2.MessageMeta{CRCExtra: 205, MinimumLength: 14, MaximumLength: 14},
	214:   mavlink2.MessageMeta{CRCExtra: 69, MinimumLength: 8, MaximumLength: 8},
	215:   mavlink2.MessageMeta{CRCExtra: 101, MinimumLength: 3, MaximumLength: 3},
	216:   mavlink2.MessageMeta{CRCExtra: 50, MinimumLength: 3, MaximumLength: 3},
	217:   mavlink2.MessageMeta{CRCExtra: 202, MinimumLength: 6, MaximumLength: 6},
	218:   mavlink2.MessageMeta{CRCExtra: 17, MinimumLength: 7, MaximumLength: 7},
	219:   mavlink2.MessageMeta{CRCExtra: 162, MinimumLength: 2, MaximumLength: 2},
	225:   mavlink2.MessageMeta{CRCExtra: 142, MinimumLength: 53, MaximumLength: 53},
	226:   mavlink2.MessageMeta{CRCExtra: 207, MinimumLength: 8, MaximumLength: 8},
	11000: mavlink2.MessageMeta{CRCExtra: 134, MinimumLength: 51, MaximumLength: 51},
	11001: mavlink2.MessageMeta{CRCExtra: 15, MinimumLength: 135, MaximumLength: 135},
	11002: mavlink2.MessageMeta{CRCExtra: 234, MinimumLength: 179, MaximumLength: 179},
	11003: mavlink2.MessageMeta{CRCExtra: 64, MinimumLength: 5, MaximumLength: 5},
	11010: mavlink2.MessageMeta{CRCExtra: 46, MinimumLength: 49, MaximumLength: 49},
	11011: mavlink2.MessageMeta{CRCExtra: 106, MinimumLength: 44, MaximumLength: 44},
	11020: mavlink2.MessageMeta{CRCExtra: 205, MinimumLength: 16, MaximumLength: 16},
	11030: mavlink2.MessageMeta{CRCExtra: 144, MinimumLength: 44, MaximumLength: 44},
	11031: mavlink2.MessageMeta{CRCExtra: 133, MinimumLength: 44, MaximumLength: 44},
	11032: mavlink2.MessageMeta{CRCExtra: 85, MinimumLength: 44, MaximumLength: 44},
}

var ardupilotmegaParsers = map[uint32]mavlink2.FrameParser{
	150: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SensorOffsets{}

		err = message.Read(frame)

		return
	}),
	151: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SetMagOffsets{}

		err = message.Read(frame)

		return
	}),
	152: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MemInfo{}

		err = message.Read(frame)

		return
	}),
	153: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &ApAdc{}

		err = message.Read(frame)

		return
	}),
	154: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DigicamConfigure{}

		err = message.Read(frame)

		return
	}),
	155: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DigicamControl{}

		err = message.Read(frame)

		return
	}),
	156: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MountConfigure{}

		err = message.Read(frame)

		return
	}),
	157: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MountControl{}

		err = message.Read(frame)

		return
	}),
	158: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MountStatus{}

		err = message.Read(frame)

		return
	}),
	160: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FencePoint{}

		err = message.Read(frame)

		return
	}),
	161: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &FenceFetchPoint{}

		err = message.Read(frame)

		return
	}),
	163: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AHRS{}

		err = message.Read(frame)

		return
	}),
	164: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &SimState{}

		err = message.Read(frame)

		return
	}),
	165: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &HWstatus{}

		err = message.Read(frame)

		return
	}),
	166: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Radio{}

		err = message.Read(frame)

		return
	}),
	167: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LimitsStatus{}

		err = message.Read(frame)

		return
	}),
	168: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Wind{}

		err = message.Read(frame)

		return
	}),
	169: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Data16{}

		err = message.Read(frame)

		return
	}),
	170: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Data32{}

		err = message.Read(frame)

		return
	}),
	171: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Data64{}

		err = message.Read(frame)

		return
	}),
	172: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Data96{}

		err = message.Read(frame)

		return
	}),
	173: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Rangefinder{}

		err = message.Read(frame)

		return
	}),
	174: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AirspeedAutocal{}

		err = message.Read(frame)

		return
	}),
	175: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RallyPoint{}

		err = message.Read(frame)

		return
	}),
	176: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RallyFetchPoint{}

		err = message.Read(frame)

		return
	}),
	177: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CompassmotStatus{}

		err = message.Read(frame)

		return
	}),
	178: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AHRS2{}

		err = message.Read(frame)

		return
	}),
	179: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraStatus{}

		err = message.Read(frame)

		return
	}),
	180: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &CameraFeedback{}

		err = message.Read(frame)

		return
	}),
	181: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Battery2{}

		err = message.Read(frame)

		return
	}),
	182: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AHRS3{}

		err = message.Read(frame)

		return
	}),
	183: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AutopilotVersionRequest{}

		err = message.Read(frame)

		return
	}),
	184: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RemoteLogDataBlock{}

		err = message.Read(frame)

		return
	}),
	185: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &RemoteLogBlockStatus{}

		err = message.Read(frame)

		return
	}),
	186: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &LedControl{}

		err = message.Read(frame)

		return
	}),
	191: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MagCalProgress{}

		err = message.Read(frame)

		return
	}),
	192: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &MagCalReport{}

		err = message.Read(frame)

		return
	}),
	193: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EKFStatusReport{}

		err = message.Read(frame)

		return
	}),
	194: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &PIDTuning{}

		err = message.Read(frame)

		return
	}),
	195: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Deepstall{}

		err = message.Read(frame)

		return
	}),
	200: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GimbalReport{}

		err = message.Read(frame)

		return
	}),
	201: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GimbalControl{}

		err = message.Read(frame)

		return
	}),
	214: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GimbalTorqueCmdReport{}

		err = message.Read(frame)

		return
	}),
	215: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GoproHeartbeat{}

		err = message.Read(frame)

		return
	}),
	216: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GoproGetRequest{}

		err = message.Read(frame)

		return
	}),
	217: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GoproGetResponse{}

		err = message.Read(frame)

		return
	}),
	218: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GoproSetRequest{}

		err = message.Read(frame)

		return
	}),
	219: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &GoproSetResponse{}

		err = message.Read(frame)

		return
	}),
	225: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EfiStatus{}

		err = message.Read(frame)

		return
	}),
	226: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &Rpm{}

		err = message.Read(frame)

		return
	}),
	11000: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DeviceOpRead{}

		err = message.Read(frame)

		return
	}),
	11001: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DeviceOpReadReply{}

		err = message.Read(frame)

		return
	}),
	11002: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DeviceOpWrite{}

		err = message.Read(frame)

		return
	}),
	11003: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &DeviceOpWriteReply{}

		err = message.Read(frame)

		return
	}),
	11010: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AdapTuning{}

		err = message.Read(frame)

		return
	}),
	11011: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &VisionPositionDelta{}

		err = message.Read(frame)

		return
	}),
	11020: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &AoaSsa{}

		err = message.Read(frame)

		return
	}),
	11030: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EscTelemetry1To4{}

		err = message.Read(frame)

		return
	}),
	11031: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EscTelemetry5To8{}

		err = message.Read(frame)

		return
	}),
	11032: mavlink2.FrameParser(func(frame mavlink2.Frame) (message mavlink2.Message, err error) {
		message = &EscTelemetry9To12{}

		err = message.Read(frame)

		return
	}),
}
