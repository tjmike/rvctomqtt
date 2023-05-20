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

// TODO we need some means to SET and push state.... this lets us see the commands
// 0x1FFE0

// 0x1FFE1
type inverterTemperatureStatus struct {
	RvcItem
	instance        byte
	fet1Temp        float64 //
	transformerTemp float64 //
	fet2Temp        float64 //
}

//func (r *inverterTemperatureStatus) GetInstanceKey() interface{} {
//	return DGNInstanceKey{r.dgn, r.instance}
//}

func (i *inverterTemperatureStatus) String() string {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return fmt.Sprintf("ivts DGN: %x(%s) SA: %d Instance: %d fet1Temp(C/f): %f/%f transformerTemp(C/f): %f/%f fet2Temp(Cf): %f/%f",
		i.dgn, i.GetName(),
		i.GetSourceAddress(),
		i.GetInstance(),
		i.fet1Temp,
		convert.CelsiusToFahrenheit(i.fet1Temp),
		i.transformerTemp,
		convert.CelsiusToFahrenheit(i.transformerTemp),
		i.fet2Temp,
		convert.CelsiusToFahrenheit(i.fet2Temp),
	)
}

func (r *inverterTemperatureStatus) GetInstance() byte {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.instance
}
func (r *inverterTemperatureStatus) GetFet1Temp() float64 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.fet1Temp
}
func (r *inverterTemperatureStatus) GetTransformerTemp() float64 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.transformerTemp
}
func (r *inverterTemperatureStatus) GetFet2Temp() float64 {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.fet2Temp
}

func (r *inverterTemperatureStatus) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.RvcItem.Init(from)
	var changed = false
	var dataBytes = &from.Data

	{
		var tmp = utils.GetByte(dataBytes, 0)
		if (*r).instance != tmp {
			changed = true
		}
		(*r).instance = tmp
	}
	{
		var tmp = convert.ToDegreesC16(utils.Getuint16(dataBytes, 1))
		if tmp != (*r).fet1Temp {
			changed = true
		}
		(*r).fet1Temp = tmp

		tmp = convert.ToDegreesC16(utils.Getuint16(dataBytes, 3))
		if tmp != (*r).transformerTemp {
			changed = true
		}
		(*r).transformerTemp = tmp

		tmp = convert.ToDegreesC16(utils.Getuint16(dataBytes, 5))
		if tmp != (*r).fet2Temp {
			changed = true
		}
		(*r).fet2Temp = tmp

	}
	if changed {
		(*r).lastChanged = (*r).timestamp
	}
}
