package can

import (
	"encoding/binary"
	"rvctomqtt/constants"
	//"rvctomqtt/pform"
	"testing"
	"time"
)

func TestCanFrame(t *testing.T) {

	var ts = time.Now()

	// 32 bit CAN_ID + EFF/RTR/ERR flags
	// bit 0-28: CAN identifier (11/29 bit)
	// bit 29: error message flag (ERR)
	// bit 30: remote transmision request (RTR)
	// bit 31: extended frame format (EFF) (MSB)
	// This is the canID INCLUDING the 3 MSB flags

	var cf Frame = Frame{
		Timestamp:    ts,
		ID:           0,
		Length:       0,
		Flags:        0,
		Res0:         0,
		Res1:         0,
		Data:         [constants.MaxFrameDataLength]uint8{},
		MessageBytes: [constants.MAX_MESSAGE]byte{0, 0, 0, 0, 8, 0, 0, 0, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
	}

	// TODO Need to move this to a platform specific metod (in pform now -but need to split that out)
	cf.BuildCanFrame(binary.LittleEndian.Uint32)

	if !ts.Equal(cf.Timestamp) {
		t.Errorf("Timestamp does not match. Expectd %s got %s ", ts, cf.Timestamp)
	}
	if !(cf.Length == 8) {
		t.Errorf("Length wrong expected %d got %d ", 8, cf.Length)
	}
	var b [constants.MaxFrameDataLength]uint8 = cf.Data
	for i := 0; i < 8; i++ {
		if int(b[i]) != i {
			t.Errorf("Array element wrong expected %d got %d ", i, b[i])

		}

	}

}
