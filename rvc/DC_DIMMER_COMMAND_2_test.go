package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	"testing"
)

func TestDCDimmer2Command(t *testing.T) {
	var dcds = DCDimmerCommand2{}
	var f = RvcFrame{}
	var raw = f.GetMessage()
	// func TestBuildCanFrame(t *testing.T) {

	// 9f db fe 99 08 00 00 00 3b ff c8 01 02 00 ff ff
	raw[0] = 0x9f
	raw[1] = 0xdb
	raw[2] = 0xfe
	raw[3] = 0x99
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0

	// deviceInstance    byte    // 0
	//	group             uint8   // 1
	//	desiredBrightness float64 // 2 (%)
	//	command           uint8   // 3
	//	delayDuration     uint8   // 4 seconds
	//
	//	interlockStatus uint2   // 5 (0,1)
	//	reserved1       uint2   // 5 (2-3)
	//	reserved2       uint2   // 5 (4-5)
	//	reserved3       uint2   // 5 (6-7)
	//	rampTime        float64 // seconds (0.1) 0-25 seconds
	//	reserved        uint8   //
	//}
	raw[8] = 0x01           // 0-instance
	raw[9] = 0x02           // 1-group
	raw[10] = 0x04          // 2-desiredBrightness(%)
	raw[11] = 0x03          // 3-command
	raw[12] = 0x01          // 4-delayDuration
	raw[13] = 0b00_01_10_11 // 5-lastCommand
	raw[14] = 0xFA          // (25.0) Ramp Time (0-250)*.1
	raw[15] = 0xff          // 7 Reserved

	f.BuildCanFrameX()
	dcds.Init(&f)

	{
		var expected uint32 = 0x1FEDB
		if dcds.GetDGN() != expected {
			t.Errorf("Wrong DGN expected %x got %x", expected, dcds.GetDGN())
		}
	}
	{
		var expected uint8 = 1
		if dcds.GetInstance() != expected {
			t.Errorf("Wrong instance expected %x got %x", expected, dcds.GetInstance())
		}
		expected = 2
		if dcds.GetGroup() != expected {
			t.Errorf("Wrong group expected %x got %x", expected, dcds.GetGroup())
		}
	}
	{
		var expected float64 = 2
		if dcds.GetDesiredBrightness() != expected {
			t.Errorf("Wrong desiredBrightness expected %f got %f", expected, dcds.GetDesiredBrightness())
		}
	}
	{
		var expected uint8 = 3
		var got = dcds.GetCommand()
		if expected != got {
			t.Errorf("Wrong command expected %x got %x", expected, got)
		}

		expected = 1
		got = dcds.GetDelayDuration()
		if expected != got {
			t.Errorf("Wrong delay expected %x got %x", expected, got)
		}
	}
	{
		var expected constants.Uint2 = 0b11
		var got = dcds.GetInterlockStatus()
		if expected != got {
			t.Errorf("Wrong interlock status expected %x got %x", expected, got)
		}

		expected = 0b10
		got = dcds.GetReserved1()
		if expected != got {
			t.Errorf("Wrong reserved1 status expected %x got %x", expected, got)
		}

		expected = 0b01
		got = dcds.GetReserved2()
		if expected != got {
			t.Errorf("Wrong reserved2 status expected %x got %x", expected, got)
		}

		expected = 0b00
		got = dcds.GetReserved3()
		if expected != got {
			t.Errorf("Wrong reserved3 status expected %x got %x", expected, got)
		}

	}

	{
		var expected = 25.0
		var got = dcds.GetRampTime()
		if expected != got {
			t.Errorf("Wrong ramp time expected %x got %x", expected, got)
		}
	}
	{
		var expected uint8 = 255
		var got = dcds.GetReserved()
		if expected != got {
			t.Errorf("Wrong reserved value expected %x got %x", expected, got)
		}
	}
	//interlockStatus   uint2 // 6 (0,1)
	//loadStatus        uint2 //6 (2-3)
	//reserved          uint2 // 6 (4-5)
	//undercurrent      uint2 // 6 (6-7)
	//raw[14] = 0b00 10 11 10

}

func TestDCDimmer2CreateFrame(t *testing.T) {

	var c = DCDimmerCommand2{}
	c.group = 0
	c.delayDuration = 0
	c.desiredBrightness = 100 //         1100 1000 128+64 192
	c.deviceInstance = 2
	c.dgn = DGN_DC_DIMMER_COMMAND_2 // 0x1FEDB
	c.name = DGNtoName[c.dgn]
	c.priority = 0x06
	c.sourceAddress = 0x10

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
