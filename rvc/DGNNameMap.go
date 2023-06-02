package rvc

var dGNtoName map[uint32]string = make(map[uint32]string)

func init() {
	dGNtoName[0x1f9c] = "THERMOSTAT_AMBIENT_STATUS"
	dGNtoName[0x1f9c] = "THERMOSTAT_AMBIENT_STATUS"
	dGNtoName[0x1ffd7] = "INVERTER_AC_STATUS_1"
	dGNtoName[0x1ffd4] = "INVERTER_STATUS"
	dGNtoName[0x17F00] = "GENERAL_RESET"
	dGNtoName[0x17E00] = "TERMINAL"
	dGNtoName[0x1FED8] = "GENERIC_CONFIGURATION_STATUS"
	dGNtoName[0x1FFFF] = "DATE_TIME_STATUS"
	dGNtoName[0x1FFFE] = "SET_DATE_TIME_COMMAND"

	dGNtoName[DGN_DC_SOURCE_STATUS_1] = "DC_SOURCE_STATUS_1"
	dGNtoName[0x1FFFC] = "DC_SOURCE_STATUS_2"
	dGNtoName[0x1FFFB] = "DC_SOURCE_STATUS_3"
	dGNtoName[0x1FEC9] = "DC_SOURCE_STATUS_4"
	dGNtoName[0x1FEC8] = "DC_SOURCE_STATUS_5"
	dGNtoName[0x1FEC7] = "DC_SOURCE_STATUS_6"

	dGNtoName[DGN_DC_SOURCE_STATUS_1_SPYDER] = "DC_SOURCE_STATUS_SPYDER"

	dGNtoName[0x1FFFA] = "COMMUNICATION_STATUS_1"
	dGNtoName[0x1FFF9] = "COMMUNICATION_STATUS_2"
	dGNtoName[0x1FFF8] = "COMMUNICATION_STATUS_3"
	dGNtoName[0x1FFF7] = "WATERHEATER_STATUS"
	dGNtoName[0x1FFF6] = "WATERHEATER_COMMAND"
	dGNtoName[0x1FFF5] = "GAS_SENSOR_STATUS"
	dGNtoName[DGN_CHASSIS_MOBILITY_STATUS] = "CHASSIS_MOBILITY_STATUS"
	dGNtoName[0x1FFF3] = "CHASSIS_MOBILITY_COMMAND"
	dGNtoName[0x1FFF2] = "AAS_CONFIG_STATUS"
	dGNtoName[0x1FFF1] = "AAS_COMMAND"
	dGNtoName[0x1FFF0] = "AAS_STATUS"
	dGNtoName[0x1FFEF] = "AAS_SENSOR_STATUS"
	dGNtoName[0x1FED1] = "SUSPENSION_AIR_PRESSURE_STATUS"
	dGNtoName[0x1FFEE] = "LEVELING_CONTROL_COMMAND"
	dGNtoName[0x1FFED] = "LEVELING_CONTROL_STATUS"
	dGNtoName[0x1FFEC] = "LEVELING_JACK_STATUS"
	dGNtoName[0x1FFEB] = "LEVELING_SENSOR_STATUS"
	dGNtoName[0x1FFEA] = "HYDRAULIC_PUMP_STATUS"
	dGNtoName[0x1FEBC] = "HYDRAULIC_PUMP_COMMAND"
	dGNtoName[0x1FFE9] = "LEVELING_AIR_STATUS"
	dGNtoName[0x1FFE8] = "SLIDE_STATUS"
	dGNtoName[0x1FFE7] = "SLIDE_COMMAND"
	dGNtoName[0x1FFE6] = "SLIDE_SENSOR_STATUS"
	dGNtoName[0x1FFE5] = "SLIDE_MOTOR_STATUS"
	dGNtoName[0x1FFE4] = "FURNACE_STATUS"
	dGNtoName[0x1FFE3] = "FURNACE_COMMAND"
	dGNtoName[0x1FFE2] = "THERMOSTAT_STATUS_1"
	dGNtoName[0x1FEFA] = "THERMOSTAT_STATUS_2"
	dGNtoName[0x1FEF9] = "THERMOSTAT_COMMAND_1"
	dGNtoName[0x1FEF8] = "THERMOSTAT_COMMAND_2"
	dGNtoName[0x1FEF7] = "THERMOSTAT_SCHEDULE_STATUS_1"
	dGNtoName[0x1FEF6] = "THERMOSTAT_SCHEDULE_STATUS_2"
	dGNtoName[0x1FEF5] = "THERMOSTAT_SCHEDULE_COMMAND_1"
	dGNtoName[0x1FEF4] = "THERMOSTAT_SCHEDULE_COMMAND_2"
	dGNtoName[0x1FF9C] = "THERMOSTAT_AMBIENT_STATUS"
	dGNtoName[DGN_AIR_CONDITIONER_STATUS] = "AIR_CONDITIONER_STATUS"
	dGNtoName[DGN_AIR_CONDITIONER_COMMAND] = "AIR_CONDITIONER_COMMAND"
	dGNtoName[0x1FF9B] = "HEAT_PUMP_STATUS"
	dGNtoName[0x1FF9A] = "HEAT_PUMP_COMMAND"
	dGNtoName[0x1FFDF] = "GENERATOR_AC_STATUS_1"
	dGNtoName[0x1FFDE] = "GENERATOR_AC_STATUS_2"
	dGNtoName[0x1FFDD] = "GENERATOR_AC_STATUS_3"
	dGNtoName[0x1FF94] = "GENERATOR_AC_STATUS_4"
	dGNtoName[0x1FEC6] = "GENERATOR_DC_STATUS_1"
	dGNtoName[0x1FEC5] = "GENERATOR_DC_CONFIGURATION_STATUS"
	dGNtoName[0x1FEC4] = "GENERATOR_DC_COMMAND"
	dGNtoName[0x1FEC3] = "GENERATOR_DC_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEC2] = "GENERATOR_DC_EQUALIZATION_STATUS"
	dGNtoName[0x1FEC1] = "GENERATOR_DC_EQUALIZATION_CONFIGURATION_STATUS"
	dGNtoName[0x1FEC0] = "GENERATOR_DC_EQUALIZATION_CONFIGURATION_COMMAND"
	dGNtoName[0x1FFDC] = "GENERATOR_STATUS_1"
	dGNtoName[0x1FFDB] = "GENERATOR_STATUS_2"
	dGNtoName[0x1FFDA] = "GENERATOR_COMMAND"
	dGNtoName[0x1FFD9] = "GENERATOR_START_CONFIG_STATUS"
	dGNtoName[0x1FFD8] = "GENERATOR_START_CONFIG_COMMAND"
	dGNtoName[0x1FFD6] = "INVERTER_AC_STATUS_2"
	dGNtoName[0x1FFD5] = "INVERTER_AC_STATUS_3"
	dGNtoName[0x1FF8F] = "INVERTER_AC_STATUS_4"
	dGNtoName[0x1FF8E] = "INVERTER_ACFAULT_CONFIGURATION_STATUS_1"
	dGNtoName[0x1FF8D] = "INVERTER_ACFAULT_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FF8C] = "INVERTER_ACFAULT_CONFIGURATION_COMMAND_1"
	dGNtoName[0x1FF8B] = "INVERTER_ACFAULT_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FFD3] = "INVERTER_COMMAND"
	dGNtoName[0x1FFD2] = "INVERTER_CONFIGURATION_STATUS_1"
	dGNtoName[0x1FFD1] = "INVERTER_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FECE] = "INVERTER_CONFIGURATION_STATUS_3"
	dGNtoName[0x1FFD0] = "INVERTER_CONFIGURATION_COMMAND_1"
	dGNtoName[0x1FFCF] = "INVERTER_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FECD] = "INVERTER_CONFIGURATION_COMMAND_3"
	dGNtoName[0x1FFCE] = "INVERTER_STATISTIC_STATUS"
	dGNtoName[0x1FFCD] = "INVERTER_APS_STATUS"
	dGNtoName[0x1FFCC] = "INVERTER_DCBUS_STATUS"
	dGNtoName[0x1FFCB] = "INVERTER_OPE_STATUS"
	dGNtoName[0x1FEE8] = "INVERTER_DC_STATUS"
	dGNtoName[DGN_INVERTER_TEMPERATURE_STATUS] = "INVERTER_TEMPERATURE_STATUS"
	dGNtoName[DGN_INVERTER_TEMPERATURE_STATUS_2] = "INVERTER_TEMPERATURE_STATUS2"
	dGNtoName[0x1FFCA] = "CHARGER_AC_STATUS_1"
	dGNtoName[0x1FFC9] = "CHARGER_AC_STATUS_2"
	dGNtoName[0x1FFC8] = "CHARGER_AC_STATUS_3"
	dGNtoName[0x1FF8A] = "CHARGER_AC_STATUS_4"
	dGNtoName[0x1FF89] = "CHARGER_ACFAULT_CONFIGURATION_STATUS_1"
	dGNtoName[0x1FF88] = "CHARGER_ACFAULT_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FF87] = "CHARGER_ACFAULT_CONFIGURATION_COMMAND_1"
	dGNtoName[0x1FF86] = "CHARGER_ACFAULT_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FFC7] = "CHARGER_STATUS"
	dGNtoName[0x1FFC6] = "CHARGER_CONFIGURATION_STATUS"
	dGNtoName[0x1FFC5] = "CHARGER_COMMAND"
	dGNtoName[0x1FFC4] = "CHARGER_CONFIGURATION_COMMAND"
	dGNtoName[0x1FF96] = "CHARGER_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FF95] = "CHARGER_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FECC] = "CHARGER_CONFIGURATION_STATUS_3"
	dGNtoName[0x1FECB] = "CHARGER_CONFIGURATION_COMMAND_3"
	dGNtoName[0x1FEBF] = "CHARGER_CONFIGURATION_STATUS_4"
	dGNtoName[0x1FEBE] = "CHARGER_CONFIGURATION_COMMAND_4"
	dGNtoName[0x1FF99] = "CHARGER_EQUALIZATION_STATUS"
	dGNtoName[0x1FF98] = "CHARGER_EQUALIZATION_CONFIGURATION_STATUS"
	dGNtoName[0x1FF97] = "CHARGER_EQUALIZATION_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEBB] = "GENERIC_AC_STATUS_1"
	dGNtoName[0x1FEBA] = "GENERIC_AC_STATUS_2"
	dGNtoName[0x1FEB9] = "GENERIC_AC_STATUS_3"
	dGNtoName[0x1FEB8] = "GENERIC_AC_STATUS_4"
	dGNtoName[0x1FEB7] = "GENERIC_ACFAULT_CONFIGURATION_STATUS_1"
	dGNtoName[0x1FEB6] = "GENERIC_ACFAULT_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FEB5] = "GENERIC_ACFAULT_CONFIGURATION_COMMAND_1"
	dGNtoName[0x1FEB4] = "GENERIC_ACFAULT_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FFBF] = "AC_LOAD_STATUS"
	dGNtoName[0x1FEDD] = "AC_LOAD_STATUS_2"
	dGNtoName[0x1FFBE] = "AC_LOAD_COMMAND"
	dGNtoName[0x1FFBD] = "DC_LOAD_STATUS"
	dGNtoName[0x1FEDC] = "DC_LOAD_STATUS_2"
	dGNtoName[0x1FFBC] = "DC_LOAD_COMMAND"
	dGNtoName[0x1FFBB] = "DC_DIMMER_STATUS_1"
	dGNtoName[0x1FFBA] = "DC_DIMMER_STATUS_2"
	dGNtoName[DGN_DC_DIMMER_STATUS_3] = "DC_DIMMER_STATUS_3"
	dGNtoName[0x1FFB9] = "DC_DIMMER_COMMAND"
	dGNtoName[DGN_DC_DIMMER_COMMAND_2] = "DC_DIMMER_COMMAND_2"
	dGNtoName[0x1FFB8] = "DIGITAL_INPUT_STATUS"
	dGNtoName[0x1FED7] = "GENERIC_INDICATOR_STATUS"
	dGNtoName[0x1FED9] = "GENERIC_INDICATOR_COMMAND"
	dGNtoName[0x1FEE0] = "DC_MOTOR_CONTROL_STATUS"
	dGNtoName[0x1FEE1] = "DC_MOTOR_CONTROL_COMMAND"
	dGNtoName[DGN_TANK_STATUS] = "TANK_STATUS"
	dGNtoName[0x1FFB6] = "TANK_CALIBRATION_COMMAND"
	dGNtoName[0x1FFB5] = "TANK_GEOMETRY_STATUS"
	dGNtoName[0x1FFB4] = "TANK_GEOMETRY_COMMAND"
	dGNtoName[0x1FFB3] = "WATER_PUMP_STATUS"
	dGNtoName[0x1FFB2] = "WATER_PUMP_COMMAND"
	dGNtoName[0x1FFB1] = "AUTOFILL_STATUS"
	dGNtoName[0x1FFB0] = "AUTOFILL_COMMAND"
	dGNtoName[0x1FFAF] = "WASTEDUMP_STATUS"
	dGNtoName[0x1FFAE] = "WASTEDUMP_COMMAND"
	dGNtoName[0x1FFAD] = "ATS_AC_STATUS_1"
	dGNtoName[0x1FFAC] = "ATS_AC_STATUS_2"
	dGNtoName[0x1FFAB] = "ATS_AC_STATUS_3"
	dGNtoName[0x1FF85] = "ATS_AC_STATUS_4"
	dGNtoName[0x1FF84] = "ATS_ACFAULT_CONFIGURATION_STATUS_1"
	dGNtoName[0x1FF83] = "ATS_ACFAULT_CONFIGURATION_STATUS_2"
	dGNtoName[0x1FF82] = "ATS_ACFAULT_CONFIGURATION_COMMAND_1"
	dGNtoName[0x1FF81] = "ATS_ACFAULT_CONFIGURATION_COMMAND_2"
	dGNtoName[0x1FFAA] = "ATS_STATUS"
	dGNtoName[0x1FFA9] = "ATS_COMMAND"
	dGNtoName[0x1FFA5] = "WEATHER_STATUS_1"
	dGNtoName[0x1FFA4] = "WEATHER_STATUS_2"
	dGNtoName[0x1FFA3] = "ALTIMETER_STATUS"
	dGNtoName[0x1FFA2] = "ALTIMETER_COMMAND"
	dGNtoName[0x1FFA1] = "WEATHER_CALIBRATE_COMMAND"
	dGNtoName[0x1FFA0] = "COMPASS_BEARING_STATUS"
	dGNtoName[0x1FF9F] = "COMPASS_CALIBRATE_COMMAND"
	dGNtoName[0x1FF80] = "GENERATOR_DEMAND_STATUS"
	dGNtoName[0x1FEFF] = "GENERATOR_DEMAND_COMMAND"
	dGNtoName[0x1FEFE] = "AGS_CRITERION_STATUS"
	dGNtoName[0x1FED2] = "AGS_CRITERION_STATUS_2"
	dGNtoName[0x1FEFD] = "AGS_CRITERION_COMMAND"
	dGNtoName[0x1FED5] = "AGS_DEMAND_CONFIGURATION_STATUS"
	dGNtoName[0x1FED4] = "AGS_DEMAND_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEFC] = "FLOOR_HEAT_STATUS"
	dGNtoName[0x1FEFB] = "FLOOR_HEAT_COMMAND"
	dGNtoName[0x1FEF1] = "TIRE_RAW_STATUS"
	dGNtoName[0x1FEF0] = "TIRE_STATUS"
	dGNtoName[0x1FEEF] = "TIRE_SLOW_LEAK_ALARM"
	dGNtoName[0x1FEEE] = "TIRE_TEMPERATURE_CONFIGURATION_STATUS"
	dGNtoName[0x1FEED] = "TIRE_PRESSURE_CONFIGURATION_STATUS"
	dGNtoName[0x1FEEC] = "TIRE_PRESSURE_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEEB] = "TIRE_TEMPERATURE_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEEA] = "TIRE_ID_STATUS"
	dGNtoName[0x1FEE9] = "TIRE_ID_COMMAND"
	dGNtoName[0x1FEF3] = "AWNING_STATUS"
	dGNtoName[0x1FEF2] = "AWNING_COMMAND"
	dGNtoName[0x1FEDE] = "WINDOW_SHADE_CONTROL_STATUS"
	dGNtoName[0x1FEDF] = "WINDOW_SHADE_CONTROL_COMMAND"
	dGNtoName[0x1FEE5] = "LOCK_STATUS"
	dGNtoName[0x1FEE4] = "LOCK_COMMAND"
	dGNtoName[0x1FEE3] = "WINDOW_STATUS"
	dGNtoName[0x1FEE2] = "WINDOW_COMMAND"
	dGNtoName[0x0FEF3] = "GPS_POSITION"
	dGNtoName[0x1FED3] = "GPS_STATUS"
	dGNtoName[0x1FED0] = "DC_DISCONNECT_STATUS"
	dGNtoName[0x1FECF] = "DC_DISCONNECT_COMMAND"
	dGNtoName[0x1FEB3] = "SOLAR_CONTROLLER_STATUS"
	dGNtoName[0x1FEB2] = "SOLAR_CONTROLLER_CONFIGURATION_STATUS"
	dGNtoName[0x1FEB1] = "SOLAR_CONTROLLER_COMMAND"
	dGNtoName[0x1FEB0] = "SOLAR_CONTROLLER_CONFIGURATION_COMMAND"
	dGNtoName[0x1FEAF] = "SOLAR_EQUALIZATION_STATUS"
	dGNtoName[0x1FEAE] = "SOLAR_EQUALIZATION_CONFIGURATION_STATUS"
	dGNtoName[0x1FEAD] = "SOLAR_EQUALIZATION_CONFIGURATION_COMMAND"
	dGNtoName[0x1FED6] = "MFG_SPECIFIC_CLAIM_REQUEST"
	dGNtoName[0x1FECA] = "DM_RV"
	dGNtoName[0x0FECA] = "DM_1"
	dGNtoName[DGN_INITIAL_PACKET] = "INITIAL_PACKET"
	dGNtoName[DGN_DATA_PACKET] = "DATA_PACKET"
	dGNtoName[DGN_ADDRESS_CLAIMED] = "ADDRESS_CLAIMED"
}

// func DGNName(rvcFrame *RvcFrame, DGN uint32) string {
func DGNName(dgn uint32) string {
	var ret = dGNtoName[dgn]

	// special cases
	if ret == "" {

		if (dgn & 0xfff00) == 0x017E00 {
			ret = "TERMINAL??"
		} else if (dgn & 0xfff00) == 0x0E800 {
			ret = "ACK"
		} else if (dgn & 0xfff00) == DGN_INFORMATION_REQUEST {
			ret = "INFORMATION REQUEST"
		} else if (dgn & 0xfff00) == 0x0EE00 {
			ret = "ADDRESS CLAIMED"
			// 1EF00h 1EFxxh proprietary DGN
		} else {
			ret = "unknown"
		}

	}
	return ret

	/*
		switch DGN {
		//case 0x1fffd:
		//	ret = "DC_SOURCE_STATUS_1"
		//	break
		//case 0x1fffc:
		//	ret = "DC_SOURCE_STATUS_2"
		//	break
		//case 0x1ffad:
		//	ret = "ATS_AC_STATUS_1"
		//	break
		//case 0x1ffbf:
		//	ret = "AC_LOAD_STATUS"
		//	break
		case 0x1f9c:
			ret = "THERMOSTAT_AMBIENT_STATUS"
			break
		//case 0x1fecc:
		//	ret = "CHARGER_CONFIGURATION_STATUS_3"
		//	break
		//case 0x1feda:
		//	ret = "DC_DIMMER_STATUS_3"
		//	break
		//case 0x1ffbe:
		//	ret = "AC_LOAD_COMMAND"
		//	break
		//case 0x1ffda:
		//	ret = "GENERATOR_COMMAND"
		//	break
		case 0x1ffd7:
			ret = "INVERTER_AC_STATUS_1"
			break
		case 0x1ffd4:
			ret = "INVERTER_STATUS"
			break
		//case 0x1ffe2:
		//	ret = "THERMOSTAT_STATUS_1"
		//	break
		case 0x17F00:
			ret = "GENERAL_RESET"
			break
		case 0x17E00:
			ret = "TERMINAL"
			break
		case 0x1FED8:
			ret = "GENERIC_CONFIGURATION_STATUS"
			break
		case 0x1FFFF:
			ret = "DATE_TIME_STATUS"
			break
		case 0x1FFFE:
			ret = "SET_DATE_TIME_COMMAND"
			break
		case 0x1FFFD:
			ret = "DC_SOURCE_STATUS_1"
			break
		case 0x1FFFC:
			ret = "DC_SOURCE_STATUS_2"
			break
		case 0x1FFFB:
			ret = "DC_SOURCE_STATUS_3"
			break
		case 0x1FEC9:
			ret = "DC_SOURCE_STATUS_4"
			break
		case 0x1FEC8:
			ret = "DC_SOURCE_STATUS_5"
			break
		case 0x1FEC7:
			ret = "DC_SOURCE_STATUS_6"
			break
		case 0x10FFD:
			ret = "DC_SOURCE_STATUS_SPYDER"
			break
		case 0x1FFFA:
			ret = "COMMUNICATION_STATUS_1"
			break
		case 0x1FFF9:
			ret = "COMMUNICATION_STATUS_2"
			break
		case 0x1FFF8:
			ret = "COMMUNICATION_STATUS_3"
			break
		case 0x1FFF7:
			ret = "WATERHEATER_STATUS"
			break
		case 0x1FFF6:
			ret = "WATERHEATER_COMMAND"
			break
		case 0x1FFF5:
			ret = "GAS_SENSOR_STATUS"
			break
		case 0x1FFF4:
			ret = "CHASSIS_MOBILITY_STATUS"
			break
		case 0x1FFF3:
			ret = "CHASSIS_MOBILITY_COMMAND"
			break
		case 0x1FFF2:
			ret = "AAS_CONFIG_STATUS"
			break
		case 0x1FFF1:
			ret = "AAS_COMMAND"
			break
		case 0x1FFF0:
			ret = "AAS_STATUS"
			break
		case 0x1FFEF:
			ret = "AAS_SENSOR_STATUS"
			break
		case 0x1FED1:
			ret = "SUSPENSION_AIR_PRESSURE_STATUS"
			break
		case 0x1FFEE:
			ret = "LEVELING_CONTROL_COMMAND"
			break
		case 0x1FFED:
			ret = "LEVELING_CONTROL_STATUS"
			break
		case 0x1FFEC:
			ret = "LEVELING_JACK_STATUS"
			break
		case 0x1FFEB:
			ret = "LEVELING_SENSOR_STATUS"
			break
		case 0x1FFEA:
			ret = "HYDRAULIC_PUMP_STATUS"
			break
		case 0x1FEBC:
			ret = "HYDRAULIC_PUMP_COMMAND"
			break
		case 0x1FFE9:
			ret = "LEVELING_AIR_STATUS"
			break
		case 0x1FFE8:
			ret = "SLIDE_STATUS"
			break
		case 0x1FFE7:
			ret = "SLIDE_COMMAND"
			break
		case 0x1FFE6:
			ret = "SLIDE_SENSOR_STATUS"
			break
		case 0x1FFE5:
			ret = "SLIDE_MOTOR_STATUS"
			break
		case 0x1FFE4:
			ret = "FURNACE_STATUS"
			break
		case 0x1FFE3:
			ret = "FURNACE_COMMAND"
			break
		case 0x1FFE2:
			ret = "THERMOSTAT_STATUS_1"
			break
		case 0x1FEFA:
			ret = "THERMOSTAT_STATUS_2"
			break
		case 0x1FEF9:
			ret = "THERMOSTAT_COMMAND_1"
			break
		case 0x1FEF8:
			ret = "THERMOSTAT_COMMAND_2"
			break
		case 0x1FEF7:
			ret = "THERMOSTAT_SCHEDULE_STATUS_1"
			break
		case 0x1FEF6:
			ret = "THERMOSTAT_SCHEDULE_STATUS_2"
			break
		case 0x1FEF5:
			ret = "THERMOSTAT_SCHEDULE_COMMAND_1"
			break
		case 0x1FEF4:
			ret = "THERMOSTAT_SCHEDULE_COMMAND_2"
			break
		case 0x1FF9C:

				//var p = utils.UintParser{ByteOffset: 1}
				//var instance = rvcFrame.Data[0]
				//
				//var tmp = uint16(rvcFrame.Data[2]) << 8
				//tmp = tmp | uint16(rvcFrame.Data[1])
				//
				//var tmp2 uint16 = binary.LittleEndian.Uint16(rvcFrame.Data[1:])
				//var tmp3 = p.ParseInt16(&rvcFrame.Data)
				//fmt.Printf("BY HHAND = %x and tooling = %x  and util= %x   ", tmp, tmp2, tmp3)
				//
				//var temp float64 = float64(tmp)
				//temp = (temp / 32) - 273       // degrees C
				//temp = (temp * 9.0 / 5.0) + 32 // C -> F
				//ret = fmt.Sprintf("THERMOSTAT_AMBIENT_STATUS instance %d tempf = %rvcFrame", instance, temp)

			ret = "THERMOSTAT_AMBIENT_STATUS"
			break
		case 0x1FFE1:
			ret = "AIR_CONDITIONER_STATUS"
			break
		case 0x1FFE0:
			ret = "AIR_CONDITIONER_COMMAND"
			break
		case 0x1FF9B:
			ret = "HEAT_PUMP_STATUS"
			break
		case 0x1FF9A:
			ret = "HEAT_PUMP_COMMAND"
			break
		case 0x1FFDF:
			ret = "GENERATOR_AC_STATUS_1"
			break
		case 0x1FFDE:
			ret = "GENERATOR_AC_STATUS_2"
			break
		case 0x1FFDD:
			ret = "GENERATOR_AC_STATUS_3"
			break
		case 0x1FF94:
			ret = "GENERATOR_AC_STATUS_4"
			break
		case 0x1FEC6:
			ret = "GENERATOR_DC_STATUS_1"
			break
		case 0x1FEC5:
			ret = "GENERATOR_DC_CONFIGURATION_STATUS"
			break
		case 0x1FEC4:
			ret = "GENERATOR_DC_COMMAND"
			break
		case 0x1FEC3:
			ret = "GENERATOR_DC_CONFIGURATION_COMMAND"
			break
		case 0x1FEC2:
			ret = "GENERATOR_DC_EQUALIZATION_STATUS"
			break
		case 0x1FEC1:
			ret = "GENERATOR_DC_EQUALIZATION_CONFIGURATION_STATUS"
			break
		case 0x1FEC0:
			ret = "GENERATOR_DC_EQUALIZATION_CONFIGURATION_COMMAND"
			break
		case 0x1FFDC:
			ret = "GENERATOR_STATUS_1"
			break
		case 0x1FFDB:
			ret = "GENERATOR_STATUS_2"
			break
		case 0x1FFDA:
			ret = "GENERATOR_COMMAND"
			break
		case 0x1FFD9:
			ret = "GENERATOR_START_CONFIG_STATUS"
			break
		case 0x1FFD8:
			ret = "GENERATOR_START_CONFIG_COMMAND"
			break
		//case 0x1FFD7:
		//	ret="INVERTER_AC_STATUS_1"
		//	break;
		case 0x1FFD6:
			ret = "INVERTER_AC_STATUS_2"
			break
		case 0x1FFD5:
			ret = "INVERTER_AC_STATUS_3"
			break
		case 0x1FF8F:
			ret = "INVERTER_AC_STATUS_4"
			break
		case 0x1FF8E:
			ret = "INVERTER_ACFAULT_CONFIGURATION_STATUS_1"
			break
		case 0x1FF8D:
			ret = "INVERTER_ACFAULT_CONFIGURATION_STATUS_2"
			break
		case 0x1FF8C:
			ret = "INVERTER_ACFAULT_CONFIGURATION_COMMAND_1"
			break
		case 0x1FF8B:
			ret = "INVERTER_ACFAULT_CONFIGURATION_COMMAND_2"
			break
		//case 0x1FFD4:
		//	ret="INVERTER_STATUS"
		//	break;
		case 0x1FFD3:
			ret = "INVERTER_COMMAND"
			break
		case 0x1FFD2:
			ret = "INVERTER_CONFIGURATION_STATUS_1"
			break
		case 0x1FFD1:
			ret = "INVERTER_CONFIGURATION_STATUS_2"
			break
		case 0x1FECE:
			ret = "INVERTER_CONFIGURATION_STATUS_3"
			break
		case 0x1FFD0:
			ret = "INVERTER_CONFIGURATION_COMMAND_1"
			break
		case 0x1FFCF:
			ret = "INVERTER_CONFIGURATION_COMMAND_2"
			break
		case 0x1FECD:
			ret = "INVERTER_CONFIGURATION_COMMAND_3"
			break
		case 0x1FFCE:
			ret = "INVERTER_STATISTIC_STATUS"
			break
		case 0x1FFCD:
			ret = "INVERTER_APS_STATUS"
			break
		case 0x1FFCC:
			ret = "INVERTER_DCBUS_STATUS"
			break
		case 0x1FFCB:
			ret = "INVERTER_OPE_STATUS"
			break
		case 0x1FEE8:
			ret = "INVERTER_DC_STATUS"
			break
		case 0x1FEBD:
			ret = "INVERTER_TEMPERATURE_STATUS"
			break
		case 0x1FDCB:
			ret = "INVERTER_TEMPERATURE_STATUS2"
			break
		case 0x1FFCA:
			ret = "CHARGER_AC_STATUS_1"
			break
		case 0x1FFC9:
			ret = "CHARGER_AC_STATUS_2"
			break
		case 0x1FFC8:
			ret = "CHARGER_AC_STATUS_3"
			break
		case 0x1FF8A:
			ret = "CHARGER_AC_STATUS_4"
			break
		case 0x1FF89:
			ret = "CHARGER_ACFAULT_CONFIGURATION_STATUS_1"
			break
		case 0x1FF88:
			ret = "CHARGER_ACFAULT_CONFIGURATION_STATUS_2"
			break
		case 0x1FF87:
			ret = "CHARGER_ACFAULT_CONFIGURATION_COMMAND_1"
			break
		case 0x1FF86:
			ret = "CHARGER_ACFAULT_CONFIGURATION_COMMAND_2"
			break
		case 0x1FFC7:
			ret = "CHARGER_STATUS"
			break
		case 0x1FFC6:
			ret = "CHARGER_CONFIGURATION_STATUS"
			break
		case 0x1FFC5:
			ret = "CHARGER_COMMAND"
			break
		case 0x1FFC4:
			ret = "CHARGER_CONFIGURATION_COMMAND"
			break
		case 0x1FF96:
			ret = "CHARGER_CONFIGURATION_STATUS_2"
			break
		case 0x1FF95:
			ret = "CHARGER_CONFIGURATION_COMMAND_2"
			break
		case 0x1FECC:
			ret = "CHARGER_CONFIGURATION_STATUS_3"
			break
		case 0x1FECB:
			ret = "CHARGER_CONFIGURATION_COMMAND_3"
			break
		case 0x1FEBF:
			ret = "CHARGER_CONFIGURATION_STATUS_4"
			break
		case 0x1FEBE:
			ret = "CHARGER_CONFIGURATION_COMMAND_4"
			break
		case 0x1FF99:
			ret = "CHARGER_EQUALIZATION_STATUS"
			break
		case 0x1FF98:
			ret = "CHARGER_EQUALIZATION_CONFIGURATION_STATUS"
			break
		case 0x1FF97:
			ret = "CHARGER_EQUALIZATION_CONFIGURATION_COMMAND"
			break
		case 0x1FEBB:
			ret = "GENERIC_AC_STATUS_1"
			break
		case 0x1FEBA:
			ret = "GENERIC_AC_STATUS_2"
			break
		case 0x1FEB9:
			ret = "GENERIC_AC_STATUS_3"
			break
		case 0x1FEB8:
			ret = "GENERIC_AC_STATUS_4"
			break
		case 0x1FEB7:
			ret = "GENERIC_ACFAULT_CONFIGURATION_STATUS_1"
			break
		case 0x1FEB6:
			ret = "GENERIC_ACFAULT_CONFIGURATION_STATUS_2"
			break
		case 0x1FEB5:
			ret = "GENERIC_ACFAULT_CONFIGURATION_COMMAND_1"
			break
		case 0x1FEB4:
			ret = "GENERIC_ACFAULT_CONFIGURATION_COMMAND_2"
			break
		case 0x1FFBF:
			ret = "AC_LOAD_STATUS"
			break
		case 0x1FEDD:
			ret = "AC_LOAD_STATUS_2"
			break
		case 0x1FFBE:
			ret = "AC_LOAD_COMMAND"
			break
		case 0x1FFBD:
			ret = "DC_LOAD_STATUS"
			break
		case 0x1FEDC:
			ret = "DC_LOAD_STATUS_2"
			break
		case 0x1FFBC:
			ret = "DC_LOAD_COMMAND"
			break
		case 0x1FFBB:
			ret = "DC_DIMMER_STATUS_1"
			break
		case 0x1FFBA:
			ret = "DC_DIMMER_STATUS_2"
			break
		case 0x1FEDA:
			ret = "DC_DIMMER_STATUS_3"
			break
		case 0x1FFB9:
			ret = "DC_DIMMER_COMMAND"
			break
		case 0x1FEDB:
			ret = "DC_DIMMER_COMMAND_2"
			break
		case 0x1FFB8:
			ret = "DIGITAL_INPUT_STATUS"
			break
		case 0x1FED7:
			ret = "GENERIC_INDICATOR_STATUS"
			break
		case 0x1FED9:
			ret = "GENERIC_INDICATOR_COMMAND"
			break
		case 0x1FEE0:
			ret = "DC_MOTOR_CONTROL_STATUS"
			break
		case 0x1FEE1:
			ret = "DC_MOTOR_CONTROL_COMMAND"
			break
		case 0x1FFB7:
			ret = "TANK_STATUS"
			break
		case 0x1FFB6:
			ret = "TANK_CALIBRATION_COMMAND"
			break
		case 0x1FFB5:
			ret = "TANK_GEOMETRY_STATUS"
			break
		case 0x1FFB4:
			ret = "TANK_GEOMETRY_COMMAND"
			break
		case 0x1FFB3:
			ret = "WATER_PUMP_STATUS"
			break
		case 0x1FFB2:
			ret = "WATER_PUMP_COMMAND"
			break
		case 0x1FFB1:
			ret = "AUTOFILL_STATUS"
			break
		case 0x1FFB0:
			ret = "AUTOFILL_COMMAND"
			break
		case 0x1FFAF:
			ret = "WASTEDUMP_STATUS"
			break
		case 0x1FFAE:
			ret = "WASTEDUMP_COMMAND"
			break
		case 0x1FFAD:
			ret = "ATS_AC_STATUS_1"
			break
		case 0x1FFAC:
			ret = "ATS_AC_STATUS_2"
			break
		case 0x1FFAB:
			ret = "ATS_AC_STATUS_3"
			break
		case 0x1FF85:
			ret = "ATS_AC_STATUS_4"
			break
		case 0x1FF84:
			ret = "ATS_ACFAULT_CONFIGURATION_STATUS_1"
			break
		case 0x1FF83:
			ret = "ATS_ACFAULT_CONFIGURATION_STATUS_2"
			break
		case 0x1FF82:
			ret = "ATS_ACFAULT_CONFIGURATION_COMMAND_1"
			break
		case 0x1FF81:
			ret = "ATS_ACFAULT_CONFIGURATION_COMMAND_2"
			break
		case 0x1FFAA:
			ret = "ATS_STATUS"
			break
		case 0x1FFA9:
			ret = "ATS_COMMAND"
			break
		case 0x1FFA5:
			ret = "WEATHER_STATUS_1"
			break
		case 0x1FFA4:
			ret = "WEATHER_STATUS_2"
			break
		case 0x1FFA3:
			ret = "ALTIMETER_STATUS"
			break
		case 0x1FFA2:
			ret = "ALTIMETER_COMMAND"
			break
		case 0x1FFA1:
			ret = "WEATHER_CALIBRATE_COMMAND"
			break
		case 0x1FFA0:
			ret = "COMPASS_BEARING_STATUS"
			break
		case 0x1FF9F:
			ret = "COMPASS_CALIBRATE_COMMAND"
			break
		case 0x1FF80:
			ret = "GENERATOR_DEMAND_STATUS"
			break
		case 0x1FEFF:
			ret = "GENERATOR_DEMAND_COMMAND"
			break
		case 0x1FEFE:
			ret = "AGS_CRITERION_STATUS"
			break
		case 0x1FED2:
			ret = "AGS_CRITERION_STATUS_2"
			break
		case 0x1FEFD:
			ret = "AGS_CRITERION_COMMAND"
			break
		case 0x1FED5:
			ret = "AGS_DEMAND_CONFIGURATION_STATUS"
			break
		case 0x1FED4:
			ret = "AGS_DEMAND_CONFIGURATION_COMMAND"
			break
		case 0x1FEFC:
			ret = "FLOOR_HEAT_STATUS"
			break
		case 0x1FEFB:
			ret = "FLOOR_HEAT_COMMAND"
			break
		case 0x1FEF1:
			ret = "TIRE_RAW_STATUS"
			break
		case 0x1FEF0:
			ret = "TIRE_STATUS"
			break
		case 0x1FEEF:
			ret = "TIRE_SLOW_LEAK_ALARM"
			break
		case 0x1FEEE:
			ret = "TIRE_TEMPERATURE_CONFIGURATION_STATUS"
			break
		case 0x1FEED:
			ret = "TIRE_PRESSURE_CONFIGURATION_STATUS"
			break
		case 0x1FEEC:
			ret = "TIRE_PRESSURE_CONFIGURATION_COMMAND"
			break
		case 0x1FEEB:
			ret = "TIRE_TEMPERATURE_CONFIGURATION_COMMAND"
			break
		case 0x1FEEA:
			ret = "TIRE_ID_STATUS"
			break
		case 0x1FEE9:
			ret = "TIRE_ID_COMMAND"
			break
		case 0x1FEF3:
			ret = "AWNING_STATUS"
			break
		case 0x1FEF2:
			ret = "AWNING_COMMAND"
			break
		case 0x1FEDE:
			ret = "WINDOW_SHADE_CONTROL_STATUS"
			break
		case 0x1FEDF:
			ret = "WINDOW_SHADE_CONTROL_COMMAND"
			break
		case 0x1FEE5:
			ret = "LOCK_STATUS"
			break
		case 0x1FEE4:
			ret = "LOCK_COMMAND"
			break
		case 0x1FEE3:
			ret = "WINDOW_STATUS"
			break
		case 0x1FEE2:
			ret = "WINDOW_COMMAND"
			break
		case 0x0FEF3:
			ret = "GPS_POSITION"
			break
		case 0x1FED3:
			ret = "GPS_STATUS"
			break
		case 0x1FED0:
			ret = "DC_DISCONNECT_STATUS"
			break
		case 0x1FECF:
			ret = "DC_DISCONNECT_COMMAND"
			break
		case 0x1FEB3:
			ret = "SOLAR_CONTROLLER_STATUS"
			break
		case 0x1FEB2:
			ret = "SOLAR_CONTROLLER_CONFIGURATION_STATUS"
			break
		case 0x1FEB1:
			ret = "SOLAR_CONTROLLER_COMMAND"
			break
		case 0x1FEB0:
			ret = "SOLAR_CONTROLLER_CONFIGURATION_COMMAND"
			break
		case 0x1FEAF:
			ret = "SOLAR_EQUALIZATION_STATUS"
			break
		case 0x1FEAE:
			ret = "SOLAR_EQUALIZATION_CONFIGURATION_STATUS"
			break
		case 0x1FEAD:
			ret = "SOLAR_EQUALIZATION_CONFIGURATION_COMMAND"
			break
		case 0x1FED6:
			ret = "MFG_SPECIFIC_CLAIM_REQUEST"
			break
		case 0x1FECA:
			ret = "DM_RV"
			break
		case 0x0FECA:
			ret = "DM_1"
			break

		default:
			// 0x17E00
			if (DGN & 0xfff00) == 0x017E00 {
				ret = "TERMINAL??"
			} else if (DGN & 0xfff00) == 0x0E800 {
				ret = "ACK"
			} else if (DGN & 0xfff00) == 0x0EA00 {
				ret = "DGN REQ"

				// 1EF00h 1EFxxh proprietary DGN
			} else {
				ret = "unknown"
			}

		}

		//fmt.Printf("DGN NAMES MAP RETURNING: %s\n", ret)
		return ret
	*/
}
