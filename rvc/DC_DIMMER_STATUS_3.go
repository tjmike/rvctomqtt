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
	{name: INSTANCE, fieldType: U8},              // 0
	{name: GROUP, fieldType: U8},                 // 1
	{name: BRIGHTNESS, fieldType: F64},           // 2
	{name: LOCK, fieldType: BIT2},                // 3
	{name: OVER_CURRENT_STATUS, fieldType: BIT2}, // 4
	{name: ENABLE_STATUS, fieldType: BIT2},       // 5
	{name: DELAY_DURATION, fieldType: BIT2},      // 6
	{name: LAST_COMMAND, fieldType: U8},          // 7
	{name: INTERLOCK_STATUS, fieldType: BIT2},    // 8
	{name: LOAD_STATUS, fieldType: BIT2},         // 9
	{name: RESERVED, fieldType: BIT2},            // 10
	{name: UNDERCURRENT, fieldType: BIT2},        // 11
	{name: MASTERMEMVAL, fieldType: U8},          // 12
}

func (i *DCDimmerStatus3) getInstanceName() string {
	var k = DGNInstanceKey{DGN: i.dgn, Instance: i.deviceInstance}
	var n, ok = DGNNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}
func (i *DCDimmerStatus3) String() string {
	var s = i.RvcItem.String()
	var iname = i.getInstanceName()
	var t = fmt.Sprintf("Instance: %d (%s) group: %d brigntness: %f lockitem: %d overcurrent: %d enable: %d delayDuration: %d last %d interlock %d, load status %d reserved: %d undercurrent: %d memval: %d",

		i.deviceInstance, iname, i.group, i.brightness, i.lockitem, i.overrideSatus, i.enableStatus, i.delayDuration,
		i.lastCommand, i.interlockStatus, i.loadStatus, i.reserved, i.undercurrent, i.masterMemoryVal,
	)
	return "ZZZXXX" + s + " " + t

}

func (r *DCDimmerStatus3) getInstance() byte {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.deviceInstance
}

/*
func (r *DCSourceStatus1) GetFieldUint8(f dataField) uint8 {
	r.lockitem.RLock()
	defer r.lockitem.RUnlock()
	switch f.name {
	case myFields[0].name:
		return r.deviceInstance
		break
	case myFields[1].name:
		return r.devicePriority
		break
	}
	return NAuint8
}
*/

func (r *DCDimmerStatus3) GetFields() *[]dataField {
	return &myFieldsDCDimmerStatus3
}

func (r *DCDimmerStatus3) GetFieldUint16(f dataField) uint16 {
	r.RvcItem.lock.RLock()
	defer r.RvcItem.lock.RUnlock()
	return 0
}

func (r *DCDimmerStatus3) GetFieldUint32(f dataField) uint32 {
	r.RvcItem.lock.RLock()
	defer r.RvcItem.lock.RUnlock()
	return 0
}

func (r *DCDimmerStatus3) GetFieldUint8(f dataField) uint8 {
	r.RvcItem.lock.RLock()
	defer r.RvcItem.lock.RUnlock() //var ret = r.RvcItem.GetFieldUint8(f)
	//if ret == 0 {
	//	return ret
	//}
	switch f {
	case myFields[0]: // TODO - no good manual index like this is prone to error
		return r.deviceInstance
	case myFields[1]:
		return r.group
	}
	return 0 // need to fix all these to be spec compliant (255?)

}
func (r *DCDimmerStatus3) GetFieldFloat64(f dataField) float64 {

	r.lock.RLock()
	defer r.lock.RUnlock()
	switch f {
	case myFields[2]:
		return r.brightness
		break
	}
	return 0
}

func (r *DCDimmerStatus3) Init(from *RvcFrame) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	r.RvcItem.Init(from)

	// TODO some test for all these bits....
	var dataBytes = &from.Data

	(*r).deviceInstance = utils.GetByte(dataBytes, 0)

	(*r).group = utils.GetByte(dataBytes, 1)

	(*r).brightness = convert.ToPercent(utils.GetByte(dataBytes, 2))
	(*r).lockitem = uint2(utils.GetBits(dataBytes, 3, 0, utils.GetMask(2)))
	(*r).overCurrentStatus = uint2(utils.GetBits(dataBytes, 3, 2, utils.GetMask(2)))
	(*r).overrideSatus = uint2(utils.GetBits(dataBytes, 3, 4, utils.GetMask(2)))
	(*r).enableStatus = uint2(utils.GetBits(dataBytes, 3, 6, utils.GetMask(2)))

	(*r).delayDuration = utils.GetByte(dataBytes, 4)

	(*r).lastCommand = utils.GetByte(dataBytes, 5)

	(*r).interlockStatus = uint2(utils.GetBits(dataBytes, 6, 0, utils.GetMask(2)))
	(*r).loadStatus = uint2(utils.GetBits(dataBytes, 6, 2, utils.GetMask(2)))
	(*r).reserved = uint2(utils.GetBits(dataBytes, 6, 4, utils.GetMask(2)))
	(*r).undercurrent = uint2(utils.GetBits(dataBytes, 6, 6, utils.GetMask(2)))
	(*r).masterMemoryVal = utils.GetByte(dataBytes, 7)
}
