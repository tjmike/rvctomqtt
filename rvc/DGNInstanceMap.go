package rvc

import "fmt"

const (
	DGN_DC_SOURCE_STATUS_1            uint32 = 0x1FFFD
	DGN_DC_SOURCE_STATUS_1_SPYDER     uint32 = 0x10FFD
	DGN_DC_DIMMER_STATUS_3            uint32 = 0x1FEDA
	DGN_DC_DIMMER_COMMAND_2           uint32 = 0x1FEDB
	DGN_TANK_STATUS                   uint32 = 0x1FFB7
	DGN_AIR_CONDITIONER_STATUS        uint32 = 0x1FFE1
	DGN_AIR_CONDITIONER_COMMAND       uint32 = 0x1FFE0
	DGN_INVERTER_TEMPERATURE_STATUS   uint32 = 0x1FEBD
	DGN_INVERTER_TEMPERATURE_STATUS_2 uint32 = 0x1FDCB
	DGN_INFORMATION_REQUEST           uint32 = 0x0EA00
	DGN_ADDRESS_CLAIMED               uint32 = 0x0EE00
	DGN_INITIAL_PACKET                uint32 = 0x0ECFF
	DGN_DATA_PACKET                   uint32 = 0x0EBFF
	DGN_CHASSIS_MOBILITY_STATUS       uint32 = 0x1FFF4

	DGN_PRODUCT_IDENTIFICATION_MESSAGE uint32 = 0xFEEB

	INSTANCE_BEDROOM_CEILING      uint8 = 12
	INSTANCE_LIGHT_PASSENGER_TASK uint8 = 14
	INSTANCE_REAR_AC_UNIT         uint8 = 1
	INSTANCE_MID_AC_UNIT          uint8 = 2
	INSTANCE_FRONT_AC_UNIT        uint8 = 3
	INSTANCE_WATER_PUMP           uint8 = 125
	INSTANCE_AQUAHOT_ELECTRIC     uint8 = 129
	INSTANCE_AQUAHOT_GAS          uint8 = 130
	INSTANCE_AQUAHOT_PREHEAT      uint8 = 131

	RVC_DATA_NOT_AVAILABLE uint8 = 255
)

/*
import "sync"
*/

// DGNInstanceKey - a key to id a particular Instance (battery bank, light, etc.)
// Not every DGN has a key and the associated Name may be different for different systems
// Some key names are part of the spec.
//type DGNInstanceKey struct {
//	DGN      uint32
//	Instance byte
//}

// var DGNInstanceNames map[DGNInstanceKey]string = *new(map[DGNInstanceKey]string)
var DGNInstanceNames map[any]string

// 1 Main Ceiling
// 2 Main Entry
// 3 passenger slide ceiling
//

func init() {
	// NOTE: Envision this is read from a config file....
	DGNInstanceNames = make(map[any]string)

	// these would be read from a config file
	// consider how to set up and load a custom DGN and if ALL DGN info can be loaded from a config file
	// It's not (yet) clear how to set up dgnHasInstances with proper data types and fields from a config file
	// Maybe a generic Number/Value type like Java Number?
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 1}] = "Main House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 2}] = "Chassis Start Battery"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 3}] = "Secondary House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1_SPYDER, Instance: 4}] = "Generator Starter Battery"

	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 1}] = "Main House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 2}] = "Chassis Start Battery"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 3}] = "Secondary House Battery Bank"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_SOURCE_STATUS_1, Instance: 4}] = "Generator Starter Battery"

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
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_BEDROOM_CEILING}] = "Bedroom Ceiling"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_LIGHT_PASSENGER_TASK}] = "Passenger Task"
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
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_WATER_PUMP}] = "Water Pump"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_AQUAHOT_ELECTRIC}] = "Aquahot Electric"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_AQUAHOT_GAS}] = "Aquahot Gas"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_DC_DIMMER_STATUS_3, Instance: INSTANCE_AQUAHOT_PREHEAT}] = "Aquahot PreHeat"

	// 47,58,59,120 seem to fire all the time...

	// 129  Aquahot ELEC (MAYBE NOT????)
	//  130 Aquahot Fuel
	// 131   Aquahot PreHeat
	// 125    Water pump:

	// (30,31)  mid bath vent lid????
	// (29) mid bath vent fan

	//  (27,26) galley vent lid
	//  (25) galley vent fan
	// CLOSE:
	// TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/26 JSON: {"Command":3,"DGN":130779,"DelayDuration":0,"DesiredBrightness":0.000000,"Group":255,"Instance":26,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:14:08.513488-04:00","LastNotified":"2023-05-30T14:14:08.513488-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":150,"Timestamp":"2023-05-30T14:14:08.513488-04:00"}
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/27 JSON: {"Command":1,"DGN":130779,"DelayDuration":20,"DesiredBrightness":100.000000,"Group":255,"Instance":27,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:14:08.517644-04:00","LastNotified":"2023-05-30T14:14:08.517644-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":150,"Timestamp":"2023-05-30T14:14:08.517644-04:00"}
	// OPEN:
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/27 JSON: {"Command":3,"DGN":130779,"DelayDuration":0,"DesiredBrightness":0.000000,"Group":255,"Instance":27,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:15:28.357532-04:00","LastNotified":"2023-05-30T14:15:28.357532-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":150,"Timestamp":"2023-05-30T14:15:28.357532-04:00"}
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/26 JSON: {"Command":1,"DGN":130779,"DelayDuration":20,"DesiredBrightness":100.000000,"Group":255,"Instance":26,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:15:28.371988-04:00","LastNotified":"2023-05-30T14:15:28.371988-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":150,"Timestamp":"2023-05-30T14:15:28.371988-04:00"}
	// OPEN: Spyder
	// TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/31 JSON: {"Command":3,"DGN":130779,"DelayDuration":0,"DesiredBrightness":0.000000,"Group":255,"Instance":31,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:32:04.647526-04:00","LastNotified":"2023-05-30T14:32:04.647526-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":156,"Timestamp":"2023-05-30T14:32:04.647526-04:00"}
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/30 JSON: {"Command":1,"DGN":130779,"DelayDuration":20,"DesiredBrightness":100.000000,"Group":255,"Instance":30,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:32:04.714473-04:00","LastNotified":"2023-05-30T14:32:04.714473-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":156,"Timestamp":"2023-05-30T14:32:04.714473-04:00"}
	// Open Spyder HOLDING and then release
	// TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/31 JSON: {"Command":3,"DGN":130779,"DelayDuration":0,"DesiredBrightness":0.000000,"Group":255,"Instance":31,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:33:42.188644-04:00","LastNotified":"2023-05-30T14:33:42.188644-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":156,"Timestamp":"2023-05-30T14:33:42.188644-04:00"}
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/30 JSON: {"Command":1,"DGN":130779,"DelayDuration":2,"DesiredBrightness":100.000000,"Group":255,"Instance":30,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:33:43.187872-04:00","LastNotified":"2023-05-30T14:33:43.187872-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":156,"Timestamp":"2023-05-30T14:33:43.187872-04:00"}
	//TOPIC: mycoach/events/DC_DIMMER_COMMAND_2/30 JSON: {"Command":3,"DGN":130779,"DelayDuration":0,"DesiredBrightness":0.000000,"Group":255,"Instance":30,"InstanceName":"","InterlockStatus":0,"LastChanged":"2023-05-30T14:33:46.737499-04:00","LastNotified":"2023-05-30T14:33:46.737499-04:00","Name":"DC_DIMMER_COMMAND_2","Priority":6,"RampTime":25.500000,"Reserved":255,"Reserved1":0,"Reserved2":0,"Reserved3":0,"SourceAddress":156,"Timestamp":"2023-05-30T14:33:46.737499-04:00"}

	// locks
	//   all
	//   door
	//   bay

	DGNInstanceNames[DGNInstanceKey{DGN: DGN_AIR_CONDITIONER_STATUS, Instance: INSTANCE_REAR_AC_UNIT}] = "Rear AC"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_AIR_CONDITIONER_STATUS, Instance: INSTANCE_MID_AC_UNIT}] = "Mid AC"
	DGNInstanceNames[DGNInstanceKey{DGN: DGN_AIR_CONDITIONER_STATUS, Instance: INSTANCE_FRONT_AC_UNIT}] = "Front AC"
	//
	{
		tmp := make(map[DGNInstanceKey]string)
		// make map for dimmer command based off status mapping
		for k, v := range DGNInstanceNames {
			kk, ok := k.(DGNInstanceKey)

			if ok {
				if kk.DGN == DGN_DC_DIMMER_STATUS_3 {
					k2 := DGNInstanceKey{DGN_DC_DIMMER_COMMAND_2, kk.Instance}
					tmp[k2] = v
				}
			} else {
				// TODO user logger
				fmt.Printf("Expected DGNInstanceKey but cast failed")
			}
		}
		// set the instances in the original map - we don't want to set values while we are iterating....
		for k, v := range tmp {
			DGNInstanceNames[k] = v
		}
	}
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

// GetInstanceName - get the Instance Name for the given key. Returns VAL/TRUE if found and ""/FALSE if not found.
func GetInstanceName(k any) (string, bool) {
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
