package rvc

import (
	"fmt"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
)

type uint2 uint8

// 0x1FEDA
type DCDimmerStatus3 struct {
	RvcItem
	deviceInstance byte    // 0
	group          uint8   // 1
	brightness     float64 // 2
	// Consider keeping this a byte and then just having getters to fetch the correct bits
	lockitem          uint2 // 3 (0,1)
	overCurrentStatus uint2 // 3 (2,3)
	overrideSatus     uint2 // 3 (4,5)
	enableStatus      uint2 // 3 (6,7)
	delayDuration     uint8 // 4
	lastCommand       uint8 // 5
	interlockStatus   uint2 // 6 (0,1)
	loadStatus        uint2 //6 (2-3)
	reserved          uint2 // 6 (4-5)
	undercurrent      uint2 // 6 (6-7)
	masterMemoryVal   uint8 // 7
}

// TODO - this may be better as a map?
//      - this seems too combersome to do over and over is there a better way to express fields

var myFieldsDCDimmerStatus3 = []dataField{
	{name: instance, fieldType: U8},            // 0
	{name: group, fieldType: U8},               // 1
	{name: brightness, fieldType: F64},         // 2 251 = Value is changing (ramp command) 252 = Output is Flashing
	{name: lock, fieldType: BIT2},              // 3
	{name: overCurrentStatus, fieldType: BIT2}, // 4
	{name: enableStatus, fieldType: BIT2},      // 5
	{name: delayDuration, fieldType: BIT2},     // 6
	{name: lastCommand, fieldType: U8},         // 7
	{name: interlockStatus, fieldType: BIT2},   // 8
	{name: loadStatus, fieldType: BIT2},        // 9
	{name: reserved, fieldType: BIT2},          // 10
	{name: undercurrent, fieldType: BIT2},      // 11
	{name: masterMemoryValue, fieldType: U8},   // 12
}

// interlockStatus   uint2 // 6 (0,1)
//	loadStatus        uint2 //6 (2-3)
//	reserved          uint2 // 6 (4-5)
//	undercurrent      uint2 // 6 (6-7)

func (i *DCDimmerStatus3) getInterlockStatus() uint2 {
	return i.interlockStatus
}

func (i *DCDimmerStatus3) getLoadStatus() uint2 {
	return i.loadStatus
}

func (i *DCDimmerStatus3) getReserved() uint2 {
	return i.reserved
}
func (i *DCDimmerStatus3) getUnderCurrent() uint2 {
	return i.undercurrent
}
func (i *DCDimmerStatus3) GetLastCommand() uint8 {
	return i.lastCommand
}
func (i *DCDimmerStatus3) GetDelayDuration() uint8 {
	return i.delayDuration
}
func (i *DCDimmerStatus3) GetEnableStatus() uint2 {
	return i.enableStatus
}
func (i *DCDimmerStatus3) GetOverrideStatus() uint2 {
	return i.overrideSatus
}
func (i *DCDimmerStatus3) GetOverCurrentStatus() uint2 {
	return i.overCurrentStatus
}
func (i *DCDimmerStatus3) GetLockItem() uint2 {
	return i.lockitem
}

func (i *DCDimmerStatus3) GetBrightness() float64 {
	return i.brightness
}

func (i *DCDimmerStatus3) GetGroup() uint8 {
	return i.group
}
func (i *DCDimmerStatus3) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.dgn, Instance: i.deviceInstance}
	var n, ok = DGNInstanceNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}
func (i *DCDimmerStatus3) String() string {
	//var s = i.RvcItem.String()
	var iname = i.GetInstanceName()
	return fmt.Sprintf("DGN: %x (%s) Instance: %d (%s) group: %d brigntness: %f lockitem: %d overcurrent: %d enable: %d delayDuration: %d last %d interlock %d, load status %d reserved: %d undercurrent: %d memval: %d",
		i.dgn, i.name,
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
			fmt.Printf("INSTANCE CHANGED!!! %d != %d\n", tmp, (*r).deviceInstance)
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
		var tmp = uint2(utils.GetBits(dataBytes, 3, 0, utils.GetMask(2)))
		if (*r).lockitem != tmp {
			changed = true
		}
		(*r).lockitem = tmp

		tmp = uint2(utils.GetBits(dataBytes, 3, 2, utils.GetMask(2)))
		if (*r).overCurrentStatus != tmp {
			changed = true
		}
		(*r).overCurrentStatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 3, 4, utils.GetMask(2)))
		if (*r).overrideSatus != tmp {
			changed = true
		}
		(*r).overrideSatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 3, 6, utils.GetMask(2)))
		if (*r).enableStatus != tmp {
			changed = true
		}
		(*r).enableStatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 6, 0, utils.GetMask(2)))
		if (*r).interlockStatus != tmp {
			changed = true
		}
		(*r).interlockStatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 6, 2, utils.GetMask(2)))
		if (*r).loadStatus != tmp {
			changed = true
		}
		(*r).loadStatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 6, 4, utils.GetMask(2)))
		if (*r).reserved != tmp {
			changed = true
		}
		(*r).reserved = tmp

		tmp = uint2(utils.GetBits(dataBytes, 6, 6, utils.GetMask(2)))
		if (*r).undercurrent != tmp {
			changed = true
		}
		(*r).undercurrent = tmp
	}

	if changed {
		r.lastChanged = r.timestamp
	}

	//(*r).group = utils.GetByte(dataBytes, 1)
	//(*r).brightness = convert.ToPercent(utils.GetByte(dataBytes, 2))
	//(*r).lockitem = uint2(utils.GetBits(dataBytes, 3, 0, utils.GetMask(2)))
	//(*r).overCurrentStatus = uint2(utils.GetBits(dataBytes, 3, 2, utils.GetMask(2)))
	//(*r).overrideSatus = uint2(utils.GetBits(dataBytes, 3, 4, utils.GetMask(2)))
	//(*r).enableStatus = uint2(utils.GetBits(dataBytes, 3, 6, utils.GetMask(2)))
	//(*r).delayDuration = utils.GetByte(dataBytes, 4)
	//(*r).lastCommand = utils.GetByte(dataBytes, 5)
	//(*r).interlockStatus = uint2(utils.GetBits(dataBytes, 6, 0, utils.GetMask(2)))
	//(*r).loadStatus = uint2(utils.GetBits(dataBytes, 6, 2, utils.GetMask(2)))
	//(*r).reserved = uint2(utils.GetBits(dataBytes, 6, 4, utils.GetMask(2)))
	//(*r).undercurrent = uint2(utils.GetBits(dataBytes, 6, 6, utils.GetMask(2)))
	//(*r).masterMemoryVal = utils.GetByte(dataBytes, 7)

}
