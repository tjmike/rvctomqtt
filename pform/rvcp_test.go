package pform

// This is the public facing interface. It will call platform specific code.
// In this example doit (lower case is different code on darwin vs linux)

import (
	"encoding/binary"
	"rvctomqtt/can"

	//"github.com/tjmike/rvctomqtt/can"
	"rvctomqtt/constants"
	"testing"
	"time"
)

// Create a frame from a raw message an ensure it's what we expect
func TestBuildCanFrame(t *testing.T) {

	// 99 FE BD 42 08 00 00 00 01 00 26 60 27 FF FF FF
	//  9   9    F    E    B    D     4    2
	//1001 1001 1111 1110 1011  1101 0100 0010
	//
	//extended
	//!rtr
	//!err
	//
	//1 1001 1111 1110 1011  1101 0100 0010
	//
	//1  9    F    E   B      D    4    2

	// We only work whith little edian

	var fullID uint32 = 0x99FEBD42
	var tstVal = [constants.MAX_MESSAGE]uint8{0x99, 0xFE, 0xBD, 0x42, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x26, 0x60, 0x27, 0xFF, 0xFF, 0xFF}
	binary.LittleEndian.PutUint32(tstVal[0:4], fullID)

	r := setUpRawTestMessage(time.Now(), tstVal)

	//var f = can.Frame{}
	r.BuildCanFrame(binary.LittleEndian.Uint32)
	//BuildCanFrame(&r)

	var id = r.ID
	var idExpected uint32 = 0x99FEBD42
	if id != idExpected {
		t.Errorf("ID Wrong %x != %x", idExpected, id)
	}

	var ext = r.IsExtended() // does not compile: f.IsExteded()
	//var ext =  can.IsExteded(f)
	//var ext bool = can.IsExtended(&f)
	if !ext {
		t.Errorf("Frame should be extended format %x", id)
	}

	var canid = r.CanID()

	var expectedCanID uint32 = 0x19FEBD42

	if !(canid == expectedCanID) {
		t.Errorf("CanID expected %x but got %x ", expectedCanID, id)

	}
}

func setUpRawTestMessage(ts time.Time, dat [constants.MAX_MESSAGE]uint8) can.Frame {
	var r = can.Frame{
		Timestamp:    ts,
		MessageBytes: [constants.MAX_MESSAGE]uint8{},
	}

	for i, v := range dat {
		r.MessageBytes[i] = v
	}
	return r
}
