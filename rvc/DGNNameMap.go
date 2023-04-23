package rvc

// func DGNName(rvcFrame *RvcFrame, dgn uint32) string {
func DGNName(dgn uint32) string {
	var ret string

	switch dgn {
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
		/*
			// A SIMPLE TEST
			// byte 0 instance
			// byte 1 0-3 operating mode,4-5 fan mode, 6-7 schedule mode
			// byte 2 fan speed
			// 3-4 heat setpoint
			// 5-6cool setpoint
			var instance uint8 = rvcFrame.Data[0]
			var opmode uint8 = rvcFrame.Data[1] & 0x0f
			var opmodeS string
			switch opmode {
			case 0:
				opmodeS = "Off"
				break
			case 1:
				opmodeS = "Cool"
				break
			case 3:
				opmodeS = "Auto Heat/Cool"
				break
			case 4:
				opmodeS = "Fan Only"
				break
			case 5:
				opmodeS = "Aux Heat"
				break
			case 6:
				opmodeS = ""
				break
			default:
				opmodeS = "huh?"

			}
			//var heatSP uint16 = binary.LittleEndian.Uint16(rvcFrame.Data[3:])
			//var heatSPFP float64 = float64(heatSP)
			//heatSPFP = (heatSPFP / 32) - 273       // degrees C
			//heatSPFP = (heatSPFP * 9.0 / 5.0) + 32 // C -> F
			//
			//var coolSP uint16 = binary.LittleEndian.Uint16(rvcFrame.Data[5:])
			//var coolSPFP float64 = float64(coolSP)
			//coolSPFP = (coolSPFP / 32) - 273       // degrees C
			//coolSPFP = (coolSPFP * 9.0 / 5.0) + 32 // C -> F

			ret = fmt.Sprintf("THERMOSTAT_STATUS_1 xx instance: %d opModel: %s",
				instance, opmodeS)
		*/
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
		/*
			var p = utils.UintParser{ByteOffset: 1}
			var instance = rvcFrame.Data[0]

			var tmp = uint16(rvcFrame.Data[2]) << 8
			tmp = tmp | uint16(rvcFrame.Data[1])

			var tmp2 uint16 = binary.LittleEndian.Uint16(rvcFrame.Data[1:])
			var tmp3 = p.ParseInt16(&rvcFrame.Data)
			fmt.Printf("BY HHAND = %x and tooling = %x  and util= %x   ", tmp, tmp2, tmp3)

			var temp float64 = float64(tmp)
			temp = (temp / 32) - 273       // degrees C
			temp = (temp * 9.0 / 5.0) + 32 // C -> F
			ret = fmt.Sprintf("THERMOSTAT_AMBIENT_STATUS instance %d tempf = %rvcFrame", instance, temp)
		*/
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
		if (dgn & 0xfff00) == 0x017E00 {
			ret = "TERMINAL??"
		} else if (dgn & 0xfff00) == 0x0E800 {
			ret = "ACK"
		} else if (dgn & 0xfff00) == 0x0EA00 {
			ret = "DGN REQ"

			// 1EF00h 1EFxxh proprietary DGN
		} else {
			ret = "unknown"
		}

	}

	//fmt.Printf("DGN NAMES MAP RETURNING: %s\n", ret)
	return ret

}
