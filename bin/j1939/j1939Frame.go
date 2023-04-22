package j1939

import (
	"encoding/binary"
	"fmt"
	"rvctomqtt/can"
)

// J1939 Frame Info
// 29 bit CAN frame
// 3 Bit priority  18 bit PGN  8 bit source address (SA)
// PGN
// 1 bit reserved,1 bit page,8 but PDU,8 bit PDU specific

const (
	SA_MASK        uint32 = 0x000000ff // source address
	PGN_MASK       uint32 = 0x03FFFF00 // PGN MASK
	PGN_SHIFT             = 8
	PRIORITY_MASK         = 0b00011100_00000000_00000000_00000000
	PRIORITY_SHIFT        = 26

	// given a PGN
)

// 00001100 12 A B C

type J1939Frame struct {
	priority byte // 3 MS bits (of 29)
	PGN
	sourceAddress byte // LS Byte
	// This is a pointer to a can frame that was already created
	// Need to maintain awareness that we are passing these data structure
	// pointers around
	can.Frame
}

func CreateJ1939(can *can.Frame) J1939Frame {
	var a = can.ID
	fmt.Printf("a = %x", a)
	var b = a & PRIORITY_MASK
	var c = b >> PRIORITY_SHIFT

	return J1939Frame{
		priority:      byte(c), // byte((can.ID & PRIORITY_MASK) >> PRIORITY_SHIFT),
		PGN:           NewPGN(can),
		sourceAddress: byte(can.ID & 0xff),
		Frame:         *can,
	}
}

func (f *J1939Frame) GetPriority() byte {
	return f.priority
}

func (f *J1939Frame) GetSourceAddress() byte {
	return f.sourceAddress
}

func (frame *J1939Frame) BuildCanFrameX() {
	// TODO this cant be here
	//fmt.Println("BUILD CAN FRAME X  (J1939) - use platform specific method")
	frame.BuildCanFrame(binary.LittleEndian.Uint32)
}

func (frame *J1939Frame) BuildCanFrame(bytesTounit func([]byte) uint32) {
	//fmt.Println("Build Can Frame (J1939)")
	frame.Frame.BuildCanFrame(bytesTounit)
	var a = frame.Frame.ID
	var b = a & PRIORITY_MASK
	var c = b >> PRIORITY_SHIFT
	frame.priority = byte(c)
	frame.PGN = NewPGN(&frame.Frame)
	frame.sourceAddress = byte(frame.Frame.ID & 0xff)
}
