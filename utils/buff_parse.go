package utils

// Utilities to manage extracting data from buffers
import (
	"encoding/binary"
	"rvctomqtt/constants"
	"rvctomqtt/rvc"
)

// DGN SPEC
// # HI
// # LOW (or 0 means only consider high)

// DATA SPEC
// b - bits
// # - byte offset
// # - #bits to grab
// B
// # - byte offset
// # - size 1,2,4

// Data Types
// unit       type     min   max     precision      special values
// %          uint8    0     125      0.5%
// instance   uint8    0     250                      0=all
// degC       uint8  -40     210       1deg c
//		      uint16  -273   1735      0.03125 °C
//		V     uint8     0     250       1 V
//	          uint16    0    3215.5    0.05 V
//	    A     uint8     0     250        1A

// MessageSpec
//
//				DGNHI
//			 DGNLOW - 00 means use only high
//		  Fields
//	     Field
//	        Byte

/*
		"messageSpec": {
			"DGNName": {
				"dgnHi": "0x111",
				"dgnLow": "0x11",
				"fields": {
					"Field1": {
						"byteOffset": 0,
						"bitOffset": 0,
						"bits": 8,
						"status": "retired",
						"engine": "Gecko",
						"engine_version": "1.7"
					}
				}
			}
		}
	}
*/
type message struct {
	dgn    uint32
	sa     byte
	fields int
}

type RVCMessageIF interface {
	getRawFrame() *rvc.RvcFrame
	getName() *string                  // comes from a mapping of the DGN
	getFields() *map[string]RVCFieldIF // comes from mapping of DGN fields
}

type RVCFieldIF interface {
	getFieldIndex() int
	getValue() float64
	getFieldName() string
	getUnits() string
}

type uintParser struct {
	byteOffset uint8
}

func (p *uintParser) parseInt8(buf *[constants.MaxFrameDataLength]byte) uint8 {
	return getByte(buf, p.byteOffset)
}
func (p *uintParser) parseInt32(buf *[constants.MaxFrameDataLength]byte) uint32 {
	return getuint32(buf, p.byteOffset)
}
func (p *uintParser) parseInt16(buf *[constants.MaxFrameDataLength]byte) uint16 {
	return getuint16(buf, p.byteOffset)
}

type bitParser struct {
	bitPosition uint8
	byteOffset  uint8
	nbits       uint8
}

func (p *bitParser) parseBits(buf *[constants.MaxFrameDataLength]byte) byte {
	var mask uint8 = 0x03 // default to 2 bits
	switch p.nbits {
	case 1:
		mask = 0x01
		break
	case 3:
		mask = 0x07
		break
	case 4:
		mask = 0x0f
		break
	case 5:
		mask = 0x1f
		break
	case 6:
		mask = 0x3f
		break
	case 7:
		mask = 0x7f
		break
	case 8:
		mask = 0xff
		break
	default:
		mask = 0x03

	}
	return getBits(buf, p.byteOffset, p.bitPosition, mask)
}

// Get the specified bits from the buffer
// getBits give an input buffer, buff, go to zero based byteoffset in the buffer and make the bit at bitPos the least
// significant bit. Then apply the specified mask. This allows us to pull bits from a byte buffer. No error checking is
// performed
func getBits(buf *[constants.MaxFrameDataLength]byte, byteOffset byte, bitPos byte, mask byte) byte {
	var ret = buf[byteOffset]
	ret = ret >> bitPos
	ret = ret & mask
	return ret
}

// getByte - get the specified byte from the buffer, no error checking is performed
func getByte(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) byte {
	return buf[byteOffset]
}

// getuint16 - get a 16 bit unsigned from the buffer. These are always byte aligned and always built assuming little edian
// as the RVC spec claims data is always little edian
func getuint16(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) uint16 {
	return binary.LittleEndian.Uint16(buf[byteOffset:])
}

// getuint16 - get a 16 bit unsigned from the buffer. These are always byte aligned and always built assuming little edian
// as the RVC spec claims data is always little edian
func getuint32(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) uint32 {
	return binary.LittleEndian.Uint32(buf[byteOffset:])
}
