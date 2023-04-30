package j1939

import (
	"encoding/binary"
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

	//PAGE_MASK         = 0b00000001_00000000_00000000_00000000
	//PAGE_MASK         = 0b00000001_00000000_00000000_00000000

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

//func CreateJ1939(can *can.Frame) J1939Frame {
//	var a = can.ID
//	fmt.Printf("a = %x", a)
//	var b = a & PRIORITY_MASK
//	var c = b >> PRIORITY_SHIFT
//
//	return J1939Frame{
//		priority:      byte(c), // byte((can.ID & PRIORITY_MASK) >> PRIORITY_SHIFT),
//		PGN:           NewPGN(can),
//		sourceAddress: byte(can.ID & 0xff),
//		Frame:         *can,
//	}
//}

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

// BuildCanFrame - build out the J1939 (and CAN) frame from the raw message bytes already in the struct
func (frame *J1939Frame) BuildCanFrame(bytesTounit func([]byte) uint32) {
	frame.Frame.BuildCanFrame(bytesTounit)

	// 29 bit message (or 11 bit if not extended. Were only considering extended right now
	var c = (frame.Frame.ID & PRIORITY_MASK) >> PRIORITY_SHIFT

	//var b = a & PRIORITY_MASK
	//var c = b >> PRIORITY_SHIFT
	frame.priority = byte(c)

	frame.PGN.SetPGN(&frame.Frame)
	frame.sourceAddress = byte(frame.Frame.ID & 0xff)

	// TODO should not be here - was this a test?
	//frame.SetCanMessage()
}

// SetCanMessage - set the CAN message from the fields of the J1939 message. This just sets toe ID bits as the
// the data payload bits are expected to be set by something that knows how to interpret them. In some cases it may
// be that SetCanMessages does not init those fields because the message is read only.
func (f *J1939Frame) SetCanMessage() {
	// assume that the data bytes are already set. Some other code must be mapping fields back to data bytes
	//
	// 3 biy priority
	var canID = (uint32(f.priority)) << PRIORITY_SHIFT
	// 18 bit PGN
	canID = canID | (f.PGN.GetPGN() << 8)
	// 8 bit source address
	canID = canID | uint32(f.sourceAddress)

	// We need ro rebuild the CANID in the frame before we call its SetCanMessage. Note that we are leaving the
	// the rew PHN value alone.
	f.Frame.ID = canID
	f.Frame.SetCanMessage()
	/*
		var v uint32 = uint32(f.sourceAddress)
		v = v | uint32(f.PGN.GetPDUSpecific())<<8
		v = v | uint32(f.PGN.GetPDUFormat())<<16

		if f.IsPageBitSet() {
			v |= PGN_PAGE_MASK << 8
		}
		if f.IsReservedBitSet() {
			v |= PGN_RESERVED_MASK << 8
		}

		v |= uint32(f.priority) << PRIORITY_SHIFT
		// set
		f.Frame.ID = v
		f.Frame.SetCanMessage()
		fmt.Printf("AFT: %x\n", f.Frame.MessageBytes)
	*/
}
