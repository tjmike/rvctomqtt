package rvc

import (
	"fmt"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
)

// todo - move to a types file

// 0x1FEDA
type DCDimmerCommand2 struct {
	RvcItem
	deviceInstance    byte    // 0
	group             uint8   // 1
	desiredBrightness float64 // 2 (%)
	command           uint8   // 3
	delayDuration     uint8   // 4 seconds

	interlockStatus uint2   // 5 (0,1)
	reserved1       uint2   // 5 (2-3)
	reserved2       uint2   // 5 (4-5)
	reserved3       uint2   // 5 (6-7)
	rampTime        float64 // seconds (0.1) 0-25 seconds
	reserved        uint8   //
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
func (i *DCDimmerCommand2) GetInterlockStatus() uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.interlockStatus
}

func (i *DCDimmerCommand2) GetDelayDuration() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.delayDuration
}

func (i *DCDimmerCommand2) GetRampTime() float64 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.rampTime
}

func (i *DCDimmerCommand2) GetReserved() uint8 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved
}

func (i *DCDimmerCommand2) GetReserved1() uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved1
}

func (i *DCDimmerCommand2) GetReserved2() uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved2
}
func (i *DCDimmerCommand2) GetReserved3() uint2 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.reserved3
}

func (i *DCDimmerCommand2) GetInstanceName() string {
	var k = DGNInstanceKey{DGN: i.dgn, Instance: i.deviceInstance}
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

	return fmt.Sprintf("%s DGN: %x (%s) SA: %d Instance: %d (%s) group: %d brigntness: %f command: %d lockitem: %d res1: %d res2: %d res3: %d rampTime %f reserved %d",
		i.GetTimestamp().Format("01-02-2006 15:04:05.000000"),
		i.dgn, i.name,
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

/*
func createMessage() *RvcFrame {

		// 29 bit message
		var priority uint8 = 1 // 3 bits
		var reserved uint8 = 0 // 1 bit

		var dgnHigh uint16 = 0 // 9 bits
		var dgnLow uint8 = 0   // 8 bits
		var sa uint8 = 0       // 8      // THIS WILL HAVE TO BE NEGOTIATED AT THE BEGINNING

		// reverse
		//deviceInstance    byte    // 0
		//group             uint8   // 1
		//desiredBrightness float64 // 2 (%)
		//command           uint8   // 3
		//delayDuration     uint8   // 4 seconds
		//
		//interlockStatus uint2   // 5 (0,1)
		//reserved1       uint2   // 5 (2-3)
		//reserved2       uint2   // 5 (4-5)
		//reserved3       uint2   // 5 (6-7)
		//rampTime        float64 // seconds (0.1) 0-25 seconds
		//reserved        uint8   //

		can.Frame{
			Timestamp:    time.Time{},
			ID:           0,
			Length:       0,
			Flags:        0,
			Res0:         0,
			Res1:         0,
			Data:         [8]uint8{},
			MessageBytes: [16]byte{},
		}

		// create an RVC message from a CAN message
		var j = j1939.J1939Frame{
			PGN:   j1939.PGN{},
			Frame: can.Frame{},
		}
		var v = RvcFrame{}
	}
*/
func (r *DCDimmerCommand2) Init(from *RvcFrame) {
	r.lock.Lock()
	defer r.lock.Unlock()
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
		var tmp = uint2(utils.GetBits(dataBytes, 5, 0, utils.GetMask(2)))
		if (*r).interlockStatus != tmp {
			changed = true
		}
		(*r).interlockStatus = tmp

		tmp = uint2(utils.GetBits(dataBytes, 5, 2, utils.GetMask(2)))
		if (*r).reserved1 != tmp {
			changed = true
		}
		(*r).reserved1 = tmp

		tmp = uint2(utils.GetBits(dataBytes, 5, 4, utils.GetMask(2)))
		if (*r).reserved2 != tmp {
			changed = true
		}
		(*r).reserved2 = tmp

		tmp = uint2(utils.GetBits(dataBytes, 5, 6, utils.GetMask(2)))
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
