package rvc

const (
	DGN_DC_SOURCE_STATUS_1        = 0x1FFFD
	DGN_DC_SOURCE_STATUS_1_SPYDER = 0x10FFD
	DGN_DC_DIMMER_STATUS_3        = 0x1FEDA
	DGN_TANK_STATUS               = 0x1FFB7
	DGN_AIR_CONDITIONER_STATUS    = 0x1FFE1
)

/*
import "sync"
*/

// DGNInstanceKey - a key to id a particular Instance (battery bank, light, etc.)
// Not every DGN has a key and the associated name may be different for different systems
// Some key names are part of the spec.
type DGNInstanceKey struct {
	DGN      uint32
	Instance byte
}

// var DGNInstanceNames map[DGNInstanceKey]string = *new(map[DGNInstanceKey]string)
var DGNInstanceNames map[DGNInstanceKey]string

// 1 Main Ceiling
// 2 Main Entry
// 3 passenger slide ceiling
//

func init() {
	// NOTE: Envision this is read from a config file....
	DGNInstanceNames = make(map[DGNInstanceKey]string)

	// these would be read from a config file
	// consider how to set up and load a custom DGN and if ALL DGN info can be loaded from a config file
	// It's not (yet) clear how to set up dgnHasInstances with proper data types and fields from a config file
	// Maybe a generic Number/Value type like Java Number?
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 1}] = "Main House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 1}] = "Main House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 2}] = "Chassis Start Battery"

	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 1}] = "Main Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 2}] = "Main Entry"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 3}] = "Passenger Side Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 4}] = "Driver Side Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 6}] = "Passenger Slide"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 7}] = "Main Sconce"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 8}] = "Driver Slide"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 9}] = "Passenger Courtesy"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 10}] = "Bedroom Courtesy"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 11}] = "Bedroom Sconce"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 12}] = "Bedroom Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 14}] = "Passenger Task"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 15}] = "Mid Bath Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 16}] = "Mid Bath Vanity"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 18}] = "TV Accent"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 19}] = "Acccent 1"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 20}] = "Acccent 2"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 21}] = "Bedroom Accent"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 28}] = "Closet  Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 73}] = "Exterior Awning"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 75}] = "Exterior Door"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 76}] = "Exterior Slide"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 85}] = "Exterior Under Coach"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 86}] = "Exterior Windshield"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 87}] = "Passenger Map"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 91}] = "Exterior Road"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 92}] = "Exterior Porch"
	// 0 — Fresh water
	//1 — Black waste (If only one waste tank is
	//present, it shall be reported as Black.)
	//2 — Gray waste
	//3 — LPG
	//16 — Second fresh water
	//17 — Second black waste
	//18 — Second gray waste
	//19 — Second LPG
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 0}] = "Fresh Water"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 1}] = "Black"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 2}] = "Gray"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 3}] = "LPG"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 16}] = "Fresh Water (2)"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 17}] = "Black (2)"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 18}] = "Gray (2)"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_TANK_STATUS, Instance: 19}] = "LPG (2)"
}

// GetInstanceName - get the Instance name for the given key. Returns VAL/TRUE if found and ""/FALSE if not found.
func GetInstanceName(k DGNInstanceKey) (string, bool) {
	var ret, found = DGNInstanceNames[k]
	if found {
		return ret, true
	} else {
		return "", false
	}
}

/*
type SystemState struct {
}

type DGNInstanceFactory struct {
	cache sync.Map
}

func (f *DGNInstanceFactory) get(DGN uint32, Instance byte) RvcItem {
	var k = DGNInstanceKey{
		DGN:      DGN,
		Instance: Instance,
	}
	var v, ok = f.cache.Load(k)
	if ok {
		var x = v.(RvcItem)
		return x
	}


}
*/
