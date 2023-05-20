package rvc

import (
	"fmt"
	"rvctomqtt/constants"
)

//package rvc
//
//import (
//"fmt"
//"rvctomqtt/constants"
//"rvctomqtt/utils"
//)

// AddressClaimed 0xEE00
type AddressClaimed struct {
	RvcItem
	serialNumber            uint32
	mfgrCode                uint32
	nodeInstance            uint8
	functionInstance        uint8
	compatabilityField1     uint8
	reserved1               uint8
	compatabilityField2     uint8
	compatabilityField3     uint8
	compatabilityField4     uint8
	arbitraryAddressCapable uint8 // Required. 1 – Node supports address claiming  0 – Node uses a fixed source address
}

//func (r *AddressClaimed) GetInstanceKey() interface{} {
//	return AddressClaimedKey{r.GetSourceAddress()}
//}

func (r *AddressClaimed) Equals(o *AddressClaimed) bool {
	if r.serialNumber != o.serialNumber {
		return false
	}
	if r.mfgrCode != o.mfgrCode {
		return false
	}
	if r.nodeInstance != o.nodeInstance {
		return false
	}
	if r.functionInstance != o.functionInstance {
		return false
	}
	if r.compatabilityField1 != o.compatabilityField1 {
		return false
	}
	if r.compatabilityField2 != o.compatabilityField2 {
		return false
	}
	if r.compatabilityField3 != o.compatabilityField3 {
		return false
	}
	if r.compatabilityField4 != o.compatabilityField4 {
		return false
	}
	if r.reserved1 != o.reserved1 {
		return false
	}
	if r.arbitraryAddressCapable != o.arbitraryAddressCapable {
		return false
	}

	return true
}

func (ac *AddressClaimed) SetCompatabilityField1(cf uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.compatabilityField1 = cf
}
func (ac *AddressClaimed) SetCompatabilityField2(cf uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.compatabilityField2 = cf & 0x7f
}
func (ac *AddressClaimed) SetCompatabilityField3(cf uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.compatabilityField3 = cf & 0x0f
}
func (ac *AddressClaimed) SetCompatabilityField4(cf uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.compatabilityField3 = cf & 0x07
}
func (ac *AddressClaimed) SetSerialNumber(snr uint32) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.serialNumber = snr & 0x00_1f_ff_ff
}
func (ac *AddressClaimed) SetManufacturerCode(mfgrCode uint32) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.mfgrCode = mfgrCode & 0x07ff
}
func (ac *AddressClaimed) SetNodeInstance(inst uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.nodeInstance = inst
}
func (ac *AddressClaimed) SetFunctionInstance(inst uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.functionInstance = inst
}
func (ac *AddressClaimed) SetArbitraryAddressCapable(c uint8) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.arbitraryAddressCapable = c & 0x01
}

// //////////////////////////////////
func (ac *AddressClaimed) SetSourceAddress(sa uint8) {
	ac.lock.RLock()
	defer ac.lock.RUnlock()
	ac.sourceAddress = sa
}
func (ac *AddressClaimed) SetDGN(dgn uint32) {
	ac.lock.RLock()
	defer ac.lock.RUnlock()
	ac.dgn = dgn
	ac.name = dGNtoName[dgn]
}
func (ac *AddressClaimed) SetPriority(p uint8) {
	ac.lock.RLock()
	defer ac.lock.RUnlock()
	ac.priority = p
}
func (ac *AddressClaimed) String() string {
	var s = ac.RvcItem.String()
	s = fmt.Sprintf("%s SNR: %x mfgr code: %x node inst: %x func inst: %x, arbitrary addres: %x\n",
		s, ac.serialNumber, ac.mfgrCode, ac.nodeInstance,
		ac.functionInstance, ac.arbitraryAddressCapable)
	return s

}
func (ac *AddressClaimed) Init(from *RvcFrame) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.RvcItem.Init(from)

	// b0 serial # lsb -- Optional. Required if multiple nodes from the same manufacturer may be present on the network.
	// b1 serial #
	// b2(0-4) serial # msb
	// b2(5-7) mfgr code LSB -- Required. Code obtained from SAE or RVIA.
	// b3  mfgr code
	// b4 (0-2) node instance -- For devices implementing multiple RV-C nodes (normally 0)
	// b4 (3-7) function instance -- Optional, intended to allow multiple instances of the same RV-C node, normally 0
	// b5 compatability field -- Optional, normally 0
	// b6 (0) reserved ZERO -- Always 0
	// b6 (1-7) compatability field -- Optional, normally 0
	// b7 (0-3) Optional, normally 0
	// b7 (4-6) Always 0
	// b7 (7) Required. 1 – Node supports address claiming 0 – Node uses a fixed source address

	ac.serialNumber = uint32(from.Data[0])
	ac.serialNumber |= uint32(from.Data[1]) << 8
	ac.serialNumber |= uint32(from.Data[2]&0b00011111) << 16

	ac.mfgrCode = uint32(from.Data[2] >> 5)
	ac.mfgrCode |= uint32(from.Data[3]) << 3

	ac.nodeInstance = from.Data[4] & 0x07
	ac.functionInstance = (from.Data[4] >> 3) & 0b00011111

	ac.compatabilityField1 = from.Data[5]

	ac.reserved1 = from.Data[6] & 0x01

	var v = (from.Data[6] >> 1)
	v = v & 0x7f
	ac.compatabilityField2 = v
	ac.compatabilityField2 = (from.Data[6] >> 1) & 0x7f

	ac.compatabilityField3 = from.Data[7] & 0x0F
	ac.compatabilityField4 = from.Data[7] >> 4 & 0x07 // always zero
	ac.arbitraryAddressCapable = from.Data[7] >> 7
}

// CreateFrame - create the data frame to be send in order to issue the command
func (ac *AddressClaimed) CreateFrame() *RvcFrame {
	var ret = RvcFrame{}
	ac.lock.Lock()
	defer ac.lock.Unlock()

	// b0 serial # lsb -- Optional. Required if multiple nodes from the same manufacturer may be present on the network.
	// b1 serial #
	// b2(0-4) serial # msb
	// b2(5-7) mfgr code LSB -- Required. Code obtained from SAE or RVIA.
	// b3  mfgr code
	// b4 (0-2) node instance -- For devices implementing multiple RV-C nodes (normally 0)
	// b4 (3-7) function instance -- Optional, intended to allow multiple instances of the same RV-C node, normally 0
	// b5 compatability field -- Optional, normally 0
	// b6 (0) reserved ZERO -- Always 0
	// b6 (1-7) compatability field -- Optional, normally 0
	// b7 (0-3) Optional, normally 0
	// b7 (4-6) Always 0
	// b7 (7) Required. 1 – Node supports address claiming 0 – Node uses a fixed source address

	ret.Data[0] = uint8(ac.serialNumber & 0x0ff)
	ret.Data[1] = uint8(ac.serialNumber >> 8 & 0x0ff)
	ret.Data[2] = uint8(ac.serialNumber >> 16 & 0x01f)
	ret.Data[2] |= uint8((ac.mfgrCode & 0x07) << 5)
	ret.Data[3] = uint8((ac.mfgrCode >> 3) & 0xff)
	ret.Data[4] = ac.nodeInstance & 0x07
	ret.Data[4] |= (ac.functionInstance & 0x01f) << 3
	ret.Data[5] = ac.compatabilityField1
	ret.Data[6] = ac.reserved1 & 0x01
	ret.Data[6] |= (ac.compatabilityField2 & 0x7f) << 1
	ret.Data[7] = ac.compatabilityField3 & 0x0f
	ret.Data[7] |= (ac.compatabilityField4 & 0x07) << 4
	ret.Data[7] |= (ac.arbitraryAddressCapable & 0x01) << 7

	// TODO we really should not have to do this here !
	ret.SetEFF_RTR_ERR_Flag(constants.CAN_EFF_FLAG2)

	// create the DGN - ignoring the SA and setting it to the desired SA
	var reqDGN = ac.dgn
	reqDGN = reqDGN & 0xffffff00
	//reqDGN = reqDGN | (uint32(ac.desiredSA))

	ret.PGN.SetPGNValue(reqDGN)
	ret.SetPriority(0x06)
	ret.Length = 8
	ret.SetSourceAddress(254) // Per spec fixed value
	ret.SetCanMessage()
	return &ret
}
