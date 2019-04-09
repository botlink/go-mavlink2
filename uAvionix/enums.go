package uavionix

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

/*UAVIONIX_ADSB_OUT_DYNAMIC_STATE - State flags for ADS-B transponder dynamic report */
const (
	/*UavionixAdsbOutDynamicStateIntentChange -  */
	UavionixAdsbOutDynamicStateIntentChange = 1
	/*UavionixAdsbOutDynamicStateAutopilotEnabled -  */
	UavionixAdsbOutDynamicStateAutopilotEnabled = 2
	/*UavionixAdsbOutDynamicStateNicbaroCrosschecked -  */
	UavionixAdsbOutDynamicStateNicbaroCrosschecked = 4
	/*UavionixAdsbOutDynamicStateOnGround -  */
	UavionixAdsbOutDynamicStateOnGround = 8
	/*UavionixAdsbOutDynamicStateIDent -  */
	UavionixAdsbOutDynamicStateIDent = 16
	/*UavionixAdsbOutDynamicStateEnumEnd -  */
	UavionixAdsbOutDynamicStateEnumEnd = 17
)

/*UAVIONIX_ADSB_OUT_RF_SELECT - Transceiver RF control flags for ADS-B transponder dynamic reports */
const (
	/*UavionixAdsbOutRfSelectStandby -  */
	UavionixAdsbOutRfSelectStandby = 0
	/*UavionixAdsbOutRfSelectRxEnabled -  */
	UavionixAdsbOutRfSelectRxEnabled = 1
	/*UavionixAdsbOutRfSelectTxEnabled -  */
	UavionixAdsbOutRfSelectTxEnabled = 2
	/*UavionixAdsbOutRfSelectEnumEnd -  */
	UavionixAdsbOutRfSelectEnumEnd = 3
)

/*UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX - Status for ADS-B transponder dynamic input */
const (
	/*UavionixAdsbOutDynamicGPSFixNone0 -  */
	UavionixAdsbOutDynamicGPSFixNone0 = 0
	/*UavionixAdsbOutDynamicGPSFixNone1 -  */
	UavionixAdsbOutDynamicGPSFixNone1 = 1
	/*UavionixAdsbOutDynamicGPSFix2D -  */
	UavionixAdsbOutDynamicGPSFix2D = 2
	/*UavionixAdsbOutDynamicGPSFix3D -  */
	UavionixAdsbOutDynamicGPSFix3D = 3
	/*UavionixAdsbOutDynamicGPSFixDGPS -  */
	UavionixAdsbOutDynamicGPSFixDGPS = 4
	/*UavionixAdsbOutDynamicGPSFixRtk -  */
	UavionixAdsbOutDynamicGPSFixRtk = 5
	/*UavionixAdsbOutDynamicGPSFixEnumEnd -  */
	UavionixAdsbOutDynamicGPSFixEnumEnd = 6
)

/*UAVIONIX_ADSB_RF_HEALTH - Status flags for ADS-B transponder dynamic output */
const (
	/*UavionixAdsbRfHealthInitializing -  */
	UavionixAdsbRfHealthInitializing = 0
	/*UavionixAdsbRfHealthOk -  */
	UavionixAdsbRfHealthOk = 1
	/*UavionixAdsbRfHealthFailTx -  */
	UavionixAdsbRfHealthFailTx = 2
	/*UavionixAdsbRfHealthFailRx -  */
	UavionixAdsbRfHealthFailRx = 16
	/*UavionixAdsbRfHealthEnumEnd -  */
	UavionixAdsbRfHealthEnumEnd = 17
)

/*UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE - Definitions for aircraft size */
const (
	/*UavionixAdsbOutCfgAiRCraftSizeNoData -  */
	UavionixAdsbOutCfgAiRCraftSizeNoData = 0
	/*UavionixAdsbOutCfgAiRCraftSizeL15MW23M -  */
	UavionixAdsbOutCfgAiRCraftSizeL15MW23M = 1
	/*UavionixAdsbOutCfgAiRCraftSizeL25MW28P5M -  */
	UavionixAdsbOutCfgAiRCraftSizeL25MW28P5M = 2
	/*UavionixAdsbOutCfgAiRCraftSizeL2534M -  */
	UavionixAdsbOutCfgAiRCraftSizeL2534M = 3
	/*UavionixAdsbOutCfgAiRCraftSizeL3533M -  */
	UavionixAdsbOutCfgAiRCraftSizeL3533M = 4
	/*UavionixAdsbOutCfgAiRCraftSizeL3538M -  */
	UavionixAdsbOutCfgAiRCraftSizeL3538M = 5
	/*UavionixAdsbOutCfgAiRCraftSizeL4539P5M -  */
	UavionixAdsbOutCfgAiRCraftSizeL4539P5M = 6
	/*UavionixAdsbOutCfgAiRCraftSizeL4545M -  */
	UavionixAdsbOutCfgAiRCraftSizeL4545M = 7
	/*UavionixAdsbOutCfgAiRCraftSizeL5545M -  */
	UavionixAdsbOutCfgAiRCraftSizeL5545M = 8
	/*UavionixAdsbOutCfgAiRCraftSizeL5552M -  */
	UavionixAdsbOutCfgAiRCraftSizeL5552M = 9
	/*UavionixAdsbOutCfgAiRCraftSizeL6559P5M -  */
	UavionixAdsbOutCfgAiRCraftSizeL6559P5M = 10
	/*UavionixAdsbOutCfgAiRCraftSizeL6567M -  */
	UavionixAdsbOutCfgAiRCraftSizeL6567M = 11
	/*UavionixAdsbOutCfgAiRCraftSizeL75W72P5M -  */
	UavionixAdsbOutCfgAiRCraftSizeL75W72P5M = 12
	/*UavionixAdsbOutCfgAiRCraftSizeL75W80M -  */
	UavionixAdsbOutCfgAiRCraftSizeL75W80M = 13
	/*UavionixAdsbOutCfgAiRCraftSizeL85W80M -  */
	UavionixAdsbOutCfgAiRCraftSizeL85W80M = 14
	/*UavionixAdsbOutCfgAiRCraftSizeL85W90M -  */
	UavionixAdsbOutCfgAiRCraftSizeL85W90M = 15
	/*UavionixAdsbOutCfgAiRCraftSizeEnumEnd -  */
	UavionixAdsbOutCfgAiRCraftSizeEnumEnd = 16
)

/*UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT - GPS lataral offset encoding */
const (
	/*UavionixAdsbOutCfgGPSOffsetLatNoData -  */
	UavionixAdsbOutCfgGPSOffsetLatNoData = 0
	/*UavionixAdsbOutCfgGPSOffsetLatLeft2M -  */
	UavionixAdsbOutCfgGPSOffsetLatLeft2M = 1
	/*UavionixAdsbOutCfgGPSOffsetLatLeft4M -  */
	UavionixAdsbOutCfgGPSOffsetLatLeft4M = 2
	/*UavionixAdsbOutCfgGPSOffsetLatLeft6M -  */
	UavionixAdsbOutCfgGPSOffsetLatLeft6M = 3
	/*UavionixAdsbOutCfgGPSOffsetLatRight0M -  */
	UavionixAdsbOutCfgGPSOffsetLatRight0M = 4
	/*UavionixAdsbOutCfgGPSOffsetLatRight2M -  */
	UavionixAdsbOutCfgGPSOffsetLatRight2M = 5
	/*UavionixAdsbOutCfgGPSOffsetLatRight4M -  */
	UavionixAdsbOutCfgGPSOffsetLatRight4M = 6
	/*UavionixAdsbOutCfgGPSOffsetLatRight6M -  */
	UavionixAdsbOutCfgGPSOffsetLatRight6M = 7
	/*UavionixAdsbOutCfgGPSOffsetLatEnumEnd -  */
	UavionixAdsbOutCfgGPSOffsetLatEnumEnd = 8
)

/*UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON - GPS longitudinal offset encoding */
const (
	/*UavionixAdsbOutCfgGPSOffsetLonNoData -  */
	UavionixAdsbOutCfgGPSOffsetLonNoData = 0
	/*UavionixAdsbOutCfgGPSOffsetLonAppliedBySensor -  */
	UavionixAdsbOutCfgGPSOffsetLonAppliedBySensor = 1
	/*UavionixAdsbOutCfgGPSOffsetLonEnumEnd -  */
	UavionixAdsbOutCfgGPSOffsetLonEnumEnd = 2
)

/*UAVIONIX_ADSB_EMERGENCY_STATUS - Emergency status encoding */
const (
	/*UavionixAdsbOutNoEmergency -  */
	UavionixAdsbOutNoEmergency = 0
	/*UavionixAdsbOutGeneralEmergency -  */
	UavionixAdsbOutGeneralEmergency = 1
	/*UavionixAdsbOutLifeguardEmergency -  */
	UavionixAdsbOutLifeguardEmergency = 2
	/*UavionixAdsbOutMinIMUmFuelEmergency -  */
	UavionixAdsbOutMinIMUmFuelEmergency = 3
	/*UavionixAdsbOutNoCommEmergency -  */
	UavionixAdsbOutNoCommEmergency = 4
	/*UavionixAdsbOutUnlawfulInterferanceEmergency -  */
	UavionixAdsbOutUnlawfulInterferanceEmergency = 5
	/*UavionixAdsbOutDownedAiRCraftEmergency -  */
	UavionixAdsbOutDownedAiRCraftEmergency = 6
	/*UavionixAdsbOutReserved -  */
	UavionixAdsbOutReserved = 7
	/*UavionixAdsbEmergencyStatusEnumEnd -  */
	UavionixAdsbEmergencyStatusEnumEnd = 8
)
