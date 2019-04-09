package ualberta

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

/*UALBERTA_AUTOPILOT_MODE - Available autopilot modes for ualberta uav */
const (
	/*ModeManualDirect - Raw input pulse widts sent to output */
	ModeManualDirect = 1
	/*ModeManualScaled - Inputs are normalized using calibration, the converted back to raw pulse widths for output */
	ModeManualScaled = 2
	/*ModeAutoPIDAtt -  dfsdfs */
	ModeAutoPIDAtt = 3
	/*ModeAutoPIDVel -  dfsfds */
	ModeAutoPIDVel = 4
	/*ModeAutoPIDPos -  dfsdfsdfs */
	ModeAutoPIDPos = 5
	/*UalbertaAutopilotModeEnumEnd -  */
	UalbertaAutopilotModeEnumEnd = 6
)

/*UALBERTA_NAV_MODE - Navigation filter mode */
const (
	/*NAVAHRSInit -  */
	NAVAHRSInit = 1
	/*NAVAHRS - AHRS mode */
	NAVAHRS = 2
	/*NAVInsGPSInit - INS/GPS initialization mode */
	NAVInsGPSInit = 3
	/*NAVInsGPS - INS/GPS mode */
	NAVInsGPS = 4
	/*UalbertaNAVModeEnumEnd -  */
	UalbertaNAVModeEnumEnd = 5
)

/*UALBERTA_PILOT_MODE - Mode currently commanded by pilot */
const (
	/*PilotManual -  sdf */
	PilotManual = 1
	/*PilotAuto -  dfs */
	PilotAuto = 2
	/*PilotRoto -  Rotomotion mode  */
	PilotRoto = 3
	/*UalbertaPilotModeEnumEnd -  */
	UalbertaPilotModeEnumEnd = 4
)
