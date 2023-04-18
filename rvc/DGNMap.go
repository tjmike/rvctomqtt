package rvc

const (
	DGN_DC_SOURCE_STATUS_1        = 0x1FFFD
	DGN_DC_SOURCE_STATUS_1_SPYDER = 0x10FFD
	DGN_DC_DIMMER_STATUS_3        = 0x1FEDA
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

// var DGNNames map[DGNInstanceKey]string = *new(map[DGNInstanceKey]string)
var DGNNames map[DGNInstanceKey]string

// 1 Main Ceiling
// 2 Main Entry
// 3 passenger slide ceiling
//

func init() {
	// NOTE: Envision this is read from a config file....
	DGNNames = make(map[DGNInstanceKey]string)

	// these would be read from a config file
	// consider how to set up and load a custom DGN and if ALL DGN info can be loaded from a config file
	// It's not (yet) clear how to set up instances with proper data types and fields from a config file
	// Maybe a generic Number/Value type like Java Number?
	DGNNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 1}] = "Main House Battery Bank"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 1}] = "Main House Battery Bank"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 2}] = "Chassis Start Battery"

	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 1}] = "Main Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 2}] = "Main Entry"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 3}] = "Passenger Side Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 4}] = "Driver Side Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 6}] = "Passenger Slide"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 7}] = "Main Sconce"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 8}] = "Driver Slide"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 9}] = "Passenger Courtesy"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 10}] = "Bedroom Courtesy"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 11}] = "Bedroom Sconce"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 12}] = "Bedroom Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 14}] = "Passenger Task"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 15}] = "Mid Bath Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 16}] = "Mid Bath Vanity"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 18}] = "TV Accent"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 19}] = "Acccent 1"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 20}] = "Acccent 2"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 21}] = "Bedroom Accent"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 28}] = "Closet  Ceiling"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 73}] = "Exterior Awning"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 75}] = "Exterior Door"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 76}] = "Exterior Slide"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 85}] = "Exterior Under Coach"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 86}] = "Exterior Windshield"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 87}] = "Passenger Map"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 91}] = "Exterior Road"
	DGNNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: 92}] = "Exterior Porch"
}

// GetInstanceName - get the Instance name for the given key. Returns VAL/TRUE if found and ""/FALSE if not found.
func GetInstanceName(k DGNInstanceKey) (string, bool) {
	var ret, found = DGNNames[k]
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
