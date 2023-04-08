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

func GetCANMessages(fromSocket, toSocket chan *rvccan.RawCanMessage) {
	fmt.Printf("IMPLEMENT ME\n")
	var m = &rvccan.RawCanMessage{}
	for {
		m.Timestamp = time.Now()
		m.CanMessage[0] = 1

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
