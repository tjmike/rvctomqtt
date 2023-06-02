package main

import (
	"rvctomqtt/intf"
	"rvctomqtt/pform"
	"rvctomqtt/rvc"
	"time"
)

func main() {

	// TODO we need the notion of a device and deviceID at some point - we might be able to get away with doing this later....

	// Listen on this to process the raw can message
	msgChan := make(chan *intf.CanFrameIF, 32)

	// set tp the sender
	go pform.SendCANMessages(msgChan)

	var desiredsa uint8 = 253

	var targets = []uint8{
		rvc.INSTANCE_LIGHT_PASSENGER_TASK,
		rvc.INSTANCE_BEDROOM_CEILING,
	}

	var idx = 0
	var sz = len(targets)

	for {
		var which = idx % sz
		idx++
		// Make a light command request
		//var dcdimmerCommand = rvc.DCDimmerCommand2{}

		// TODO should start from a command and have it automatically generate the frame

		//RAW FRAME: 04-30-2023 18:44:07.330792 99 db fe 99 08 00 00 00 0e ff c8 03 ff 00 ff ff
		//04-30-2023 18:44:07.330792 DGN: 1fedb (DC_DIMMER_COMMAND_2) SA: 153 Instance: 14 (Passenger Task) group: 255 brigntness: 100.000000 command: 3 lockitem: 0 res1: 0 res2: 0 res3: 0 rampTime 25.500000 reserved 255

		if desiredsa == 253 {
			cif := createDCDimmerFrame(targets[which])
			msgChan <- &cif
		}
		time.Sleep(time.Millisecond * 100)
		//{
		//	cif := createReqAddFrame(desiredsa)
		//	msgChan <- &cif
		//}
		print("Sleep\n")
		time.Sleep(time.Second * 1)

		//desiredsa -= 1
		//if desiredsa == 1 {
		//	break
		//}
	}

}

func createDCDimmerFrame(inst uint8) intf.CanFrameIF {
	var dmr = rvc.DCDimmerCommand2{}

	// TODO - this "device" should probably negotiate  a source address
	dmr.SetSourceAddress(153)

	// TODO -  need command constants
	dmr.SetCommand(5)

	dmr.SetDesiredBringhtness(100)
	dmr.SetGroup(rvc.RVC_DATA_NOT_AVAILABLE)
	dmr.SetDesiredBringhtness(100)
	dmr.SetInstance(inst)

	dmr.SetPriority(0x06)
	dmr.SetDGN(rvc.DGN_DC_DIMMER_COMMAND_2)
	dmr.SetDelayDuration(rvc.RVC_DATA_NOT_AVAILABLE)
	dmr.SetReserved(rvc.RVC_DATA_NOT_AVAILABLE)
	dmr.SetRampTime(25.5)

	var f = dmr.CreateFrame()

	var cif intf.CanFrameIF = intf.CanFrameIF(f)
	return cif
}

func createReqAddFrame(desiredSA uint8) intf.CanFrameIF {
	var dmr = rvc.InformationRequest{}
	dmr.SetAsAddressClaim(desiredSA)
	dmr.SetDGN(rvc.DGN_INFORMATION_REQUEST)
	dmr.SetPriority(6)
	var f = dmr.CreateFrame()
	var cif intf.CanFrameIF = intf.CanFrameIF(f)
	return cif
}
