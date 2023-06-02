package rvc

import (
	"rvctomqtt/constants"
	"testing"
)

func TestChassisMobilityStatus(t *testing.T) {
	var cms = ChassisMobilityStatus{isTiffinFlavor: false}
	var f = RvcFrame{}
	var raw = f.GetMessage()
	// func TestBuildCanFrame(t *testing.T) {

	//9f f4 ff 99 08 00 00 00 f100000000000000
	raw[0] = 0x9f
	raw[1] = 0xf4
	raw[2] = 0xff
	raw[3] = 0x99
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0

	//engineRPM                float64         // 0-1 uint16 (8191.875 = 65535 = UNKNOWN)
	//speed                    float64         // 2-3 uint16 Precision 1/256 (255.99609375 = Unknown)
	//parkBrakeStatus          constants.Uint2 //  4 (0-1) 00 released, 01 engaged
	//transmissionLockStatus   constants.Uint2 //  4 (2-3) 00 not locked, 01 locked - will not go into gear
	//engineLockStatus         constants.Uint2 //  4 (4-5) 00 free to start, 01 locked - will not start
	//ignitionSwitchStatus     constants.Uint2 //  5 (0-1) 00 off 01 0n
	//accessorySwitchStatus    constants.Uint2 //  5 (2-3) 00 off 01 on NOTE: Spec says bits 2-4 - assume thats wrong
	//transmissionCurrentGear  uint8           // 124 = Reverse 1 125 = Neutral  126 = Forward 1 127 = Forward 2  etc.  251 = Park (Note: Most heavy transmissions do  not have a 'Park', but are parked in Neutral.) (255->130 = unavailable/unknown)
	//transmissionSelectedGear uint8           /

	raw[8] = 0x00 // 0-1 engine rpm (1/8)
	raw[9] = 0x01 //
	raw[10] = 0x0 // 2-3 speed (1/256)
	raw[11] = 0x02
	raw[12] = 0b11_10_00_01 // 4-// brake(0-1) transLock(2-3), englock (4-5),
	raw[13] = 0b11_11_10_01 // 5-ign sw(0-1) acc sw(2-3)
	raw[14] = 126           // trans gear
	raw[15] = 125           // trans selected gear

	f.BuildCanFrameX()
	cms.Init(&f)

	if cms.GetDGN() != DGN_CHASSIS_MOBILITY_STATUS {
		t.Errorf("Wrong DGN expected %x got %x", DGN_CHASSIS_MOBILITY_STATUS, cms.GetDGN())
	}

	if cms.GetEngineRPM() != 32.0 {
		t.Errorf("Wrong RPM expected %f got %f", 32.0, cms.GetEngineRPM())
	}

	//  2*0.621371192 1.242742384
	if cms.GetSpeed() != constants.KMPHtoMPH(2.0) {
		t.Errorf("Wrong RPM expected %f got %f", constants.KMPHtoMPH(2.0), cms.GetSpeed())
	}

	if cms.GetParkBrakeStatus() != 1 {
		t.Errorf("Wrong brake status, expected %d got %d", 1, cms.GetParkBrakeStatus())
	}
	if cms.GetTransmissionLockStatus() != 0 {
		t.Errorf("Wrong trans lock status, expected %d got %d", 0, cms.GetTransmissionLockStatus())
	}
	if cms.GetEngineLockStatus() != 2 {
		t.Errorf("Wrong engine lock status, expected %d got %d", 2, cms.GetEngineLockStatus())
	}
	if cms.GetIgnitionSwitchStatus() != 1 {
		t.Errorf("Wrong ignition switch status, expected %d got %d", 1, cms.GetIgnitionSwitchStatus())
	}
	if cms.GetAccessorySwitchStatus() != 2 {
		t.Errorf("Wrong accessory switch status, expected %d got %d", 2, cms.GetAccessorySwitchStatus())
	}

	if cms.GetTransmissionCurrentGear() != 1 {
		t.Errorf("Wrong trans gear, expected %d got %d", 1, cms.GetTransmissionCurrentGear())
	}
}

/*
func TestDCDimmer2CreateFrame(t *testing.T) {

	var c = DCDimmerCommand2{}
	c.group = 0
	c.delayDuration = 0
	c.desiredBrightness = 100 //         1100 1000 128+64 192
	c.deviceInstance = 2
	c.DGN = DGN_DC_DIMMER_COMMAND_2 // 0x1FEDB
	c.Name = DGNName(c.DGN)
	c.Priority = 0x06
	c.SourceAddress = 0x10

	// DGN : 0x1FEDB
	// PRI  : 6
	//    1    9   f    e    d    b
	// xxx1 1001 1111 1110 1101 1011
	var f = c.CreateFrame()
	// 00 00 00 00 00 00 00 00          02 00 c8 00 00 00 00 00

	var f2 = RvcFrame{}
	var bytes = f2.GetMessage()
	for i := 0; i < len(bytes); i++ {
		f2.GetMessage()[i] = f.GetMessage()[i]
	}

	f2.BuildCanFrameX()
	var c2 = DCDimmerCommand2{}
	c2.Init(&f2)

	fmt.Printf("OUT = %x\n", f.MessageBytes)
	fmt.Printf("OUT = %x\n", f2.MessageBytes)

	if !c.Equals(&c2) {
		t.Errorf("Initial and reconstructed Dimmer commands dont match %s != %s\n", c.String(), c2.String())
	}

}
*/
