package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/utils"
	"strings"
	"unicode"
)

// TODO : we really need to utilize message length and packet count - for now we just fetch the data
type DataPacket struct {
	RvcItem
	packetNumber uint8 // 1-255
	message      string
}

// GetInstanceKey - this identifies a particular Data Packet Instance for a given source address. Ideally this means
// that a cache can retain all the data packets for each device (source address). It's still possible to retain bad data
// as each device might send items based on each instance - see the INITAL PACKET command to understand the issue.
// At the very least the initial packet should be used to determine what packets to read here
//
//	func (r *DataPacket) GetInstanceKey() interface{} {
//		r.lock.RLock()
//		defer r.lock.RUnlock()
//		return DataPacketKey{
//			r.sourceAddress,
//			r.packetNumber,
//		}
//	}
func (r *DataPacket) String() string {
	var s = r.RvcItem.String()
	s = fmt.Sprintf("%s number: %d message: %s\n",
		s,
		r.packetNumber,
		r.message,
	)
	return s
}
func (r *DataPacket) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.RvcItem.Init(from)
	var dataBytes = &from.Data

	r.packetNumber = utils.GetByte(dataBytes, 0)

	var sb strings.Builder
	for i := 1; i < int(constants.MaxFrameDataLength); i++ {
		var b = dataBytes[i]
		if b < unicode.MaxASCII {
			sb.WriteByte(dataBytes[i])
		}
	}
	r.message = sb.String()
}
