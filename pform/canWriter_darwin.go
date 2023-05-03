package pform

import (
	"fmt"
	"rvctomqtt/intf"
)

func SendCANMessages(message chan *intf.CanFrameIF) {

	// Forever
	for {

		//var dcd = rvc.DCDimmerCommand2{}

		// get the request to send
		var canFrame *intf.CanFrameIF = <-message
		// get the raw byte buffer
		var bytes = (*canFrame).GetMessage()

		// The "send" is a simple prinnt
		fmt.Printf("SENDING %x\n", bytes)
	}
}
