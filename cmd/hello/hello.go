package main

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"rvctomqtt/can"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	svc "rvctomqtt/reg"
	"time"
	"unsafe"

	//"github.com/tjmike/rvctomqtt/internal/hello"
	//"github.com/tjmike/rvctomqtt/pform"
	//"github.com/tjmike/rvctomqtt/reg"
	"rvctomqtt/internal/hello"
	"rvctomqtt/pform"
)

func main() {

	fmt.Println(hello.Hello())

	//svc.single ss = svc.GetInstance()
	var ss *svc.Single = svc.GetInstance()
	//svc.Single foo  = &Single{};

	fmt.Printf("%#v\n", ss)

	pform.Doit()

	var i uint32 = 0x01020304
	u := unsafe.Pointer(&i)
	//var foo = [4]byte;

	pb := (*([4]uint8))(u)

	//b := (*pb)[0]
	fmt.Printf("\n0=%x \n", (*pb)[0])
	fmt.Printf("\n1=%x \n", (*pb)[1])
	fmt.Printf("\n2=%x \n", (*pb)[2])
	fmt.Printf("\n3=%x \n", (*pb)[3])

	var xx = binary.LittleEndian.Uint32((*pb)[0:])
	fmt.Printf("\nrebuild=%x \n", xx)

	var little = make([]byte, 4)
	var big = make([]byte, 4)
	binary.LittleEndian.PutUint32(little, i)
	binary.BigEndian.PutUint32(big, i)

	for i, v := range little {
		fmt.Printf("Little %d = %d\n", i, v)
	}
	for i, v := range big {
		fmt.Printf("Big %d = %d\n", i, v)
	}

	print("MAX PROCS=")
	print(runtime.GOMAXPROCS(0))
	print("\n")

	// Listen on this to process the raw can message
	fromSocket := make(chan *intf.CanThing, 32)

	// When done with the message - give it back to the socket listener
	toSocket := make(chan *intf.CanThing, 32)

	var ff = &can.CanFrameFactory{}
	var zzz = ff.Create()
	(*zzz).GetMessage()

	var aa = &can.MyCanThingFactory{}
	var abc intf.CanThingFactory = aa

	var p = pool.NewPool(&abc, 10)

	go pform.GetRVCMessages(p, fromSocket, toSocket)
	go can.CanMessageHandler(fromSocket, toSocket)

	for {
		print("Sleep\n")
		time.Sleep(time.Second * 30)
	}
	//return (b == 0x04)

}
