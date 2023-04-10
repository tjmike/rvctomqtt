package can

import (
	"fmt"
	"rvctomqtt/intf"
)

func CanMessageHandler(fromSocket, toSocket chan *intf.CanFrameIF) {
	fmt.Printf("############################### HANDLER #####################\n")
	var nmsg uint32 = 0
	for {
		data := <-fromSocket
		fmt.Println((*data).ToString())
		toSocket <- data
		nmsg++
	}
}
