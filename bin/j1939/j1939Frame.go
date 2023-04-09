package j1939

import (
	"fmt"
	"rvctomqtt/can"
)

const (
	SA_MASK        uint32 = 0x000000ff // source address
	PGN_MASK       uint32 = 0x03FFFF00 // PGN MASK
	PGN_SHIFT             = 8
	PRIORITY_MASK         = 0b00011100_00000000_00000000_00000000
	PRIORITY_SHIFT        = 26

	// given a PGN
)

// 00001100 12 A B C

type j1939Frame struct {
	priority byte // 3 MS bits (of 29)
	PGN
	sourceAddress byte // LS Byte
	// This is a pointer to a can frame that was already created
	// Need to maintain awareness that we are passing these data structure
	// pointers around
	*can.Frame
}

func CreateJ1939(can *can.Frame) j1939Frame {
	var a = can.ID
	fmt.Printf("a = %x", a)
	var b = a & PRIORITY_MASK
	var c = b >> PRIORITY_SHIFT

	return j1939Frame{
		priority:      byte(c), // byte((can.ID & PRIORITY_MASK) >> PRIORITY_SHIFT),
		PGN:           NewPGN(can),
		sourceAddress: byte(can.ID & 0xff),
		Frame:         can,
	}
}

func (f *j1939Frame) GetPriority() byte {
	return f.priority
}

func (f *j1939Frame) GetSourceAddress() byte {
	return f.sourceAddress
}
