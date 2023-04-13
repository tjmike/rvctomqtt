package utils

import (
	"rvctomqtt/constants"
	"testing"
)

// C    0
//1100 0000
// //10   10000
//010000
//
// 0101 0000
// 0110 0000

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
		bitPosition: 2,
		byteOffset:  6,
		nbits:       1,
	}
	bit = bp.parseBits(&foo)
	if bit != 1 {
		t.Errorf("Expected single bit value to be  %x but got %x", 1, bit)
	}

	bp = bitParser{
		bitPosition: 2,
		byteOffset:  6,
		nbits:       3,
	}
	bit = bp.parseBits(&foo)
	if bit != 1 {
		t.Errorf("Expected single bit value to be  %x but got %x", 1, bit)
	}

	bp = bitParser{
		bitPosition: 3,
		byteOffset:  2,
		nbits:       5,
	}
	bit = bp.parseBits(&foo)
	if bit != 5 {
		t.Errorf("Expected single bit value to be  %x but got %x", 5, bit)
	}

	bp = bitParser{
		bitPosition: 2,
		byteOffset:  2,
		nbits:       6,
	}
	bit = bp.parseBits(&foo)
	if bit != 10 {
		t.Errorf("Expected single bit value to be  %x but got %x", 10, bit)
	}

	bp = bitParser{
		bitPosition: 1,
		byteOffset:  1,
		nbits:       6,
	}
	bit = bp.parseBits(&foo)
	if bit != 32 {
		t.Errorf("Expected single bit value to be  %x but got %x", 16, bit)
	}

	bp = bitParser{
		bitPosition: 1,
		byteOffset:  1,
		nbits:       7,
	}
	bit = bp.parseBits(&foo)
	if bit != 0x60 {
		t.Errorf("Expected single bit value to be  %x but got %x", 16, bit)
	}

	bp = bitParser{
		bitPosition: 0,
		byteOffset:  1,
		nbits:       8,
	}
	bit = bp.parseBits(&foo)
	if bit != 0xC0 {
		t.Errorf("Expected single bit value to be  %x but got %x", 0xc0, bit)
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

	var bp = uintParser{byteOffset: 2}
	var b = bp.parseInt8(&foo)
	//var b = getByte(&foo, 2)
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
	var d = ip.parseInt32(&foo)
	if d != 0x0500ffc0 {
		t.Errorf("uint32 value wrong expected %x but got %x", 0x0f00ffc0, d)
	}

}
