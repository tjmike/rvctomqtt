package j1939

import "rvctomqtt/can"

const (
	PGN_RESERVED_MASK uint32 = 0x020000
	PGN_PAGE_MASK     uint32 = 0x010000
)

// PGN is 18 bits
// reserved (1bit) page (1bit) pduFormat(byte) pduSpecific(byte)
type PGN struct {
	//reserved    byte // really 1 bit - we could NOT store and use a method against pgn
	//page        byte // really one bit - we could NOT store and use a method against pgn
	pduFormat   byte
	pduSpecific byte
	pgn         uint32 // the full PGN word
}

// func NewPGN(rawPGN uint32) PGN {
func NewPGN(canFrame *can.Frame) PGN {
	var rawPGN = (canFrame.ID & 0x03_ff_ff_00) >> 8

	var pduSpecific byte = byte(rawPGN & 0x0ff)
	var pduFormat byte = byte(rawPGN >> 8 & 0x0ff)
	var ret = PGN{
		//reserved:    reserved,
		//page:        page,
		pduFormat:   pduFormat,
		pduSpecific: pduSpecific,
		pgn:         rawPGN,
	}
	return ret
}

func (pgn *PGN) IsReservedBitSet() bool {
	return (pgn.pgn & PGN_RESERVED_MASK) != 0
	//return pgn.reserved != 0
}

func (pgn *PGN) IsPageBitSet() bool {
	return (pgn.pgn & PGN_PAGE_MASK) != 0
}
