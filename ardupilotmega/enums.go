package ardupilotmega

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

/*ACCELCAL_VEHICLE_POS -  */
const (
	/*AccelcalVehiclePosLevel -  */
	AccelcalVehiclePosLevel = 1
	/*AccelcalVehiclePosLeft -  */
	AccelcalVehiclePosLeft = 2
	/*AccelcalVehiclePosRight -  */
	AccelcalVehiclePosRight = 3
	/*AccelcalVehiclePosNosedown -  */
	AccelcalVehiclePosNosedown = 4
	/*AccelcalVehiclePosNoseup -  */
	AccelcalVehiclePosNoseup = 5
	/*AccelcalVehiclePosBack -  */
	AccelcalVehiclePosBack = 6
	/*AccelcalVehiclePosSuccess -  */
	AccelcalVehiclePosSuccess = 16777215
	/*AccelcalVehiclePosFailed -  */
	AccelcalVehiclePosFailed = 16777216
	/*AccelcalVehiclePosEnumEnd -  */
	AccelcalVehiclePosEnumEnd = 16777217
)

/*LIMITS_STATE -  */
const (
	/*LimitsInit - Pre-initialization. */
	LimitsInit = 0
	/*LimitsDisabled - Disabled. */
	LimitsDisabled = 1
	/*LimitsEnabled - Checking limits. */
	LimitsEnabled = 2
	/*LimitsTriggered - A limit has been breached. */
	LimitsTriggered = 3
	/*LimitsRecovering - Taking action e.g. Return/RTL. */
	LimitsRecovering = 4
	/*LimitsRecovered - We're no longer in breach of a limit. */
	LimitsRecovered = 5
	/*LimitsStateEnumEnd -  */
	LimitsStateEnumEnd = 6
)

/*LIMIT_MODULE -  */
const (
	/*LimitGPSlock - Pre-initialization. */
	LimitGPSlock = 1
	/*LimitGeofence - Disabled. */
	LimitGeofence = 2
	/*LimitAltitude - Checking limits. */
	LimitAltitude = 4
	/*LimitModuleEnumEnd -  */
	LimitModuleEnumEnd = 5
)

/*RALLY_FLAGS - Flags in RALLY_POINT message. */
const (
	/*FavorableWind - Flag set when requiring favorable winds for landing. */
	FavorableWind = 1
	/*LandImmediately - Flag set when plane is to immediately descend to break altitude and land without GCS intervention. Flag not set when plane is to loiter at Rally point until commanded to land. */
	LandImmediately = 2
	/*RallyFlagsEnumEnd -  */
	RallyFlagsEnumEnd = 3
)

/*PARACHUTE_ACTION -  */
const (
	/*ParachuteDisable - Disable parachute release. */
	ParachuteDisable = 0
	/*ParachuteEnable - Enable parachute release. */
	ParachuteEnable = 1
	/*ParachuteRelease - Release parachute. */
	ParachuteRelease = 2
	/*ParachuteActionEnumEnd -  */
	ParachuteActionEnumEnd = 3
)

/*GRIPPER_ACTIONS - Gripper actions. */
const (
	/*GripperActionRelease - Gripper release cargo. */
	GripperActionRelease = 0
	/*GripperActionGrab - Gripper grab onto cargo. */
	GripperActionGrab = 1
	/*GripperActionsEnumEnd -  */
	GripperActionsEnumEnd = 2
)

/*WINCH_ACTIONS - Winch actions. */
const (
	/*WinchRelaxed - Relax winch. */
	WinchRelaxed = 0
	/*WinchRelativeLengthControl - Winch unwinds or winds specified length of cable optionally using specified rate. */
	WinchRelativeLengthControl = 1
	/*WinchRateControl - Winch unwinds or winds cable at specified rate in meters/seconds. */
	WinchRateControl = 2
	/*WinchActionsEnumEnd -  */
	WinchActionsEnumEnd = 3
)

/*CAMERA_STATUS_TYPES -  */
const (
	/*CameraStatusTypeHeartbeat - Camera heartbeat, announce camera component ID at 1Hz. */
	CameraStatusTypeHeartbeat = 0
	/*CameraStatusTypeTrigger - Camera image triggered. */
	CameraStatusTypeTrigger = 1
	/*CameraStatusTypeDisconnect - Camera connection lost. */
	CameraStatusTypeDisconnect = 2
	/*CameraStatusTypeError - Camera unknown error. */
	CameraStatusTypeError = 3
	/*CameraStatusTypeLowbatt - Camera battery low. Parameter p1 shows reported voltage. */
	CameraStatusTypeLowbatt = 4
	/*CameraStatusTypeLowstore - Camera storage low. Parameter p1 shows reported shots remaining. */
	CameraStatusTypeLowstore = 5
	/*CameraStatusTypeLowstorev - Camera storage low. Parameter p1 shows reported video minutes remaining. */
	CameraStatusTypeLowstorev = 6
	/*CameraStatusTypesEnumEnd -  */
	CameraStatusTypesEnumEnd = 7
)

/*CAMERA_FEEDBACK_FLAGS -  */
const (
	/*CameraFeedbackPhoto - Shooting photos, not video. */
	CameraFeedbackPhoto = 0
	/*CameraFeedbackVIDeo - Shooting video, not stills. */
	CameraFeedbackVIDeo = 1
	/*CameraFeedbackBadexposure - Unable to achieve requested exposure (e.g. shutter speed too low). */
	CameraFeedbackBadexposure = 2
	/*CameraFeedbackClosedloop - Closed loop feedback from camera, we know for sure it has successfully taken a picture. */
	CameraFeedbackClosedloop = 3
	/*CameraFeedbackOpenloop - Open loop camera, an image trigger has been requested but we can't know for sure it has successfully taken a picture. */
	CameraFeedbackOpenloop = 4
	/*CameraFeedbackFlagsEnumEnd -  */
	CameraFeedbackFlagsEnumEnd = 5
)

/*MAV_MODE_GIMBAL -  */
const (
	/*MavModeGimbalUninitialized - Gimbal is powered on but has not started initializing yet. */
	MavModeGimbalUninitialized = 0
	/*MavModeGimbalCalibratingPitch - Gimbal is currently running calibration on the pitch axis. */
	MavModeGimbalCalibratingPitch = 1
	/*MavModeGimbalCalibratingRoll - Gimbal is currently running calibration on the roll axis. */
	MavModeGimbalCalibratingRoll = 2
	/*MavModeGimbalCalibratingYaw - Gimbal is currently running calibration on the yaw axis. */
	MavModeGimbalCalibratingYaw = 3
	/*MavModeGimbalInitialized - Gimbal has finished calibrating and initializing, but is relaxed pending reception of first rate command from copter. */
	MavModeGimbalInitialized = 4
	/*MavModeGimbalActive - Gimbal is actively stabilizing. */
	MavModeGimbalActive = 5
	/*MavModeGimbalRateCmdTimeout - Gimbal is relaxed because it missed more than 10 expected rate command messages in a row. Gimbal will move back to active mode when it receives a new rate command. */
	MavModeGimbalRateCmdTimeout = 6
	/*MavModeGimbalEnumEnd -  */
	MavModeGimbalEnumEnd = 7
)

/*GIMBAL_AXIS -  */
const (
	/*GimbalAxisYaw - Gimbal yaw axis. */
	GimbalAxisYaw = 0
	/*GimbalAxisPitch - Gimbal pitch axis. */
	GimbalAxisPitch = 1
	/*GimbalAxisRoll - Gimbal roll axis. */
	GimbalAxisRoll = 2
	/*GimbalAxisEnumEnd -  */
	GimbalAxisEnumEnd = 3
)

/*GIMBAL_AXIS_CALIBRATION_STATUS -  */
const (
	/*GimbalAxisCalibrationStatusInProgress - Axis calibration is in progress. */
	GimbalAxisCalibrationStatusInProgress = 0
	/*GimbalAxisCalibrationStatusSucceeded - Axis calibration succeeded. */
	GimbalAxisCalibrationStatusSucceeded = 1
	/*GimbalAxisCalibrationStatusFailed - Axis calibration failed. */
	GimbalAxisCalibrationStatusFailed = 2
	/*GimbalAxisCalibrationStatusEnumEnd -  */
	GimbalAxisCalibrationStatusEnumEnd = 3
)

/*GIMBAL_AXIS_CALIBRATION_REQUIRED -  */
const (
	/*GimbalAxisCalibrationRequiredUnknown - Whether or not this axis requires calibration is unknown at this time. */
	GimbalAxisCalibrationRequiredUnknown = 0
	/*GimbalAxisCalibrationRequiredTrue - This axis requires calibration. */
	GimbalAxisCalibrationRequiredTrue = 1
	/*GimbalAxisCalibrationRequiredFalse - This axis does not require calibration. */
	GimbalAxisCalibrationRequiredFalse = 2
	/*GimbalAxisCalibrationRequiredEnumEnd -  */
	GimbalAxisCalibrationRequiredEnumEnd = 3
)

/*GOPRO_HEARTBEAT_STATUS -  */
const (
	/*GoproHeartbeatStatusDisconnected - No GoPro connected. */
	GoproHeartbeatStatusDisconnected = 0
	/*GoproHeartbeatStatusIncompatible - The detected GoPro is not HeroBus compatible. */
	GoproHeartbeatStatusIncompatible = 1
	/*GoproHeartbeatStatusConnected - A HeroBus compatible GoPro is connected. */
	GoproHeartbeatStatusConnected = 2
	/*GoproHeartbeatStatusError - An unrecoverable error was encountered with the connected GoPro, it may require a power cycle. */
	GoproHeartbeatStatusError = 3
	/*GoproHeartbeatStatusEnumEnd -  */
	GoproHeartbeatStatusEnumEnd = 4
)

/*GOPRO_HEARTBEAT_FLAGS -  */
const (
	/*GoproFlagRecording - GoPro is currently recording. */
	GoproFlagRecording = 1
	/*GoproHeartbeatFlagsEnumEnd -  */
	GoproHeartbeatFlagsEnumEnd = 2
)

/*GOPRO_REQUEST_STATUS -  */
const (
	/*GoproRequestSuccess - The write message with ID indicated succeeded. */
	GoproRequestSuccess = 0
	/*GoproRequestFailed - The write message with ID indicated failed. */
	GoproRequestFailed = 1
	/*GoproRequestStatusEnumEnd -  */
	GoproRequestStatusEnumEnd = 2
)

/*GOPRO_COMMAND -  */
const (
	/*GoproCommandPower - (Get/Set). */
	GoproCommandPower = 0
	/*GoproCommandCaptureMode - (Get/Set). */
	GoproCommandCaptureMode = 1
	/*GoproCommandShutter - (___/Set). */
	GoproCommandShutter = 2
	/*GoproCommandBattery - (Get/___). */
	GoproCommandBattery = 3
	/*GoproCommandModel - (Get/___). */
	GoproCommandModel = 4
	/*GoproCommandVIDeoSettings - (Get/Set). */
	GoproCommandVIDeoSettings = 5
	/*GoproCommandLowLight - (Get/Set). */
	GoproCommandLowLight = 6
	/*GoproCommandPhotoResolution - (Get/Set). */
	GoproCommandPhotoResolution = 7
	/*GoproCommandPhotoBurstRate - (Get/Set). */
	GoproCommandPhotoBurstRate = 8
	/*GoproCommandProtune - (Get/Set). */
	GoproCommandProtune = 9
	/*GoproCommandProtuneWhiteBalance - (Get/Set) Hero 3+ Only. */
	GoproCommandProtuneWhiteBalance = 10
	/*GoproCommandProtuneColour - (Get/Set) Hero 3+ Only. */
	GoproCommandProtuneColour = 11
	/*GoproCommandProtuneGain - (Get/Set) Hero 3+ Only. */
	GoproCommandProtuneGain = 12
	/*GoproCommandProtuneSharpness - (Get/Set) Hero 3+ Only. */
	GoproCommandProtuneSharpness = 13
	/*GoproCommandProtuneExposure - (Get/Set) Hero 3+ Only. */
	GoproCommandProtuneExposure = 14
	/*GoproCommandTime - (Get/Set). */
	GoproCommandTime = 15
	/*GoproCommandCharging - (Get/Set). */
	GoproCommandCharging = 16
	/*GoproCommandEnumEnd -  */
	GoproCommandEnumEnd = 17
)

/*GOPRO_CAPTURE_MODE -  */
const (
	/*GoproCaptureModeVIDeo - Video mode. */
	GoproCaptureModeVIDeo = 0
	/*GoproCaptureModePhoto - Photo mode. */
	GoproCaptureModePhoto = 1
	/*GoproCaptureModeBurst - Burst mode, Hero 3+ only. */
	GoproCaptureModeBurst = 2
	/*GoproCaptureModeTimeLapse - Time lapse mode, Hero 3+ only. */
	GoproCaptureModeTimeLapse = 3
	/*GoproCaptureModeMultiShot - Multi shot mode, Hero 4 only. */
	GoproCaptureModeMultiShot = 4
	/*GoproCaptureModePlayback - Playback mode, Hero 4 only, silver only except when LCD or HDMI is connected to black. */
	GoproCaptureModePlayback = 5
	/*GoproCaptureModeSetup - Playback mode, Hero 4 only. */
	GoproCaptureModeSetup = 6
	/*GoproCaptureModeUnknown - Mode not yet known. */
	GoproCaptureModeUnknown = 255
	/*GoproCaptureModeEnumEnd -  */
	GoproCaptureModeEnumEnd = 256
)

/*GOPRO_RESOLUTION -  */
const (
	/*GoproResolution480P - 848 x 480 (480p). */
	GoproResolution480P = 0
	/*GoproResolution720P - 1280 x 720 (720p). */
	GoproResolution720P = 1
	/*GoproResolution960P - 1280 x 960 (960p). */
	GoproResolution960P = 2
	/*GoproResolution1080P - 1920 x 1080 (1080p). */
	GoproResolution1080P = 3
	/*GoproResolution1440P - 1920 x 1440 (1440p). */
	GoproResolution1440P = 4
	/*GoproResolution27K179 - 2704 x 1440 (2.7k-17:9). */
	GoproResolution27K179 = 5
	/*GoproResolution27K169 - 2704 x 1524 (2.7k-16:9). */
	GoproResolution27K169 = 6
	/*GoproResolution27K43 - 2704 x 2028 (2.7k-4:3). */
	GoproResolution27K43 = 7
	/*GoproResolution4K169 - 3840 x 2160 (4k-16:9). */
	GoproResolution4K169 = 8
	/*GoproResolution4K179 - 4096 x 2160 (4k-17:9). */
	GoproResolution4K179 = 9
	/*GoproResolution720PSuperview - 1280 x 720 (720p-SuperView). */
	GoproResolution720PSuperview = 10
	/*GoproResolution1080PSuperview - 1920 x 1080 (1080p-SuperView). */
	GoproResolution1080PSuperview = 11
	/*GoproResolution27KSuperview - 2704 x 1520 (2.7k-SuperView). */
	GoproResolution27KSuperview = 12
	/*GoproResolution4KSuperview - 3840 x 2160 (4k-SuperView). */
	GoproResolution4KSuperview = 13
	/*GoproResolutionEnumEnd -  */
	GoproResolutionEnumEnd = 14
)

/*GOPRO_FRAME_RATE -  */
const (
	/*GoproFrameRate12 - 12 FPS. */
	GoproFrameRate12 = 0
	/*GoproFrameRate15 - 15 FPS. */
	GoproFrameRate15 = 1
	/*GoproFrameRate24 - 24 FPS. */
	GoproFrameRate24 = 2
	/*GoproFrameRate25 - 25 FPS. */
	GoproFrameRate25 = 3
	/*GoproFrameRate30 - 30 FPS. */
	GoproFrameRate30 = 4
	/*GoproFrameRate48 - 48 FPS. */
	GoproFrameRate48 = 5
	/*GoproFrameRate50 - 50 FPS. */
	GoproFrameRate50 = 6
	/*GoproFrameRate60 - 60 FPS. */
	GoproFrameRate60 = 7
	/*GoproFrameRate80 - 80 FPS. */
	GoproFrameRate80 = 8
	/*GoproFrameRate90 - 90 FPS. */
	GoproFrameRate90 = 9
	/*GoproFrameRate100 - 100 FPS. */
	GoproFrameRate100 = 10
	/*GoproFrameRate120 - 120 FPS. */
	GoproFrameRate120 = 11
	/*GoproFrameRate240 - 240 FPS. */
	GoproFrameRate240 = 12
	/*GoproFrameRate125 - 12.5 FPS. */
	GoproFrameRate125 = 13
	/*GoproFrameRateEnumEnd -  */
	GoproFrameRateEnumEnd = 14
)

/*GOPRO_FIELD_OF_VIEW -  */
const (
	/*GoproFieldOfViewWIDe - 0x00: Wide. */
	GoproFieldOfViewWIDe = 0
	/*GoproFieldOfViewMedium - 0x01: Medium. */
	GoproFieldOfViewMedium = 1
	/*GoproFieldOfViewNarrow - 0x02: Narrow. */
	GoproFieldOfViewNarrow = 2
	/*GoproFieldOfViewEnumEnd -  */
	GoproFieldOfViewEnumEnd = 3
)

/*GOPRO_VIDEO_SETTINGS_FLAGS -  */
const (
	/*GoproVIDeoSettingsTvMode - 0=NTSC, 1=PAL. */
	GoproVIDeoSettingsTvMode = 1
	/*GoproVIDeoSettingsFlagsEnumEnd -  */
	GoproVIDeoSettingsFlagsEnumEnd = 2
)

/*GOPRO_PHOTO_RESOLUTION -  */
const (
	/*GoproPhotoResolution5MpMedium - 5MP Medium. */
	GoproPhotoResolution5MpMedium = 0
	/*GoproPhotoResolution7MpMedium - 7MP Medium. */
	GoproPhotoResolution7MpMedium = 1
	/*GoproPhotoResolution7MpWIDe - 7MP Wide. */
	GoproPhotoResolution7MpWIDe = 2
	/*GoproPhotoResolution10MpWIDe - 10MP Wide. */
	GoproPhotoResolution10MpWIDe = 3
	/*GoproPhotoResolution12MpWIDe - 12MP Wide. */
	GoproPhotoResolution12MpWIDe = 4
	/*GoproPhotoResolutionEnumEnd -  */
	GoproPhotoResolutionEnumEnd = 5
)

/*GOPRO_PROTUNE_WHITE_BALANCE -  */
const (
	/*GoproProtuneWhiteBalanceAuto - Auto. */
	GoproProtuneWhiteBalanceAuto = 0
	/*GoproProtuneWhiteBalance3000K - 3000K. */
	GoproProtuneWhiteBalance3000K = 1
	/*GoproProtuneWhiteBalance5500K - 5500K. */
	GoproProtuneWhiteBalance5500K = 2
	/*GoproProtuneWhiteBalance6500K - 6500K. */
	GoproProtuneWhiteBalance6500K = 3
	/*GoproProtuneWhiteBalanceRaw - Camera Raw. */
	GoproProtuneWhiteBalanceRaw = 4
	/*GoproProtuneWhiteBalanceEnumEnd -  */
	GoproProtuneWhiteBalanceEnumEnd = 5
)

/*GOPRO_PROTUNE_COLOUR -  */
const (
	/*GoproProtuneColourStandard - Auto. */
	GoproProtuneColourStandard = 0
	/*GoproProtuneColourNeutral - Neutral. */
	GoproProtuneColourNeutral = 1
	/*GoproProtuneColourEnumEnd -  */
	GoproProtuneColourEnumEnd = 2
)

/*GOPRO_PROTUNE_GAIN -  */
const (
	/*GoproProtuneGain400 - ISO 400. */
	GoproProtuneGain400 = 0
	/*GoproProtuneGain800 - ISO 800 (Only Hero 4). */
	GoproProtuneGain800 = 1
	/*GoproProtuneGain1600 - ISO 1600. */
	GoproProtuneGain1600 = 2
	/*GoproProtuneGain3200 - ISO 3200 (Only Hero 4). */
	GoproProtuneGain3200 = 3
	/*GoproProtuneGain6400 - ISO 6400. */
	GoproProtuneGain6400 = 4
	/*GoproProtuneGainEnumEnd -  */
	GoproProtuneGainEnumEnd = 5
)

/*GOPRO_PROTUNE_SHARPNESS -  */
const (
	/*GoproProtuneSharpnessLow - Low Sharpness. */
	GoproProtuneSharpnessLow = 0
	/*GoproProtuneSharpnessMedium - Medium Sharpness. */
	GoproProtuneSharpnessMedium = 1
	/*GoproProtuneSharpnessHigh - High Sharpness. */
	GoproProtuneSharpnessHigh = 2
	/*GoproProtuneSharpnessEnumEnd -  */
	GoproProtuneSharpnessEnumEnd = 3
)

/*GOPRO_PROTUNE_EXPOSURE -  */
const (
	/*GoproProtuneExposureNeg50 - -5.0 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg50 = 0
	/*GoproProtuneExposureNeg45 - -4.5 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg45 = 1
	/*GoproProtuneExposureNeg40 - -4.0 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg40 = 2
	/*GoproProtuneExposureNeg35 - -3.5 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg35 = 3
	/*GoproProtuneExposureNeg30 - -3.0 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg30 = 4
	/*GoproProtuneExposureNeg25 - -2.5 EV (Hero 3+ Only). */
	GoproProtuneExposureNeg25 = 5
	/*GoproProtuneExposureNeg20 - -2.0 EV. */
	GoproProtuneExposureNeg20 = 6
	/*GoproProtuneExposureNeg15 - -1.5 EV. */
	GoproProtuneExposureNeg15 = 7
	/*GoproProtuneExposureNeg10 - -1.0 EV. */
	GoproProtuneExposureNeg10 = 8
	/*GoproProtuneExposureNeg05 - -0.5 EV. */
	GoproProtuneExposureNeg05 = 9
	/*GoproProtuneExposureZero - 0.0 EV. */
	GoproProtuneExposureZero = 10
	/*GoproProtuneExposurePos05 - +0.5 EV. */
	GoproProtuneExposurePos05 = 11
	/*GoproProtuneExposurePos10 - +1.0 EV. */
	GoproProtuneExposurePos10 = 12
	/*GoproProtuneExposurePos15 - +1.5 EV. */
	GoproProtuneExposurePos15 = 13
	/*GoproProtuneExposurePos20 - +2.0 EV. */
	GoproProtuneExposurePos20 = 14
	/*GoproProtuneExposurePos25 - +2.5 EV (Hero 3+ Only). */
	GoproProtuneExposurePos25 = 15
	/*GoproProtuneExposurePos30 - +3.0 EV (Hero 3+ Only). */
	GoproProtuneExposurePos30 = 16
	/*GoproProtuneExposurePos35 - +3.5 EV (Hero 3+ Only). */
	GoproProtuneExposurePos35 = 17
	/*GoproProtuneExposurePos40 - +4.0 EV (Hero 3+ Only). */
	GoproProtuneExposurePos40 = 18
	/*GoproProtuneExposurePos45 - +4.5 EV (Hero 3+ Only). */
	GoproProtuneExposurePos45 = 19
	/*GoproProtuneExposurePos50 - +5.0 EV (Hero 3+ Only). */
	GoproProtuneExposurePos50 = 20
	/*GoproProtuneExposureEnumEnd -  */
	GoproProtuneExposureEnumEnd = 21
)

/*GOPRO_CHARGING -  */
const (
	/*GoproChargingDisabled - Charging disabled. */
	GoproChargingDisabled = 0
	/*GoproChargingEnabled - Charging enabled. */
	GoproChargingEnabled = 1
	/*GoproChargingEnumEnd -  */
	GoproChargingEnumEnd = 2
)

/*GOPRO_MODEL -  */
const (
	/*GoproModelUnknown - Unknown gopro model. */
	GoproModelUnknown = 0
	/*GoproModelHero3PlusSilver - Hero 3+ Silver (HeroBus not supported by GoPro). */
	GoproModelHero3PlusSilver = 1
	/*GoproModelHero3PlusBlack - Hero 3+ Black. */
	GoproModelHero3PlusBlack = 2
	/*GoproModelHero4Silver - Hero 4 Silver. */
	GoproModelHero4Silver = 3
	/*GoproModelHero4Black - Hero 4 Black. */
	GoproModelHero4Black = 4
	/*GoproModelEnumEnd -  */
	GoproModelEnumEnd = 5
)

/*GOPRO_BURST_RATE -  */
const (
	/*GoproBurstRate3In1Second - 3 Shots / 1 Second. */
	GoproBurstRate3In1Second = 0
	/*GoproBurstRate5In1Second - 5 Shots / 1 Second. */
	GoproBurstRate5In1Second = 1
	/*GoproBurstRate10In1Second - 10 Shots / 1 Second. */
	GoproBurstRate10In1Second = 2
	/*GoproBurstRate10In2Second - 10 Shots / 2 Second. */
	GoproBurstRate10In2Second = 3
	/*GoproBurstRate10In3Second - 10 Shots / 3 Second (Hero 4 Only). */
	GoproBurstRate10In3Second = 4
	/*GoproBurstRate30In1Second - 30 Shots / 1 Second. */
	GoproBurstRate30In1Second = 5
	/*GoproBurstRate30In2Second - 30 Shots / 2 Second. */
	GoproBurstRate30In2Second = 6
	/*GoproBurstRate30In3Second - 30 Shots / 3 Second. */
	GoproBurstRate30In3Second = 7
	/*GoproBurstRate30In6Second - 30 Shots / 6 Second. */
	GoproBurstRate30In6Second = 8
	/*GoproBurstRateEnumEnd -  */
	GoproBurstRateEnumEnd = 9
)

/*LED_CONTROL_PATTERN -  */
const (
	/*LedControlPatternOff - LED patterns off (return control to regular vehicle control). */
	LedControlPatternOff = 0
	/*LedControlPatternFirmwareupdate - LEDs show pattern during firmware update. */
	LedControlPatternFirmwareupdate = 1
	/*LedControlPatternCustom - Custom Pattern using custom bytes fields. */
	LedControlPatternCustom = 255
	/*LedControlPatternEnumEnd -  */
	LedControlPatternEnumEnd = 256
)

/*EKF_STATUS_FLAGS - Flags in EKF_STATUS message. */
const (
	/*EKFAttitude - Set if EKF's attitude estimate is good. */
	EKFAttitude = 1
	/*EKFVelocityHoriz - Set if EKF's horizontal velocity estimate is good. */
	EKFVelocityHoriz = 2
	/*EKFVelocityVert - Set if EKF's vertical velocity estimate is good. */
	EKFVelocityVert = 4
	/*EKFPosHorizRel - Set if EKF's horizontal position (relative) estimate is good. */
	EKFPosHorizRel = 8
	/*EKFPosHorizAbs - Set if EKF's horizontal position (absolute) estimate is good. */
	EKFPosHorizAbs = 16
	/*EKFPosVertAbs - Set if EKF's vertical position (absolute) estimate is good. */
	EKFPosVertAbs = 32
	/*EKFPosVertAgl - Set if EKF's vertical position (above ground) estimate is good. */
	EKFPosVertAgl = 64
	/*EKFConstPosMode - EKF is in constant position mode and does not know it's absolute or relative position. */
	EKFConstPosMode = 128
	/*EKFPredPosHorizRel - Set if EKF's predicted horizontal position (relative) estimate is good. */
	EKFPredPosHorizRel = 256
	/*EKFPredPosHorizAbs - Set if EKF's predicted horizontal position (absolute) estimate is good. */
	EKFPredPosHorizAbs = 512
	/*EKFStatusFlagsEnumEnd -  */
	EKFStatusFlagsEnumEnd = 513
)

/*PID_TUNING_AXIS -  */
const (
	/*PIDTuningRoll -  */
	PIDTuningRoll = 1
	/*PIDTuningPitch -  */
	PIDTuningPitch = 2
	/*PIDTuningYaw -  */
	PIDTuningYaw = 3
	/*PIDTuningAccz -  */
	PIDTuningAccz = 4
	/*PIDTuningSteer -  */
	PIDTuningSteer = 5
	/*PIDTuningLanding -  */
	PIDTuningLanding = 6
	/*PIDTuningAxisEnumEnd -  */
	PIDTuningAxisEnumEnd = 7
)

/*MAG_CAL_STATUS -  */
const (
	/*MagCalNotStarted -  */
	MagCalNotStarted = 0
	/*MagCalWaitingToStart -  */
	MagCalWaitingToStart = 1
	/*MagCalRunningStepOne -  */
	MagCalRunningStepOne = 2
	/*MagCalRunningStepTwo -  */
	MagCalRunningStepTwo = 3
	/*MagCalSuccess -  */
	MagCalSuccess = 4
	/*MagCalFailed -  */
	MagCalFailed = 5
	/*MagCalBadOrientation -  */
	MagCalBadOrientation = 6
	/*MagCalStatusEnumEnd -  */
	MagCalStatusEnumEnd = 7
)

/*MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS - Special ACK block numbers control activation of dataflash log streaming. */
const (
	/*MavRemoteLogDataBlockStop - UAV to stop sending DataFlash blocks. */
	MavRemoteLogDataBlockStop = 2147483645
	/*MavRemoteLogDataBlockStart - UAV to start sending DataFlash blocks. */
	MavRemoteLogDataBlockStart = 2147483646
	/*MavRemoteLogDataBlockCommandsEnumEnd -  */
	MavRemoteLogDataBlockCommandsEnumEnd = 2147483647
)

/*MAV_REMOTE_LOG_DATA_BLOCK_STATUSES - Possible remote log data block statuses. */
const (
	/*MavRemoteLogDataBlockNack - This block has NOT been received. */
	MavRemoteLogDataBlockNack = 0
	/*MavRemoteLogDataBlockAck - This block has been received. */
	MavRemoteLogDataBlockAck = 1
	/*MavRemoteLogDataBlockStatusesEnumEnd -  */
	MavRemoteLogDataBlockStatusesEnumEnd = 2
)

/*DEVICE_OP_BUSTYPE - Bus types for device operations. */
const (
	/*DeviceOpBustypeI2C - I2C Device operation. */
	DeviceOpBustypeI2C = 0
	/*DeviceOpBustypeSpi - SPI Device operation. */
	DeviceOpBustypeSpi = 1
	/*DeviceOpBustypeEnumEnd -  */
	DeviceOpBustypeEnumEnd = 2
)

/*DEEPSTALL_STAGE - Deepstall flight stage. */
const (
	/*DeepstallStageFlyToLanding - Flying to the landing point. */
	DeepstallStageFlyToLanding = 0
	/*DeepstallStageEstimateWind - Building an estimate of the wind. */
	DeepstallStageEstimateWind = 1
	/*DeepstallStageWaitForBreakout - Waiting to breakout of the loiter to fly the approach. */
	DeepstallStageWaitForBreakout = 2
	/*DeepstallStageFlyToARC - Flying to the first arc point to turn around to the landing point. */
	DeepstallStageFlyToARC = 3
	/*DeepstallStageARC - Turning around back to the deepstall landing point. */
	DeepstallStageARC = 4
	/*DeepstallStageApproach - Approaching the landing point. */
	DeepstallStageApproach = 5
	/*DeepstallStageLand - Stalling and steering towards the land point. */
	DeepstallStageLand = 6
	/*DeepstallStageEnumEnd -  */
	DeepstallStageEnumEnd = 7
)

/*PLANE_MODE - A mapping of plane flight modes for custom_mode field of heartbeat. */
const (
	/*PlaneModeManual -  */
	PlaneModeManual = 0
	/*PlaneModeCiRCle -  */
	PlaneModeCiRCle = 1
	/*PlaneModeStabilize -  */
	PlaneModeStabilize = 2
	/*PlaneModeTraining -  */
	PlaneModeTraining = 3
	/*PlaneModeAcro -  */
	PlaneModeAcro = 4
	/*PlaneModeFlyByWireA -  */
	PlaneModeFlyByWireA = 5
	/*PlaneModeFlyByWireB -  */
	PlaneModeFlyByWireB = 6
	/*PlaneModeCruise -  */
	PlaneModeCruise = 7
	/*PlaneModeAutotune -  */
	PlaneModeAutotune = 8
	/*PlaneModeAuto -  */
	PlaneModeAuto = 10
	/*PlaneModeRtl -  */
	PlaneModeRtl = 11
	/*PlaneModeLoiter -  */
	PlaneModeLoiter = 12
	/*PlaneModeAvoIDAdsb -  */
	PlaneModeAvoIDAdsb = 14
	/*PlaneModeGUIDed -  */
	PlaneModeGUIDed = 15
	/*PlaneModeInitializing -  */
	PlaneModeInitializing = 16
	/*PlaneModeQstabilize -  */
	PlaneModeQstabilize = 17
	/*PlaneModeQhover -  */
	PlaneModeQhover = 18
	/*PlaneModeQloiter -  */
	PlaneModeQloiter = 19
	/*PlaneModeQland -  */
	PlaneModeQland = 20
	/*PlaneModeQrtl -  */
	PlaneModeQrtl = 21
	/*PlaneModeEnumEnd -  */
	PlaneModeEnumEnd = 22
)

/*COPTER_MODE - A mapping of copter flight modes for custom_mode field of heartbeat. */
const (
	/*CopterModeStabilize -  */
	CopterModeStabilize = 0
	/*CopterModeAcro -  */
	CopterModeAcro = 1
	/*CopterModeAltHold -  */
	CopterModeAltHold = 2
	/*CopterModeAuto -  */
	CopterModeAuto = 3
	/*CopterModeGUIDed -  */
	CopterModeGUIDed = 4
	/*CopterModeLoiter -  */
	CopterModeLoiter = 5
	/*CopterModeRtl -  */
	CopterModeRtl = 6
	/*CopterModeCiRCle -  */
	CopterModeCiRCle = 7
	/*CopterModeLand -  */
	CopterModeLand = 9
	/*CopterModeDrift -  */
	CopterModeDrift = 11
	/*CopterModeSport -  */
	CopterModeSport = 13
	/*CopterModeFlip -  */
	CopterModeFlip = 14
	/*CopterModeAutotune -  */
	CopterModeAutotune = 15
	/*CopterModePoshold -  */
	CopterModePoshold = 16
	/*CopterModeBrake -  */
	CopterModeBrake = 17
	/*CopterModeThrow -  */
	CopterModeThrow = 18
	/*CopterModeAvoIDAdsb -  */
	CopterModeAvoIDAdsb = 19
	/*CopterModeGUIDedNoGPS -  */
	CopterModeGUIDedNoGPS = 20
	/*CopterModeSmartRtl -  */
	CopterModeSmartRtl = 21
	/*CopterModeEnumEnd -  */
	CopterModeEnumEnd = 22
)

/*SUB_MODE - A mapping of sub flight modes for custom_mode field of heartbeat. */
const (
	/*SubModeStabilize -  */
	SubModeStabilize = 0
	/*SubModeAcro -  */
	SubModeAcro = 1
	/*SubModeAltHold -  */
	SubModeAltHold = 2
	/*SubModeAuto -  */
	SubModeAuto = 3
	/*SubModeGUIDed -  */
	SubModeGUIDed = 4
	/*SubModeCiRCle -  */
	SubModeCiRCle = 7
	/*SubModeSurface -  */
	SubModeSurface = 9
	/*SubModePoshold -  */
	SubModePoshold = 16
	/*SubModeManual -  */
	SubModeManual = 19
	/*SubModeEnumEnd -  */
	SubModeEnumEnd = 20
)

/*ROVER_MODE - A mapping of rover flight modes for custom_mode field of heartbeat. */
const (
	/*RoverModeManual -  */
	RoverModeManual = 0
	/*RoverModeAcro -  */
	RoverModeAcro = 1
	/*RoverModeSteering -  */
	RoverModeSteering = 3
	/*RoverModeHold -  */
	RoverModeHold = 4
	/*RoverModeLoiter -  */
	RoverModeLoiter = 5
	/*RoverModeAuto -  */
	RoverModeAuto = 10
	/*RoverModeRtl -  */
	RoverModeRtl = 11
	/*RoverModeSmartRtl -  */
	RoverModeSmartRtl = 12
	/*RoverModeGUIDed -  */
	RoverModeGUIDed = 15
	/*RoverModeInitializing -  */
	RoverModeInitializing = 16
	/*RoverModeEnumEnd -  */
	RoverModeEnumEnd = 17
)

/*TRACKER_MODE - A mapping of antenna tracker flight modes for custom_mode field of heartbeat. */
const (
	/*TrackerModeManual -  */
	TrackerModeManual = 0
	/*TrackerModeStop -  */
	TrackerModeStop = 1
	/*TrackerModeScan -  */
	TrackerModeScan = 2
	/*TrackerModeServoTest -  */
	TrackerModeServoTest = 3
	/*TrackerModeAuto -  */
	TrackerModeAuto = 10
	/*TrackerModeInitializing -  */
	TrackerModeInitializing = 16
	/*TrackerModeEnumEnd -  */
	TrackerModeEnumEnd = 17
)
