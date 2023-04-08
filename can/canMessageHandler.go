package can

import "fmt"

func CanMessageHandler(fromSocket, toSocket chan *Frame) {
	fmt.Printf("############################### HANDLER #####################\n")
	var nmsg uint32 = 0
	for {
		data := <-fromSocket
		toSocket <- data
		nmsg++
	}
}
