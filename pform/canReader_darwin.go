package pform

import (
	//"github.com/golang-collections/collections/stack"
	//"net"
	//"os"
	//"syscall"
	//"time"
	"fmt"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"time"
)

func GetCANMessages(messagePool *pool.Pool, fromSocket, toSocket chan *intf.CanFrameIF) {
	fmt.Printf("IMPLEMENT ME\n")
	//XXX var m = &rvccan.Frame{}

	//var mm intf.CanFrameIF = &rvccan.Frame{}
	for {
		{
			var f = messagePool.Get()
			(*f).SetTimeStamp(time.Now())
			(*f).GetMessage()[0] = 1
			(*f).GetMessage()[1] = 2
			(*f).GetMessage()[2] = 4
			(*f).GetMessage()[3] = 8
			(*f).GetMessage()[4] = 8
			(*f).BuildCanFrameX()
			//m.Timestamp = time.Now()
			//m.MessageBytes[0] = 1
			// we've read the frame and are now sending to the channel
			fromSocket <- f
			time.Sleep(time.Millisecond * 5)

		}
		// we've handed to data off to the from socket channel when the processing is done it will be sent
		// back on the to channel. This ensures that all the pool management is done in a single thread.

	bufloop:
		for {
			select {
			// we get the frame from the channel
			case item := <-toSocket:
				{
					// process the frame
					time.Sleep(time.Second)
					// and give it back to the pool
					messagePool.ReturnToPool(item)
					//mm = *item
				}
			default:
				break bufloop
			}

		}

	}

}
