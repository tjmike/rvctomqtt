package utils

// Utilities to manage extracting data from buffers
import (
	"encoding/binary"
	"rvctomqtt/constants"
)

// b - bits
// # - byte offset
// # - #bits to grab
// B
// # - byte offst
// # - size 1,2,4

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
