package rvc

import "testing"

func TestDCDimmerCommand(t *testing.T) {
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
		var expected uint2 = 0b11
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
