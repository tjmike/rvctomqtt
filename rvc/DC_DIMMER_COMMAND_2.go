package rvc

import (
	"encoding/json"
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
)

type DCDimmerCommand2 struct {
	//rvcItem             RvcItem
	RvcItem
	//timestamp         time.Time // last timestamp that we saw a command
	//lastChanged       time.Time // last time this instance actually changed
	deviceInstance    byte    // 0
	group             uint8   // 1
	desiredBrightness float64 // 2 (%) this does NOT mean on/of - see command
	command           uint8   // 3 - see the spec 1-on, 2 on-delay,
	delayDuration     uint8   // 4 seconds

	interlockStatus constants.Uint2 // 5 (0,1)
	reserved1       constants.Uint2 // 5 (2-3)
	reserved2       constants.Uint2 // 5 (4-5)
	reserved3       constants.Uint2 // 5 (6-7)

	rampTime float64 // seconds (0.1) 0-25 seconds
	reserved uint8   //
	//sourceAddress   byte    //
	//lock            sync.RWMutex
}

func (r *DCDimmerCommand2) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DeviceInstance uint8 `json:"deviceInstance"`
		Group          uint8 `json:"group"`
	}{
		DeviceInstance: r.deviceInstance,
		Group:          r.group,
	})
}

//	func (r *DCDimmerCommand2) GetInstanceKey() interface{} {
//		r.lock.RLock()
//		defer r.lock.RUnlock()
//		return DGNInstanceKey{
//			r.dgn,
//			r.deviceInstance,
//		}
//	}
func (i *DCDimmerCommand2) Equals(o *DCDimmerCommand2) bool {
	// Nasty - do we really want to get two locks? What happens if we get a ref to ourself (eg both args point to same things)
	// NOT locking for now ....

	if !i.RvcItem.Equals(&o.RvcItem) {
		return false
	}

	if i.deviceInstance != o.deviceInstance {
		return false
	}
	if i.group != o.group {
		return false
	}
	if i.desiredBrightness != o.desiredBrightness {
		return false
	}
	if i.command != o.command {
		return false
	}
	if i.delayDuration != o.delayDuration {
		return false
	}
	if i.interlockStatus != o.interlockStatus {
		return false
	}
	if i.reserved1 != o.reserved1 {
		return false
	}
	if i.reserved2 != o.reserved2 {
		return false
	}
	if i.reserved3 != o.reserved3 {
		return false
	}
	if i.rampTime != o.rampTime {
		return false
	}
	if i.reserved != o.reserved {
		return false
	}

	return true
}

//func (i *DCDimmerCommand2) GetDGN() uint32 {
//	return DGN_DC_DIMMER_COMMAND_2
//}
//func (i *DCDimmerCommand2) GetTimestamp() time.Time {
//	return i.timestamp
//}

func (i *DCDimmerCommand2) SetDGN(dgn uint32) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.dgn = dgn
	i.name = DGNName(dgn)
}

func (i *DCDimmerCommand2) SetSourceAddress(sa uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.sourceAddress = sa
}

func (i *DCDimmerCommand2) SetPriority(p uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.priority = p
}

func (i *DCDimmerCommand2) SetInstance(inst uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.deviceInstance = inst
}
func (i *DCDimmerCommand2) SetGroup(g uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.group = g
}
func (i *DCDimmerCommand2) SetDesiredBringhtness(b float64) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.desiredBrightness = b
}
func (i *DCDimmerCommand2) SetCommand(cmd uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.command = cmd
}
func (i *DCDimmerCommand2) GetInstance() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.deviceInstance
}

func (i *DCDimmerCommand2) GetGroup() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.group
}

func (i *DCDimmerCommand2) GetDesiredBrightness() float64 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.desiredBrightness
}

func (i *DCDimmerCommand2) GetCommand() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.command
}
func (i *DCDimmerCommand2) GetInterlockStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.interlockStatus
}

func (i *DCDimmerCommand2) SetInterlockStatus(s constants.Uint2) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.interlockStatus = s
}

func (i *DCDimmerCommand2) SetReserved1(s constants.Uint2) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.reserved1 = s
}

func (i *DCDimmerCommand2) SetReserved2(s constants.Uint2) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.reserved2 = s
}

func (i *DCDimmerCommand2) SetReserved3(s constants.Uint2) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.reserved3 = s
}

func (i *DCDimmerCommand2) GetDelayDuration() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.delayDuration
}
func (i *DCDimmerCommand2) SetDelayDuration(d uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.delayDuration = d
}
func (i *DCDimmerCommand2) GetRampTime() float64 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.rampTime
}

func (i *DCDimmerCommand2) SetRampTime(rt float64) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.rampTime = rt
}

func (i *DCDimmerCommand2) SetReserved(r uint8) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.reserved = r
}
func (i *DCDimmerCommand2) GetReserved() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved
}

func (i *DCDimmerCommand2) GetReserved1() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved1
}

func (i *DCDimmerCommand2) GetReserved2() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved2
}
func (i *DCDimmerCommand2) GetReserved3() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved3
}

func (i *DCDimmerCommand2) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.GetDGN(), Instance: i.deviceInstance}
	var n, ok = DGNInstanceNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}
func (i *DCDimmerCommand2) String() string {
	//var s = i.RvcItem.String()

	// deviceInstance    byte    // 0
	//	group             uint8   // 1
	//	desiredBrightness float64 // 2 (%)
	//	command           uint8   // 3
	//	delayDuration     uint8   // 4 seconds
	//
	//	interlockStatus uint2   // 5 (0,1)
	//	reserved1       uint2   // 5 (2-3)
	//	reserved2       uint2   // 5 (4-5)
	//	reserved3       uint2   // 5 (6-7)
	//	rampTime        float64 // seconds (0.1) 0-25 seconds
	//	reserved        uint8   //

	return fmt.Sprintf("TS: %s DGN: %x (%s) SA: %d Instance: %d (%s) group: %d brigntness: %f command: %d lockitem: %d res1: %d res2: %d res3: %d rampTime %f reserved %d",
		i.GetTimestamp().Format("01-02-2006 15:04:05.000000"),
		i.GetDGN(), DGNName(i.GetDGN()),
		i.GetSourceAddress(),
		i.deviceInstance,
		i.GetInstanceName(),
		i.group, i.desiredBrightness, i.command,
		i.interlockStatus,
		i.reserved1,
		i.reserved2,
		i.reserved3,
		i.rampTime,
		i.reserved,
	)
	//return "ZZZXXX" + s + " " + t

}

func (r *DCDimmerCommand2) GetSourceAddress() uint8 {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.sourceAddress
}
func (r *DCDimmerCommand2) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
	//r.timestamp = (*from).GetTimeStamp()
	//r.sourceAddress = (*from).GetSourceAddress()
	r.RvcItem.Init(from)

	// TODO some test for all these bits....
	var dataBytes = &from.Data
	var changed = false

	//
	// test for bug - did instance change?
	{
		var tmp = utils.GetByte(dataBytes, 0)
		if (*r).deviceInstance != tmp {
			changed = true
		}
		(*r).deviceInstance = tmp

		tmp = utils.GetByte(dataBytes, 1)
		if tmp != (*r).group {
			changed = true
		}
		(*r).group = tmp
	}
	{

		var tmp = convert.ToPercent(utils.GetByte(dataBytes, 2))
		if tmp != (*r).desiredBrightness {
			changed = true
		}
		(*r).desiredBrightness = tmp
	}
	{

		var tmp = utils.GetByte(dataBytes, 3)
		if tmp != (*r).command {
			changed = true
		}
		(*r).command = tmp

		tmp = utils.GetByte(dataBytes, 4)
		if tmp != (*r).delayDuration {
			changed = true
		}
		(*r).delayDuration = tmp
	}

	{
		var tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 0, utils.GetMask(2)))
		if (*r).interlockStatus != tmp {
			changed = true
		}
		(*r).interlockStatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 2, utils.GetMask(2)))
		if (*r).reserved1 != tmp {
			changed = true
		}
		(*r).reserved1 = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 4, utils.GetMask(2)))
		if (*r).reserved2 != tmp {
			changed = true
		}
		(*r).reserved2 = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 6, utils.GetMask(2)))
		if (*r).reserved3 != tmp {
			changed = true
		}
		(*r).reserved3 = tmp
	}
	{
		var tmp = float64(utils.GetByte(dataBytes, 6)) / 10
		if (*r).rampTime != tmp {
			changed = true
		}
		(*r).rampTime = tmp
	}

	{
		var tmp = utils.GetByte(dataBytes, 7)
		if (*r).reserved != tmp {
			changed = true
		}
		(*r).reserved = tmp
	}

	if changed {
		r.lastChanged = r.timestamp
	}

}

// CreateFrame - create the data frame to be send in order to issue the command
func (r *DCDimmerCommand2) CreateFrame() *RvcFrame {
	var ret = RvcFrame{}
	r.lock.Lock()
	defer r.lock.Unlock()

	// TODO we really should not have to do this here !
	ret.SetEFF_RTR_ERR_Flag(constants.CAN_EFF_FLAG2)
	utils.SetByte(&ret.Data, 0, r.deviceInstance)
	utils.SetByte(&ret.Data, 1, r.group)
	utils.SetByte(&ret.Data, 2, convert.FromPercent(r.desiredBrightness))
	utils.SetByte(&ret.Data, 3, r.command)

	utils.SetByte(&ret.Data, 4, r.delayDuration)

	utils.SetBits(&ret.Data, 5, 0, utils.GetMask(2), r.interlockStatus)
	utils.SetBits(&ret.Data, 5, 2, utils.GetMask(2), r.reserved1)
	utils.SetBits(&ret.Data, 5, 4, utils.GetMask(2), r.reserved2)
	utils.SetBits(&ret.Data, 5, 6, utils.GetMask(2), r.reserved3)
	{
		var tmp = r.rampTime * 10
		if tmp > 255 {
			tmp = 255
		} else if tmp < 0 {
			tmp = 0
		}
		utils.SetByte(&ret.Data, 6, uint8(tmp))
	}
	utils.SetByte(&ret.Data, 7, r.reserved)

	ret.PGN.SetPGNValue(r.dgn)
	ret.SetPriority(0x06)
	ret.Length = 8                        //
	ret.SetSourceAddress(r.sourceAddress) // Always 8 bits
	ret.SetCanMessage()
	return &ret
}
