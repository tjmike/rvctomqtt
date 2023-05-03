package utils

// Utilities to manage extracting data from buffers
import (
	"encoding/binary"
	"rvctomqtt/constants"
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
						"ByteOffset": 0,
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

type UintParser struct {
	ByteOffset uint8
}

func (p *UintParser) parseInt8(buf *[constants.MaxFrameDataLength]byte) uint8 {
	return GetByte(buf, p.ByteOffset)
}
func (p *UintParser) ParseInt32(buf *[constants.MaxFrameDataLength]byte) uint32 {
	return Getuint32(buf, p.ByteOffset)
}
func (p *UintParser) ParseInt16(buf *[constants.MaxFrameDataLength]byte) uint16 {
	return Getuint16(buf, p.ByteOffset)
}

type bitParser struct {
	bitPosition uint8
	byteOffset  uint8
	nbits       uint8
}

func (p *bitParser) parseBits(buf *[constants.MaxFrameDataLength]byte) byte {
	var mask uint8 = GetMask(p.nbits)
	return GetBits(buf, p.byteOffset, p.bitPosition, mask)
}

// GetMask - get the 8 bit mask for the provided number of bits. A request for more that 8 buts will result in
// a full 8 bit mask 0xff
func GetMask(bits uint8) uint8 {
	var mask uint8 // default to 2 bits
	switch bits {
	case 1:
		mask = 0x01
		break
	case 2:
		mask = 0x03
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
	default: // clip at at 8 bits
		mask = 0xff
	}
	return mask
}

// Get the specified bits from the buffer
// GetBits give an input buffer, buff, go to zero based byteoffset in the buffer and make the bit at bitPos the least
// significant bit. Then apply the specified mask. This allows us to pull bits from a byte buffer. No error checking is
// performed
func GetBits(buf *[constants.MaxFrameDataLength]byte, byteOffset byte, bitPos byte, mask byte) byte {
	var ret = buf[byteOffset]
	ret = ret >> bitPos
	ret = ret & mask
	return ret
}

// take the provided value and apply the provided mask and shift the bits to the left
func SetBits[V uint8 | constants.Uint2](buf *[constants.MaxFrameDataLength]byte, byteOffset byte, bitPos byte, mask byte, val V) {

	var masked = (byte(val) & mask) << bitPos
	buf[byteOffset] = buf[byteOffset] | masked
}

// GetByte - get the specified byte from the buffer, no error checking is performed
func GetByte(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) byte {
	return buf[byteOffset]
}

func SetByte(buf *[constants.MaxFrameDataLength]byte, byteOffset byte, val byte) {
	buf[byteOffset] = val
}

// Getuint16 - get a 16 bit unsigned from the buffer. These are always byte aligned and always built assuming little edian
// as the RVC spec claims data is always little edian
func Getuint16(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) uint16 {
	return binary.LittleEndian.Uint16(buf[byteOffset:])
}

// Getuint16 - get a 16 bit unsigned from the buffer. These are always byte aligned and always built assuming little edian
// as the RVC spec claims data is always little edian
func Getuint32(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) uint32 {
	return binary.LittleEndian.Uint32(buf[byteOffset:])
}
