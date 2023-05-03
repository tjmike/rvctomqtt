package j1939

import (
	"rvctomqtt/can"
)

const (
	PGN_RESERVED_MASK uint32 = 0x020000
	PGN_PAGE_MASK     uint32 = 0x010000
)

// TODO this should be in the types file
type bit uint8

// PGN is 18 bits
// reserved (1bit) page (1bit) pduFormat(byte) pduSpecific(byte)
type PGN struct {
	reserved    bit // really 1 bit - we could NOT store and use a method against pgn
	page        bit // really one bit - we could NOT store and use a method against pgn
	pduFormat   byte
	pduSpecific byte
	pgn         uint32 // the full PGN 18 bit word
}

// SetCanMessage - TODO this name needs to change  - this is just setting the PGN field from its data components
func (f *PGN) SetCanMessage() {
	f.pgn = uint32(f.pduSpecific)
	f.pgn |= uint32(f.pduFormat) << 8
	if f.IsReservedBitSet() {
		f.pgn |= PGN_RESERVED_MASK
	}
	if f.IsPageBitSet() {
		f.pgn |= PGN_PAGE_MASK
	}
}
func (pgn *PGN) GetPDUFormat() byte {
	return pgn.pduFormat
}
func (pgn *PGN) GetPDUSpecific() byte {
	return pgn.pduSpecific
}
func (pgn *PGN) GetPGN() uint32 {
	return pgn.pgn
}

// SetPGN - set the PGN from the info in the CAN frame
func (pgn *PGN) SetPGN(canFrame *can.Frame) {

	var tmp = (canFrame.ID & 0x03_ff_ff_00) >> 8
	pgn.SetPGNValue(tmp)

}

func (pgn *PGN) SetPGNValue(tmp uint32) {
	pgn.pgn = tmp
	if (pgn.pgn & PGN_RESERVED_MASK) != 0 {
		pgn.reserved = 1
	} else {
		pgn.reserved = 0
	}
	if (pgn.pgn & PGN_PAGE_MASK) != 0 {
		pgn.page = 1
	} else {
		pgn.page = 0
	}
	pgn.pduSpecific = byte(pgn.pgn & 0x0ff)
	pgn.pduFormat = byte(pgn.pgn >> 8 & 0x0ff)
}

func (pgn *PGN) IsReservedBitSet() bool {
	return pgn.reserved != 0
	//return (pgn.pgn & PGN_RESERVED_MASK) != 0
	//return pgn.reserved != 0
}

func (pgn *PGN) IsPageBitSet() bool {
	return pgn.page != 0
	//return (pgn.pgn & PGN_PAGE_MASK) != 0
}
