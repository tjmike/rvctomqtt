package can

import (
	"encoding/binary"
	"time"
)

type RawCanMessage struct {
	Timestamp  time.Time
	CanMessage [MAX_MESSAGE]byte
}

const (
	// Starting index of the can frame payload
	canPayloadStart = 8
)

//func (msg *RawCanMessage) toString() string {
//	return fmt.Sprintf("%d - canID=%X ERR=%t EXT=%t DGN=%X SRCADDR=%X len=%d - % X",
//		msg.timestamp.UnixNano(),
//		msg.getCANIDNoFlags(),
//		msg.isErrorFrame(),
//		msg.isExtendedFrame(),
//		msg.DGN(),
//		msg.SourceAddress(),
//		msg.payloadSize(),
//		msg.canMessage)
//}

func (msg *RawCanMessage) GetCANID() uint32 {
	// TODO this is whats works for Raspberry PI, but it's not been tested on other platforms
	return binary.LittleEndian.Uint32((*msg).CanMessage[0:])
}
