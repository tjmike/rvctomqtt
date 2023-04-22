package rvc

import (
	"fmt"
	"rvctomqtt/utils"
)

// tankStatus - 0x1FFB7
type tankStatus struct {
	RvcItem
	instance      uint8
	relativeLevel uint8 // Level = Relative Level / Resolution
	resolution    uint8
	absoluteLevel uint16 // liter Value range = 0 to 65 530 l
	tankSize      uint16 // liter Value range = 0 to 65 530 l
}

func (ts *tankStatus) GetInstance() uint8 {
	return ts.instance
}

func (ts *tankStatus) GetRelativeLevel() uint8 {
	return ts.relativeLevel
}

// TODO - interesting problem is is not a "field" but should be listed in the fields list
func (ts *tankStatus) GetLevel() float64 {
	return 100.0 * float64(ts.relativeLevel) / float64(ts.resolution)
}
func (ts *tankStatus) GetAbsoluteLevel() uint16 {
	return ts.absoluteLevel
}
func (ts *tankStatus) GetTankSize() uint16 {
	return ts.tankSize
}

func (r *tankStatus) Init(from *RvcFrame) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	r.RvcItem.Init(from)

	var changed = false
	// TODO some test for all these bits....
	var dataBytes = &from.Data
	(*r).instance = utils.GetByte(dataBytes, 0)

	{
		var tmp = utils.GetByte(dataBytes, 1)
		if (*r).relativeLevel != tmp {
			changed = true
		}
		(*r).relativeLevel = tmp

		tmp = utils.GetByte(dataBytes, 2)
		if (*r).resolution != tmp {
			changed = true
		}
		(*r).resolution = utils.GetByte(dataBytes, 2)

	}
	{
		var tmp = utils.Getuint16(dataBytes, 3)
		if (*r).absoluteLevel != tmp {
			changed = true
		}
		(*r).absoluteLevel = tmp

		tmp = utils.Getuint16(dataBytes, 5)
		if (*r).tankSize != tmp {
			changed = true
		}
		(*r).tankSize = tmp
	}
	if changed {
		(*r).lastChanged = (*r).timestamp
	}
}

func (i *tankStatus) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.dgn, Instance: i.instance}
	var n, ok = DGNInstanceNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}

func (i *tankStatus) String() string {
	//var s = i.RvcItem.String()
	var iname = i.GetInstanceName()

	return fmt.Sprintf("DGN: %x (%s) Instance: %d (%s) level %f relativeLevel: %d resolution: %d absoluteLevel: %d tanSize: %d ",
		i.dgn, i.name,
		i.instance, iname,
		i.GetLevel(),
		i.relativeLevel, i.resolution, i.absoluteLevel, i.tankSize)

}
