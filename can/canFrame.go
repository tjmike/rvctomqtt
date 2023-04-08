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

// The lifecycle of a frame is as as follows:
// 1) Create or pull from pool (in can reader thread)
// 2) Populate message bytes (in can reader thread)
// 3) Send over channel
// 4) Consumer
type Frame struct {
	Timestamp time.Time
	//type frame struct {
	// 32 bit CAN_ID + EFF/RTR/ERR flags
	// bit 0-28: CAN identifier (11/29 bit)
	// bit 29: error message flag (ERR)
	// bit 30: remote transmision request (RTR)
	// bit 31: extended frame format (EFF)
	// This is the canID INCLUDING the 3 MSB flags
	ID uint32
	// Data length (0-8)
	Length uint8
	// these three bytes not used
	Flags uint8
	Res0  uint8
	Res1  uint8
	// data bytes - can have zero to max bytes
	Data [MaxFrameDataLength]uint8

	// These are the raw message bytes. This is what we send to the driver as fread takes bytes. It may be possible
	// send a struct of the right type that is (unsafely?) cast as bytes. Even if that can be done, maybe using
	// byte[] is more readable/understandable?
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

// BuildCanFrame Build the can frame from the raw data bytes. Conversion of byte[] to uint32 is platform specific
// so we pass in (a pointer to?) a function that provides that conversion for us.
func (frame *Frame) BuildCanFrame(bytesTounit func([]byte) uint32) {

	// Raspberry PI May be this
	// The test passed the other way on PI - so this means that something may be messed up
	// Leave this comment and keep an eye out. It could be  that the bytes coming from
	// the socket are messed up.
	//frame.ID = binary.LittleEndian.Uint32((*msg).CanMessage[0:])
	// MAC does this
	//frame.ID = binary.BigEndian.Uint32((*msg).CanMessage[0:])
	//frame.Timestamp = (*msg).Timestamp
	frame.ID = bytesTounit(frame.MessageBytes[0:])
	// setFrameID(frame)
	frame.Length = (*frame).MessageBytes[4]
	frame.Flags = (*frame).MessageBytes[5]
	frame.Res0 = (*frame).MessageBytes[6]
	frame.Res1 = (*frame).MessageBytes[7]

	// not needed if we set the length and all the bytes
	for i := range frame.Data {
		frame.Data[i] = 0xff
	}
	{
		var fl int = int(frame.Length)
		for i, v := range (*frame).MessageBytes[8:15] {
			if i < fl {
				frame.Data[i] = v
			} else {
				frame.Data[i] = 0xff
			}
		}
	}

}
