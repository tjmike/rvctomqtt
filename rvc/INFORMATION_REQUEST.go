package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/utils"
)

const DGN_DESIRED_ADDRESS_CLAIM uint32 = 238 << 8
const SA_ADDRESS_CLAIM uint8 = 254

// NOTE This is just a special case of information request!

// TODO this DGN is not part of "state" we can use the cache to prevent memory alloc - but we're not really
//
//	maintaining state...
type InformationRequest struct {
	RvcItem
	destinationAddress uint8 // DGN low ff=global , 254 = no address, for address claim request
	isAddressClaim     bool
	desiredDGN         uint32
	instance           uint8
	instance2          uint8
}

// GetInstanceKey - we just keep the last request for each destination address
//func (r *InformationRequest) GetInstanceKey() interface{} {
//	r.lock.RLock()
//	defer r.lock.RUnlock()
//	return InformationRequestKey{r.destinationAddress}
//}

func (i *InformationRequest) SetDesiredDGN(desiredDGN uint32) {
	i.desiredDGN = desiredDGN

}

func (i *InformationRequest) GetDesiredDGN() uint32 {
	return i.desiredDGN
}

func (i *InformationRequest) SetInstance(inst uint8) {
	i.instance = inst
}

func (i *InformationRequest) SetInstance2(inst uint8) {
	i.instance2 = inst
}

// SetAsAddressClaim - set the request fields to be address claim including the desiredSA
func (i *InformationRequest) SetAsAddressClaim(desiredSA uint8) {
	i.destinationAddress = desiredSA
	i.DGN = DGN_INFORMATION_REQUEST | uint32(desiredSA) // TODO maybe not needed here... done when create frame
	i.desiredDGN = DGN_DESIRED_ADDRESS_CLAIM
	i.SourceAddress = SA_ADDRESS_CLAIM
	i.isAddressClaim = true
	i.instance = 0xff
	i.instance2 = 0xff
}
func (i *InformationRequest) GetDestinationAddress() uint8 {
	return i.destinationAddress

}

func (i *InformationRequest) SetSourceAddress(sa uint8) {
	i.SourceAddress = sa
}
func (i *InformationRequest) SetDestinationAddress(sa uint8) {
	i.destinationAddress = sa
	// The destination request  also impacts the DGN
	var updated = i.DGN & 0xffffff00
	updated |= uint32(sa)
	i.DGN = updated

}

// SetDGN - set the DGN including the source address. If you want
// to set the source address separately then do it AFTER calling this
// method
func (i *InformationRequest) SetDGN(dgn uint32) {
	i.DGN = dgn
	i.Name = DGNName(dgn)
}

func (i *InformationRequest) SetPriority(p uint8) {
	i.Priority = p
}

func (i *InformationRequest) String() string {
	var s = i.RvcItem.String()
	// destinationAddress uint8 // DGN low ff=global , 254 = no address, for address claim request
	//	isAddressClaim     bool
	//	desiredDGN         uint32
	//	instance           uint8
	//	instance2          uint8

	s = fmt.Sprintf("%s destinationAddr: %x isAddreClaim %t desiredDGN=%x instance=%x instance2=%x ",
		s, i.destinationAddress, i.isAddressClaim, i.desiredDGN,
		i.instance, i.instance2,
	)
	return s

}

func (r *InformationRequest) Init(from *RvcFrame) {
	r.RvcItem.Init(from)

	// TODO some test for all these bits....
	var destAddr = uint8(r.DGN & 0x0ff)

	var changed = false
	if destAddr != r.destinationAddress {
		changed = true
		r.destinationAddress = destAddr
	}

	// byte 0 is LSB
	var desDGN uint32 = uint32(from.Data[0])
	desDGN |= uint32(from.Data[1]) << 8
	desDGN |= uint32(from.Data[2]&0x01) << 16

	if desDGN != r.desiredDGN {
		changed = true
		r.desiredDGN = desDGN
	}
	if r.desiredDGN == DGN_DESIRED_ADDRESS_CLAIM {
		r.isAddressClaim = true
	} else {
		r.isAddressClaim = false
	}

	var inst = from.Data[3]
	if inst != r.instance {
		changed = true
		r.instance = inst
	}
	var inst2 = from.Data[4]
	if inst2 != r.instance2 {
		changed = true
		r.instance2 = inst
	}

	if changed {
		r.LastChanged = r.Timestamp

	}
}

// CreateFrame - create the data frame to be sent in order to issue the command
func (r *InformationRequest) CreateFrame() *RvcFrame {
	var ret = RvcFrame{}

	// TODO we really should not have to do this here !
	ret.SetEFF_RTR_ERR_Flag(constants.CAN_EFF_FLAG2)

	utils.SetByte(&ret.Data, 0, uint8(r.desiredDGN&0x0ff))      // Per spec fixed value
	utils.SetByte(&ret.Data, 1, uint8((r.desiredDGN>>8)&0x0ff)) // Per spec fixed value
	utils.SetByte(&ret.Data, 2, uint8((r.desiredDGN>>16)&0x01)) // Per spec fixed value
	utils.SetByte(&ret.Data, 3, r.instance)
	utils.SetByte(&ret.Data, 4, r.instance2)

	{
		var i byte
		for i = 5; i < constants.MaxFrameDataLength; i++ {
			utils.SetByte(&ret.Data, i, 0xff)
		}
	}

	var pgnval = DGN_INFORMATION_REQUEST | uint32(r.destinationAddress)
	ret.SetPGNValue(pgnval)
	ret.SetPriority(0x06)
	ret.Length = 8
	ret.SetSourceAddress(r.SourceAddress)
	ret.SetCanMessage()
	return &ret
}

func (r *InformationRequest) Equals(o *InformationRequest) bool {

	if r.destinationAddress != o.destinationAddress {
		return false
	}
	if r.isAddressClaim != o.isAddressClaim {
		return false
	}
	if r.desiredDGN != o.desiredDGN {
		return false
	}
	if r.instance != o.instance {
		return false
	}
	if r.instance2 != o.instance2 {
		return false
	}
	return true
}
