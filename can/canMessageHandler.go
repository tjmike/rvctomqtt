package can

import (
	"fmt"
	"rvctomqtt/intf"
)

func CanMessageHandler(fromSocket, toSocket chan *intf.CanThing) {
	fmt.Printf("############################### HANDLER #####################\n")
	var nmsg uint32 = 0
	for {
		data := <-fromSocket
		toSocket <- data
		nmsg++
	}
}
