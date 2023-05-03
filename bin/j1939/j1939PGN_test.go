package j1939

import (
	"rvctomqtt/can"
	"testing"
)

func TestJ1939PGN(t *testing.T) {
	// A canFrame with all PGN bits set to zero and others set to 1
	var in uint32 = 0b11111100_00000000_00000000_11111111
	var canFrame = can.Frame{ID: in}

	var pgn PGN = PGN{}
	pgn.SetPGN(&canFrame)

	// only 18 bits should be considered - every PGN value should be zero
	if pgn.pduFormat != 0 {
		t.Errorf("pduFormat error expected  %x go %x", 0, pgn.pduFormat)
	}
	if pgn.pduSpecific != 0 {
		t.Errorf("pduSpecific error expected  %x go %x", 0, pgn.pduSpecific)
	}
	if pgn.IsReservedBitSet() {
		t.Error("reserved error expected bit unset but was set")
	}
	if pgn.IsPageBitSet() {
		t.Errorf("page error expected bit unset but was set")
	}

	// Single bit tests to make sure things are in the right place
	in = 0b00000000_10000000_00000000_00000000
	canFrame = can.Frame{ID: in}

	pgn = PGN{}
	pgn.SetPGN(&canFrame)
	if pgn.pduFormat != 0x80 {
		t.Errorf("pduFormat error expected  %x go %x", 0x80, pgn.pduFormat)
	}

	in = 0b00000000_00000000_10000000_00000000
	canFrame = can.Frame{ID: in}
	pgn = PGN{}
	pgn.SetPGN(&canFrame)
	if pgn.pduSpecific != 0x80 {
		t.Errorf("pduSpecific error expected  %x go %x", 0x80, pgn.pduSpecific)
	}

	in = 0b00000010_00000000_00000000_00000000
	canFrame = can.Frame{ID: in}
	pgn = PGN{}
	pgn.SetPGN(&canFrame)
	if !pgn.IsReservedBitSet() {
		t.Error("reserved error expected  true but got false")
	}

	in = 0b00000001_00000000_00000000_00000000
	canFrame = can.Frame{ID: in}
	pgn = PGN{}
	pgn.SetPGN(&canFrame)
	if !pgn.IsPageBitSet() {
		t.Error("page error expected true but got false")
	}

	// Test byte order , shifting and masking
	in = 0xff_02_04_ff
	var expected uint32 = 0x030204
	canFrame = can.Frame{ID: in}
	pgn = PGN{}
	pgn.SetPGN(&canFrame)
	var p = pgn.GetPGN()
	if p != expected {
		t.Errorf("Wrong PGN, expected %x got %x", expected, in)

	}
}

func TestSetCanMessageJ1939PGN(t *testing.T) {

	var p = PGN{}
	// expect 0x01 11 ee
	p.page = 1
	p.pduSpecific = 0xee
	p.pduFormat = 0x11
	p.SetCanMessage()
	var expected uint32 = 0x0111ee
	var pgn = p.GetPGN()
	if pgn != expected {
		t.Errorf("Wrong PGN, expected %x got %x", expected, pgn)
	}

	p.pgn = 0
	p.page = 0
	p.reserved = 1
	p.SetCanMessage()
	expected = 0x0211ee
	pgn = p.GetPGN()
	if pgn != expected {
		t.Errorf("Wrong PGN (page), expected %x got %x", expected, pgn)
	}
}
