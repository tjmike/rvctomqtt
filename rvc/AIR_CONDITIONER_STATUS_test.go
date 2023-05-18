package rvc

import (
	"testing"
)

func TestAirConditionerStatus(t *testing.T) {

	var acstatus = airConditionerStatus{}
	var f = RvcFrame{}
	var raw = f.GetMessage()

	//fmt.Printf("AC Status = %s", raw)
	//fmt.Printf("AC Status = %s", acstatus.String())

	//                          0  1  2  3  4  5  6  7
	// 97 e1 ff 99 08 00 00 00 01 00 ff ff 00 00 ff 00
	raw[0] = 0x97
	raw[1] = 0xe1
	raw[2] = 0xff
	raw[3] = 0x99
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0
	raw[8] = 0x11  // 0-instance
	raw[9] = 0x01  // 1-operating mode 0/1
	raw[10] = 0x02 // max fan speed (%) 0-125
	raw[11] = 0x04 // max ac output level (%)
	raw[12] = 0x06 // fan speed (%)
	raw[13] = 0x08 // ac output level %
	raw[14] = 0x01 // dead band (degrees c) (0.1)
	raw[15] = 0x0a // 2nd dead band (deg c) (1.0)

	f.BuildCanFrameX()

	acstatus.Init(&f)

	{
		var expected uint32 = 0x1FFE1
		if acstatus.GetDGN() != expected {
			t.Errorf("Wrong DGN expected %x got %x", expected, acstatus.GetDGN())
		}
	}
	{
		var expected uint8 = 17
		if acstatus.GetInstance() != expected {
			t.Errorf("Wrong instance expected %x got %x", expected, acstatus.GetInstance())
		}

		expected = 1
		if acstatus.GetOperationgMode() != expected {
			t.Errorf("Wrong operating model %x got %x", expected, acstatus.GetOperationgMode())
		}
	}
	{
		var expected float64 = 1
		if acstatus.GetMaxFanSpeed() != expected {
			t.Errorf("Wrong max fan speed expected %f got %f", expected, acstatus.GetMaxFanSpeed())
		}

		expected = 2
		if acstatus.GetMaxOutputLevel() != expected {
			t.Errorf("Wrong max output level expected %f got %f", expected, acstatus.GetMaxOutputLevel())
		}
		expected = 3
		if acstatus.GetFanSpeed() != expected {
			t.Errorf("Wrong fan speed expected %f got %f", expected, acstatus.GetFanSpeed())
		}
		expected = 4
		if acstatus.GetACOutputLevel() != expected {
			t.Errorf("Wrong ac output level expected %f got %f", expected, acstatus.GetACOutputLevel())
		}

		expected = 0.1
		if acstatus.GetDeadBand() != expected {
			t.Errorf("Wrong dead band expected %f got %f", expected, acstatus.GetDeadBand())
		}

		expected = 1
		if acstatus.GetSecondStageDeadBand() != expected {
			t.Errorf("Wrong second dead band expected %f got %f", expected, acstatus.GetSecondStageDeadBand())
		}

	}

}
