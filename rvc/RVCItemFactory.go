package rvc

import (
	"fmt"
	"rvctomqtt/constants"
	sync2 "sync"
)

// A cache of RVCItems  - one per DGN/Instance key
// This cache also knows how to create items
// Will have to address thread safety for cache and individual items
// The chache will just give a reference to the item or create if needed
// items never go away as the cache represents the current state of everything it sees
// an item can be updated from a new packet, or queried for its current state, we also should eventually
// consider change listeners.

// a single cache for the whole app

var rvcItemMap map[DGNInstanceKey]*RvcItemIF

// dgnHasInstances - a singleton map that tells us if this DGN has an instance field. Since we maintain the state of everything
// we need to know if a particular DGN has multiple instances. We need to know this BEFORE we create the instance.

var dgnHasInstances map[uint32]bool
var locker = sync2.RWMutex{}

// GetRVCItem - get the RVC item for the given frame. If it does not exist create one, cache it and return it
func GetRVCItem(f *RvcFrame) (*RvcItemIF, bool) {

	var key = getInstanceKey(f)
	locker.Lock()
	defer locker.Unlock()
	var ret, ok = rvcItemMap[key]
	if ok {
		return ret, true
	} else {
		var created, ok = createRVCItem(f)
		if ok {
			rvcItemMap[key] = &created
			return &created, true
		}
	}
	return nil, false

}

func init() {
	rvcItemMap = make(map[DGNInstanceKey]*RvcItemIF)

	dgnHasInstances = make(map[uint32]bool)
	dgnHasInstances[DGN_DC_DIMMER_STATUS_3] = true
	dgnHasInstances[DGN_DC_SOURCE_STATUS_1_SPYDER] = true
	dgnHasInstances[DGN_DC_SOURCE_STATUS_1] = true
	dgnHasInstances[DGN_TANK_STATUS] = true
	dgnHasInstances[DGN_AIR_CONDITIONER_STATUS] = true
	dgnHasInstances[DGN_AIR_CONDITIONER_COMMAND] = true
}

// getInstanceKey - get the instance key for this frame. It will pull the instatnceID if we have one
func getInstanceKey(f *RvcFrame) DGNInstanceKey {
	var dgn = f.DGN()
	var hasInstance = dgnHasInstances[dgn]
	if hasInstance {
		return DGNInstanceKey{dgn, f.Data[0]}
	} else {
		return DGNInstanceKey{dgn, constants.DataNotAvailableUint8}
	}

}

// createRVCItem - create a new RVC Item from the supplied frame
func createRVCItem(f *RvcFrame) (RvcItemIF, bool) {
	var dgn = f.DGN()

	switch dgn {
	case DGN_DC_SOURCE_STATUS_1_SPYDER:
		{
			//var name = DGNInstanceNames[DGN_DC_SOURCE_STATUS_1_SPYDER]
			var ret RvcItemIF
			ret = &DCSourceStatus1{}
			return ret, true
		}
	case DGN_DC_DIMMER_STATUS_3:
		{
			var ret RvcItemIF
			ret = &DCDimmerStatus3{}
			return ret, true
		}
	case DGN_DC_DIMMER_COMMAND_2:
		{
			var ret RvcItemIF
			ret = &DCDimmerCommand2{}
			return ret, true
		}
	case DGN_DC_SOURCE_STATUS_1:
		{
			var ret RvcItemIF
			ret = &DCSourceStatus1{}
			return ret, true
		}
	case DGN_TANK_STATUS:
		{
			var ret RvcItemIF
			ret = &tankStatus{}
			return ret, true
		}

	case DGN_AIR_CONDITIONER_STATUS:
		{
			var ret RvcItemIF
			ret = &airConditionerStatus{}
			return ret, true
		}

	case DGN_AIR_CONDITIONER_COMMAND:
		{
			var ret RvcItemIF
			ret = &airConditionerCommand{}
			return ret, true
		}

	case DGN_INVERTER_TEMPERATURE_STATUS:
		{
			var ret RvcItemIF
			ret = &inverterTemperatureStatus{}
			return ret, true
		}
	case DGN_INVERTER_TEMPERATURE_STATUS_2:
		{
			var ret RvcItemIF
			ret = &inverterTemperatureStatus2{}
			return ret, true
		}

	case DGN_ADDRESS_CLAIMED:
		{
			var ret RvcItemIF
			ret = &AddressClaimed{}
			return ret, true
		}

	}

	// special case = the lower dgn is the desired address
	if (dgn & 0xff00) == DGN_ADDRESS_CLAIM {
		fmt.Printf("TRY ADDRESS CLAIM: %x", DGN_ADDRESS_CLAIM)
		var ret RvcItemIF
		ret = &AddressClaim{}
		return ret, true
	}

	// NOTE: This should really be 0xeeff - but some SA's have bugs (ee42)
	if (dgn & 0xff00) == DGN_ADDRESS_CLAIMED {
		fmt.Printf("TRY ADDRESS CLAIMED:  %x for %x", DGN_ADDRESS_CLAIMED, dgn)
		var ret RvcItemIF
		ret = &AddressClaimed{}
		return ret, true
	}
	return nil, false
}
