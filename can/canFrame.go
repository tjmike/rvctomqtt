package can

import (
	"fmt"
	"time"
)

const (
	MAX_MESSAGE        int32 = 16
	MaxFrameDataLength uint8 = 8
	// CAN_EFF_FLAG we want these flags to be seen on all platforms so we dont use the
	// ones from the unix package
	CAN_EFF_FLAG uint32 = 0x80000000 /* EFF/SFF is set in the MSB */
	CAN_RTR_FLAG uint32 = 0x40000000 /* remote transmission request */
	CAN_ERR_FLAG uint32 = 0x20000000 /* error message frame */
	CAN_SFF_MASK uint32 = 0x000007FF /* standard frame format (SFF) */
	CAN_EFF_MASK uint32 = 0x1FFFFFFF /* extended frame format (EFF) */

)

// Frame represents a standard CAN data frame

//struct can_frame {
//canid_t can_id;  /* 32 bit CAN_ID + EFF/RTR/ERR flags */
//union {
///* CAN frame payload length in byte (0 .. CAN_MAX_DLEN)
// * was previously named can_dlc so we need to carry that
// * name for legacy support
// */
//__u8 len;
//__u8 can_dlc; /* deprecated */
//} __attribute__((packed)); /* disable padding added in some ABIs */
//__u8 __pad; /* padding */
//__u8 __res0; /* reserved / padding */
//__u8 len8_dlc; /* optional DLC for 8 byte payload length (9 .. 15) */
//__u8 data[CAN_MAX_DLEN] __attribute__((aligned(8)));
//};

type Frame struct {
	Timestamp time.Time
	//type frame struct {
	// 32 bit CAN_ID + EFF/RTR/ERR flags
	// bit 0-28: CAN identifier (11/29 bit)
	// bit 29: error message flag (ERR)
	// bit 30: remote transmision request (RTR)
	// bit 31: extended frame format (EFF)
	ID uint32
	// Data length
	Length uint8
	// these three bytes not used
	Flags uint8
	Res0  uint8
	Res1  uint8
	// data bytes
	Data [MaxFrameDataLength]uint8

	MessageBytes [MAX_MESSAGE]byte
}

func (f *Frame) IsExtended() bool {
	return (f.ID & CAN_EFF_FLAG) != 0
}
func (f *Frame) IsRTR() bool {
	return (f.ID & CAN_RTR_FLAG) != 0
}

func (f *Frame) IsERR() bool {
	return (f.ID & CAN_ERR_FLAG) != 0
}

// This is just the canID without the additional flags
func (f *Frame) CanID() uint32 {
	if f.IsExtended() {
		return f.ID & CAN_EFF_MASK
	} else {
		return f.ID & CAN_SFF_MASK
	}
}

func (msg *Frame) ToString() string {
	return fmt.Sprintf("TS:%d - canID=%X ERR=%t RTR=%t EXT=%t ", // DGN=%X SRCADDR=%X len=%d - % X" ,
		msg.Timestamp.UnixNano(),
		msg.CanID(),
		msg.IsERR(),
		msg.IsRTR(),
		msg.IsExtended())
	//msg.SourceAddress(),
	//msg.payloadSize(),
	//msg.canMessage )

}
