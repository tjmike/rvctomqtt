package can

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"rvctomqtt/constants"
	"time"
)

// We want the CAN_XXX flags to be seen on all platforms so we don't use the
// ones from the unix package.

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
// 4) Consumer uses the frame
// 5) Sends frame BACK over channel to be recycled

type Frame struct {
	Timestamp time.Time
	// 32 bit CAN_ID + EFF/RTR/ERR flags
	// bit 0-28: CAN identifier (11/29 bit)
	// bit 29: error message flag (ERR)
	// bit 30: remote transmision request (RTR)
	// bit 31: extended frame format (EFF) (MSB)
	// This is the canID INCLUDING the 3 MSB flags
	ID uint32
	// Data length (0-8)
	Length uint8
	// these three bytes not used
	Flags uint8
	Res0  uint8
	Res1  uint8
	// data bytes - can have zero to max bytes
	Data [constants.MaxFrameDataLength]uint8

	// These are the raw message bytes. This is what we send to the driver as fread() takes bytes. It may be possible
	// send a struct of the right type that is (unsafely?) cast as bytes. Even intf that can be done, maybe using
	// byte[] is more readable/understandable...
	MessageBytes [constants.MAX_MESSAGE]byte
}

// IsExtended - true intf this frame is extended format.
func (f *Frame) IsExtended() bool {
	return (f.ID & constants.CAN_EFF_FLAG) != 0
}

// IsRTR - true intf this is a remote transmission request
func (f *Frame) IsRTR() bool {
	return (f.ID & constants.CAN_RTR_FLAG) != 0
}

// IsERR - true intf this is an error frame
func (f *Frame) IsERR() bool {
	return (f.ID & constants.CAN_ERR_FLAG) != 0
}

// CanID - This is just the canID without the additional flags. Either a 29 or 11 bit value.
func (f *Frame) CanID() uint32 {
	if f.IsExtended() {
		return f.ID & constants.CAN_EFF_MASK
	} else {
		return f.ID & constants.CAN_SFF_MASK
	}
}

func (msg *Frame) ToString() string {
	return fmt.Sprintf("TS:%d - canID=%X ERR=%t RTR=%t EXT=%t Data Len=%d Data=%s RAW=%s", // DGN=%X SRCADDR=%X len=%d - % X" ,
		msg.Timestamp.UnixNano(),
		msg.CanID(),
		msg.IsERR(),
		msg.IsRTR(),
		msg.IsExtended(),
		msg.DataLength(),
		hex.EncodeToString(msg.Data[0:]),
		hex.EncodeToString(msg.MessageBytes[0:]),
	)
}

// DataLength - Data length for this frame
func (msg *Frame) DataLength() uint8 {
	return msg.Length
}

// BuildCanFrame Build the can frame from the raw data bytes. Conversion of byte[] to uint32 is platform specific
// so we pass in (a pointer to?) a function that provides that conversion for us.
func (frame *Frame) BuildCanFrame(bytesTounit func([]byte) uint32) {
	frame.ID = bytesTounit(frame.MessageBytes[0:])
	frame.Length = (*frame).MessageBytes[4]
	frame.Flags = (*frame).MessageBytes[5]
	frame.Res0 = (*frame).MessageBytes[6]
	frame.Res1 = (*frame).MessageBytes[7]

	// Set the data bytes, making any unused slots 0xff
	{
		// Get the data bytes
		var i uint8 = 0
		for ; i < frame.Length; i++ {
			frame.Data[i] = frame.MessageBytes[8+i]
		}
		// set unused bytes to all ones
		i = frame.Length // should already be this, but easier to understand...
		for ; i < constants.MaxFrameDataLength; i++ {
			frame.Data[i] = 0xFF
		}

	}
}

func (frame *Frame) GetMessage() *[16]byte {
	return &frame.MessageBytes
}
func (frame *Frame) SetTimeStamp(t time.Time) {
	frame.Timestamp = t
}
func (frame *Frame) BuildCanFrameX() {
	// TODO this cant be here
	fmt.Println("BUILD CAN FRAME EX - use platform specific method")
	frame.BuildCanFrame(binary.LittleEndian.Uint32)
}
