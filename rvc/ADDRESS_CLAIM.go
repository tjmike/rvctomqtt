package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/utils"
)

// NOTE This is just a special case of information request!

type AddressClaim struct {
	RvcItem
	desiredSA uint8
}

func (i *AddressClaim) GetDesiredSourceAddress() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.desiredSA

}

func (i *AddressClaim) SetDesiredSourceAddress(sa uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.desiredSA = sa
	// The desired SA also impacts the DGN
	var updated = i.dgn & 0xffffff00
	updated |= uint32(sa)
	i.dgn = updated

}
func (i *AddressClaim) SetDGN(dgn uint32) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.dgn = dgn
	i.name = dGNtoName[dgn]
}

func (i *AddressClaim) SetPriority(p uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.priority = p
}

func (i *AddressClaim) String() string {
	var s = i.RvcItem.String()
	s = fmt.Sprintf("%s desiredSA: %x", s, i.desiredSA)
	return s

}

func (r *AddressClaim) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.RvcItem.Init(from)

	// TODO some test for all these bits....
	var desiredSA = uint8(r.dgn & 0x0ff)

	if desiredSA != r.desiredSA {
		r.desiredSA = desiredSA
		r.lastChanged = r.timestamp
	}

}

// CreateFrame - create the data frame to be send in order to issue the command
func (r *AddressClaim) CreateFrame() *RvcFrame {
	var ret = RvcFrame{}
	r.lock.Lock()
	defer r.lock.Unlock()

	// TODO we really should not have to do this here !
	ret.SetEFF_RTR_ERR_Flag(constants.CAN_EFF_FLAG2)
	utils.SetByte(&ret.Data, 0, 0)   // Per spec fixed value
	utils.SetByte(&ret.Data, 1, 238) // Per spec fixed value
	utils.SetByte(&ret.Data, 2, 0)   // Per spec fixed value

	{
		var i byte
		for i = 3; i < constants.MaxFrameDataLength; i++ {
			utils.SetByte(&ret.Data, i, 0xff)
		}
	}

	// create the DGN - ignoring the SA and setting it to the desired SA
	var reqDGN = r.dgn
	reqDGN = reqDGN & 0xffffff00
	reqDGN = reqDGN | (uint32(r.desiredSA))

	ret.PGN.SetPGNValue(reqDGN)
	ret.SetPriority(0x06)
	ret.Length = 8
	ret.SetSourceAddress(254) // Per spec fixed value
	ret.SetCanMessage()
	return &ret
}

func (r *AddressClaim) Equals(o *AddressClaim) bool {
	if r.desiredSA != o.desiredSA {
		return false
	}
	if r.dgn != o.dgn {
		return false
	}
	return true
}
