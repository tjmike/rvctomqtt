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

	var dataBytes uint8 = 4
	var cf = Frame{
		timestamp: ts,
		ID:        0,
		Length:    0,
		Flags:     0,
		Res0:      0,
		Res1:      0,
		Data:      [constants.MaxFrameDataLength]uint8{},
		// for a test we want the 1st for raw bytes to be BACKWARDS - MS Byte first
		// if all is well the ID will end up with the bits flipped.
		// It's possible that this test may fail due to a platform change
		MessageBytes: [constants.MAX_MESSAGE]byte{1, 0, 0, 0xA0, dataBytes, 0, 0, 0, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
	}

	// TODO Need to move this to a platform specific metod (in pform now -but need to split that out)
	cf.BuildCanFrame(binary.LittleEndian.Uint32)

	checkTheFrame(t, cf, ts, dataBytes)

	cf.BuildCanFrameX()

	checkTheFrame(t, cf, ts, dataBytes)

	cf.MessageBytes[0] = 0xff
	cf.MessageBytes[1] = 0xff
	cf.MessageBytes[3] = 0
	cf.BuildCanFrameX()
	var canID = cf.CanID()
	// the canID should be 11Buts
	var expected uint32 = 0b0111_1111_1111
	if expected != canID {
		t.Errorf("11 but can message = expected = %x got %x ", expected, canID)

	}

}

func checkTheFrame(t *testing.T, cf Frame, ts time.Time, dataBytes uint8) {
	var msg = cf.GetMessage()

	if *msg != cf.MessageBytes {
		t.Errorf("message bytes don't match.  %x got %x ", msg, cf.MessageBytes)
	}

	if !ts.Equal(cf.GetTimeStamp()) {
		t.Errorf("timestamp does not match. Expectd %s got %s ", ts, cf.GetTimeStamp())
	}

	ts = ts.Add(time.Second)
	cf.SetTimeStamp(ts)

	if !ts.Equal(cf.GetTimeStamp()) {
		t.Errorf("timestamp does not match. Expectd %s got %s ", ts, cf.GetTimeStamp())
	}

	if cf.DataLength() != dataBytes {
		t.Errorf("Length wrong expected %d got %d ", 8, cf.Length)
	}
	var b [constants.MaxFrameDataLength]uint8 = cf.Data
	for i := 0; i < int(dataBytes); i++ {
		if int(b[i]) != i {
			t.Errorf("Array element wrong expected %d got %d ", i, b[i])
		}
	}

	for i := int(dataBytes); i < int(constants.MaxFrameDataLength); i++ {
		if b[i] != 0xff {
			t.Errorf("Array element wrong expected %d got %d ", i, b[i])
		}
	}
	if !cf.IsExtended() {
		t.Errorf("Expected extended frame but  was not set")
	}
	if cf.IsRTR() {
		t.Errorf("Expected NO RTR frame but  was not set")
	}
	if !cf.IsERR() {
		t.Errorf("Expected ERROR frame but  was not set")
	}

	// filtered bytes
	// BACKWARDS
	var expectedBytes = [...]byte{1, 0, 0, 0}

	// TODO use a generic method to convert
	// binary.LittleEndian.Uint32
	var expected = binary.LittleEndian.Uint32(expectedBytes[0:])
	if cf.CanID() != expected {
		t.Errorf("Expected CAN ID of 1 but got %x", cf.CanID())
	}
}

func TestRebuildCanFrame(t *testing.T) {

	var ts = time.Now()
	var dataBytes uint8 = 8
	var cf = Frame{
		timestamp:        ts,
		eFF_RTR_ERR_Flag: 0,
		ID:               0,
		Length:           0,
		Flags:            0,
		Res0:             0,
		Res1:             0,
		Data:             [constants.MaxFrameDataLength]uint8{},
		// for a test we want the 1st for raw bytes to be BACKWARDS - MS Byte first
		// if all is well the ID will end up with the bits flipped.
		// It's possible that this test may fail due to a platform change
		MessageBytes: [constants.MAX_MESSAGE]byte{1, 0x36, 0x92, 0x41, dataBytes, 0, 0, 0, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
	}
	// TODO Need to move this to a platform specific metod (in pform now -but need to split that out)
	cf.BuildCanFrame(binary.LittleEndian.Uint32)

	rebuilt := initCanFrameState(cf, dataBytes)
	// Build the raw message from the data in the frame and then check the results
	rebuilt.SetCanMessage()
	ensureMessageBytesMatch(t, rebuilt, cf)

	cf.MessageBytes[3] = 0x81
	cf.BuildCanFrame(binary.LittleEndian.Uint32)
	rebuilt = initCanFrameState(cf, dataBytes)
	rebuilt.SetCanMessage()
	ensureMessageBytesMatch(t, rebuilt, cf)

	cf.MessageBytes[3] = 0x21
	cf.BuildCanFrame(binary.LittleEndian.Uint32)
	rebuilt = initCanFrameState(cf, dataBytes)
	rebuilt.SetCanMessage()
	ensureMessageBytesMatch(t, rebuilt, cf)

}

func initCanFrameState(cf Frame, dataBytes uint8) Frame {
	var rebuilt = Frame{
		timestamp:        cf.GetTimeStamp(),
		eFF_RTR_ERR_Flag: cf.eFF_RTR_ERR_Flag,
		ID:               cf.ID,
		Length:           cf.Length,
		Flags:            cf.Flags,
		Res0:             cf.Res0,
		Res1:             cf.Res1,
		Data:             [8]uint8{},
		MessageBytes:     [16]byte{},
	}
	for i := 0; i < int(dataBytes); i++ {
		rebuilt.Data[i] = cf.Data[i]
	}
	return rebuilt
}

func ensureMessageBytesMatch(t *testing.T, rebuilt Frame, cf Frame) {
	for i := 0; i < len(rebuilt.MessageBytes); i++ {
		if rebuilt.MessageBytes[i] != cf.MessageBytes[i] {
			t.Errorf("Message index %d expected %x got %x expectedFrame: %x rebuilt: %x",
				i, cf.MessageBytes[i], rebuilt.MessageBytes[i],
				cf.MessageBytes, rebuilt.MessageBytes)
		}
	}
}
