package j1939

import (
	"rvctomqtt/can"
	"testing"
)

func TestJ1939Frame(t *testing.T) {

	var cf = can.Frame{
		// 0000 0000 0000 0000 0000 0000 0000 0000
		// 0001 1100
		// 0x1C
		ID:           0x1C_00_00_FF, // 29 bits (3 pri) (18 pgn) (8 sa)
		Length:       8,
		Flags:        0,
		Res0:         0,
		Res1:         0,
		Data:         [8]uint8{},
		MessageBytes: [16]byte{},
	}

	var j1939 = CreateJ1939(&cf)
	var pri = j1939.GetPriority()
	if !(pri == 7) {
		t.Errorf("priority error, expected  %x go %x", 7, pri)
	}

	var sa = j1939.GetSourceAddress()
	if !(sa == 255) {
		t.Errorf("source address error, expected  %x go %x", 255, sa)

	}

}
