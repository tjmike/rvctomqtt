package constants

// We want to use the constants but don't want to drag in implementation details - so they are stuffed here fir now
const (
	MAX_MESSAGE        int32 = 16 // largest size a can message can be
	MaxFrameDataLength uint8 = 8  // largest sie a data fram3 can be

	CAN_EFF_FLAG uint32 = 0x80000000 /* EFF/SFF is set in the MSB */
	CAN_RTR_FLAG uint32 = 0x40000000 /* remote transmission request */
	CAN_ERR_FLAG uint32 = 0x20000000 /* error message frame */
	CAN_SFF_MASK uint32 = 0x000007FF /* standard frame format (SFF) */
	CAN_EFF_MASK uint32 = 0x1FFFFFFF /* extended frame format (EFF) */

)
