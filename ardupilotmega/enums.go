package ardupilotmega

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

/*MAV_CMD - Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. See https://mavlink.io/en/guide/xml_schema.html#MAV_CMD for information about the structure of the MAV_CMD entries */
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
	/*MavCmdNAVLand - Land at location. */
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
	/*MavCmdDoFollow - Begin following a target */
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
	/*MavCmdNAVVtolTakeoff - Takeoff from ground using VTOL mode, and transition to forward flight with specified heading. */
	MavCmdNAVVtolTakeoff = 84
	/*MavCmdNAVVtolLand - Land using VTOL mode */
	MavCmdNAVVtolLand = 85
	/*MavCmdNAVGUIDedEnable - hand control over to an external controller */
	MavCmdNAVGUIDedEnable = 92
	/*MavCmdNAVDelay - Delay the next navigation command a number of seconds or until a specified time */
	MavCmdNAVDelay = 93
	/*MavCmdNAVPayloadPlace - Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload. */
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
	/*MavCmdDoSetRoiSysID - Mount tracks system with specified system ID. Determination of target vehicle position may be done with GLOBAL_POSITION_INT or any other means. */
	MavCmdDoSetRoiSysID = 198
	/*MavCmdDoControlVIDeo - Control onboard camera system. */
	MavCmdDoControlVIDeo = 200
	/*MavCmdDoSetRoi - Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicles control system to control the vehicle attitude and the attitude of various sensors such as cameras. */
	MavCmdDoSetRoi = 201
	/*MavCmdDoDigicamConfigure - Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ). */
	MavCmdDoDigicamConfigure = 202
	/*MavCmdDoDigicamControl - Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ). */
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
	/*MavCmdDoMotorTest - Mission command to perform motor test. */
	MavCmdDoMotorTest = 209
	/*MavCmdDoInvertedFlight - Change to/from inverted flight. */
	MavCmdDoInvertedFlight = 210
	/*MavCmdDoGripper - Mission command to operate EPM gripper. */
	MavCmdDoGripper = 211
	/*MavCmdDoAutotuneEnable - Enable/disable autotune. */
	MavCmdDoAutotuneEnable = 212
	/*MavCmdNAVSetYawSpeed - Sets a desired vehicle turn angle and speed change. */
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
	/*MavCmdOverrIDeGoto - Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position. */
	MavCmdOverrIDeGoto = 252
	/*MavCmdMissionStart - start running a mission */
	MavCmdMissionStart = 300
	/*MavCmdComponentArmDisarm - Arms / Disarms a component */
	MavCmdComponentArmDisarm = 400
	/*MavCmdIlluminatorOnOff - Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light). */
	MavCmdIlluminatorOnOff = 405
	/*MavCmdGetHomePosition - Request the home position from the vehicle. */
	MavCmdGetHomePosition = 410
	/*MavCmdStartRxPair - Starts receiver pairing. */
	MavCmdStartRxPair = 500
	/*MavCmdGetMessageInterval - Request the interval between messages for a particular MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message. */
	MavCmdGetMessageInterval = 510
	/*MavCmdSetMessageInterval - Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM. */
	MavCmdSetMessageInterval = 511
	/*MavCmdRequestMessage - Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL). */
	MavCmdRequestMessage = 512
	/*MavCmdRequestProtocolVersion - Request MAVLink protocol version compatibility */
	MavCmdRequestProtocolVersion = 519
	/*MavCmdRequestAutopilotCapabilities - Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message */
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
	/*MavCmdSetCameraMode - Set camera running mode. Use NaN for reserved values. GCS will send a MAV_CMD_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming. */
	MavCmdSetCameraMode = 530
	/*MavCmdSetCameraZoom - Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values. */
	MavCmdSetCameraZoom = 531
	/*MavCmdSetCameraFocus - Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values. */
	MavCmdSetCameraFocus = 532
	/*MavCmdJumpTag - Tagged jump target. Can be jumped to with MAV_CMD_DO_JUMP_TAG. */
	MavCmdJumpTag = 600
	/*MavCmdDoJumpTag - Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number. */
	MavCmdDoJumpTag = 601
	/*MavCmdImageStartCapture - Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values. */
	MavCmdImageStartCapture = 2000
	/*MavCmdImageStopCapture - Stop image capture sequence Use NaN for reserved values. */
	MavCmdImageStopCapture = 2001
	/*MavCmdRequestCameraImageCapture - Re-request a CAMERA_IMAGE_CAPTURE message. Use NaN for reserved values. */
	MavCmdRequestCameraImageCapture = 2002
	/*MavCmdDoTriggerControl - Enable or disable on-board camera triggering system. */
	MavCmdDoTriggerControl = 2003
	/*MavCmdVIDeoStartCapture - Starts video capture (recording). Use NaN for reserved values. */
	MavCmdVIDeoStartCapture = 2500
	/*MavCmdVIDeoStopCapture - Stop the current video capture (recording). Use NaN for reserved values. */
	MavCmdVIDeoStopCapture = 2501
	/*MavCmdVIDeoStartStreaming - Start video streaming */
	MavCmdVIDeoStartStreaming = 2502
	/*MavCmdVIDeoStopStreaming - Stop the given video stream */
	MavCmdVIDeoStopStreaming = 2503
	/*MavCmdRequestVIDeoStreamInformation - Request video stream information (VIDEO_STREAM_INFORMATION) */
	MavCmdRequestVIDeoStreamInformation = 2504
	/*MavCmdRequestVIDeoStreamStatus - Request video stream status (VIDEO_STREAM_STATUS) */
	MavCmdRequestVIDeoStreamStatus = 2505
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
	/*MavCmdFixedMagCalYaw - Magnetometer calibration based on provided known yaw. This allows for fast calibration using WMM field tables in the vehicle, given only the known yaw of the vehicle. If Latitude and longitude are both zero then use the current vehicle location. */
	MavCmdFixedMagCalYaw = 42006
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
	/*MavCmdBatteryReset - Reset battery capacity for batteries that accumulate consumed battery via integration. */
	MavCmdBatteryReset = 42651
	/*MavCmdEnumEnd -  */
	MavCmdEnumEnd = 42652
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
	/*MagCalBadRadius -  */
	MagCalBadRadius = 7
	/*MagCalStatusEnumEnd -  */
	MagCalStatusEnumEnd = 8
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
	/*PlaneModeTakeoff -  */
	PlaneModeTakeoff = 13
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
	/*PlaneModeQautotune -  */
	PlaneModeQautotune = 22
	/*PlaneModeEnumEnd -  */
	PlaneModeEnumEnd = 23
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
