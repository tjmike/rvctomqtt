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
//fields []string = {"Instance", "priority", "voltage", "current"}
//)
// The idea - we have this file called DC source status - the Instance gets updated in real time
// so there's a single Instance that exists and represents the current state at any given point in
// time. We don't keep a handle to the RVC frame because we want to return it

type DCSourceStatus1 struct {
	RvcItem
	deviceInstance byte
	devicePriority uint8
	voltage        float64
	current        float64
}

var myFields = []dataField{
	{name: instance, fieldType: U8},
	{name: PRIORITY, fieldType: U8},
	{name: VOLTAGE, fieldType: F64},
	{name: CURRENT, fieldType: F64},
}

func (i *DCSourceStatus1) String() string {
	//var s = i.RvcItem.String()

	return fmt.Sprintf("DGN: %x(%s) Instance: %d priority: %d voltage: %f current: %f",
		i.dgn, i.GetName(), i.deviceInstance, i.devicePriority, i.voltage, i.current)
	//return "ZZZXXX" + s + " " + t

}

func (r *DCSourceStatus1) getInstance() byte {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.deviceInstance
}

/*
func (r *DCSourceStatus1) FieldUint8(f dataField) uint8 {
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

func (r *DCSourceStatus1) Fields() *[]dataField {
	return &myFields
}

func (r *DCSourceStatus1) FieldUint16(f dataField) uint16 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return 0
}
func (r *DCSourceStatus1) FieldUint32(f dataField) uint32 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return 0
}
func (r *DCSourceStatus1) FieldUint8(f dataField) uint8 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	//var ret = r.RvcItem.FieldUint8(f)
	//if ret == 0 {
	//	return ret
	//}
	switch f {
	case myFields[0]:
		return r.deviceInstance
		break
	case myFields[1]:
		return r.devicePriority
		break
	}
	return 0 // need to fix all these to be spec compliant (255?)

}
func (r *DCSourceStatus1) FieldFloat64(f dataField) float64 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	switch f {
	case myFields[2]:
		return r.voltage
		break
	case myFields[3]:
		return r.current
		break
	}
	return 0
}

func (r *DCSourceStatus1) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
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
		(*r).lastChanged = (*r).timestamp
	}
}
