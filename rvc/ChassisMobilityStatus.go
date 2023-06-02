package rvc

import (
	"encoding/json"
	"fmt"
	"rvctomqtt/constants"
	"rvctomqtt/utils"
)

const TRANS_NETURAL uint8 = 125
const UNKNOWN_ENGINERPM = 8191.875
const UNKNOWN_SPEED = 255.99609375

// 0x1FFF4
type ChassisMobilityStatus struct {
	//rvcItem             RvcItem
	RvcItem
	//Timestamp         time.Time // last Timestamp that we saw a command
	//LastChanged       time.Time // last time this instance actually changed
	engineRPM                float64         // 0-1 uint16 (8191.875 = 65535 = UNKNOWN)
	speed                    float64         // 2-3 uint16 Precision 1/256 (255.99609375 = Unknown)
	parkBrakeStatus          constants.Uint2 //  4 (0-1) 00 released, 01 engaged
	transmissionLockStatus   constants.Uint2 //  4 (2-3) 00 not locked, 01 locked - will not go into gear
	engineLockStatus         constants.Uint2 //  4 (4-5) 00 free to start, 01 locked - will not start
	ignitionSwitchStatus     constants.Uint2 //  5 (0-1) 00 off 01 0n
	accessorySwitchStatus    constants.Uint2 //  5 (2-3) 00 off 01 on NOTE: Spec says bits 2-4 - assume thats wrong
	transmissionCurrentGear  uint8           // 124 = Reverse 1 125 = Neutral  126 = Forward 1 127 = Forward 2  etc.  251 = Park (Note: Most heavy transmissions do  not have a 'Park', but are parked in Neutral.) (255->130 = unavailable/unknown)
	transmissionSelectedGear uint8           // see current gear
	isTiffinFlavor           bool            // is this a "tiffin" flavor that does not implement the spec corectly
}

func (r *ChassisMobilityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DGN uint32 `json:"DGN"`
	}{
		DGN: r.GetDGN(),
	})
}

func (i *ChassisMobilityStatus) Equals(o *ChassisMobilityStatus) bool {

	if !i.RvcItem.Equals(&o.RvcItem) {
		return false
	}

	if i.engineRPM != o.engineRPM {
		return false
	}
	if i.speed != o.speed {
		return false
	}
	if i.parkBrakeStatus != o.parkBrakeStatus {
		return false
	}
	if i.transmissionLockStatus != o.transmissionLockStatus {
		return false
	}
	if i.engineLockStatus != o.engineLockStatus {
		return false
	}
	if i.ignitionSwitchStatus != o.ignitionSwitchStatus {
		return false
	}
	if i.accessorySwitchStatus != o.accessorySwitchStatus {
		return false
	}
	if i.transmissionCurrentGear != o.transmissionCurrentGear {
		return false
	}
	if i.transmissionSelectedGear != o.transmissionSelectedGear {
		return false
	}

	return true
}

//func (i *DCDimmerCommand2) GetDGN() uint32 {
//	return DGN_DC_DIMMER_COMMAND_2
//}
//func (i *DCDimmerCommand2) GetTimestamp() time.Time {
//	return i.Timestamp
//}

func (i *ChassisMobilityStatus) SetDGN(dgn uint32) {
	i.DGN = dgn
	i.Name = DGNName(dgn)
}

func (i *ChassisMobilityStatus) SetSourceAddress(sa uint8) {
	i.SourceAddress = sa
}

func (i *ChassisMobilityStatus) SetPriority(p uint8) {
	i.Priority = p
}

func (i *ChassisMobilityStatus) GetInstanceName() string {
	return ""
}
func (i *ChassisMobilityStatus) String() string {
	//var s = i.RvcItem.String()

	return fmt.Sprintf("TS: %s DGN: %x (%s) SA: %d rpm: %f speed: %f parkBrake: %d transmissionLockStatus: %d engineLockStatus: %d ignitionSwitchStatus: %d accesorySwitchStatus: %d transmissionCurrentGear %d transmissionSelectedGear %d",
		i.GetTimestamp().Format("01-02-2006 15:04:05.000000"),
		i.GetDGN(), DGNName(i.GetDGN()),
		i.GetSourceAddress(),
		i.engineRPM,
		i.speed,
		i.parkBrakeStatus,
		i.transmissionLockStatus,
		i.engineLockStatus,
		i.ignitionSwitchStatus,
		i.accessorySwitchStatus,
		(TRANS_NETURAL - i.transmissionCurrentGear),
		(TRANS_NETURAL - i.transmissionSelectedGear),
	)
	//return "ZZZXXX" + s + " " + t

}

func (r *ChassisMobilityStatus) GetEngineRPM() float64 {
	return r.engineRPM
}
func (r *ChassisMobilityStatus) GetSpeed() float64 {
	return r.speed
}
func (r *ChassisMobilityStatus) GetParkBrakeStatus() constants.Uint2 {
	return r.parkBrakeStatus
}

func (r *ChassisMobilityStatus) GetTransmissionLockStatus() constants.Uint2 {
	return r.transmissionLockStatus
}
func (r *ChassisMobilityStatus) GetEngineLockStatus() constants.Uint2 {
	return r.engineLockStatus
}

func (r *ChassisMobilityStatus) GetIgnitionSwitchStatus() constants.Uint2 {
	return r.ignitionSwitchStatus
}

func (r *ChassisMobilityStatus) GetTransmissionCurrentGear() uint8 {
	return r.transmissionCurrentGear - TRANS_NETURAL
}

func (r *ChassisMobilityStatus) GetTransmissionSelectedGear() uint8 {
	return r.transmissionSelectedGear - TRANS_NETURAL
}
func (r *ChassisMobilityStatus) GetAccessorySwitchStatus() constants.Uint2 {
	return r.accessorySwitchStatus
}
func (r *ChassisMobilityStatus) GetSourceAddress() uint8 {
	return r.SourceAddress
}

const TIFFIN_IGNITION_STATUS uint8 = 0xf1

const TIFFIN_PARKING_STATUS uint8 = 0xf2

func (r *ChassisMobilityStatus) Init(from *RvcFrame) {
	r.RvcItem.Init(from)

	// It seems as if Tiffin(2018) diverges from the spec - or there was some older version of the spec.
	// With the Tiffin flavor byte zero is a flag (F1=ignition status, or F2=parking brake status)
	// As far as I can tell no other fields are set. This brings up the issue of how to deal with divergences
	// from the spec. So far the code assumes a well adhered spec....
	// Tiffin Data:
	//f2 00 00 00 01 00 00 00
	//f1 00 00 00 00 01 00 00

	//
	//TODO - think about this....If that turns out to not be the
	// case we may need to move to a more generic approach where everything is defined in config files.
	// This is super flexible and cool - but it seemed as if having a firmer model would be better - eg
	// the struct is the doc for the particular dgn. In the more generic case I guess the external def
	// could be made to be just as good. The generic case -- I think -- would have more abstraction and be harder
	// to debug - it could be the complete opposite though less code in the generic case - easier to test and once
	// it's working just add a new data definition. The cached data struct could just be the byte array and we could
	// check for changes by just inspecting the raw bytes. getSpeed() might then be changed to get("speed") or
	// something like that.

	var dataBytes = &from.Data
	var changed = false

	if r.isTiffinFlavor {
		r.engineRPM = UNKNOWN_ENGINERPM
		r.speed = UNKNOWN_SPEED
		r.engineLockStatus = constants.Uint2NA
		r.transmissionLockStatus = constants.Uint2NA
		r.accessorySwitchStatus = constants.Uint2NA
		r.transmissionCurrentGear = constants.DataNotAvailableUint8
		r.transmissionSelectedGear = constants.DataNotAvailableUint8
		// Read the 1st byte
		var tmp = utils.GetByte(dataBytes, 0)

		if tmp == TIFFIN_IGNITION_STATUS {
			var tmpStatus = constants.Uint2(utils.GetBits(dataBytes, 5, 0, utils.GetMask(2)))
			if r.ignitionSwitchStatus != tmpStatus {
				r.ignitionSwitchStatus = tmpStatus
				changed = true
			}

		} else if tmp == TIFFIN_PARKING_STATUS {
			var tmpStatus = constants.Uint2(utils.GetBits(dataBytes, 4, 0, utils.GetMask(2)))
			if r.parkBrakeStatus != tmpStatus {
				r.parkBrakeStatus = tmpStatus
				changed = true
			}
		} else {
			// May need to emit an error OR switch to spec based mode
		}

	} else {
		{
			var tmp = float64(utils.Getuint16(dataBytes, 0))
			tmp = tmp / 8
			if (*r).engineRPM != tmp {
				changed = true
			}
			(*r).engineRPM = tmp

			tmp = float64(utils.Getuint16(dataBytes, 2))
			tmp = tmp / 256.0 // KPH
			tmp = constants.KMPHtoMPH(tmp)
			if tmp != (*r).speed {
				changed = true
			}
			(*r).speed = tmp
		}

		{

			var tmp = constants.Uint2(utils.GetBits(dataBytes, 4, 0, utils.GetMask(2)))
			if tmp != (*r).parkBrakeStatus {
				changed = true
			}
			(*r).parkBrakeStatus = tmp

			tmp = constants.Uint2(utils.GetBits(dataBytes, 4, 2, utils.GetMask(2)))
			if tmp != (*r).transmissionLockStatus {
				changed = true
			}
			(*r).transmissionLockStatus = tmp

			tmp = constants.Uint2(utils.GetBits(dataBytes, 4, 4, utils.GetMask(2)))
			if tmp != (*r).engineLockStatus {
				changed = true
			}
			(*r).engineLockStatus = tmp

			tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 0, utils.GetMask(2)))
			if tmp != (*r).ignitionSwitchStatus {
				changed = true
			}
			(*r).ignitionSwitchStatus = tmp

			tmp = constants.Uint2(utils.GetBits(dataBytes, 5, 2, utils.GetMask(2)))
			if tmp != (*r).accessorySwitchStatus {
				changed = true
			}
			(*r).accessorySwitchStatus = tmp

		}
		{
			var tmp = utils.GetByte(dataBytes, 6)
			if tmp != (*r).transmissionCurrentGear {
				changed = true
			}
			(*r).transmissionCurrentGear = tmp

			tmp = utils.GetByte(dataBytes, 7)
			if tmp != (*r).transmissionSelectedGear {
				changed = true
			}
			(*r).transmissionSelectedGear = tmp
		}
	}
	if changed {
		r.LastChanged = r.Timestamp
	}

}
