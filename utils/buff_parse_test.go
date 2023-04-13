package utils

import (
	"rvctomqtt/constants"
	"testing"
)

func TestBusParseBits(t *testing.T) {

	var foo = [constants.MaxFrameDataLength]byte{0x10, 0xC0, 0b00101010, 4, 5, 6, 7, 0}
	var bar = [constants.MaxFrameDataLength]byte{0x10, 0xC0, 0b00101010, 4, 5, 6, 7, 0}

	var bp = bitParser{
		bitPosition: 4,
		byteOffset:  0,
		nbits:       2,
	}

	var bit = bp.parseBits(&foo)
	if bit != 1 {
		t.Errorf("Expected bit value to be  %x but got %x", 1, bit)
	}

	bp = bitParser{
		bitPosition: 6,
		byteOffset:  1,
		nbits:       2,
	}
	bit = bp.parseBits(&foo)
	if bit != 3 {
		t.Errorf("Expected bit value to be  %x but got %x", 3, bit)
	}

	bp = bitParser{
		bitPosition: 2,
		byteOffset:  2,
		nbits:       4,
	}
	bit = bp.parseBits(&foo)
	//bit = getBits(&foo, 2, 2, 0x0f)
	if bit != 0x0A {
		t.Errorf("Expected bit value to be  %x but got %x", 0x0A, bit)
	}

	if foo != bar {
		t.Errorf("Source array changed unexpectedly Expected  %x but got %x", bar, foo)
	}

}

func TestInParse(t *testing.T) {

	var foo = [constants.MaxFrameDataLength]byte{0x10, 0xC0, 0xff, 0x00, 0x05, 6, 7, 0}
	//var bar = [constants.MaxFrameDataLength]byte{0x10, 0xC0, 0b00101010, 4, 5, 6, 7, 0}

	var b = getByte(&foo, 2)
	if b != 0xff {
		t.Errorf("byte value wrong expected %x but got %x", 0b00101010, b)
	}

	var sp = uintParser{
		byteOffset: 0,
	}
	var c = sp.parseInt16(&foo)
	if c != 0xC010 {
		t.Errorf("uint16 value wrong expected %x but got %x", 0xC010, c)
	}

	sp = uintParser{
		byteOffset: 1,
	}
	c = sp.parseInt16(&foo)
	if c != 0xffC0 {
		t.Errorf("uint16 value wrong expected %x but got %x", 0xffC0, c)
	}

	var ip = uintParser{byteOffset: 1}
	//var d = getuint32(&foo, 1)
	var d = ip.parseInt32(&foo)
	if d != 0x0500ffc0 {
		t.Errorf("uint32 value wrong expected %x but got %x", 0x0f00ffc0, d)
	}

}
