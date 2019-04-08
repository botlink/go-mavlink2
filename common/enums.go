package common

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

/*MAV_AUTOPILOT - Micro air vehicle / autopilot classes. This identifies the individual model. */
const (
	/*MavAutopilotGeneric - Generic autopilot, full support for everything */
	MavAutopilotGeneric = 0
	/*MavAutopilotReserved - Reserved for future use. */
	MavAutopilotReserved = 1
	/*MavAutopilotSlugs - SLUGS autopilot, http://slugsuav.soe.ucsc.edu */
	MavAutopilotSlugs = 2
	/*MavAutopilotArdupilotmega - ArduPilot - Plane/Copter/Rover/Sub/Tracker, http://ardupilot.org */
	MavAutopilotArdupilotmega = 3
	/*MavAutopilotOpenpilot - OpenPilot, http://openpilot.org */
	MavAutopilotOpenpilot = 4
	/*MavAutopilotGenericWaypointsOnly - Generic autopilot only supporting simple waypoints */
	MavAutopilotGenericWaypointsOnly = 5
	/*MavAutopilotGenericWaypointsAndSimpleNAVigationOnly - Generic autopilot supporting waypoints and other simple navigation commands */
	MavAutopilotGenericWaypointsAndSimpleNAVigationOnly = 6
	/*MavAutopilotGenericMissionFull - Generic autopilot supporting the full mission command set */
	MavAutopilotGenericMissionFull = 7
	/*MavAutopilotInvalID - No valid autopilot, e.g. a GCS or other MAVLink component */
	MavAutopilotInvalID = 8
	/*MavAutopilotPpz - PPZ UAV - http://nongnu.org/paparazzi */
	MavAutopilotPpz = 9
	/*MavAutopilotUdb - UAV Dev Board */
	MavAutopilotUdb = 10
	/*MavAutopilotFp - FlexiPilot */
	MavAutopilotFp = 11
	/*MavAutopilotPx4 - PX4 Autopilot - http://px4.io/ */
	MavAutopilotPx4 = 12
	/*MavAutopilotSmaccmpilot - SMACCMPilot - http://smaccmpilot.org */
	MavAutopilotSmaccmpilot = 13
	/*MavAutopilotAutoquad - AutoQuad -- http://autoquad.org */
	MavAutopilotAutoquad = 14
	/*MavAutopilotArmazila - Armazila -- http://armazila.com */
	MavAutopilotArmazila = 15
	/*MavAutopilotAerob - Aerob -- http://aerob.ru */
	MavAutopilotAerob = 16
	/*MavAutopilotAsluav - ASLUAV autopilot -- http://www.asl.ethz.ch */
	MavAutopilotAsluav = 17
	/*MavAutopilotSmartap - SmartAP Autopilot - http://sky-drones.com */
	MavAutopilotSmartap = 18
	/*MavAutopilotAirrails - AirRails - http://uaventure.com */
	MavAutopilotAirrails = 19
	/*MavAutopilotEnumEnd -  */
	MavAutopilotEnumEnd = 20
)

/*MAV_TYPE -  */
const (
	/*MavTypeGeneric - Generic micro air vehicle. */
	MavTypeGeneric = 0
	/*MavTypeFixedWing - Fixed wing aircraft. */
	MavTypeFixedWing = 1
	/*MavTypeQuadrotor - Quadrotor */
	MavTypeQuadrotor = 2
	/*MavTypeCoaxial - Coaxial helicopter */
	MavTypeCoaxial = 3
	/*MavTypeHelicopter - Normal helicopter with tail rotor. */
	MavTypeHelicopter = 4
	/*MavTypeAntennaTracker - Ground installation */
	MavTypeAntennaTracker = 5
	/*MavTypeGcs - Operator control unit / ground control station */
	MavTypeGcs = 6
	/*MavTypeAirship - Airship, controlled */
	MavTypeAirship = 7
	/*MavTypeFreeBalloon - Free balloon, uncontrolled */
	MavTypeFreeBalloon = 8
	/*MavTypeRocket - Rocket */
	MavTypeRocket = 9
	/*MavTypeGroundRover - Ground rover */
	MavTypeGroundRover = 10
	/*MavTypeSurfaceBoat - Surface vessel, boat, ship */
	MavTypeSurfaceBoat = 11
	/*MavTypeSubmarine - Submarine */
	MavTypeSubmarine = 12
	/*MavTypeHexarotor - Hexarotor */
	MavTypeHexarotor = 13
	/*MavTypeOctorotor - Octorotor */
	MavTypeOctorotor = 14
	/*MavTypeTricopter - Tricopter */
	MavTypeTricopter = 15
	/*MavTypeFlappingWing - Flapping wing */
	MavTypeFlappingWing = 16
	/*MavTypeKite - Kite */
	MavTypeKite = 17
	/*MavTypeOnboardController - Onboard companion controller */
	MavTypeOnboardController = 18
	/*MavTypeVtolDuorotor - Two-rotor VTOL using control surfaces in vertical operation in addition. Tailsitter. */
	MavTypeVtolDuorotor = 19
	/*MavTypeVtolQuadrotor - Quad-rotor VTOL using a V-shaped quad config in vertical operation. Tailsitter. */
	MavTypeVtolQuadrotor = 20
	/*MavTypeVtolTiltrotor - Tiltrotor VTOL */
	MavTypeVtolTiltrotor = 21
	/*MavTypeVtolReserved2 - VTOL reserved 2 */
	MavTypeVtolReserved2 = 22
	/*MavTypeVtolReserved3 - VTOL reserved 3 */
	MavTypeVtolReserved3 = 23
	/*MavTypeVtolReserved4 - VTOL reserved 4 */
	MavTypeVtolReserved4 = 24
	/*MavTypeVtolReserved5 - VTOL reserved 5 */
	MavTypeVtolReserved5 = 25
	/*MavTypeGimbal - Onboard gimbal */
	MavTypeGimbal = 26
	/*MavTypeAdsb - Onboard ADSB peripheral */
	MavTypeAdsb = 27
	/*MavTypeParafoil - Steerable, nonrigid airfoil */
	MavTypeParafoil = 28
	/*MavTypeDodecarotor - Dodecarotor */
	MavTypeDodecarotor = 29
	/*MavTypeCamera - Camera */
	MavTypeCamera = 30
	/*MavTypeChargingStation - Charging station */
	MavTypeChargingStation = 31
	/*MavTypeFlarm - Onboard FLARM collision avoidance system */
	MavTypeFlarm = 32
	/*MavTypeEnumEnd -  */
	MavTypeEnumEnd = 33
)

/*FIRMWARE_VERSION_TYPE - These values define the type of firmware release.  These values indicate the first version or release of this type.  For example the first alpha release would be 64, the second would be 65. */
const (
	/*FirmwareVersionTypeDev - development release */
	FirmwareVersionTypeDev = 0
	/*FirmwareVersionTypeAlpha - alpha release */
	FirmwareVersionTypeAlpha = 64
	/*FirmwareVersionTypeBeta - beta release */
	FirmwareVersionTypeBeta = 128
	/*FirmwareVersionTypeRC - release candidate */
	FirmwareVersionTypeRC = 192
	/*FirmwareVersionTypeOfficial - official stable release */
	FirmwareVersionTypeOfficial = 255
	/*FirmwareVersionTypeEnumEnd -  */
	FirmwareVersionTypeEnumEnd = 256
)

/*HL_FAILURE_FLAG - Flags to report failure cases over the high latency telemtry. */
const (
	/*HlFailureFlagGPS - GPS failure. */
	HlFailureFlagGPS = 1
	/*HlFailureFlagDifferentialPressure - Differential pressure sensor failure. */
	HlFailureFlagDifferentialPressure = 2
	/*HlFailureFlagAbsolutePressure - Absolute pressure sensor failure. */
	HlFailureFlagAbsolutePressure = 4
	/*HlFailureFlag3DAccel - Accelerometer sensor failure. */
	HlFailureFlag3DAccel = 8
	/*HlFailureFlag3DGyro - Gyroscope sensor failure. */
	HlFailureFlag3DGyro = 16
	/*HlFailureFlag3DMag - Magnetometer sensor failure. */
	HlFailureFlag3DMag = 32
	/*HlFailureFlagTerrain - Terrain subsystem failure. */
	HlFailureFlagTerrain = 64
	/*HlFailureFlagBattery - Battery failure/critical low battery. */
	HlFailureFlagBattery = 128
	/*HlFailureFlagRCReceiver - RC receiver failure/no rc connection. */
	HlFailureFlagRCReceiver = 256
	/*HlFailureFlagOffboardLink - Offboard link failure. */
	HlFailureFlagOffboardLink = 512
	/*HlFailureFlagEngine - Engine failure. */
	HlFailureFlagEngine = 1024
	/*HlFailureFlagGeofence - Geofence violation. */
	HlFailureFlagGeofence = 2048
	/*HlFailureFlagEstimator - Estimator failure, for example measurement rejection or large variances. */
	HlFailureFlagEstimator = 4096
	/*HlFailureFlagMission - Mission failure. */
	HlFailureFlagMission = 8192
	/*HlFailureFlagEnumEnd -  */
	HlFailureFlagEnumEnd = 8193
)

/*MAV_MODE_FLAG - These flags encode the MAV mode. */
const (
	/*MavModeFlagCustomModeEnabled - 0b00000001 Reserved for future use. */
	MavModeFlagCustomModeEnabled = 1
	/*MavModeFlagTestEnabled - 0b00000010 system has a test mode enabled. This flag is intended for temporary system tests and should not be used for stable implementations. */
	MavModeFlagTestEnabled = 2
	/*MavModeFlagAutoEnabled - 0b00000100 autonomous mode enabled, system finds its own goal positions. Guided flag can be set or not, depends on the actual implementation. */
	MavModeFlagAutoEnabled = 4
	/*MavModeFlagGUIDedEnabled - 0b00001000 guided mode enabled, system flies waypoints / mission items. */
	MavModeFlagGUIDedEnabled = 8
	/*MavModeFlagStabilizeEnabled - 0b00010000 system stabilizes electronically its attitude (and optionally position). It needs however further control inputs to move around. */
	MavModeFlagStabilizeEnabled = 16
	/*MavModeFlagHilEnabled - 0b00100000 hardware in the loop simulation. All motors / actuators are blocked, but internal software is full operational. */
	MavModeFlagHilEnabled = 32
	/*MavModeFlagManualInputEnabled - 0b01000000 remote control input is enabled. */
	MavModeFlagManualInputEnabled = 64
	/*MavModeFlagSafetyArmed - 0b10000000 MAV safety set to armed. Motors are enabled / running / can start. Ready to fly. Additional note: this flag is to be ignore when sent in the command MAV_CMD_DO_SET_MODE and MAV_CMD_COMPONENT_ARM_DISARM shall be used instead. The flag can still be used to report the armed state. */
	MavModeFlagSafetyArmed = 128
	/*MavModeFlagEnumEnd -  */
	MavModeFlagEnumEnd = 129
)

/*MAV_MODE_FLAG_DECODE_POSITION - These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not. */
const (
	/*MavModeFlagDecodePositionCustomMode - Eighth bit: 00000001 */
	MavModeFlagDecodePositionCustomMode = 1
	/*MavModeFlagDecodePositionTest - Seventh bit: 00000010 */
	MavModeFlagDecodePositionTest = 2
	/*MavModeFlagDecodePositionAuto - Sixt bit:   00000100 */
	MavModeFlagDecodePositionAuto = 4
	/*MavModeFlagDecodePositionGUIDed - Fifth bit:  00001000 */
	MavModeFlagDecodePositionGUIDed = 8
	/*MavModeFlagDecodePositionStabilize - Fourth bit: 00010000 */
	MavModeFlagDecodePositionStabilize = 16
	/*MavModeFlagDecodePositionHil - Third bit:  00100000 */
	MavModeFlagDecodePositionHil = 32
	/*MavModeFlagDecodePositionManual - Second bit: 01000000 */
	MavModeFlagDecodePositionManual = 64
	/*MavModeFlagDecodePositionSafety - First bit:  10000000 */
	MavModeFlagDecodePositionSafety = 128
	/*MavModeFlagDecodePositionEnumEnd -  */
	MavModeFlagDecodePositionEnumEnd = 129
)

/*MAV_GOTO - Override command, pauses current mission execution and moves immediately to a position */
const (
	/*MavGotoDoHold - Hold at the current position. */
	MavGotoDoHold = 0
	/*MavGotoDoContinue - Continue with the next item in mission execution. */
	MavGotoDoContinue = 1
	/*MavGotoHoldAtCurrentPosition - Hold at the current position of the system */
	MavGotoHoldAtCurrentPosition = 2
	/*MavGotoHoldAtSpecifiedPosition - Hold at the position specified in the parameters of the DO_HOLD action */
	MavGotoHoldAtSpecifiedPosition = 3
	/*MavGotoEnumEnd -  */
	MavGotoEnumEnd = 4
)

/*MAV_MODE - These defines are predefined OR-combined mode flags. There is no need to use values from this enum, but it
  simplifies the use of the mode flags. Note that manual input is enabled in all modes as a safety override. */
const (
	/*MavModePreflight - System is not ready to fly, booting, calibrating, etc. No flag is set. */
	MavModePreflight = 0
	/*MavModeManualDisarmed - System is allowed to be active, under manual (RC) control, no stabilization */
	MavModeManualDisarmed = 64
	/*MavModeTestDisarmed - UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only. */
	MavModeTestDisarmed = 66
	/*MavModeStabilizeDisarmed - System is allowed to be active, under assisted RC control. */
	MavModeStabilizeDisarmed = 80
	/*MavModeGUIDedDisarmed - System is allowed to be active, under autonomous control, manual setpoint */
	MavModeGUIDedDisarmed = 88
	/*MavModeAutoDisarmed - System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints) */
	MavModeAutoDisarmed = 92
	/*MavModeManualArmed - System is allowed to be active, under manual (RC) control, no stabilization */
	MavModeManualArmed = 192
	/*MavModeTestArmed - UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only. */
	MavModeTestArmed = 194
	/*MavModeStabilizeArmed - System is allowed to be active, under assisted RC control. */
	MavModeStabilizeArmed = 208
	/*MavModeGUIDedArmed - System is allowed to be active, under autonomous control, manual setpoint */
	MavModeGUIDedArmed = 216
	/*MavModeAutoArmed - System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints) */
	MavModeAutoArmed = 220
	/*MavModeEnumEnd -  */
	MavModeEnumEnd = 221
)

/*MAV_STATE -  */
const (
	/*MavStateUninit - Uninitialized system, state is unknown. */
	MavStateUninit = 0
	/*MavStateBoot - System is booting up. */
	MavStateBoot = 1
	/*MavStateCalibrating - System is calibrating and not flight-ready. */
	MavStateCalibrating = 2
	/*MavStateStandby - System is grounded and on standby. It can be launched any time. */
	MavStateStandby = 3
	/*MavStateActive - System is active and might be already airborne. Motors are engaged. */
	MavStateActive = 4
	/*MavStateCritical - System is in a non-normal flight mode. It can however still navigate. */
	MavStateCritical = 5
	/*MavStateEmergency - System is in a non-normal flight mode. It lost control over parts or over the whole airframe. It is in mayday and going down. */
	MavStateEmergency = 6
	/*MavStatePoweroff - System just initialized its power-down sequence, will shut down now. */
	MavStatePoweroff = 7
	/*MavStateFlightTermination - System is terminating itself. */
	MavStateFlightTermination = 8
	/*MavStateEnumEnd -  */
	MavStateEnumEnd = 9
)

/*MAV_COMPONENT -  */
const (
	/*MavCompIDAll -  */
	MavCompIDAll = 0
	/*MavCompIDAutopilot1 -  */
	MavCompIDAutopilot1 = 1
	/*MavCompIDCamera -  */
	MavCompIDCamera = 100
	/*MavCompIDCamera2 -  */
	MavCompIDCamera2 = 101
	/*MavCompIDCamera3 -  */
	MavCompIDCamera3 = 102
	/*MavCompIDCamera4 -  */
	MavCompIDCamera4 = 103
	/*MavCompIDCamera5 -  */
	MavCompIDCamera5 = 104
	/*MavCompIDCamera6 -  */
	MavCompIDCamera6 = 105
	/*MavCompIDServo1 -  */
	MavCompIDServo1 = 140
	/*MavCompIDServo2 -  */
	MavCompIDServo2 = 141
	/*MavCompIDServo3 -  */
	MavCompIDServo3 = 142
	/*MavCompIDServo4 -  */
	MavCompIDServo4 = 143
	/*MavCompIDServo5 -  */
	MavCompIDServo5 = 144
	/*MavCompIDServo6 -  */
	MavCompIDServo6 = 145
	/*MavCompIDServo7 -  */
	MavCompIDServo7 = 146
	/*MavCompIDServo8 -  */
	MavCompIDServo8 = 147
	/*MavCompIDServo9 -  */
	MavCompIDServo9 = 148
	/*MavCompIDServo10 -  */
	MavCompIDServo10 = 149
	/*MavCompIDServo11 -  */
	MavCompIDServo11 = 150
	/*MavCompIDServo12 -  */
	MavCompIDServo12 = 151
	/*MavCompIDServo13 -  */
	MavCompIDServo13 = 152
	/*MavCompIDServo14 -  */
	MavCompIDServo14 = 153
	/*MavCompIDGimbal -  */
	MavCompIDGimbal = 154
	/*MavCompIDLog -  */
	MavCompIDLog = 155
	/*MavCompIDAdsb -  */
	MavCompIDAdsb = 156
	/*MavCompIDOsd - On Screen Display (OSD) devices for video links */
	MavCompIDOsd = 157
	/*MavCompIDPeripheral - Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter sub-protocol */
	MavCompIDPeripheral = 158
	/*MavCompIDQx1Gimbal -  */
	MavCompIDQx1Gimbal = 159
	/*MavCompIDFlarm -  */
	MavCompIDFlarm = 160
	/*MavCompIDMapper -  */
	MavCompIDMapper = 180
	/*MavCompIDMissionplanner -  */
	MavCompIDMissionplanner = 190
	/*MavCompIDPathplanner -  */
	MavCompIDPathplanner = 195
	/*MavCompIDIMU -  */
	MavCompIDIMU = 200
	/*MavCompIDIMU2 -  */
	MavCompIDIMU2 = 201
	/*MavCompIDIMU3 -  */
	MavCompIDIMU3 = 202
	/*MavCompIDGPS -  */
	MavCompIDGPS = 220
	/*MavCompIDGPS2 -  */
	MavCompIDGPS2 = 221
	/*MavCompIDUDPBrIDge -  */
	MavCompIDUDPBrIDge = 240
	/*MavCompIDUartBrIDge -  */
	MavCompIDUartBrIDge = 241
	/*MavCompIDSystemControl -  */
	MavCompIDSystemControl = 250
	/*MavComponentEnumEnd -  */
	MavComponentEnumEnd = 251
)

/*MAV_SYS_STATUS_SENSOR - These encode the sensors whose status is sent as part of the SYS_STATUS message. */
const (
	/*MavSysStatusSensor3DGyro - 0x01 3D gyro */
	MavSysStatusSensor3DGyro = 1
	/*MavSysStatusSensor3DAccel - 0x02 3D accelerometer */
	MavSysStatusSensor3DAccel = 2
	/*MavSysStatusSensor3DMag - 0x04 3D magnetometer */
	MavSysStatusSensor3DMag = 4
	/*MavSysStatusSensorAbsolutePressure - 0x08 absolute pressure */
	MavSysStatusSensorAbsolutePressure = 8
	/*MavSysStatusSensorDifferentialPressure - 0x10 differential pressure */
	MavSysStatusSensorDifferentialPressure = 16
	/*MavSysStatusSensorGPS - 0x20 GPS */
	MavSysStatusSensorGPS = 32
	/*MavSysStatusSensorOpticalFlow - 0x40 optical flow */
	MavSysStatusSensorOpticalFlow = 64
	/*MavSysStatusSensorVisionPosition - 0x80 computer vision position */
	MavSysStatusSensorVisionPosition = 128
	/*MavSysStatusSensorLaserPosition - 0x100 laser based position */
	MavSysStatusSensorLaserPosition = 256
	/*MavSysStatusSensorExternalGroundTruth - 0x200 external ground truth (Vicon or Leica) */
	MavSysStatusSensorExternalGroundTruth = 512
	/*MavSysStatusSensorAngularRateControl - 0x400 3D angular rate control */
	MavSysStatusSensorAngularRateControl = 1024
	/*MavSysStatusSensorAttitudeStabilization - 0x800 attitude stabilization */
	MavSysStatusSensorAttitudeStabilization = 2048
	/*MavSysStatusSensorYawPosition - 0x1000 yaw position */
	MavSysStatusSensorYawPosition = 4096
	/*MavSysStatusSensorZAltitudeControl - 0x2000 z/altitude control */
	MavSysStatusSensorZAltitudeControl = 8192
	/*MavSysStatusSensorXyPositionControl - 0x4000 x/y position control */
	MavSysStatusSensorXyPositionControl = 16384
	/*MavSysStatusSensorMotorOutputs - 0x8000 motor outputs / control */
	MavSysStatusSensorMotorOutputs = 32768
	/*MavSysStatusSensorRCReceiver - 0x10000 rc receiver */
	MavSysStatusSensorRCReceiver = 65536
	/*MavSysStatusSensor3DGyro2 - 0x20000 2nd 3D gyro */
	MavSysStatusSensor3DGyro2 = 131072
	/*MavSysStatusSensor3DAccel2 - 0x40000 2nd 3D accelerometer */
	MavSysStatusSensor3DAccel2 = 262144
	/*MavSysStatusSensor3DMag2 - 0x80000 2nd 3D magnetometer */
	MavSysStatusSensor3DMag2 = 524288
	/*MavSysStatusGeofence - 0x100000 geofence */
	MavSysStatusGeofence = 1048576
	/*MavSysStatusAHRS - 0x200000 AHRS subsystem health */
	MavSysStatusAHRS = 2097152
	/*MavSysStatusTerrain - 0x400000 Terrain subsystem health */
	MavSysStatusTerrain = 4194304
	/*MavSysStatusReverseMotor - 0x800000 Motors are reversed */
	MavSysStatusReverseMotor = 8388608
	/*MavSysStatusLogging - 0x1000000 Logging */
	MavSysStatusLogging = 16777216
	/*MavSysStatusSensorBattery - 0x2000000 Battery */
	MavSysStatusSensorBattery = 33554432
	/*MavSysStatusSensorProximity - 0x4000000 Proximity */
	MavSysStatusSensorProximity = 67108864
	/*MavSysStatusSensorSatcom - 0x8000000 Satellite Communication  */
	MavSysStatusSensorSatcom = 134217728
	/*MavSysStatusSensorEnumEnd -  */
	MavSysStatusSensorEnumEnd = 134217729
)

/*MAV_FRAME -  */
const (
	/*MavFrameGlobal - Global coordinate frame, WGS84 coordinate system. First value / x: latitude, second value / y: longitude, third value / z: positive altitude over mean sea level (MSL). */
	MavFrameGlobal = 0
	/*MavFrameLocalNed - Local coordinate frame, Z-down (x: north, y: east, z: down). */
	MavFrameLocalNed = 1
	/*MavFrameMission - NOT a coordinate frame, indicates a mission command. */
	MavFrameMission = 2
	/*MavFrameGlobalRelativeAlt - Global coordinate frame, WGS84 coordinate system, relative altitude over ground with respect to the home position. First value / x: latitude, second value / y: longitude, third value / z: positive altitude with 0 being at the altitude of the home location. */
	MavFrameGlobalRelativeAlt = 3
	/*MavFrameLocalEnu - Local coordinate frame, Z-up (x: east, y: north, z: up). */
	MavFrameLocalEnu = 4
	/*MavFrameGlobalInt - Global coordinate frame, WGS84 coordinate system. First value / x: latitude in degrees*1.0e-7, second value / y: longitude in degrees*1.0e-7, third value / z: positive altitude over mean sea level (MSL). */
	MavFrameGlobalInt = 5
	/*MavFrameGlobalRelativeAltInt - Global coordinate frame, WGS84 coordinate system, relative altitude over ground with respect to the home position. First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude with 0 being at the altitude of the home location. */
	MavFrameGlobalRelativeAltInt = 6
	/*MavFrameLocalOffsetNed - Offset to the current local frame. Anything expressed in this frame should be added to the current local frame position. */
	MavFrameLocalOffsetNed = 7
	/*MavFrameBodyNed - Setpoint in body NED frame. This makes sense if all position control is externalized - e.g. useful to command 2 m/s^2 acceleration to the right. */
	MavFrameBodyNed = 8
	/*MavFrameBodyOffsetNed - Offset in body NED frame. This makes sense if adding setpoints to the current flight path, to avoid an obstacle - e.g. useful to command 2 m/s^2 acceleration to the east. */
	MavFrameBodyOffsetNed = 9
	/*MavFrameGlobalTerrainAlt - Global (WGS84) coordinate frame with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees, second value / y: longitude in degrees, third value / z: positive altitude in meters with 0 being at ground level in terrain model. */
	MavFrameGlobalTerrainAlt = 10
	/*MavFrameGlobalTerrainAltInt - Global (WGS84) coordinate frame with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude in meters with 0 being at ground level in terrain model. */
	MavFrameGlobalTerrainAltInt = 11
	/*MavFrameBodyFrd - Body fixed frame of reference, Z-down (x: forward, y: right, z: down). */
	MavFrameBodyFrd = 12
	/*MavFrameBodyFlu - Body fixed frame of reference, Z-up (x: forward, y: left, z: up). */
	MavFrameBodyFlu = 13
	/*MavFrameMocapNed - Odometry local coordinate frame of data given by a motion capture system, Z-down (x: north, y: east, z: down). */
	MavFrameMocapNed = 14
	/*MavFrameMocapEnu - Odometry local coordinate frame of data given by a motion capture system, Z-up (x: east, y: north, z: up). */
	MavFrameMocapEnu = 15
	/*MavFrameVisionNed - Odometry local coordinate frame of data given by a vision estimation system, Z-down (x: north, y: east, z: down). */
	MavFrameVisionNed = 16
	/*MavFrameVisionEnu - Odometry local coordinate frame of data given by a vision estimation system, Z-up (x: east, y: north, z: up). */
	MavFrameVisionEnu = 17
	/*MavFrameEstimNed - Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-down (x: north, y: east, z: down). */
	MavFrameEstimNed = 18
	/*MavFrameEstimEnu - Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-up (x: east, y: noth, z: up). */
	MavFrameEstimEnu = 19
	/*MavFrameEnumEnd -  */
	MavFrameEnumEnd = 20
)

/*MAVLINK_DATA_STREAM_TYPE -  */
const (
	/*MavlinkDataStreamImgJpeg -  */
	MavlinkDataStreamImgJpeg = 1
	/*MavlinkDataStreamImgBmp -  */
	MavlinkDataStreamImgBmp = 2
	/*MavlinkDataStreamImgRaw8U -  */
	MavlinkDataStreamImgRaw8U = 3
	/*MavlinkDataStreamImgRaw32U -  */
	MavlinkDataStreamImgRaw32U = 4
	/*MavlinkDataStreamImgPgm -  */
	MavlinkDataStreamImgPgm = 5
	/*MavlinkDataStreamImgPng -  */
	MavlinkDataStreamImgPng = 6
	/*MavlinkDataStreamTypeEnumEnd -  */
	MavlinkDataStreamTypeEnumEnd = 7
)

/*FENCE_ACTION -  */
const (
	/*FenceActionNone - Disable fenced mode */
	FenceActionNone = 0
	/*FenceActionGUIDed - Switched to guided mode to return point (fence point 0) */
	FenceActionGUIDed = 1
	/*FenceActionReport - Report fence breach, but don't take action */
	FenceActionReport = 2
	/*FenceActionGUIDedThrPass - Switched to guided mode to return point (fence point 0) with manual throttle control */
	FenceActionGUIDedThrPass = 3
	/*FenceActionRtl - Switch to RTL (return to launch) mode and head for the return point. */
	FenceActionRtl = 4
	/*FenceActionEnumEnd -  */
	FenceActionEnumEnd = 5
)

/*FENCE_BREACH -  */
const (
	/*FenceBreachNone - No last fence breach */
	FenceBreachNone = 0
	/*FenceBreachMinalt - Breached minimum altitude */
	FenceBreachMinalt = 1
	/*FenceBreachMaxalt - Breached maximum altitude */
	FenceBreachMaxalt = 2
	/*FenceBreachBoundary - Breached fence boundary */
	FenceBreachBoundary = 3
	/*FenceBreachEnumEnd -  */
	FenceBreachEnumEnd = 4
)

/*MAV_MOUNT_MODE - Enumeration of possible mount operation modes */
const (
	/*MavMountModeRetract - Load and keep safe position (Roll,Pitch,Yaw) from permant memory and stop stabilization */
	MavMountModeRetract = 0
	/*MavMountModeNeutral - Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory. */
	MavMountModeNeutral = 1
	/*MavMountModeMavlinkTargeting - Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization */
	MavMountModeMavlinkTargeting = 2
	/*MavMountModeRCTargeting - Load neutral position and start RC Roll,Pitch,Yaw control with stabilization */
	MavMountModeRCTargeting = 3
	/*MavMountModeGPSPoint - Load neutral position and start to point to Lat,Lon,Alt */
	MavMountModeGPSPoint = 4
	/*MavMountModeEnumEnd -  */
	MavMountModeEnumEnd = 5
)

/*UAVCAN_NODE_HEALTH - Generalized UAVCAN node health */
const (
	/*UavcanNodeHealthOk - The node is functioning properly. */
	UavcanNodeHealthOk = 0
	/*UavcanNodeHealthWarning - A critical parameter went out of range or the node has encountered a minor failure. */
	UavcanNodeHealthWarning = 1
	/*UavcanNodeHealthError - The node has encountered a major failure. */
	UavcanNodeHealthError = 2
	/*UavcanNodeHealthCritical - The node has suffered a fatal malfunction. */
	UavcanNodeHealthCritical = 3
	/*UavcanNodeHealthEnumEnd -  */
	UavcanNodeHealthEnumEnd = 4
)

/*UAVCAN_NODE_MODE - Generalized UAVCAN node mode */
const (
	/*UavcanNodeModeOperational - The node is performing its primary functions. */
	UavcanNodeModeOperational = 0
	/*UavcanNodeModeInitialization - The node is initializing; this mode is entered immediately after startup. */
	UavcanNodeModeInitialization = 1
	/*UavcanNodeModeMaintenance - The node is under maintenance. */
	UavcanNodeModeMaintenance = 2
	/*UavcanNodeModeSoftwareUpdate - The node is in the process of updating its software. */
	UavcanNodeModeSoftwareUpdate = 3
	/*UavcanNodeModeOffline - The node is no longer available online. */
	UavcanNodeModeOffline = 7
	/*UavcanNodeModeEnumEnd -  */
	UavcanNodeModeEnumEnd = 8
)

/*MAV_CMD - Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. */
const (
	/*MavCmdNAVWaypoint - Navigate to waypoint. */
	MavCmdNAVWaypoint = 16
	/*MavCmdNAVLoiterUnlim - Loiter around this waypoint an unlimited amount of time */
	MavCmdNAVLoiterUnlim = 17
	/*MavCmdNAVLoiterTurns - Loiter around this waypoint for X turns */
	MavCmdNAVLoiterTurns = 18
	/*MavCmdNAVLoiterTime - Loiter around this waypoint for X seconds */
	MavCmdNAVLoiterTime = 19
	/*MavCmdNAVReturnToLaunch - Return to launch location */
	MavCmdNAVReturnToLaunch = 20
	/*MavCmdNAVLand - Land at location */
	MavCmdNAVLand = 21
	/*MavCmdNAVTakeoff - Takeoff from ground / hand */
	MavCmdNAVTakeoff = 22
	/*MavCmdNAVLandLocal - Land at local position (local frame only) */
	MavCmdNAVLandLocal = 23
	/*MavCmdNAVTakeoffLocal - Takeoff from local position (local frame only) */
	MavCmdNAVTakeoffLocal = 24
	/*MavCmdNAVFollow - Vehicle following, i.e. this waypoint represents the position of a moving vehicle */
	MavCmdNAVFollow = 25
	/*MavCmdNAVContinueAndChangeAlt - Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached. */
	MavCmdNAVContinueAndChangeAlt = 30
	/*MavCmdNAVLoiterToAlt - Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached.  Additionally, if the Heading Required parameter is non-zero the  aircraft will not leave the loiter until heading toward the next waypoint. */
	MavCmdNAVLoiterToAlt = 31
	/*MavCmdDoFollow - Being following a target */
	MavCmdDoFollow = 32
	/*MavCmdDoFollowReposition - Reposition the MAV after a follow target command has been sent */
	MavCmdDoFollowReposition = 33
	/*MavCmdDoOrbit - Start orbiting on the circumference of a circle defined by the parameters. Setting any value NaN results in using defaults. */
	MavCmdDoOrbit = 34
	/*MavCmdNAVRoi - Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdNAVRoi = 80
	/*MavCmdNAVPathplanning - Control autonomous path planning on the MAV. */
	MavCmdNAVPathplanning = 81
	/*MavCmdNAVSplineWaypoint - Navigate to waypoint using a spline path. */
	MavCmdNAVSplineWaypoint = 82
	/*MavCmdNAVAltitudeWait - Mission command to wait for an altitude or downwards vertical speed. This is meant for high altitude balloon launches, allowing the aircraft to be idle until either an altitude is reached or a negative vertical speed is reached (indicating early balloon burst). The wiggle time is how often to wiggle the control surfaces to prevent them seizing up. */
	MavCmdNAVAltitudeWait = 83
	/*MavCmdNAVVtolTakeoff - Takeoff from ground using VTOL mode */
	MavCmdNAVVtolTakeoff = 84
	/*MavCmdNAVVtolLand - Land using VTOL mode */
	MavCmdNAVVtolLand = 85
	/*MavCmdNAVGUIDedEnable - hand control over to an external controller */
	MavCmdNAVGUIDedEnable = 92
	/*MavCmdNAVDelay - Delay the next navigation command a number of seconds or until a specified time */
	MavCmdNAVDelay = 93
	/*MavCmdNAVPayloadPlace - Descend and place payload.  Vehicle descends until it detects a hanging payload has reached the ground, the gripper is opened to release the payload */
	MavCmdNAVPayloadPlace = 94
	/*MavCmdNAVLast - NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration */
	MavCmdNAVLast = 95
	/*MavCmdConditionDelay - Delay mission state machine. */
	MavCmdConditionDelay = 112
	/*MavCmdConditionChangeAlt - Ascend/descend at rate.  Delay mission state machine until desired altitude reached. */
	MavCmdConditionChangeAlt = 113
	/*MavCmdConditionDistance - Delay mission state machine until within desired distance of next NAV point. */
	MavCmdConditionDistance = 114
	/*MavCmdConditionYaw - Reach a certain target angle. */
	MavCmdConditionYaw = 115
	/*MavCmdConditionLast - NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration */
	MavCmdConditionLast = 159
	/*MavCmdDoSetMode - Set system mode. */
	MavCmdDoSetMode = 176
	/*MavCmdDoJump - Jump to the desired command in the mission list.  Repeat this action only the specified number of times */
	MavCmdDoJump = 177
	/*MavCmdDoChangeSpeed - Change speed and/or throttle set points. */
	MavCmdDoChangeSpeed = 178
	/*MavCmdDoSetHome - Changes the home location either to the current location or a specified location. */
	MavCmdDoSetHome = 179
	/*MavCmdDoSetParameter - Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter. */
	MavCmdDoSetParameter = 180
	/*MavCmdDoSetRelay - Set a relay to a condition. */
	MavCmdDoSetRelay = 181
	/*MavCmdDoRepeatRelay - Cycle a relay on and off for a desired number of cycles with a desired period. */
	MavCmdDoRepeatRelay = 182
	/*MavCmdDoSetServo - Set a servo to a desired PWM value. */
	MavCmdDoSetServo = 183
	/*MavCmdDoRepeatServo - Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period. */
	MavCmdDoRepeatServo = 184
	/*MavCmdDoFlighttermination - Terminate flight immediately */
	MavCmdDoFlighttermination = 185
	/*MavCmdDoChangeAltitude - Change altitude set point. */
	MavCmdDoChangeAltitude = 186
	/*MavCmdDoLandStart - Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence. */
	MavCmdDoLandStart = 189
	/*MavCmdDoRallyLand - Mission command to perform a landing from a rally point. */
	MavCmdDoRallyLand = 190
	/*MavCmdDoGoAround - Mission command to safely abort an autonomous landing. */
	MavCmdDoGoAround = 191
	/*MavCmdDoReposition - Reposition the vehicle to a specific WGS84 global position. */
	MavCmdDoReposition = 192
	/*MavCmdDoPauseContinue - If in a GPS controlled position mode, hold the current position or continue. */
	MavCmdDoPauseContinue = 193
	/*MavCmdDoSetReverse - Set moving direction to forward or reverse. */
	MavCmdDoSetReverse = 194
	/*MavCmdDoSetRoiLocation - Sets the region of interest (ROI) to a location. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdDoSetRoiLocation = 195
	/*MavCmdDoSetRoiWpnextOffset - Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdDoSetRoiWpnextOffset = 196
	/*MavCmdDoSetRoiNone - Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdDoSetRoiNone = 197
	/*MavCmdDoControlVIDeo - Control onboard camera system. */
	MavCmdDoControlVIDeo = 200
	/*MavCmdDoSetRoi - Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdDoSetRoi = 201
	/*MavCmdDoDigicamConfigure - THIS INTERFACE IS DEPRECATED since 2018-01. Please use PARAM_EXT_XXX messages and the camera definition format described in https://mavlink.io/en/protocol/camera_def.html. */
	MavCmdDoDigicamConfigure = 202
	/*MavCmdDoDigicamControl - THIS INTERFACE IS DEPRECATED since 2018-01. Please use PARAM_EXT_XXX messages and the camera definition format described in https://mavlink.io/en/protocol/camera_def.html. */
	MavCmdDoDigicamControl = 203
	/*MavCmdDoMountConfigure - Mission command to configure a camera or antenna mount */
	MavCmdDoMountConfigure = 204
	/*MavCmdDoMountControl - Mission command to control a camera or antenna mount */
	MavCmdDoMountControl = 205
	/*MavCmdDoSetCamTriggDist - Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera. */
	MavCmdDoSetCamTriggDist = 206
	/*MavCmdDoFenceEnable - Mission command to enable the geofence */
	MavCmdDoFenceEnable = 207
	/*MavCmdDoParachute - Mission command to trigger a parachute */
	MavCmdDoParachute = 208
	/*MavCmdDoMotorTest - Mission command to perform motor test */
	MavCmdDoMotorTest = 209
	/*MavCmdDoInvertedFlight - Change to/from inverted flight */
	MavCmdDoInvertedFlight = 210
	/*MavCmdDoGripper - Mission command to operate EPM gripper. */
	MavCmdDoGripper = 211
	/*MavCmdDoAutotuneEnable - Enable/disable autotune. */
	MavCmdDoAutotuneEnable = 212
	/*MavCmdNAVSetYawSpeed - Sets a desired vehicle turn angle and speed change */
	MavCmdNAVSetYawSpeed = 213
	/*MavCmdDoSetCamTriggInterval - Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera. */
	MavCmdDoSetCamTriggInterval = 214
	/*MavCmdDoMountControlQuat - Mission command to control a camera or antenna mount, using a quaternion as reference. */
	MavCmdDoMountControlQuat = 220
	/*MavCmdDoGUIDedMaster - set id of master controller */
	MavCmdDoGUIDedMaster = 221
	/*MavCmdDoGUIDedLimits - Set limits for external control */
	MavCmdDoGUIDedLimits = 222
	/*MavCmdDoEngineControl - Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines */
	MavCmdDoEngineControl = 223
	/*MavCmdDoSetMissionCurrent - Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between). */
	MavCmdDoSetMissionCurrent = 224
	/*MavCmdDoLast - NOP - This command is only used to mark the upper limit of the DO commands in the enumeration */
	MavCmdDoLast = 240
	/*MavCmdPreflightCalibration - Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero. */
	MavCmdPreflightCalibration = 241
	/*MavCmdPreflightSetSensorOffsets - Set sensor offsets. This command will be only accepted if in pre-flight mode. */
	MavCmdPreflightSetSensorOffsets = 242
	/*MavCmdPreflightUavcan - Trigger UAVCAN config. This command will be only accepted if in pre-flight mode. */
	MavCmdPreflightUavcan = 243
	/*MavCmdPreflightStorage - Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode. */
	MavCmdPreflightStorage = 245
	/*MavCmdPreflightRebootShutdown - Request the reboot or shutdown of system components. */
	MavCmdPreflightRebootShutdown = 246
	/*MavCmdOverrIDeGoto - Hold / continue the current action */
	MavCmdOverrIDeGoto = 252
	/*MavCmdMissionStart - start running a mission */
	MavCmdMissionStart = 300
	/*MavCmdComponentArmDisarm - Arms / Disarms a component */
	MavCmdComponentArmDisarm = 400
	/*MavCmdGetHomePosition - Request the home position from the vehicle. */
	MavCmdGetHomePosition = 410
	/*MavCmdStartRxPair - Starts receiver pairing */
	MavCmdStartRxPair = 500
	/*MavCmdGetMessageInterval - Request the interval between messages for a particular MAVLink message ID */
	MavCmdGetMessageInterval = 510
	/*MavCmdSetMessageInterval - Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM */
	MavCmdSetMessageInterval = 511
	/*MavCmdRequestProtocolVersion - Request MAVLink protocol version compatibility */
	MavCmdRequestProtocolVersion = 519
	/*MavCmdRequestAutopilotCapabilities - Request autopilot capabilities */
	MavCmdRequestAutopilotCapabilities = 520
	/*MavCmdRequestCameraInformation - Request camera information (CAMERA_INFORMATION). */
	MavCmdRequestCameraInformation = 521
	/*MavCmdRequestCameraSettings - Request camera settings (CAMERA_SETTINGS). */
	MavCmdRequestCameraSettings = 522
	/*MavCmdRequestStorageInformation - Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage. */
	MavCmdRequestStorageInformation = 525
	/*MavCmdStorageFormat - Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage. */
	MavCmdStorageFormat = 526
	/*MavCmdRequestCameraCaptureStatus - Request camera capture status (CAMERA_CAPTURE_STATUS) */
	MavCmdRequestCameraCaptureStatus = 527
	/*MavCmdRequestFlightInformation - Request flight information (FLIGHT_INFORMATION) */
	MavCmdRequestFlightInformation = 528
	/*MavCmdResetCameraSettings - Reset all camera settings to Factory Default */
	MavCmdResetCameraSettings = 529
	/*MavCmdSetCameraMode - Set camera running mode. Use NAN for reserved values. */
	MavCmdSetCameraMode = 530
	/*MavCmdImageStartCapture - Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NAN for reserved values. */
	MavCmdImageStartCapture = 2000
	/*MavCmdImageStopCapture - Stop image capture sequence Use NAN for reserved values. */
	MavCmdImageStopCapture = 2001
	/*MavCmdRequestCameraImageCapture - Re-request a CAMERA_IMAGE_CAPTURE packet. Use NAN for reserved values. */
	MavCmdRequestCameraImageCapture = 2002
	/*MavCmdDoTriggerControl - Enable or disable on-board camera triggering system. */
	MavCmdDoTriggerControl = 2003
	/*MavCmdVIDeoStartCapture - Starts video capture (recording). Use NAN for reserved values. */
	MavCmdVIDeoStartCapture = 2500
	/*MavCmdVIDeoStopCapture - Stop the current video capture (recording). Use NAN for reserved values. */
	MavCmdVIDeoStopCapture = 2501
	/*MavCmdVIDeoStartStreaming - Start video streaming */
	MavCmdVIDeoStartStreaming = 2502
	/*MavCmdVIDeoStopStreaming - Stop the current video streaming */
	MavCmdVIDeoStopStreaming = 2503
	/*MavCmdRequestVIDeoStreamInformation - Request video stream information (VIDEO_STREAM_INFORMATION) */
	MavCmdRequestVIDeoStreamInformation = 2504
	/*MavCmdLoggingStart - Request to start streaming logging data over MAVLink (see also LOGGING_DATA message) */
	MavCmdLoggingStart = 2510
	/*MavCmdLoggingStop - Request to stop streaming log data over MAVLink */
	MavCmdLoggingStop = 2511
	/*MavCmdAirframeConfiguration -  */
	MavCmdAirframeConfiguration = 2520
	/*MavCmdControlHighLatency - Request to start/stop transmitting over the high latency telemetry */
	MavCmdControlHighLatency = 2600
	/*MavCmdPanoramaCreate - Create a panorama at the current position */
	MavCmdPanoramaCreate = 2800
	/*MavCmdDoVtolTransition - Request VTOL transition */
	MavCmdDoVtolTransition = 3000
	/*MavCmdArmAuthorizationRequest - Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	 */
	MavCmdArmAuthorizationRequest = 3001
	/*MavCmdSetGUIDedSubmodeStandard - This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	 */
	MavCmdSetGUIDedSubmodeStandard = 4000
	/*MavCmdSetGUIDedSubmodeCiRCle - This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	 */
	MavCmdSetGUIDedSubmodeCiRCle = 4001
	/*MavCmdConditionGate - Delay mission state machine until gate has been reached. */
	MavCmdConditionGate = 4501
	/*MavCmdNAVFenceReturnPoint - Fence return point. There can only be one fence return point.
	 */
	MavCmdNAVFenceReturnPoint = 5000
	/*MavCmdNAVFencePolygonVertexInclusion - Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	 */
	MavCmdNAVFencePolygonVertexInclusion = 5001
	/*MavCmdNAVFencePolygonVertexExclusion - Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	 */
	MavCmdNAVFencePolygonVertexExclusion = 5002
	/*MavCmdNAVFenceCiRCleInclusion - Circular fence area. The vehicle must stay inside this area.
	 */
	MavCmdNAVFenceCiRCleInclusion = 5003
	/*MavCmdNAVFenceCiRCleExclusion - Circular fence area. The vehicle must stay outside this area.
	 */
	MavCmdNAVFenceCiRCleExclusion = 5004
	/*MavCmdNAVRallyPoint - Rally point. You can have multiple rally points defined.
	 */
	MavCmdNAVRallyPoint = 5100
	/*MavCmdUavcanGetNodeInfo - Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages. */
	MavCmdUavcanGetNodeInfo = 5200
	/*MavCmdPayloadPrepareDeploy - Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity. */
	MavCmdPayloadPrepareDeploy = 30001
	/*MavCmdPayloadControlDeploy - Control the payload deployment. */
	MavCmdPayloadControlDeploy = 30002
	/*MavCmdWaypointUser1 - User defined waypoint item. Ground Station will show the Vehicle as flying through this item. */
	MavCmdWaypointUser1 = 31000
	/*MavCmdWaypointUser2 - User defined waypoint item. Ground Station will show the Vehicle as flying through this item. */
	MavCmdWaypointUser2 = 31001
	/*MavCmdWaypointUser3 - User defined waypoint item. Ground Station will show the Vehicle as flying through this item. */
	MavCmdWaypointUser3 = 31002
	/*MavCmdWaypointUser4 - User defined waypoint item. Ground Station will show the Vehicle as flying through this item. */
	MavCmdWaypointUser4 = 31003
	/*MavCmdWaypointUser5 - User defined waypoint item. Ground Station will show the Vehicle as flying through this item. */
	MavCmdWaypointUser5 = 31004
	/*MavCmdSpatialUser1 - User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. */
	MavCmdSpatialUser1 = 31005
	/*MavCmdSpatialUser2 - User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. */
	MavCmdSpatialUser2 = 31006
	/*MavCmdSpatialUser3 - User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. */
	MavCmdSpatialUser3 = 31007
	/*MavCmdSpatialUser4 - User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. */
	MavCmdSpatialUser4 = 31008
	/*MavCmdSpatialUser5 - User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item. */
	MavCmdSpatialUser5 = 31009
	/*MavCmdUser1 - User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. */
	MavCmdUser1 = 31010
	/*MavCmdUser2 - User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. */
	MavCmdUser2 = 31011
	/*MavCmdUser3 - User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. */
	MavCmdUser3 = 31012
	/*MavCmdUser4 - User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. */
	MavCmdUser4 = 31013
	/*MavCmdUser5 - User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item. */
	MavCmdUser5 = 31014
	/*MavCmdPowerOffInitiated - A system wide power-off event has been initiated. */
	MavCmdPowerOffInitiated = 42000
	/*MavCmdSoloBtnFlyClick - FLY button has been clicked. */
	MavCmdSoloBtnFlyClick = 42001
	/*MavCmdSoloBtnFlyHold - FLY button has been held for 1.5 seconds. */
	MavCmdSoloBtnFlyHold = 42002
	/*MavCmdSoloBtnPauseClick - PAUSE button has been clicked. */
	MavCmdSoloBtnPauseClick = 42003
	/*MavCmdFixedMagCal - Magnetometer calibration based on fixed position
	  in earth field given by inclination, declination and intensity. */
	MavCmdFixedMagCal = 42004
	/*MavCmdFixedMagCalField - Magnetometer calibration based on fixed expected field values in milliGauss. */
	MavCmdFixedMagCalField = 42005
	/*MavCmdDoStartMagCal - Initiate a magnetometer calibration. */
	MavCmdDoStartMagCal = 42424
	/*MavCmdDoAcceptMagCal - Initiate a magnetometer calibration. */
	MavCmdDoAcceptMagCal = 42425
	/*MavCmdDoCancelMagCal - Cancel a running magnetometer calibration. */
	MavCmdDoCancelMagCal = 42426
	/*MavCmdSetFactoryTestMode - Command autopilot to get into factory test/diagnostic mode. */
	MavCmdSetFactoryTestMode = 42427
	/*MavCmdDoSendBanner - Reply with the version banner. */
	MavCmdDoSendBanner = 42428
	/*MavCmdAccelcalVehiclePos - Used when doing accelerometer calibration. When sent to the GCS tells it what position to put the vehicle in. When sent to the vehicle says what position the vehicle is in. */
	MavCmdAccelcalVehiclePos = 42429
	/*MavCmdGimbalReset - Causes the gimbal to reset and boot as if it was just powered on. */
	MavCmdGimbalReset = 42501
	/*MavCmdGimbalAxisCalibrationStatus - Reports progress and success or failure of gimbal axis calibration procedure. */
	MavCmdGimbalAxisCalibrationStatus = 42502
	/*MavCmdGimbalRequestAxisCalibration - Starts commutation calibration on the gimbal. */
	MavCmdGimbalRequestAxisCalibration = 42503
	/*MavCmdGimbalFullReset - Erases gimbal application and parameters. */
	MavCmdGimbalFullReset = 42505
	/*MavCmdDoWinch - Command to operate winch. */
	MavCmdDoWinch = 42600
	/*MavCmdFlashBootloader - Update the bootloader */
	MavCmdFlashBootloader = 42650
	/*MavCmdEnumEnd -  */
	MavCmdEnumEnd = 42651
)

/*MAV_DATA_STREAM - A data stream is not a fixed set of messages, but rather a
  recommendation to the autopilot software. Individual autopilots may or may not obey
  the recommended messages. */
const (
	/*MavDataStreamAll - Enable all data streams */
	MavDataStreamAll = 0
	/*MavDataStreamRawSensors - Enable IMU_RAW, GPS_RAW, GPS_STATUS packets. */
	MavDataStreamRawSensors = 1
	/*MavDataStreamExtendedStatus - Enable GPS_STATUS, CONTROL_STATUS, AUX_STATUS */
	MavDataStreamExtendedStatus = 2
	/*MavDataStreamRCChannels - Enable RC_CHANNELS_SCALED, RC_CHANNELS_RAW, SERVO_OUTPUT_RAW */
	MavDataStreamRCChannels = 3
	/*MavDataStreamRawController - Enable ATTITUDE_CONTROLLER_OUTPUT, POSITION_CONTROLLER_OUTPUT, NAV_CONTROLLER_OUTPUT. */
	MavDataStreamRawController = 4
	/*MavDataStreamPosition - Enable LOCAL_POSITION, GLOBAL_POSITION/GLOBAL_POSITION_INT messages. */
	MavDataStreamPosition = 6
	/*MavDataStreamExtra1 - Dependent on the autopilot */
	MavDataStreamExtra1 = 10
	/*MavDataStreamExtra2 - Dependent on the autopilot */
	MavDataStreamExtra2 = 11
	/*MavDataStreamExtra3 - Dependent on the autopilot */
	MavDataStreamExtra3 = 12
	/*MavDataStreamEnumEnd -  */
	MavDataStreamEnumEnd = 13
)

/*MAV_ROI - The ROI (region of interest) for the vehicle. This can be
  be used by the vehicle for camera/vehicle attitude alignment (see
  MAV_CMD_NAV_ROI). */
const (
	/*MavRoiNone - No region of interest. */
	MavRoiNone = 0
	/*MavRoiWpnext - Point toward next waypoint, with optional pitch/roll/yaw offset. */
	MavRoiWpnext = 1
	/*MavRoiWpindex - Point toward given waypoint. */
	MavRoiWpindex = 2
	/*MavRoiLocation - Point toward fixed location. */
	MavRoiLocation = 3
	/*MavRoiTarget - Point toward of given id. */
	MavRoiTarget = 4
	/*MavRoiEnumEnd -  */
	MavRoiEnumEnd = 5
)

/*MAV_CMD_ACK - ACK / NACK / ERROR values as a result of MAV_CMDs and for mission item transmission. */
const (
	/*MavCmdAckOk - Command / mission item is ok. */
	MavCmdAckOk = 1
	/*MavCmdAckErrFail - Generic error message if none of the other reasons fails or if no detailed error reporting is implemented. */
	MavCmdAckErrFail = 2
	/*MavCmdAckErrAccessDenied - The system is refusing to accept this command from this source / communication partner. */
	MavCmdAckErrAccessDenied = 3
	/*MavCmdAckErrNotSupported - Command or mission item is not supported, other commands would be accepted. */
	MavCmdAckErrNotSupported = 4
	/*MavCmdAckErrCoordinateFrameNotSupported - The coordinate frame of this command / mission item is not supported. */
	MavCmdAckErrCoordinateFrameNotSupported = 5
	/*MavCmdAckErrCoordinatesOutOfRange - The coordinate frame of this command is ok, but he coordinate values exceed the safety limits of this system. This is a generic error, please use the more specific error messages below if possible. */
	MavCmdAckErrCoordinatesOutOfRange = 6
	/*MavCmdAckErrXLatOutOfRange - The X or latitude value is out of range. */
	MavCmdAckErrXLatOutOfRange = 7
	/*MavCmdAckErrYLonOutOfRange - The Y or longitude value is out of range. */
	MavCmdAckErrYLonOutOfRange = 8
	/*MavCmdAckErrZAltOutOfRange - The Z or altitude value is out of range. */
	MavCmdAckErrZAltOutOfRange = 9
	/*MavCmdAckEnumEnd -  */
	MavCmdAckEnumEnd = 10
)

/*MAV_PARAM_TYPE - Specifies the datatype of a MAVLink parameter. */
const (
	/*MavParamTypeUint8 - 8-bit unsigned integer */
	MavParamTypeUint8 = 1
	/*MavParamTypeInt8 - 8-bit signed integer */
	MavParamTypeInt8 = 2
	/*MavParamTypeUint16 - 16-bit unsigned integer */
	MavParamTypeUint16 = 3
	/*MavParamTypeInt16 - 16-bit signed integer */
	MavParamTypeInt16 = 4
	/*MavParamTypeUint32 - 32-bit unsigned integer */
	MavParamTypeUint32 = 5
	/*MavParamTypeInt32 - 32-bit signed integer */
	MavParamTypeInt32 = 6
	/*MavParamTypeUint64 - 64-bit unsigned integer */
	MavParamTypeUint64 = 7
	/*MavParamTypeInt64 - 64-bit signed integer */
	MavParamTypeInt64 = 8
	/*MavParamTypeReal32 - 32-bit floating-point */
	MavParamTypeReal32 = 9
	/*MavParamTypeReal64 - 64-bit floating-point */
	MavParamTypeReal64 = 10
	/*MavParamTypeEnumEnd -  */
	MavParamTypeEnumEnd = 11
)

/*MAV_PARAM_EXT_TYPE - Specifies the datatype of a MAVLink extended parameter. */
const (
	/*MavParamExtTypeUint8 - 8-bit unsigned integer */
	MavParamExtTypeUint8 = 1
	/*MavParamExtTypeInt8 - 8-bit signed integer */
	MavParamExtTypeInt8 = 2
	/*MavParamExtTypeUint16 - 16-bit unsigned integer */
	MavParamExtTypeUint16 = 3
	/*MavParamExtTypeInt16 - 16-bit signed integer */
	MavParamExtTypeInt16 = 4
	/*MavParamExtTypeUint32 - 32-bit unsigned integer */
	MavParamExtTypeUint32 = 5
	/*MavParamExtTypeInt32 - 32-bit signed integer */
	MavParamExtTypeInt32 = 6
	/*MavParamExtTypeUint64 - 64-bit unsigned integer */
	MavParamExtTypeUint64 = 7
	/*MavParamExtTypeInt64 - 64-bit signed integer */
	MavParamExtTypeInt64 = 8
	/*MavParamExtTypeReal32 - 32-bit floating-point */
	MavParamExtTypeReal32 = 9
	/*MavParamExtTypeReal64 - 64-bit floating-point */
	MavParamExtTypeReal64 = 10
	/*MavParamExtTypeCustom - Custom Type */
	MavParamExtTypeCustom = 11
	/*MavParamExtTypeEnumEnd -  */
	MavParamExtTypeEnumEnd = 12
)

/*MAV_RESULT - result from a mavlink command */
const (
	/*MavResultAccepted - Command ACCEPTED and EXECUTED */
	MavResultAccepted = 0
	/*MavResultTemporarilyRejected - Command TEMPORARY REJECTED/DENIED */
	MavResultTemporarilyRejected = 1
	/*MavResultDenied - Command PERMANENTLY DENIED */
	MavResultDenied = 2
	/*MavResultUnsupported - Command UNKNOWN/UNSUPPORTED */
	MavResultUnsupported = 3
	/*MavResultFailed - Command executed, but failed */
	MavResultFailed = 4
	/*MavResultInProgress - WIP: Command being executed */
	MavResultInProgress = 5
	/*MavResultEnumEnd -  */
	MavResultEnumEnd = 6
)

/*MAV_MISSION_RESULT - result in a MAVLink mission ack */
const (
	/*MavMissionAccepted - mission accepted OK */
	MavMissionAccepted = 0
	/*MavMissionError - generic error / not accepting mission commands at all right now */
	MavMissionError = 1
	/*MavMissionUnsupportedFrame - coordinate frame is not supported */
	MavMissionUnsupportedFrame = 2
	/*MavMissionUnsupported - command is not supported */
	MavMissionUnsupported = 3
	/*MavMissionNoSpace - mission item exceeds storage space */
	MavMissionNoSpace = 4
	/*MavMissionInvalID - one of the parameters has an invalid value */
	MavMissionInvalID = 5
	/*MavMissionInvalIDParam1 - param1 has an invalid value */
	MavMissionInvalIDParam1 = 6
	/*MavMissionInvalIDParam2 - param2 has an invalid value */
	MavMissionInvalIDParam2 = 7
	/*MavMissionInvalIDParam3 - param3 has an invalid value */
	MavMissionInvalIDParam3 = 8
	/*MavMissionInvalIDParam4 - param4 has an invalid value */
	MavMissionInvalIDParam4 = 9
	/*MavMissionInvalIDParam5X - x/param5 has an invalid value */
	MavMissionInvalIDParam5X = 10
	/*MavMissionInvalIDParam6Y - y/param6 has an invalid value */
	MavMissionInvalIDParam6Y = 11
	/*MavMissionInvalIDParam7 - param7 has an invalid value */
	MavMissionInvalIDParam7 = 12
	/*MavMissionInvalIDSequence - received waypoint out of sequence */
	MavMissionInvalIDSequence = 13
	/*MavMissionDenied - not accepting any mission commands from this communication partner */
	MavMissionDenied = 14
	/*MavMissionResultEnumEnd -  */
	MavMissionResultEnumEnd = 15
)

/*MAV_SEVERITY - Indicates the severity level, generally used for status messages to indicate their relative urgency. Based on RFC-5424 using expanded definitions at: http://www.kiwisyslog.com/kb/info:-syslog-message-levels/. */
const (
	/*MavSeverityEmergency - System is unusable. This is a "panic" condition. */
	MavSeverityEmergency = 0
	/*MavSeverityAlert - Action should be taken immediately. Indicates error in non-critical systems. */
	MavSeverityAlert = 1
	/*MavSeverityCritical - Action must be taken immediately. Indicates failure in a primary system. */
	MavSeverityCritical = 2
	/*MavSeverityError - Indicates an error in secondary/redundant systems. */
	MavSeverityError = 3
	/*MavSeverityWarning - Indicates about a possible future error if this is not resolved within a given timeframe. Example would be a low battery warning. */
	MavSeverityWarning = 4
	/*MavSeverityNotice - An unusual event has occurred, though not an error condition. This should be investigated for the root cause. */
	MavSeverityNotice = 5
	/*MavSeverityInfo - Normal operational messages. Useful for logging. No action is required for these messages. */
	MavSeverityInfo = 6
	/*MavSeverityDebug - Useful non-operational messages that can assist in debugging. These should not occur during normal operation. */
	MavSeverityDebug = 7
	/*MavSeverityEnumEnd -  */
	MavSeverityEnumEnd = 8
)

/*MAV_POWER_STATUS - Power supply status flags (bitmask) */
const (
	/*MavPowerStatusBrickValID - main brick power supply valid */
	MavPowerStatusBrickValID = 1
	/*MavPowerStatusServoValID - main servo power supply valid for FMU */
	MavPowerStatusServoValID = 2
	/*MavPowerStatusUsbConnected - USB power is connected */
	MavPowerStatusUsbConnected = 4
	/*MavPowerStatusPeriphOveRCurrent - peripheral supply is in over-current state */
	MavPowerStatusPeriphOveRCurrent = 8
	/*MavPowerStatusPeriphHipowerOveRCurrent - hi-power peripheral supply is in over-current state */
	MavPowerStatusPeriphHipowerOveRCurrent = 16
	/*MavPowerStatusChanged - Power status has changed since boot */
	MavPowerStatusChanged = 32
	/*MavPowerStatusEnumEnd -  */
	MavPowerStatusEnumEnd = 33
)

/*SERIAL_CONTROL_DEV - SERIAL_CONTROL device types */
const (
	/*SerialControlDevTelem1 - First telemetry port */
	SerialControlDevTelem1 = 0
	/*SerialControlDevTelem2 - Second telemetry port */
	SerialControlDevTelem2 = 1
	/*SerialControlDevGPS1 - First GPS port */
	SerialControlDevGPS1 = 2
	/*SerialControlDevGPS2 - Second GPS port */
	SerialControlDevGPS2 = 3
	/*SerialControlDevShell - system shell */
	SerialControlDevShell = 10
	/*SerialControlDevEnumEnd -  */
	SerialControlDevEnumEnd = 11
)

/*SERIAL_CONTROL_FLAG - SERIAL_CONTROL flags (bitmask) */
const (
	/*SerialControlFlagReply - Set if this is a reply */
	SerialControlFlagReply = 1
	/*SerialControlFlagRespond - Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message */
	SerialControlFlagRespond = 2
	/*SerialControlFlagExclusive - Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set */
	SerialControlFlagExclusive = 4
	/*SerialControlFlagBlocking - Block on writes to the serial port */
	SerialControlFlagBlocking = 8
	/*SerialControlFlagMulti - Send multiple replies until port is drained */
	SerialControlFlagMulti = 16
	/*SerialControlFlagEnumEnd -  */
	SerialControlFlagEnumEnd = 17
)

/*MAV_DISTANCE_SENSOR - Enumeration of distance sensor types */
const (
	/*MavDistanceSensorLaser - Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units */
	MavDistanceSensorLaser = 0
	/*MavDistanceSensorUltrasound - Ultrasound rangefinder, e.g. MaxBotix units */
	MavDistanceSensorUltrasound = 1
	/*MavDistanceSensorInfrared - Infrared rangefinder, e.g. Sharp units */
	MavDistanceSensorInfrared = 2
	/*MavDistanceSensorRadar - Radar type, e.g. uLanding units */
	MavDistanceSensorRadar = 3
	/*MavDistanceSensorUnknown - Broken or unknown type, e.g. analog units */
	MavDistanceSensorUnknown = 4
	/*MavDistanceSensorEnumEnd -  */
	MavDistanceSensorEnumEnd = 5
)

/*MAV_SENSOR_ORIENTATION - Enumeration of sensor orientation, according to its rotations */
const (
	/*MavSensorRotationNone - Roll: 0, Pitch: 0, Yaw: 0 */
	MavSensorRotationNone = 0
	/*MavSensorRotationYaw45 - Roll: 0, Pitch: 0, Yaw: 45 */
	MavSensorRotationYaw45 = 1
	/*MavSensorRotationYaw90 - Roll: 0, Pitch: 0, Yaw: 90 */
	MavSensorRotationYaw90 = 2
	/*MavSensorRotationYaw135 - Roll: 0, Pitch: 0, Yaw: 135 */
	MavSensorRotationYaw135 = 3
	/*MavSensorRotationYaw180 - Roll: 0, Pitch: 0, Yaw: 180 */
	MavSensorRotationYaw180 = 4
	/*MavSensorRotationYaw225 - Roll: 0, Pitch: 0, Yaw: 225 */
	MavSensorRotationYaw225 = 5
	/*MavSensorRotationYaw270 - Roll: 0, Pitch: 0, Yaw: 270 */
	MavSensorRotationYaw270 = 6
	/*MavSensorRotationYaw315 - Roll: 0, Pitch: 0, Yaw: 315 */
	MavSensorRotationYaw315 = 7
	/*MavSensorRotationRoll180 - Roll: 180, Pitch: 0, Yaw: 0 */
	MavSensorRotationRoll180 = 8
	/*MavSensorRotationRoll180Yaw45 - Roll: 180, Pitch: 0, Yaw: 45 */
	MavSensorRotationRoll180Yaw45 = 9
	/*MavSensorRotationRoll180Yaw90 - Roll: 180, Pitch: 0, Yaw: 90 */
	MavSensorRotationRoll180Yaw90 = 10
	/*MavSensorRotationRoll180Yaw135 - Roll: 180, Pitch: 0, Yaw: 135 */
	MavSensorRotationRoll180Yaw135 = 11
	/*MavSensorRotationPitch180 - Roll: 0, Pitch: 180, Yaw: 0 */
	MavSensorRotationPitch180 = 12
	/*MavSensorRotationRoll180Yaw225 - Roll: 180, Pitch: 0, Yaw: 225 */
	MavSensorRotationRoll180Yaw225 = 13
	/*MavSensorRotationRoll180Yaw270 - Roll: 180, Pitch: 0, Yaw: 270 */
	MavSensorRotationRoll180Yaw270 = 14
	/*MavSensorRotationRoll180Yaw315 - Roll: 180, Pitch: 0, Yaw: 315 */
	MavSensorRotationRoll180Yaw315 = 15
	/*MavSensorRotationRoll90 - Roll: 90, Pitch: 0, Yaw: 0 */
	MavSensorRotationRoll90 = 16
	/*MavSensorRotationRoll90Yaw45 - Roll: 90, Pitch: 0, Yaw: 45 */
	MavSensorRotationRoll90Yaw45 = 17
	/*MavSensorRotationRoll90Yaw90 - Roll: 90, Pitch: 0, Yaw: 90 */
	MavSensorRotationRoll90Yaw90 = 18
	/*MavSensorRotationRoll90Yaw135 - Roll: 90, Pitch: 0, Yaw: 135 */
	MavSensorRotationRoll90Yaw135 = 19
	/*MavSensorRotationRoll270 - Roll: 270, Pitch: 0, Yaw: 0 */
	MavSensorRotationRoll270 = 20
	/*MavSensorRotationRoll270Yaw45 - Roll: 270, Pitch: 0, Yaw: 45 */
	MavSensorRotationRoll270Yaw45 = 21
	/*MavSensorRotationRoll270Yaw90 - Roll: 270, Pitch: 0, Yaw: 90 */
	MavSensorRotationRoll270Yaw90 = 22
	/*MavSensorRotationRoll270Yaw135 - Roll: 270, Pitch: 0, Yaw: 135 */
	MavSensorRotationRoll270Yaw135 = 23
	/*MavSensorRotationPitch90 - Roll: 0, Pitch: 90, Yaw: 0 */
	MavSensorRotationPitch90 = 24
	/*MavSensorRotationPitch270 - Roll: 0, Pitch: 270, Yaw: 0 */
	MavSensorRotationPitch270 = 25
	/*MavSensorRotationPitch180Yaw90 - Roll: 0, Pitch: 180, Yaw: 90 */
	MavSensorRotationPitch180Yaw90 = 26
	/*MavSensorRotationPitch180Yaw270 - Roll: 0, Pitch: 180, Yaw: 270 */
	MavSensorRotationPitch180Yaw270 = 27
	/*MavSensorRotationRoll90Pitch90 - Roll: 90, Pitch: 90, Yaw: 0 */
	MavSensorRotationRoll90Pitch90 = 28
	/*MavSensorRotationRoll180Pitch90 - Roll: 180, Pitch: 90, Yaw: 0 */
	MavSensorRotationRoll180Pitch90 = 29
	/*MavSensorRotationRoll270Pitch90 - Roll: 270, Pitch: 90, Yaw: 0 */
	MavSensorRotationRoll270Pitch90 = 30
	/*MavSensorRotationRoll90Pitch180 - Roll: 90, Pitch: 180, Yaw: 0 */
	MavSensorRotationRoll90Pitch180 = 31
	/*MavSensorRotationRoll270Pitch180 - Roll: 270, Pitch: 180, Yaw: 0 */
	MavSensorRotationRoll270Pitch180 = 32
	/*MavSensorRotationRoll90Pitch270 - Roll: 90, Pitch: 270, Yaw: 0 */
	MavSensorRotationRoll90Pitch270 = 33
	/*MavSensorRotationRoll180Pitch270 - Roll: 180, Pitch: 270, Yaw: 0 */
	MavSensorRotationRoll180Pitch270 = 34
	/*MavSensorRotationRoll270Pitch270 - Roll: 270, Pitch: 270, Yaw: 0 */
	MavSensorRotationRoll270Pitch270 = 35
	/*MavSensorRotationRoll90Pitch180Yaw90 - Roll: 90, Pitch: 180, Yaw: 90 */
	MavSensorRotationRoll90Pitch180Yaw90 = 36
	/*MavSensorRotationRoll90Yaw270 - Roll: 90, Pitch: 0, Yaw: 270 */
	MavSensorRotationRoll90Yaw270 = 37
	/*MavSensorRotationRoll90Pitch68Yaw293 - Roll: 90, Pitch: 68, Yaw: 293 */
	MavSensorRotationRoll90Pitch68Yaw293 = 38
	/*MavSensorRotationPitch315 - Pitch: 315 */
	MavSensorRotationPitch315 = 39
	/*MavSensorRotationRoll90Pitch315 - Roll: 90, Pitch: 315 */
	MavSensorRotationRoll90Pitch315 = 40
	/*MavSensorRotationCustom - Custom orientation */
	MavSensorRotationCustom = 100
	/*MavSensorOrientationEnumEnd -  */
	MavSensorOrientationEnumEnd = 101
)

/*MAV_PROTOCOL_CAPABILITY - Bitmask of (optional) autopilot capabilities (64 bit). If a bit is set, the autopilot supports this capability. */
const (
	/*MavProtocolCapabilityMissionFloat - Autopilot supports MISSION float message type. */
	MavProtocolCapabilityMissionFloat = 1
	/*MavProtocolCapabilityParamFloat - Autopilot supports the new param float message type. */
	MavProtocolCapabilityParamFloat = 2
	/*MavProtocolCapabilityMissionInt - Autopilot supports MISSION_INT scaled integer message type. */
	MavProtocolCapabilityMissionInt = 4
	/*MavProtocolCapabilityCommandInt - Autopilot supports COMMAND_INT scaled integer message type. */
	MavProtocolCapabilityCommandInt = 8
	/*MavProtocolCapabilityParamUnion - Autopilot supports the new param union message type. */
	MavProtocolCapabilityParamUnion = 16
	/*MavProtocolCapabilityFtp - Autopilot supports the new FILE_TRANSFER_PROTOCOL message type. */
	MavProtocolCapabilityFtp = 32
	/*MavProtocolCapabilitySetAttitudeTarget - Autopilot supports commanding attitude offboard. */
	MavProtocolCapabilitySetAttitudeTarget = 64
	/*MavProtocolCapabilitySetPositionTargetLocalNed - Autopilot supports commanding position and velocity targets in local NED frame. */
	MavProtocolCapabilitySetPositionTargetLocalNed = 128
	/*MavProtocolCapabilitySetPositionTargetGlobalInt - Autopilot supports commanding position and velocity targets in global scaled integers. */
	MavProtocolCapabilitySetPositionTargetGlobalInt = 256
	/*MavProtocolCapabilityTerrain - Autopilot supports terrain protocol / data handling. */
	MavProtocolCapabilityTerrain = 512
	/*MavProtocolCapabilitySetActuatorTarget - Autopilot supports direct actuator control. */
	MavProtocolCapabilitySetActuatorTarget = 1024
	/*MavProtocolCapabilityFlightTermination - Autopilot supports the flight termination command. */
	MavProtocolCapabilityFlightTermination = 2048
	/*MavProtocolCapabilityCompassCalibration - Autopilot supports onboard compass calibration. */
	MavProtocolCapabilityCompassCalibration = 4096
	/*MavProtocolCapabilityMavlink2 - Autopilot supports MAVLink version 2. */
	MavProtocolCapabilityMavlink2 = 8192
	/*MavProtocolCapabilityMissionFence - Autopilot supports mission fence protocol. */
	MavProtocolCapabilityMissionFence = 16384
	/*MavProtocolCapabilityMissionRally - Autopilot supports mission rally point protocol. */
	MavProtocolCapabilityMissionRally = 32768
	/*MavProtocolCapabilityFlightInformation - Autopilot supports the flight information protocol. */
	MavProtocolCapabilityFlightInformation = 65536
	/*MavProtocolCapabilityEnumEnd -  */
	MavProtocolCapabilityEnumEnd = 65537
)

/*MAV_MISSION_TYPE - Type of mission items being requested/sent in mission protocol. */
const (
	/*MavMissionTypeMission - Items are mission commands for main mission. */
	MavMissionTypeMission = 0
	/*MavMissionTypeFence - Specifies GeoFence area(s). Items are MAV_CMD_FENCE_ GeoFence items. */
	MavMissionTypeFence = 1
	/*MavMissionTypeRally - Specifies the rally points for the vehicle. Rally points are alternative RTL points. Items are MAV_CMD_RALLY_POINT rally point items. */
	MavMissionTypeRally = 2
	/*MavMissionTypeAll - Only used in MISSION_CLEAR_ALL to clear all mission types. */
	MavMissionTypeAll = 255
	/*MavMissionTypeEnumEnd -  */
	MavMissionTypeEnumEnd = 256
)

/*MAV_ESTIMATOR_TYPE - Enumeration of estimator types */
const (
	/*MavEstimatorTypeNaive - This is a naive estimator without any real covariance feedback. */
	MavEstimatorTypeNaive = 1
	/*MavEstimatorTypeVision - Computer vision based estimate. Might be up to scale. */
	MavEstimatorTypeVision = 2
	/*MavEstimatorTypeVio - Visual-inertial estimate. */
	MavEstimatorTypeVio = 3
	/*MavEstimatorTypeGPS - Plain GPS estimate. */
	MavEstimatorTypeGPS = 4
	/*MavEstimatorTypeGPSIns - Estimator integrating GPS and inertial sensing. */
	MavEstimatorTypeGPSIns = 5
	/*MavEstimatorTypeEnumEnd -  */
	MavEstimatorTypeEnumEnd = 6
)

/*MAV_BATTERY_TYPE - Enumeration of battery types */
const (
	/*MavBatteryTypeUnknown - Not specified. */
	MavBatteryTypeUnknown = 0
	/*MavBatteryTypeLipo - Lithium polymer battery */
	MavBatteryTypeLipo = 1
	/*MavBatteryTypeLife - Lithium-iron-phosphate battery */
	MavBatteryTypeLife = 2
	/*MavBatteryTypeLion - Lithium-ION battery */
	MavBatteryTypeLion = 3
	/*MavBatteryTypeNimh - Nickel metal hydride battery */
	MavBatteryTypeNimh = 4
	/*MavBatteryTypeEnumEnd -  */
	MavBatteryTypeEnumEnd = 5
)

/*MAV_BATTERY_FUNCTION - Enumeration of battery functions */
const (
	/*MavBatteryFunctionUnknown - Battery function is unknown */
	MavBatteryFunctionUnknown = 0
	/*MavBatteryFunctionAll - Battery supports all flight systems */
	MavBatteryFunctionAll = 1
	/*MavBatteryFunctionPropulsion - Battery for the propulsion system */
	MavBatteryFunctionPropulsion = 2
	/*MavBatteryFunctionAvionics - Avionics battery */
	MavBatteryFunctionAvionics = 3
	/*MavBatteryTypePayload - Payload battery */
	MavBatteryTypePayload = 4
	/*MavBatteryFunctionEnumEnd -  */
	MavBatteryFunctionEnumEnd = 5
)

/*MAV_BATTERY_CHARGE_STATE - Enumeration for low battery states. */
const (
	/*MavBatteryChargeStateUndefined - Low battery state is not provided */
	MavBatteryChargeStateUndefined = 0
	/*MavBatteryChargeStateOk - Battery is not in low state. Normal operation. */
	MavBatteryChargeStateOk = 1
	/*MavBatteryChargeStateLow - Battery state is low, warn and monitor close. */
	MavBatteryChargeStateLow = 2
	/*MavBatteryChargeStateCritical - Battery state is critical, return or abort immediately. */
	MavBatteryChargeStateCritical = 3
	/*MavBatteryChargeStateEmergency - Battery state is too low for ordinary abort sequence. Perform fastest possible emergency stop to prevent damage. */
	MavBatteryChargeStateEmergency = 4
	/*MavBatteryChargeStateFailed - Battery failed, damage unavoidable. */
	MavBatteryChargeStateFailed = 5
	/*MavBatteryChargeStateUnhealthy - Battery is diagnosed to be defective or an error occurred, usage is discouraged / prohibited. */
	MavBatteryChargeStateUnhealthy = 6
	/*MavBatteryChargeStateEnumEnd -  */
	MavBatteryChargeStateEnumEnd = 7
)

/*MAV_VTOL_STATE - Enumeration of VTOL states */
const (
	/*MavVtolStateUndefined - MAV is not configured as VTOL */
	MavVtolStateUndefined = 0
	/*MavVtolStateTransitionToFw - VTOL is in transition from multicopter to fixed-wing */
	MavVtolStateTransitionToFw = 1
	/*MavVtolStateTransitionToMc - VTOL is in transition from fixed-wing to multicopter */
	MavVtolStateTransitionToMc = 2
	/*MavVtolStateMc - VTOL is in multicopter state */
	MavVtolStateMc = 3
	/*MavVtolStateFw - VTOL is in fixed-wing state */
	MavVtolStateFw = 4
	/*MavVtolStateEnumEnd -  */
	MavVtolStateEnumEnd = 5
)

/*MAV_LANDED_STATE - Enumeration of landed detector states */
const (
	/*MavLandedStateUndefined - MAV landed state is unknown */
	MavLandedStateUndefined = 0
	/*MavLandedStateOnGround - MAV is landed (on ground) */
	MavLandedStateOnGround = 1
	/*MavLandedStateInAir - MAV is in air */
	MavLandedStateInAir = 2
	/*MavLandedStateTakeoff - MAV currently taking off */
	MavLandedStateTakeoff = 3
	/*MavLandedStateLanding - MAV currently landing */
	MavLandedStateLanding = 4
	/*MavLandedStateEnumEnd -  */
	MavLandedStateEnumEnd = 5
)

/*ADSB_ALTITUDE_TYPE - Enumeration of the ADSB altimeter types */
const (
	/*AdsbAltitudeTypePressureQnh - Altitude reported from a Baro source using QNH reference */
	AdsbAltitudeTypePressureQnh = 0
	/*AdsbAltitudeTypeGeometric - Altitude reported from a GNSS source */
	AdsbAltitudeTypeGeometric = 1
	/*AdsbAltitudeTypeEnumEnd -  */
	AdsbAltitudeTypeEnumEnd = 2
)

/*ADSB_EMITTER_TYPE - ADSB classification for the type of vehicle emitting the transponder signal */
const (
	/*AdsbEmitterTypeNoInfo -  */
	AdsbEmitterTypeNoInfo = 0
	/*AdsbEmitterTypeLight -  */
	AdsbEmitterTypeLight = 1
	/*AdsbEmitterTypeSmall -  */
	AdsbEmitterTypeSmall = 2
	/*AdsbEmitterTypeLarge -  */
	AdsbEmitterTypeLarge = 3
	/*AdsbEmitterTypeHighVortexLarge -  */
	AdsbEmitterTypeHighVortexLarge = 4
	/*AdsbEmitterTypeHeavy -  */
	AdsbEmitterTypeHeavy = 5
	/*AdsbEmitterTypeHighlyManuv -  */
	AdsbEmitterTypeHighlyManuv = 6
	/*AdsbEmitterTypeRotocraft -  */
	AdsbEmitterTypeRotocraft = 7
	/*AdsbEmitterTypeUnassigned -  */
	AdsbEmitterTypeUnassigned = 8
	/*AdsbEmitterTypeGlIDer -  */
	AdsbEmitterTypeGlIDer = 9
	/*AdsbEmitterTypeLighterAir -  */
	AdsbEmitterTypeLighterAir = 10
	/*AdsbEmitterTypeParachute -  */
	AdsbEmitterTypeParachute = 11
	/*AdsbEmitterTypeUltraLight -  */
	AdsbEmitterTypeUltraLight = 12
	/*AdsbEmitterTypeUnassigned2 -  */
	AdsbEmitterTypeUnassigned2 = 13
	/*AdsbEmitterTypeUav -  */
	AdsbEmitterTypeUav = 14
	/*AdsbEmitterTypeSpace -  */
	AdsbEmitterTypeSpace = 15
	/*AdsbEmitterTypeUnassgined3 -  */
	AdsbEmitterTypeUnassgined3 = 16
	/*AdsbEmitterTypeEmergencySurface -  */
	AdsbEmitterTypeEmergencySurface = 17
	/*AdsbEmitterTypeServiceSurface -  */
	AdsbEmitterTypeServiceSurface = 18
	/*AdsbEmitterTypePointObstacle -  */
	AdsbEmitterTypePointObstacle = 19
	/*AdsbEmitterTypeEnumEnd -  */
	AdsbEmitterTypeEnumEnd = 20
)

/*ADSB_FLAGS - These flags indicate status such as data validity of each data source. Set = data valid */
const (
	/*AdsbFlagsValIDCoords -  */
	AdsbFlagsValIDCoords = 1
	/*AdsbFlagsValIDAltitude -  */
	AdsbFlagsValIDAltitude = 2
	/*AdsbFlagsValIDHeading -  */
	AdsbFlagsValIDHeading = 4
	/*AdsbFlagsValIDVelocity -  */
	AdsbFlagsValIDVelocity = 8
	/*AdsbFlagsValIDCallsign -  */
	AdsbFlagsValIDCallsign = 16
	/*AdsbFlagsValIDSquawk -  */
	AdsbFlagsValIDSquawk = 32
	/*AdsbFlagsSIMUlated -  */
	AdsbFlagsSIMUlated = 64
	/*AdsbFlagsEnumEnd -  */
	AdsbFlagsEnumEnd = 65
)

/*MAV_DO_REPOSITION_FLAGS - Bitmap of options for the MAV_CMD_DO_REPOSITION */
const (
	/*MavDoRepositionFlagsChangeMode - The aircraft should immediately transition into guided. This should not be set for follow me applications */
	MavDoRepositionFlagsChangeMode = 1
	/*MavDoRepositionFlagsEnumEnd -  */
	MavDoRepositionFlagsEnumEnd = 2
)

/*ESTIMATOR_STATUS_FLAGS - Flags in EKF_STATUS message */
const (
	/*EstimatorAttitude - True if the attitude estimate is good */
	EstimatorAttitude = 1
	/*EstimatorVelocityHoriz - True if the horizontal velocity estimate is good */
	EstimatorVelocityHoriz = 2
	/*EstimatorVelocityVert - True if the  vertical velocity estimate is good */
	EstimatorVelocityVert = 4
	/*EstimatorPosHorizRel - True if the horizontal position (relative) estimate is good */
	EstimatorPosHorizRel = 8
	/*EstimatorPosHorizAbs - True if the horizontal position (absolute) estimate is good */
	EstimatorPosHorizAbs = 16
	/*EstimatorPosVertAbs - True if the vertical position (absolute) estimate is good */
	EstimatorPosVertAbs = 32
	/*EstimatorPosVertAgl - True if the vertical position (above ground) estimate is good */
	EstimatorPosVertAgl = 64
	/*EstimatorConstPosMode - True if the EKF is in a constant position mode and is not using external measurements (eg GPS or optical flow) */
	EstimatorConstPosMode = 128
	/*EstimatorPredPosHorizRel - True if the EKF has sufficient data to enter a mode that will provide a (relative) position estimate */
	EstimatorPredPosHorizRel = 256
	/*EstimatorPredPosHorizAbs - True if the EKF has sufficient data to enter a mode that will provide a (absolute) position estimate */
	EstimatorPredPosHorizAbs = 512
	/*EstimatorGPSGlitch - True if the EKF has detected a GPS glitch */
	EstimatorGPSGlitch = 1024
	/*EstimatorAccelError - True if the EKF has detected bad accelerometer data */
	EstimatorAccelError = 2048
	/*EstimatorStatusFlagsEnumEnd -  */
	EstimatorStatusFlagsEnumEnd = 2049
)

/*MOTOR_TEST_ORDER -  */
const (
	/*MotorTestOrderDefault - default autopilot motor test method */
	MotorTestOrderDefault = 0
	/*MotorTestOrderSequence - motor numbers are specified as their index in a predefined vehicle-specific sequence */
	MotorTestOrderSequence = 1
	/*MotorTestOrderBoard - motor numbers are specified as the output as labeled on the board */
	MotorTestOrderBoard = 2
	/*MotorTestOrderEnumEnd -  */
	MotorTestOrderEnumEnd = 3
)

/*MOTOR_TEST_THROTTLE_TYPE -  */
const (
	/*MotorTestThrottlePeRCent - throttle as a percentage from 0 ~ 100 */
	MotorTestThrottlePeRCent = 0
	/*MotorTestThrottlePwm - throttle as an absolute PWM value (normally in range of 1000~2000) */
	MotorTestThrottlePwm = 1
	/*MotorTestThrottlePilot - throttle pass-through from pilot's transmitter */
	MotorTestThrottlePilot = 2
	/*MotorTestCompassCal - per-motor compass calibration test */
	MotorTestCompassCal = 3
	/*MotorTestThrottleTypeEnumEnd -  */
	MotorTestThrottleTypeEnumEnd = 4
)

/*GPS_INPUT_IGNORE_FLAGS -  */
const (
	/*GPSInputIgnoreFlagAlt - ignore altitude field */
	GPSInputIgnoreFlagAlt = 1
	/*GPSInputIgnoreFlagHdop - ignore hdop field */
	GPSInputIgnoreFlagHdop = 2
	/*GPSInputIgnoreFlagVdop - ignore vdop field */
	GPSInputIgnoreFlagVdop = 4
	/*GPSInputIgnoreFlagVelHoriz - ignore horizontal velocity field (vn and ve) */
	GPSInputIgnoreFlagVelHoriz = 8
	/*GPSInputIgnoreFlagVelVert - ignore vertical velocity field (vd) */
	GPSInputIgnoreFlagVelVert = 16
	/*GPSInputIgnoreFlagSpeedAccuracy - ignore speed accuracy field */
	GPSInputIgnoreFlagSpeedAccuracy = 32
	/*GPSInputIgnoreFlagHorizontalAccuracy - ignore horizontal accuracy field */
	GPSInputIgnoreFlagHorizontalAccuracy = 64
	/*GPSInputIgnoreFlagVerticalAccuracy - ignore vertical accuracy field */
	GPSInputIgnoreFlagVerticalAccuracy = 128
	/*GPSInputIgnoreFlagsEnumEnd -  */
	GPSInputIgnoreFlagsEnumEnd = 129
)

/*MAV_COLLISION_ACTION - Possible actions an aircraft can take to avoid a collision. */
const (
	/*MavCollisionActionNone - Ignore any potential collisions */
	MavCollisionActionNone = 0
	/*MavCollisionActionReport - Report potential collision */
	MavCollisionActionReport = 1
	/*MavCollisionActionAscendOrDescend - Ascend or Descend to avoid threat */
	MavCollisionActionAscendOrDescend = 2
	/*MavCollisionActionMoveHorizontally - Move horizontally to avoid threat */
	MavCollisionActionMoveHorizontally = 3
	/*MavCollisionActionMovePerpendicular - Aircraft to move perpendicular to the collision's velocity vector */
	MavCollisionActionMovePerpendicular = 4
	/*MavCollisionActionRtl - Aircraft to fly directly back to its launch point */
	MavCollisionActionRtl = 5
	/*MavCollisionActionHover - Aircraft to stop in place */
	MavCollisionActionHover = 6
	/*MavCollisionActionEnumEnd -  */
	MavCollisionActionEnumEnd = 7
)

/*MAV_COLLISION_THREAT_LEVEL - Aircraft-rated danger from this threat. */
const (
	/*MavCollisionThreatLevelNone - Not a threat */
	MavCollisionThreatLevelNone = 0
	/*MavCollisionThreatLevelLow - Craft is mildly concerned about this threat */
	MavCollisionThreatLevelLow = 1
	/*MavCollisionThreatLevelHigh - Craft is panicing, and may take actions to avoid threat */
	MavCollisionThreatLevelHigh = 2
	/*MavCollisionThreatLevelEnumEnd -  */
	MavCollisionThreatLevelEnumEnd = 3
)

/*MAV_COLLISION_SRC - Source of information about this collision. */
const (
	/*MavCollisionSRCAdsb - ID field references ADSB_VEHICLE packets */
	MavCollisionSRCAdsb = 0
	/*MavCollisionSRCMavlinkGPSGlobalInt - ID field references MAVLink SRC ID */
	MavCollisionSRCMavlinkGPSGlobalInt = 1
	/*MavCollisionSRCEnumEnd -  */
	MavCollisionSRCEnumEnd = 2
)

/*GPS_FIX_TYPE - Type of GPS fix */
const (
	/*GPSFixTypeNoGPS - No GPS connected */
	GPSFixTypeNoGPS = 0
	/*GPSFixTypeNoFix - No position information, GPS is connected */
	GPSFixTypeNoFix = 1
	/*GPSFixType2DFix - 2D position */
	GPSFixType2DFix = 2
	/*GPSFixType3DFix - 3D position */
	GPSFixType3DFix = 3
	/*GPSFixTypeDGPS - DGPS/SBAS aided 3D position */
	GPSFixTypeDGPS = 4
	/*GPSFixTypeRtkFloat - RTK float, 3D position */
	GPSFixTypeRtkFloat = 5
	/*GPSFixTypeRtkFixed - RTK Fixed, 3D position */
	GPSFixTypeRtkFixed = 6
	/*GPSFixTypeStatic - Static fixed, typically used for base stations */
	GPSFixTypeStatic = 7
	/*GPSFixTypePpp - PPP, 3D position. */
	GPSFixTypePpp = 8
	/*GPSFixTypeEnumEnd -  */
	GPSFixTypeEnumEnd = 9
)

/*RTK_BASELINE_COORDINATE_SYSTEM - RTK GPS baseline coordinate system, used for RTK corrections */
const (
	/*RtkBaselineCoordinateSystemEcef - Earth-centered, Earth-fixed */
	RtkBaselineCoordinateSystemEcef = 0
	/*RtkBaselineCoordinateSystemNed - North, East, Down */
	RtkBaselineCoordinateSystemNed = 1
	/*RtkBaselineCoordinateSystemEnumEnd -  */
	RtkBaselineCoordinateSystemEnumEnd = 2
)

/*LANDING_TARGET_TYPE - Type of landing target */
const (
	/*LandingTargetTypeLightBeacon - Landing target signaled by light beacon (ex: IR-LOCK) */
	LandingTargetTypeLightBeacon = 0
	/*LandingTargetTypeRadioBeacon - Landing target signaled by radio beacon (ex: ILS, NDB) */
	LandingTargetTypeRadioBeacon = 1
	/*LandingTargetTypeVisionFIDucial - Landing target represented by a fiducial marker (ex: ARTag) */
	LandingTargetTypeVisionFIDucial = 2
	/*LandingTargetTypeVisionOther - Landing target represented by a pre-defined visual shape/feature (ex: X-marker, H-marker, square) */
	LandingTargetTypeVisionOther = 3
	/*LandingTargetTypeEnumEnd -  */
	LandingTargetTypeEnumEnd = 4
)

/*VTOL_TRANSITION_HEADING - Direction of VTOL transition */
const (
	/*VtolTransitionHeadingVehicleDefault - Respect the heading configuration of the vehicle. */
	VtolTransitionHeadingVehicleDefault = 0
	/*VtolTransitionHeadingNextWaypoint - Use the heading pointing towards the next waypoint. */
	VtolTransitionHeadingNextWaypoint = 1
	/*VtolTransitionHeadingTakeoff - Use the heading on takeoff (while sitting on the ground). */
	VtolTransitionHeadingTakeoff = 2
	/*VtolTransitionHeadingSpecified - Use the specified heading in parameter 4. */
	VtolTransitionHeadingSpecified = 3
	/*VtolTransitionHeadingAny - Use the current heading when reaching takeoff altitude (potentially facing the wind when weather-vaning is active). */
	VtolTransitionHeadingAny = 4
	/*VtolTransitionHeadingEnumEnd -  */
	VtolTransitionHeadingEnumEnd = 5
)

/*CAMERA_CAP_FLAGS - Camera capability flags (Bitmap). */
const (
	/*CameraCapFlagsCaptureVIDeo - Camera is able to record video. */
	CameraCapFlagsCaptureVIDeo = 1
	/*CameraCapFlagsCaptureImage - Camera is able to capture images. */
	CameraCapFlagsCaptureImage = 2
	/*CameraCapFlagsHasModes - Camera has separate Video and Image/Photo modes (MAV_CMD_SET_CAMERA_MODE) */
	CameraCapFlagsHasModes = 4
	/*CameraCapFlagsCanCaptureImageInVIDeoMode - Camera can capture images while in video mode */
	CameraCapFlagsCanCaptureImageInVIDeoMode = 8
	/*CameraCapFlagsCanCaptureVIDeoInImageMode - Camera can capture videos while in Photo/Image mode */
	CameraCapFlagsCanCaptureVIDeoInImageMode = 16
	/*CameraCapFlagsHasImageSurveyMode - Camera has image survey mode (MAV_CMD_SET_CAMERA_MODE) */
	CameraCapFlagsHasImageSurveyMode = 32
	/*CameraCapFlagsEnumEnd -  */
	CameraCapFlagsEnumEnd = 33
)

/*PARAM_ACK - Result from a PARAM_EXT_SET message. */
const (
	/*ParamAckAccepted - Parameter value ACCEPTED and SET */
	ParamAckAccepted = 0
	/*ParamAckValueUnsupported - Parameter value UNKNOWN/UNSUPPORTED */
	ParamAckValueUnsupported = 1
	/*ParamAckFailed - Parameter failed to set */
	ParamAckFailed = 2
	/*ParamAckInProgress - Parameter value received but not yet validated or set. A subsequent PARAM_EXT_ACK will follow once operation is completed with the actual result. These are for parameters that may take longer to set. Instead of waiting for an ACK and potentially timing out, you will immediately receive this response to let you know it was received. */
	ParamAckInProgress = 3
	/*ParamAckEnumEnd -  */
	ParamAckEnumEnd = 4
)

/*CAMERA_MODE - Camera Modes. */
const (
	/*CameraModeImage - Camera is in image/photo capture mode. */
	CameraModeImage = 0
	/*CameraModeVIDeo - Camera is in video capture mode. */
	CameraModeVIDeo = 1
	/*CameraModeImageSurvey - Camera is in image survey capture mode. It allows for camera controller to do specific settings for surveys. */
	CameraModeImageSurvey = 2
	/*CameraModeEnumEnd -  */
	CameraModeEnumEnd = 3
)

/*MAV_ARM_AUTH_DENIED_REASON -  */
const (
	/*MavArmAuthDeniedReasonGeneric - Not a specific reason */
	MavArmAuthDeniedReasonGeneric = 0
	/*MavArmAuthDeniedReasonNone - Authorizer will send the error as string to GCS */
	MavArmAuthDeniedReasonNone = 1
	/*MavArmAuthDeniedReasonInvalIDWaypoint - At least one waypoint have a invalid value */
	MavArmAuthDeniedReasonInvalIDWaypoint = 2
	/*MavArmAuthDeniedReasonTimeout - Timeout in the authorizer process(in case it depends on network) */
	MavArmAuthDeniedReasonTimeout = 3
	/*MavArmAuthDeniedReasonAirspaceInUse - Airspace of the mission in use by another vehicle, second result parameter can have the waypoint id that caused it to be denied. */
	MavArmAuthDeniedReasonAirspaceInUse = 4
	/*MavArmAuthDeniedReasonBadWeather - Weather is not good to fly */
	MavArmAuthDeniedReasonBadWeather = 5
	/*MavArmAuthDeniedReasonEnumEnd -  */
	MavArmAuthDeniedReasonEnumEnd = 6
)

/*RC_TYPE - RC type */
const (
	/*RCTypeSpektrumDsm2 - Spektrum DSM2 */
	RCTypeSpektrumDsm2 = 0
	/*RCTypeSpektrumDsmx - Spektrum DSMX */
	RCTypeSpektrumDsmx = 1
	/*RCTypeEnumEnd -  */
	RCTypeEnumEnd = 2
)

/*POSITION_TARGET_TYPEMASK - Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b0000000000000000 or 0b0000001000000000 indicates that none of the setpoint dimensions should be ignored. If bit 9 is set the floats afx afy afz should be interpreted as force instead of acceleration. */
const (
	/*PositionTargetTypemaskXIgnore - Ignore position x */
	PositionTargetTypemaskXIgnore = 1
	/*PositionTargetTypemaskYIgnore - Ignore position y */
	PositionTargetTypemaskYIgnore = 2
	/*PositionTargetTypemaskZIgnore - Ignore position z */
	PositionTargetTypemaskZIgnore = 4
	/*PositionTargetTypemaskVxIgnore - Ignore velocity x */
	PositionTargetTypemaskVxIgnore = 8
	/*PositionTargetTypemaskVyIgnore - Ignore velocity y */
	PositionTargetTypemaskVyIgnore = 16
	/*PositionTargetTypemaskVzIgnore - Ignore velocity z */
	PositionTargetTypemaskVzIgnore = 32
	/*PositionTargetTypemaskAxIgnore - Ignore acceleration x */
	PositionTargetTypemaskAxIgnore = 64
	/*PositionTargetTypemaskAyIgnore - Ignore acceleration y */
	PositionTargetTypemaskAyIgnore = 128
	/*PositionTargetTypemaskAzIgnore - Ignore acceleration z */
	PositionTargetTypemaskAzIgnore = 256
	/*PositionTargetTypemaskFoRCeSet - Use force instead of acceleration */
	PositionTargetTypemaskFoRCeSet = 512
	/*PositionTargetTypemaskYawIgnore - Ignore yaw */
	PositionTargetTypemaskYawIgnore = 1024
	/*PositionTargetTypemaskYawRateIgnore - Ignore yaw rate */
	PositionTargetTypemaskYawRateIgnore = 2048
	/*PositionTargetTypemaskEnumEnd -  */
	PositionTargetTypemaskEnumEnd = 2049
)
