package pform

import (
	//"github.com/golang-collections/collections/stack"
	//"net"
	//"os"
	//"syscall"
	//"time"
	"fmt"
	rvccan "rvctomqtt/can"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"time"
)

func GetCANMessages(messagePool *pool.Pool, fromSocket, toSocket chan *intf.CanFrameIF) {
	fmt.Printf("IMPLEMENT ME\n")
	//XXX var m = &rvccan.Frame{}
	var mm intf.CanFrameIF = &rvccan.Frame{}
	for {
		mm.SetTimeStamp(time.Now())
		//m.Timestamp = time.Now()
		//m.MessageBytes[0] = 1

		//fromSocket <- m
		fromSocket <- &mm

		time.Sleep(time.Second)
	bufloop:
		for {
			select {
			case item := <-toSocket:
				{
					mm = *item
					fmt.Printf("FOUND\n")
				}
			default:
				break bufloop
			}

		}

	}

}
