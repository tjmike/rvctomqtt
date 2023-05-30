package rvc

import (
	"fmt"
	"rvctomqtt/utils"
)

type InitialPacket struct {
	RvcItem
	reserved      uint8  // always zero
	messageLength uint16 // 0-1785
	packetCount   uint8  // NOTE: 255 (0xff) really means 255 packets
	dgn           uint32 // byte5 LSB byte6 lsb byte7 < MSB
}

func (r *InitialPacket) getMessageLength() uint16 {
	return r.messageLength
}
func (r *InitialPacket) GetPacketCount() uint8 {
	return r.packetCount
}
func (r *InitialPacket) GetRequestedDGN() uint32 {
	return r.dgn
}

// GetInstanceKey - we can really only loosely tie initial packet to DATA_PACKET to that end we just keep
// track of the source address and use this to tie the initial packet to data packets.
// the
//func (r *InitialPacket) GetInstanceKey() interface{} {
//	return InitialPacketKey{r.SourceAddress}
//}

func (r *InitialPacket) String() string {
	var s = r.RvcItem.String()
	s = fmt.Sprintf("%s reserved: %x messageLength: %d packetCount: %d DGN: %x\n",
		s,
		r.reserved,
		r.messageLength,
		r.packetCount,
		r.dgn,
	)
	return s
}
func (r *InitialPacket) Init(from *RvcFrame) {
	r.RvcItem.Init(from)
	var dataBytes = &from.Data

	{
		r.reserved = utils.GetByte(dataBytes, 0)
	}
	{
		var tmp uint16 = uint16(utils.GetByte(dataBytes, 1))
		tmp |= uint16(utils.GetByte(dataBytes, 2)) << 8
		r.messageLength = tmp
	}
	r.packetCount = utils.GetByte(dataBytes, 3)

	// no byte 4?????
	{
		var dgn uint32 = uint32(utils.GetByte(dataBytes, 5))
		dgn |= uint32(utils.GetByte(dataBytes, 6)) << 8
		dgn |= uint32(utils.GetByte(dataBytes, 7)) << 16
	}
}
