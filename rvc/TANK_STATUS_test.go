package rvc

import (
	"testing"
)

func TestTankStatus(t *testing.T) {
	var dcds = tankStatus{}
	var f = RvcFrame{}
	var raw = f.GetMessage()

	// 0x1FFB7
	// func TestBuildCanFrame(t *testing.T) {

	// RAW FRAME: af b7 ff 99 08 00 00 00 00 02 14 ff ff ff ff ff
	//DCSS1 VALUE = DGN: 1ffb7 (TANK_STATUS) Instance: 0 (Fresh Water) level 10.000000 relativeLevel: 2 resolution: 20 absoluteLevel: 65535 tanSize: 65535

	raw[0] = 0xaf
	raw[1] = 0xb7
	raw[2] = 0xff
	raw[3] = 0x99
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0
	raw[8] = 0x0   // instance
	raw[9] = 0x02  // relative level
	raw[10] = 0x14 //  resolution
	raw[11] = 0xff // absolute level uint16
	raw[12] = 0xff
	raw[13] = 0xff // tanksize uint16
	raw[14] = 0xff
	raw[15] = 0xff

	f.BuildCanFrameX()
	dcds.Init(&f)

	{
		var expected uint32 = 0x1FFB7
		var got = dcds.GetDGN()
		if expected != got {
			t.Errorf("Wrong DGN expected %x got %x", expected, dcds.GetDGN())
		}
	}
	{
		var expected uint8 = 0
		var got = dcds.GetInstance()
		if got != expected {
			t.Errorf("Wrong instance expected %x got %x", expected, dcds.GetInstance())
		}
	}

	// TODO: It may not make sense to test the parsing of every field as we've already tested that (or should have)
	//       Maybe it would be better to test the Fields and each associated method from that
	/*
		{
			var expected uint8 = 17
			if dcds.GetGroup() != expected {
				t.Errorf("Wrong group expected %x got %x", expected, dcds.GetGroup())
			}
		}
		{
			var expected float64 = 0.5
			if dcds.GetBrightness() != expected {
				t.Errorf("Wrong desiredBrightness expected %f got %f", expected, dcds.GetBrightness())
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
	*/
}
