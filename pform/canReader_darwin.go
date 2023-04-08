package pform

import (
	//"github.com/golang-collections/collections/stack"
	//"net"
	//"os"
	//"syscall"
	//"time"
	"fmt"
	rvccan "rvctomqtt/can"
	"time"
)

func GetCANMessages(fromSocket, toSocket chan *rvccan.Frame) {
	fmt.Printf("IMPLEMENT ME\n")
	var m = &rvccan.Frame{}
	for {
		m.Timestamp = time.Now()
		m.MessageBytes[0] = 1

		fromSocket <- m

		time.Sleep(time.Second)
	bufloop:
		for {
			select {
			case item := <-toSocket:
				{
					m = item
					fmt.Printf("FOUND\n")
				}
			default:
				break bufloop
			}

		}

	}

}
