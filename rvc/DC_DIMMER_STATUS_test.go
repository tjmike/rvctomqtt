package rvc

import "testing"

func TestDCDimmerStatus3(t *testing.T) {
	var dcds = DCDimmerStatus3{}
	var f = RvcFrame{}
	var raw = f.GetMessage()
	// func TestBuildCanFrame(t *testing.T) {
	raw[0] = 0x8c
	raw[1] = 0xda
	raw[2] = 0xfe
	raw[3] = 0x99
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0
	raw[8] = 0x01  // 0-instance
	raw[9] = 0x11  // 1-group
	raw[10] = 0x01 // 2-brightness

	//lockitem          uint2 // 3-(0,1)
	//overCurrentStatus uint2 // 3-(2,3)
	//overrideSatus     uint2 // 3-(4,5)
	//enableStatus      uint2 // 3-(6,7)
	raw[11] = 0b11001001

	raw[12] = 0x22 // 4-delayDuration
	raw[13] = 0x33 // 5-lastCommand

	//interlockStatus   uint2 // 6 (0,1)
	//loadStatus        uint2 //6 (2-3)
	//reserved          uint2 // 6 (4-5)
	//undercurrent      uint2 // 6 (6-7)
	raw[14] = 0b00101110

	raw[15] = 0xff
	f.BuildCanFrameX()
	dcds.Init(&f)

	{
		var expected uint32 = 0x1FEDA
		if dcds.GetDGN() != expected {
			t.Errorf("Wrong DGN expected %x got %x", expected, dcds.GetDGN())
		}
	}
	{
		var expected uint8 = 1
		if dcds.GetInstance() != expected {
			t.Errorf("Wrong instance expected %x got %x", expected, dcds.GetInstance())
		}
	}
	{
		var expected uint8 = 17
		if dcds.GetGroup() != expected {
			t.Errorf("Wrong group expected %x got %x", expected, dcds.GetGroup())
		}
	}
	{
		var expected float64 = 0.5
		if dcds.GetBrightness() != expected {
			t.Errorf("Wrong brightness expected %f got %f", expected, dcds.GetBrightness())
		}
	}
	{
		var expected uint2 = 1
		var got = dcds.GetLockItem()
		if expected != got {
			t.Errorf("Wrong Lock Item expected %x got %x", expected, got)
		}

		expected = 2
		got = dcds.GetOverCurrentStatus()
		if expected != got {
			t.Errorf("Wrong overcurrent status expected %x got %x", expected, got)
		}

		expected = 0
		got = dcds.GetOverrideStatus()
		if expected != got {
			t.Errorf("Wrong override status expected %x got %x", expected, got)
		}
		expected = 3
		got = dcds.GetEnableStatus()
		if expected != got {
			t.Errorf("Wrong enabled status expected %x got %x", expected, got)
		}

	}
	{
		var expected uint8 = 0x22
		var got = dcds.GetDelayDuration()
		if expected != got {
			t.Errorf("Wrong delay duration expected %x got %x", expected, got)
		}
	}
	{
		var expected uint8 = 0x33
		var got = dcds.GetLastCommand()
		if expected != got {
			t.Errorf("Wrong last command expected %x got %x", expected, got)
		}
	}

	//interlockStatus   uint2 // 6 (0,1)
	//loadStatus        uint2 //6 (2-3)
	//reserved          uint2 // 6 (4-5)
	//undercurrent      uint2 // 6 (6-7)
	//raw[14] = 0b00 10 11 10
	{
		var expected uint2 = 2
		var got = dcds.getInterlockStatus()
		if expected != got {
			t.Errorf("Wrong interlock status expected %x got %x", expected, got)
		}

		expected = 3
		got = dcds.getLoadStatus()
		if expected != got {
			t.Errorf("Wrong load status expected %x got %x", expected, got)
		}

		expected = 2
		got = dcds.getReserved()
		if expected != got {
			t.Errorf("Wrong reserved status expected %x got %x", expected, got)
		}
		expected = 0
		got = dcds.getUnderCurrent()
		if expected != got {
			t.Errorf("Wrong undercurrent expected %x got %x", expected, got)
		}

	}

}
