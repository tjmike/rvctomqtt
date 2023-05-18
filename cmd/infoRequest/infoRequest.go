package main

import (
	"rvctomqtt/intf"
	"rvctomqtt/pform"
	"rvctomqtt/rvc"

	"time"
)

func main() {

	var targets [23]uint8 = [23]uint8{
		0xb8,
		0x9f,
		0x9e,
		0x9d,
		0x9c,
		0x9b,
		0x9a,
		0x99,
		0x98,
		0x97,
		0x96,
		0x95,
		0x94,
		0x93,
		0x92,
		0x91,
		0x90,
		0x8f,
		0x8e,
		0x8d,
		0x8c,
		0x8b,
		0x42,
	}

	// Listen on this to process the raw can message
	msgChan := make(chan *intf.CanFrameIF, 32)

	// set tp the sender
	go pform.SendCANMessages(msgChan)

	//var destAddress uint8 = 0xff
	var myAddres uint8 = 222
	var dgnRequest = rvc.DGN_PRODUCT_IDENTIFICATION_MESSAGE

	for _, destAddress := range targets {

		var dmr = rvc.InformationRequest{}

		dmr.SetPriority(0x06)

		dmr.SetDGN(rvc.DGN_INFORMATION_REQUEST)
		dmr.SetDestinationAddress(destAddress)
		dmr.SetDesiredDGN(dgnRequest)
		dmr.SetInstance(0xff)
		dmr.SetInstance2(0xff)
		dmr.SetSourceAddress(myAddres)
		var f = dmr.CreateFrame()
		var cif intf.CanFrameIF = intf.CanFrameIF(f)
		msgChan <- &cif
		print("Sleep\n")
		time.Sleep(time.Second * 1)
		//desired--
		//if desired < 1 {
		//	break
		//}
	}

}
