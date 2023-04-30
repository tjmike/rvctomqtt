package j1939

import (
	"rvctomqtt/can"
	"testing"
)

func TestJ1939Frame(t *testing.T) {

	var jf = J1939Frame{}

	// // bit 0-28: CAN identifier (11/29 bit)
	//	// bit 29: error message flag (ERR)
	//	// bit 30: remote transmision request (RTR)
	//	// bit 31: extended frame format (EFF) (MSB)
	//	// This is the canID INCLUDING the 3 MSB flags

	jf.MessageBytes[0] = 0x66 // SA
	jf.MessageBytes[1] = 0x22 // PDU Specific
	jf.MessageBytes[2] = 0x44 // PDU Format
	jf.MessageBytes[3] = 0x91 // EFF
	jf.MessageBytes[4] = 0x08 // length
	jf.MessageBytes[8] = 0x01
	jf.MessageBytes[9] = 0x02
	jf.MessageBytes[10] = 0x04
	jf.MessageBytes[11] = 0x08
	jf.MessageBytes[12] = 0x10
	jf.MessageBytes[13] = 0x20
	jf.MessageBytes[14] = 0x40
	jf.MessageBytes[15] = 0x80
	// build the frame from the raw message bytes
	jf.BuildCanFrameX()

	if jf.GetPDUFormat() != 0x44 {
		t.Errorf("Expected format %x got %x ", 0x44, jf.GetPDUFormat())
	}

	if jf.GetPDUSpecific() != 0x22 {
		t.Errorf("Expected specific %x got %x ", 0x22, jf.GetPDUSpecific())
	}

	if !jf.IsExtended() {
		t.Errorf("Expected extended frame ")
	}
	if jf.GetSourceAddress() != 0x66 {
		t.Errorf("Wrong source address - expectted %x got %x ", 0x66, jf.GetSourceAddress())
	}
	if jf.Length != 8 {
		t.Errorf("Wrong length expected %d got %d", 8, jf.Length)
	}

}

func TestBackwards(t *testing.T) {

	var f = J1939Frame{
		priority: 1,
		PGN: PGN{
			reserved:    0,
			page:        1,
			pduFormat:   0x44,
			pduSpecific: 0x22,
			pgn:         0,
		},
		sourceAddress: 0x66,
		Frame:         can.Frame{},
	}

	if f.GetPriority() != 1 {
		t.Errorf("Wrong format, expected %d got %d\n", 1, f.GetPriority())
	}
}
