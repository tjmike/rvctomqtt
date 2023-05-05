package pform

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"rvctomqtt/intf"
	//"rvctomqtt/can"
	"rvctomqtt/pool"
	"syscall"
	"time"
	"unsafe"
)

// TODO: there is  J1939 support in the Linux kernel
/**
 *
 * Listen for CAN messages on the can interface
 * Parse the messages and place fromSocket channel.
 * fromSocket - socket messages will be sent FROM this via this channel
 * toSocket   - when a socket message can be reclaimed it will be returned via this channel
 *
 */
func GetCANMessages(messagePool *pool.Pool, fromSocket, toSocket chan *intf.CanFrameIF) {

	var socketInterface = "can0"
	//var socketInterface = "vcan0"

	// Find the interface we are intersted in by name
	iface, err := net.InterfaceByName(socketInterface)

	if err != nil {
		fmt.Errorf("Error finding interface: %s", socketInterface)
		return
	}

	// TODO error handler
	// socket number and error returned
	// _ says to ignore the error
	s, err := syscall.Socket(syscall.AF_CAN, syscall.SOCK_RAW, unix.CAN_RAW)
	if err != nil {
		fmt.Errorf("Error opening socket on interface: %s", socketInterface)
		return
	}
	addr := &unix.SockaddrCAN{Ifindex: iface.Index}

	// Bind the socket and return intf there's an error
	// TODO better error handling
	if err := unix.Bind(s, addr); err != nil {
		fmt.Errorf("Error opening socket on interface: %s", socketInterface)
		return
	}

	// This is our file descriptor
	f := os.NewFile(uintptr(s), fmt.Sprintf("fd %d", s))

	fmt.Println("Start socket loop forever")

	var pktTime time.Time = time.Now()
	recvTime := syscall.Timeval{}

	// Forever
	for {

		//  START READ FRAME

		// get a canframe from the pool and get the by buffer
		var canFrame *intf.CanFrameIF = (*messagePool).Get()

		var bytes = (*canFrame).GetMessage()
		// Get the next message into the can raw bytes
		// TODO This should be a goroutine - otherwise this whole loop is blocked waiting for the next can message

		_, err := f.Read(bytes[0:])

		// TODO Handle Errors
		// We may want to allow a single error or a few errors in a row
		// We could keep error stats too - total errors / recent errors etc
		if err != nil {
			fmt.Println(err)
			break
		}

		{
			// TODO from https://www.kernel.org/doc/Documentation/networking/can.txt
			// Get the associated timestamp
			_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(f.Fd()), uintptr(syscall.SIOCGSTAMP),
				uintptr(unsafe.Pointer(&recvTime)))

			if errno == 0 {
				pktTime = time.Unix(0, recvTime.Nano())
			} else {
				// consider logging if we can't use Unix time - BUT if we do don't log every packet
				pktTime = time.Now()
			}
			(*canFrame).SetTimeStamp(pktTime)
		}

		fmt.Printf("READ: %x\n", (*canFrame).GetMessage())

		(*canFrame).BuildCanFrameX()
		//fmt.Println((*canFrame).String())

		// Send the frame to the channel
		fromSocket <- canFrame
		/// XXX END PROCESS FRAOM
	bufloop:
		for {
			select {
			case item := <-toSocket:
				messagePool.ReturnToPool(item)
			default:
				break bufloop
			}
		}
	}
	f.Close()
}
