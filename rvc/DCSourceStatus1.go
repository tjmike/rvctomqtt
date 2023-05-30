package rvc

import (
	"fmt"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
)

//type fieldName string
//const (
//	Instance string = "123"
//df dataField =
//})
//fields []string = {"Instance", "Priority", "voltage", "current"}
//)
// The idea - we have this file called DC source status - the Instance gets updated in real time
// so there's a single Instance that exists and represents the current state at any given point in
// time. We don't keep a handle to the RVC frame because we want to return it

type DCSourceStatus1Spyder struct {
	DCSourceStatus1
}
type DCSourceStatus1 struct {
	RvcItem
	deviceInstance byte
	devicePriority uint8
	voltage        float64
	current        float64
}

func (i *DCSourceStatus1) GetVoltage() float64 {
	return i.voltage
}

func (i *DCSourceStatus1) GetCurrent() float64 {
	return i.current
}

//func (r *DCSourceStatus1) GetInstanceKey() interface{} {
//	r.lock.RLock()
//	defer r.lock.RUnlock()
//	return DGNInstanceKey{
//		r.DGN,
//		r.deviceInstance,
//	}
//}

func (i *DCSourceStatus1) String() string {
	//var s = i.RvcItem.String()

	return fmt.Sprintf("dcssDGN: %x(%s) SA: %d Instance: %d Priority: %d voltage: %f current: %f",
		i.DGN, i.GetName(), i.GetSourceAddress(), i.deviceInstance, i.devicePriority, i.voltage, i.current)
	//return "ZZZXXX" + s + " " + t

}

func (r *DCSourceStatus1) GetInstance() byte {
	return r.deviceInstance
}

func (i *DCSourceStatus1) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.DGN, Instance: i.deviceInstance}
	var n, ok = GetInstanceName(k)
	if ok {
		return n
	} else {
		return ""
	}
}

/*
func (r *DCSourceStatus1) FieldUint8(f dataField) uint8 {
	r.lockitem.RLock()
	defer r.lockitem.RUnlock()
	switch f.Name {
	case myFields[0].Name:
		return r.deviceInstance
		break
	case myFields[1].Name:
		return r.devicePriority
		break
	}
	return NAuint8
}
*/

func (r *DCSourceStatus1) Init(from *RvcFrame) {
	r.RvcItem.Init(from)
	var changed = false
	var dataBytes = &from.Data

	{
		var tmp = utils.GetByte(dataBytes, 0)
		if (*r).deviceInstance != tmp {
			changed = true
		}
		(*r).deviceInstance = tmp

		tmp = utils.GetByte(dataBytes, 1)
		if (*r).devicePriority != tmp {
			changed = true
		}
		(*r).devicePriority = tmp
	}
	{
		var tmp = convert.ToVolts(utils.Getuint16(dataBytes, 2))
		if tmp != (*r).voltage {
			changed = true
		}
		(*r).voltage = tmp

		tmp = convert.ToCurrent(utils.Getuint32(dataBytes, 4))
		if tmp != (*r).current {
			changed = true
		}
		(*r).current = tmp
	}
	if changed {
		(*r).LastChanged = (*r).Timestamp
	}
}
