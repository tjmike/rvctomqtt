package rvc

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"rvctomqtt/bin/j1939"
)

// J1939 Frame Info
// 29 bit CAN frame
// 3 Bit priority  18 bit PGN  8 bit source address (SA)
// PGN
// 1 bit reserved,1 bit page,8 but PDU,8 bit PDU specific
// bits 28-0
// 28-26 priority
// 25 reserved
// 24 page
// 23-16 PDU
// 15-8 PDU Specific
// 0-7 SA

// From RVC spec
// 29 bit can frame
// BASE ID - 13 most significant bits
// Extended ID - 2 least significant bytes
// For that 29 bit message
// Bits ----
// 28-26 Priority 111 Highest - 000 Lowest (j1939 priority)
//    25 Reserved - always zero
// 24-18 DGN (Data Group Number) High
// 17-16      Identifies how the data packet should be parsed, possibly in combination with the DGN-Low
// 15-8 DGN low - Either determines the target node for the message, or with the DGN-High determines how the data packet should be parsed.
//  7-0 Source Address(SA)

// bits 28-0

// Data  = always 8 bytes

type RvcFrame struct {
	j1939.J1939Frame
}

// For priority we already have GetPriority()

func (msg *RvcFrame) DGN() uint32 {
	var ret uint32 = uint32(msg.DGNHigh())
	ret = ret << 8
	ret |= uint32(msg.DGNLow())
	return ret
}
func (msg *RvcFrame) DGNHigh() uint16 {

	// TODO we don't want to compute this every single time...

	var ret = uint16(msg.PGN.GetPDUFormat())

	// set MSBit if page bit is set
	if msg.PGN.IsPageBitSet() {
		ret |= 0x100
	}

	return ret
}
func (msg *RvcFrame) DGNLow() uint8 {
	return msg.PGN.GetPDUSpecific()
}

func (msg *RvcFrame) String() string {
	return fmt.Sprintf("TS:%d - canID=%X DGNH = %x DGNL=%x ERR=%t RTR=%t EXT=%t Data Len=%d Data=%s RAW=%s", // DGN=%X SRCADDR=%X len=%d - % X" ,
		msg.Timestamp.UnixNano(),
		msg.CanID(),
		msg.DGNHigh(),
		msg.DGNLow(),
		msg.IsERR(),
		msg.IsRTR(),
		msg.IsExtended(),
		msg.DataLength(),
		hex.EncodeToString(msg.Data[0:]),
		hex.EncodeToString(msg.MessageBytes[0:]),
	)
}

func (frame *RvcFrame) BuildCanFrameX() {
	// TODO this cant be here
	//fmt.Println("BUILD CAN FRAME X  (RVC) - use platform specific method")
	frame.BuildCanFrame(binary.LittleEndian.Uint32)
}

func (frame *RvcFrame) BuildCanFrame(bytesTounit func([]byte) uint32) {
	//fmt.Println("Build Can Frame (RVC)")
	frame.J1939Frame.BuildCanFrame(bytesTounit)
}

// We need some means to read a value from the byte array
// A value can be:
//   * 2 bit field - always within a byte
//   * 4 but field
//   * 8 bit value
//   * 16 but value
//   * 32 bit value
// maybe we can have a function that takes a byte offset (slice) and

type RVCMessageIF interface {
	getRawFrame() *RvcFrame
	getName() *string                  // comes from a mapping of the DGN
	getFields() *map[string]RVCFieldIF // comes from mapping of DGN fields
}

type RVCFieldIF interface {
	getFieldIndex() int
	getValue() float64
	getFieldName() string
	getUnits() string
}
