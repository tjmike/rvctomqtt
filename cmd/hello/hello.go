package main

import (
	"fmt"
	"runtime"
	"rvctomqtt/handler"
	"rvctomqtt/intf"
	"rvctomqtt/pform"
	"rvctomqtt/pool"
	"rvctomqtt/rvc"
	"time"
)

func main() {
	print("MAX PROCS=")
	print(runtime.GOMAXPROCS(0))
	print("\n")

	// Listen on this to process the raw can message
	fromSocket := make(chan *intf.CanFrameIF, 32)

	// seems like we must be explicit with the interface - we can't pass the item
	// that implements it
	//var frameFactoryInterface intf.CanFrameFactory = &can.CanFrameFactory{}
	var frameFactoryInterface intf.CanFrameFactory = &rvc.RVCFrameFactory{}
	var p = pool.NewPool(&frameFactoryInterface)

	go pform.GetRVCMessages(p, fromSocket)
	go handler.RVCMessageHandler(fromSocket, p)

	for {
		fmt.Printf("Sleep # goRoutines = %d\n", runtime.NumGoroutine())

		time.Sleep(time.Second * 30)
	}

}
