package main

import (
	"rvctomqtt/intf"
	"rvctomqtt/pform"
	"rvctomqtt/rvc"

	"time"
)

func main() {

	// Listen on this to process the raw can message
	msgChan := make(chan *intf.CanFrameIF, 32)

	// set tp the sender
	go pform.SendCANMessages(msgChan)

	var desired uint8 = 250

	for {
		var dmr = rvc.InformationRequest{}

		dmr.SetPriority(0x06)
		dmr.SetDGN(rvc.DGN_INFORMATION_REQUEST)
		dmr.SetAsAddressClaim(desired)
		var f = dmr.CreateFrame()
		var cif intf.CanFrameIF = intf.CanFrameIF(f)
		msgChan <- &cif
		print("Sleep\n")
		time.Sleep(time.Second * 1)
		desired--
		if desired < 1 {
			break
		}
	}

}
