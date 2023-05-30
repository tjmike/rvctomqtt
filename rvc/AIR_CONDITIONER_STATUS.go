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
type airConditionerCommand airConditionerStatus

// 0x1FFE1
type airConditionerStatus struct {
	RvcItem
	instance            byte
	operatingMode       uint8   //0 automatic, manual
	maxFanSpeed         float64 // %
	maxOutputLevel      float64 // %
	fanSpeed            float64 // %
	acOutputLevel       float64 // %
	deadBand            float64 // Precision = 0.1 °C This is the amount over and under the set point that the AC will tolerate. A larger value reducescycling.
	secondStageDeadBand float64 // Value range = 0 to 25.0 °C  Precision = 0.1 °C  This is the amount over the set point that will trigger a higher A/C output
}

//func (r *airConditionerStatus) GetInstanceKey() interface{} {
//	r.lock.RLock()
//	defer r.lock.RUnlock()
//	return DGNInstanceKey{
//		r.DGN,
//		r.instance,
//	}
//}

func (i *airConditionerStatus) String() string {
	return fmt.Sprintf("acss DGN: %x(%s) SA: %d Instance: %d operatingModel: %d maxFanSpeed: %f maxOutputLevel: %f fanSpeed: %f acOutputLevel: %f deadBand %f, 2ndStagDeadBand2 %f",
		i.DGN, i.GetName(),
		i.GetSourceAddress(),
		i.instance, i.operatingMode, i.maxFanSpeed, i.maxOutputLevel, i.fanSpeed,
		i.acOutputLevel, i.deadBand, i.secondStageDeadBand)
}

func (r *airConditionerStatus) GetInstance() byte {
	return r.instance
}
func (i *airConditionerStatus) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.GetDGN(), Instance: i.instance}
	var n, ok = DGNInstanceNames[k]
	if ok {
		return n
	} else {
		return ""
	}
}

func (r *airConditionerStatus) GetOperationgMode() uint8 {
	return r.operatingMode
}
func (r *airConditionerStatus) GetMaxFanSpeed() float64 {
	return r.maxFanSpeed
}
func (r *airConditionerStatus) GetMaxOutputLevel() float64 {
	return r.maxOutputLevel
}

func (r *airConditionerStatus) GetFanSpeed() float64 {
	return r.fanSpeed
}
func (r *airConditionerStatus) GetACOutputLevel() float64 {
	return r.acOutputLevel
}

func (r *airConditionerStatus) GetDeadBand() float64 {
	return r.deadBand
}

func (r *airConditionerStatus) GetSecondStageDeadBand() float64 {
	return r.secondStageDeadBand
}

func (r *airConditionerStatus) Init(from *RvcFrame) {
	r.RvcItem.Init(from)
	var changed = false
	var dataBytes = &from.Data

	{

		var tmp = utils.GetByte(dataBytes, 0)
		if (*r).instance != tmp {
			changed = true
		}
		(*r).instance = tmp
		tmp = utils.GetByte(dataBytes, 1)
		if (*r).operatingMode != tmp {
			changed = true
		}
		(*r).operatingMode = tmp
	}
	{
		var tmp = convert.ToPercent(utils.GetByte(dataBytes, 2))
		if tmp != (*r).maxFanSpeed {
			changed = true
		}
		(*r).maxFanSpeed = tmp

		tmp = convert.ToPercent(utils.GetByte(dataBytes, 3))
		if tmp != (*r).maxOutputLevel {
			changed = true
		}
		(*r).maxOutputLevel = tmp

		tmp = convert.ToPercent(utils.GetByte(dataBytes, 4))
		if tmp != (*r).fanSpeed {
			changed = true
		}
		(*r).fanSpeed = tmp

		tmp = convert.ToPercent(utils.GetByte(dataBytes, 5))
		if tmp != (*r).acOutputLevel {
			changed = true
		}
		(*r).acOutputLevel = tmp

		tmp = convert.ToDegreesC(utils.GetByte(dataBytes, 6))
		if tmp != (*r).deadBand {
			changed = true
		}
		(*r).deadBand = tmp

		tmp = convert.ToDegreesC(utils.GetByte(dataBytes, 7))
		if tmp != (*r).secondStageDeadBand {
			changed = true
		}
		(*r).secondStageDeadBand = tmp

	}
	if changed {
		(*r).LastChanged = (*r).Timestamp
	}
}
