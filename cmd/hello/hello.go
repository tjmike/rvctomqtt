package main

import (
	"runtime"
	"rvctomqtt/handler"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"rvctomqtt/rvc"
	"time"

	"rvctomqtt/pform"
)

func main() {
	print("MAX PROCS=")
	print(runtime.GOMAXPROCS(0))
	print("\n")

	// Listen on this to process the raw can message
	fromSocket := make(chan *intf.CanFrameIF, 32)

	// When done with the message - give it back to the socket listener
	toSocket := make(chan *intf.CanFrameIF, 32)

	// seems like we must be explicit with the interface - we can't pass the item
	// that implements it
	//var frameFactoryInterface intf.CanFrameFactory = &can.CanFrameFactory{}
	var frameFactoryInterface intf.CanFrameFactory = &rvc.RVCFrameFactory{}

	var p = pool.NewPool(&frameFactoryInterface, 10)

	go pform.GetRVCMessages(p, fromSocket, toSocket)
	go handler.CanMessageHandler(fromSocket, toSocket)

	for {
		print("Sleep\n")
		time.Sleep(time.Second * 30)
	}

}
