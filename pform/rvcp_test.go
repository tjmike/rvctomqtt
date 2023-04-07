package pform

// This is the public facing interface. It will call platform specific code.
// In this example doit (lower case is different code on darwin vs linux)

import (
	//"github.com/tjmike/rvctomqtt/can"
	"rvctomqtt/can"
	"testing"
	"time"
)

// Create a frame from a raw message an ensure it's what we expect
func TestBuildCanFrame(t *testing.T) {

	r := setUpRawTestMessage()

	var f = can.Frame{}
	BuildCanFrame(&f, &r)

	var id = f.ID
	var idExpected uint32 = 0x99FEBD42
	if id != idExpected {
		t.Errorf("ID Wrong %x != %x", idExpected, id)
	}

	var ext = can.IsExtended(&f) // does not compile: f.IsExteded()
	//var ext =  can.IsExteded(f)
	//var ext bool = can.IsExtended(&f)
	if !ext {
		t.Errorf("Frame should be extended format %x", id)
	}

	var canid = can.CanID(&f)

	var expectedCanID uint32 = 0x19FEBD42

	if !(canid == expectedCanID) {
		t.Errorf("CanID expected %x but got %x ", expectedCanID, id)

	}
}

func setUpRawTestMessage() can.RawCanMessage {
	var r = can.RawCanMessage{
		Timestamp:  time.Now(),
		CanMessage: [16]uint8{},
	}

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

	var i uint8 = 0
	// can ID
	r.CanMessage[i] = 0x99
	i++
	r.CanMessage[i] = 0xFE
	i++
	r.CanMessage[i] = 0xBD
	i++
	r.CanMessage[i] = 0x42
	i++
	// size
	r.CanMessage[i] = 0x08
	i++
	// reserved
	r.CanMessage[i] = 0x00
	i++
	// reserved
	r.CanMessage[i] = 0x00
	i++
	// reserved
	r.CanMessage[i] = 0x00
	i++

	// 8 bytes of data
	r.CanMessage[i] = 0x01
	i++
	r.CanMessage[i] = 0x00
	i++
	r.CanMessage[i] = 0x26
	i++
	r.CanMessage[i] = 0x60
	i++
	r.CanMessage[i] = 0x27
	i++
	r.CanMessage[i] = 0xFF
	i++
	r.CanMessage[i] = 0xFF
	i++
	r.CanMessage[i] = 0xFF
	i++
	return r
}
