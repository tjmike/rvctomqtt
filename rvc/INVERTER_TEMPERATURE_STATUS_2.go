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

// TODO we need some means to SET and push state.... this lets us see the commands
// 0x1FFE0

// 0x1FFE1
type inverterTemperatureStatus2 struct {
	RvcItem
	instance         byte
	controlBoardTemp float64 //
	capacitorTemp    float64 //
	ambientTemp      float64 //
}

//func (r *inverterTemperatureStatus2) GetInstanceKey() interface{} {
//	return DGNInstanceKey{r.DGN, r.instance}
//}

func (i *inverterTemperatureStatus2) String() string {
	return fmt.Sprintf("DGN: %x(%s) SA: %d Instance: %d controlBoardTemp(C/f): %f/%f capTemp(C/f): %f/%f ambientTemp(Cf): %f/%f",
		i.DGN, i.GetName(),
		i.GetSourceAddress(),
		i.GetInstance(),
		i.controlBoardTemp,
		convert.CelsiusToFahrenheit(i.controlBoardTemp),
		i.capacitorTemp,
		convert.CelsiusToFahrenheit(i.capacitorTemp),
		i.ambientTemp,
		convert.CelsiusToFahrenheit(i.ambientTemp),
	)
}

func (r *inverterTemperatureStatus2) GetInstance() byte {
	return r.instance
}
func (r *inverterTemperatureStatus2) GetCapacitorTemperature() float64 {
	return r.capacitorTemp
}
func (r *inverterTemperatureStatus2) GetAmbientTemp() float64 {
	return r.ambientTemp
}
func (r *inverterTemperatureStatus2) Init(from *RvcFrame) {
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
		if tmp != (*r).controlBoardTemp {
			changed = true
		}
		(*r).controlBoardTemp = tmp

		tmp = convert.ToDegreesC16(utils.Getuint16(dataBytes, 3))
		if tmp != (*r).capacitorTemp {
			changed = true
		}
		(*r).capacitorTemp = tmp

		tmp = convert.ToDegreesC16(utils.Getuint16(dataBytes, 5))
		if tmp != (*r).ambientTemp {
			changed = true
		}
		(*r).ambientTemp = tmp

	}
	if changed {
		(*r).LastChanged = (*r).Timestamp
	}
}
