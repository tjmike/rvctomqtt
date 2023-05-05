package addrClaim

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

	for {
		var dmr = rvc.AddressClaim{}

		dmr.SetDesiredSourceAddress(222)
		dmr.SetPriority(0x06)
		dmr.SetDGN(rvc.DGN_ADDRESS_CLAIM)
		var f = dmr.CreateFrame()
		var cif intf.CanFrameIF = intf.CanFrameIF(f)
		msgChan <- &cif
		print("Sleep\n")
		time.Sleep(time.Second * 10)
	}

}
