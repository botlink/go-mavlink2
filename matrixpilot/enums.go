package matrixpilot

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

/*MAV_PREFLIGHT_STORAGE_ACTION - Action required when performing CMD_PREFLIGHT_STORAGE */
const (
	/*MavPfsCmdReadAll - Read all parameters from storage */
	MavPfsCmdReadAll = 0
	/*MavPfsCmdWriteAll - Write all parameters to storage */
	MavPfsCmdWriteAll = 1
	/*MavPfsCmdClearAll - Clear all  parameters in storage */
	MavPfsCmdClearAll = 2
	/*MavPfsCmdReadSpecific - Read specific parameters from storage */
	MavPfsCmdReadSpecific = 3
	/*MavPfsCmdWriteSpecific - Write specific parameters to storage */
	MavPfsCmdWriteSpecific = 4
	/*MavPfsCmdClearSpecific - Clear specific parameters in storage */
	MavPfsCmdClearSpecific = 5
	/*MavPfsCmdDoNothing - do nothing */
	MavPfsCmdDoNothing = 6
	/*MavPreflightStorageActionEnumEnd -  */
	MavPreflightStorageActionEnumEnd = 7
)

/*MAV_CMD - Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. */
const (
	/*MavCmdPreflightStorageAdvanced - Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode. */
	MavCmdPreflightStorageAdvanced = 0
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
	/*MavCmdGetHomePosition - Request the home position from the vehicle. */
	MavCmdGetHomePosition = 410
	/*MavCmdStartRxPair - Starts receiver pairing. */
	MavCmdStartRxPair = 500
	/*MavCmdGetMessageInterval - Request the interval between messages for a particular MAVLink message ID */
	MavCmdGetMessageInterval = 510
	/*MavCmdSetMessageInterval - Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM */
	MavCmdSetMessageInterval = 511
	/*MavCmdRequestMessage - Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL). */
	MavCmdRequestMessage = 512
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
	/*MavCmdEnumEnd -  */
	MavCmdEnumEnd = 31015
)
