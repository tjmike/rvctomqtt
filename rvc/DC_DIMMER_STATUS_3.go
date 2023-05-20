package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
)

// todo - move to a types file

// 0x1FEDA
type DCDimmerStatus3 struct {
	RvcItem
	deviceInstance byte    // 0
	group          uint8   // 1
	brightness     float64 // 2
	// Consider keeping this a byte and then just having getters to fetch the correct bits
	lockitem          constants.Uint2 // 3 (0,1)
	overCurrentStatus constants.Uint2 // 3 (2,3)
	overrideSatus     constants.Uint2 // 3 (4,5)
	enableStatus      constants.Uint2 // 3 (6,7)
	delayDuration     uint8           // 4
	lastCommand       uint8           // 5
	interlockStatus   constants.Uint2 // 6 (0,1)
	loadStatus        constants.Uint2 //6 (2-3)
	reserved          constants.Uint2 // 6 (4-5)
	undercurrent      constants.Uint2 // 6 (6-7)
	masterMemoryVal   uint8           // 7
}

//	func (r *DCDimmerStatus3) GetInstanceKey() interface{} {
//		r.lock.RLock()
//		defer r.lock.RUnlock()
//		return DGNInstanceKey{
//			r.dgn,
//			r.deviceInstance,
//		}
//	}
func (i *DCDimmerStatus3) getInterlockStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.interlockStatus
}

func (i *DCDimmerStatus3) getLoadStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.loadStatus
}

func (i *DCDimmerStatus3) getReserved() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved
}
func (i *DCDimmerStatus3) getUnderCurrent() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.undercurrent
}
func (i *DCDimmerStatus3) GetLastCommand() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.lastCommand
}
func (i *DCDimmerStatus3) GetDelayDuration() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.delayDuration
}
func (i *DCDimmerStatus3) GetEnableStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.enableStatus
}
func (i *DCDimmerStatus3) GetOverrideStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.overrideSatus
}
func (i *DCDimmerStatus3) GetOverCurrentStatus() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.overCurrentStatus
}
func (i *DCDimmerStatus3) GetLockItem() constants.Uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.lockitem
}

func (i *DCDimmerStatus3) GetBrightness() float64 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.brightness
}

func (i *DCDimmerStatus3) GetGroup() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.group
}
func (i *DCDimmerStatus3) GetInstanceName() string {
	i.lock.RLock()
	defer i.lock.RUnlock()
	var k = DGNInstanceKey{DGN: i.dgn, Instance: i.deviceInstance}
	var n, ok = DGNInstanceNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}
func (i *DCDimmerStatus3) String() string {
	i.lock.RLock()
	defer i.lock.RUnlock()
	//var s = i.RvcItem.String()
	var iname = i.GetInstanceName()
	return fmt.Sprintf("%s DGN: %x (%s) SA: %x Instance: %d (%s) group: %d brigntness: %f lockitem: %d overcurrent: %d enable: %d delayDuration: %d last %d interlock %d, load status %d reserved: %d undercurrent: %d memval: %d",
		i.GetTimestamp().Format("01-02-2006 15:04:05.000000"),
		i.dgn, i.name,
		i.GetSourceAddress(),
		i.deviceInstance, iname, i.group, i.brightness, i.lockitem, i.overrideSatus, i.enableStatus, i.delayDuration,
		i.lastCommand, i.interlockStatus, i.loadStatus, i.reserved, i.undercurrent, i.masterMemoryVal,
	)
	//return "ZZZXXX" + s + " " + t

}

func (r *DCDimmerStatus3) GetInstance() byte {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.deviceInstance
}

func (r *DCDimmerStatus3) Init(from *RvcFrame) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	r.RvcItem.Init(from)

	// TODO some test for all these bits....
	var dataBytes = &from.Data

	//
	// test for bug - did instance change?
	{
		var tmp = utils.GetByte(dataBytes, 0)
		if (*r).deviceInstance != tmp {
			fmt.Printf("DCDimmerStatus3 INSTANCE CHANGED!!! %d != %d\n", tmp, (*r).deviceInstance)
		}
		(*r).deviceInstance = utils.GetByte(dataBytes, 0)
	}
	// test for changed and if changed set last changed timestamp
	var changed = false
	{
		var tmp = utils.GetByte(dataBytes, 1)
		if tmp != (*r).group {
			changed = true
		}
		(*r).group = tmp

		tmp = utils.GetByte(dataBytes, 4)
		if tmp != (*r).delayDuration {
			changed = true
		}
		(*r).delayDuration = tmp

		tmp = utils.GetByte(dataBytes, 5)
		if tmp != (*r).lastCommand {
			changed = true
		}
		(*r).lastCommand = tmp

		tmp = utils.GetByte(dataBytes, 7)
		if tmp != (*r).masterMemoryVal {
			changed = true
		}
		(*r).masterMemoryVal = tmp
	}
	{
		var tmp = convert.ToPercent(utils.GetByte(dataBytes, 2))
		if tmp != (*r).brightness {
			changed = true
		}
		(*r).brightness = tmp
	}

	{
		var tmp = constants.Uint2(utils.GetBits(dataBytes, 3, 0, utils.GetMask(2)))
		if (*r).lockitem != tmp {
			changed = true
		}
		(*r).lockitem = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 3, 2, utils.GetMask(2)))
		if (*r).overCurrentStatus != tmp {
			changed = true
		}
		(*r).overCurrentStatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 3, 4, utils.GetMask(2)))
		if (*r).overrideSatus != tmp {
			changed = true
		}
		(*r).overrideSatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 3, 6, utils.GetMask(2)))
		if (*r).enableStatus != tmp {
			changed = true
		}
		(*r).enableStatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 6, 0, utils.GetMask(2)))
		if (*r).interlockStatus != tmp {
			changed = true
		}
		(*r).interlockStatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 6, 2, utils.GetMask(2)))
		if (*r).loadStatus != tmp {
			changed = true
		}
		(*r).loadStatus = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 6, 4, utils.GetMask(2)))
		if (*r).reserved != tmp {
			changed = true
		}
		(*r).reserved = tmp

		tmp = constants.Uint2(utils.GetBits(dataBytes, 6, 6, utils.GetMask(2)))
		if (*r).undercurrent != tmp {
			changed = true
		}
		(*r).undercurrent = tmp
	}

	if changed {
		r.lastChanged = r.timestamp
	}

}
