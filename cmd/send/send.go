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

	for {

		// Make a light command request
		//var dcdimmerCommand = rvc.DCDimmerCommand2{}

		// TODO should start from a command and have it automatically generate the frame

		//RAW FRAME: 04-30-2023 18:44:07.330792 99 db fe 99 08 00 00 00 0e ff c8 03 ff 00 ff ff
		//04-30-2023 18:44:07.330792 DGN: 1fedb (DC_DIMMER_COMMAND_2) SA: 153 Instance: 14 (Passenger Task) group: 255 brigntness: 100.000000 command: 3 lockitem: 0 res1: 0 res2: 0 res3: 0 rampTime 25.500000 reserved 255

		var dmr = rvc.DCDimmerCommand2{}

		// TODO - this "device" should probably get a source address
		dmr.SetSourceAddress(153)

		// TODO -  need command constants
		dmr.SetCommand(5)

		dmr.SetDesiredBringhtness(100)
		dmr.SetGroup(rvc.RVC_DATA_NOT_AVAILABLE)
		dmr.SetDesiredBringhtness(100)
		dmr.SetInstance(rvc.INSTANCE_LIGHT_PASSENGER_TASK)

		dmr.SetPriority(0x06)
		dmr.SetDGN(rvc.DGN_DC_DIMMER_COMMAND_2)
		dmr.SetDelayDuration(rvc.RVC_DATA_NOT_AVAILABLE)
		dmr.SetReserved(rvc.RVC_DATA_NOT_AVAILABLE)
		dmr.SetRampTime(25.5)

		var f = dmr.CreateFrame()

		var cif intf.CanFrameIF = intf.CanFrameIF(f)

		msgChan <- &cif
		print("Sleep\n")
		time.Sleep(time.Second * 10)
	}

}
