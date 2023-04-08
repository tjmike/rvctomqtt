package pform

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"rvctomqtt/can"
	"syscall"
	"time"
)

/**
 *
 * Listen for CAN messages on the can0 (TODO make flexible) interface
 * Parse the messages and place fromSocket channel.
 * fromSocket - socket messages will be sent FROM this via this channel
 * toSocket   - when a socket message can be reclaimed it will be returned via this channel
 *
 */
func GetCANMessages(fromSocket, toSocket chan *can.RawCanMessage) {
	// We do some byte swswapping and need a temporary byte to store data. We allocate the byte up
	// front and pass it to the swap function so that an allocation isn't needed for every buffer read.
	// This stack is is our buffer pool for processing
	// we push a number of pointers onto the stack
	// We pop one off for processing, pass it down the line
	// and it should eventually make it back to use via a channel
	myStack := stack.New()

	// Create a pool of CAN Messages
	// The general idea here is the a can message is pulled from the pool , populated and then
	// sent to the listener.
	for i := 0; i < 10; i++ {
		var xxx [can.MAX_MESSAGE]byte
		rawCAN := can.RawCanMessage{
			time.Now(),
			xxx}
		myStack.Push(&rawCAN)
	}

	// Find the interface we are intersted in by name
	// TODO - make this a param
	iface, err := net.InterfaceByName("can0")

	// TODO - learn to handle/print errors, etc.
	if err != nil {
		// return nil, err
		return
	}

	// TODO error handler
	// socket number and error returned
	// _ says to ignore the error
	s, _ := syscall.Socket(syscall.AF_CAN, syscall.SOCK_RAW, unix.CAN_RAW)
	// s, _ := syscall.Socket(syscall.AF_CAN, syscall.SOCK_RAW, unix.CAN_RAW_FD_FRAMES)

	// This is (I THINK) just strict to hold the interface id
	addr := &unix.SockaddrCAN{Ifindex: iface.Index}

	// Bind the socket and return if there's an error
	// TODO better error handling
	if err := unix.Bind(s, addr); err != nil {
		// return nil, err
		return
	}

	// This is our file descriptor
	f := os.NewFile(uintptr(s), fmt.Sprintf("fd %d", s))

	// pointer to CAN message pulled from our pool
	var rawPointer *can.RawCanMessage

	fmt.Println("Start socket loop forever")
	// Forever
	for {
		fmt.Println("LOOP PASS")

		nBuffers := myStack.Len()

		// make sure we have something in the pool, if not create one
		// for now we can be sure that the stack will not shrink before we pop it
		//
		if nBuffers > 0 {
			tmp := myStack.Pop()
			rawPointer = tmp.(*can.RawCanMessage)
		} else {
			// Create a new message, don't put on the stack - this message is already "popped"
			// We need some instrumentation - query buffer sizes/etc
			var xxx [can.MAX_MESSAGE]byte
			rawCAN := can.RawCanMessage{
				time.Now(),
				xxx}
			rawPointer = &rawCAN
		}

		fmt.Println("READ")

		// var zzz [16]byte
		_, err := f.Read((*rawPointer).CanMessage[0:can.MAX_MESSAGE])
		// nRead, err := f.Read((*rawPointer).canMessage[0:16])
		// nRead, err := f.Read(zzz[0:16])

		// TODO from https://www.kernel.org/doc/Documentation/networking/can.txt
		// We can get a more accurate timestamps using ioctl
		//
		// An accurate timestamp can be obtained with an ioctl(2) call after reading
		// a message from the socket:
		// struct timeval tv;
		// ioctl(s, SIOCGSTAMP, &tv);
		//
		// The timestamp has a resolution of one microsecond and is set automatically
		// at the reception of a CAN frame.

		(*rawPointer).Timestamp = time.Now()

		// TODO Handle Errors
		// We may want to allow a single error or a few errors in a row
		// We could keep error stats too - total errors / recent errors etc
		if err != nil {
			fmt.Println(err)
			break
		}

		//fmt.Printf("GOT: %x\n", binary.LittleEndian.Uint32((*rawPointer).CanMessage[0:]))

		var f = can.Frame{}
		BuildCanFrame(&f, rawPointer)
		fmt.Printf("GOT: %s\n", f.ToString())
		//fmt.Printf("GOT: %x\n", binary.LittleEndian.Uint32((*rawPointer).CanMessage[0:]))

		// ((*rawPointer).canMessage[0])
		// idBEF := binary.LittleEndian.Uint32(((*rawPointer).canMessage[0:]))
		// idBEF2 := (*rawPointer).getCANID()
		// print(fmt.Sprintf("\nID BEF =  %X bef2 = %X \n" ,  idBEF, idBEF2  ))
		// FIX Bytes - The Raspberry PI seems to swap the first 4 bytes. This is, I assume because the kernel is putting the 1st 4 bytes into an int
		// and this causes the bytes to be backwards.
		// IMPORTANT - ANY FIX HERE WILL LIKELY NEED TO BE DONE ON THE SEND SIDE AS WELL

		// print("READ ")
		// print(nRead)
		// print("\n")
		// print(rawPointer.toString())
		// print("\n")
		// var canIDA  uint32 = (uint32) ((*rawPointer).canMessage[0]) << 24
		// canIDA  = canIDA | (uint32)((*rawPointer).canMessage[1])<< 16
		// canIDA  = canIDA | (uint32)( (*rawPointer).canMessage[2]) <<  8
		// canIDA  = canIDA | (uint32) ( (*rawPointer).canMessage[3] )
		// var isErr bool = unix.CAN_ERR_FLAG & canIDA != 0
		// var isExtended bool = unix.CAN_EFF_FLAG & canIDA != 0
		// print(fmt.Sprintf("\nisError = %t isExtended = %t \n" ,  isErr, isExtended  ))
		// print("\n")
		// print(fmt.Sprintf("canIDA = %X\n" ,  canIDA  ))
		// print("canIDA = ")
		// print( canIDA)
		// print("\n")
		// print("\n")
		// print("\n")
		// print("\n")

		// Sentthe data out on a channel
		// This data is no longer ours. It should be returned via the toSocket channel , but we are done, don't access (log/print for example)

		fmt.Println("SEND")

		fromSocket <- rawPointer
		// slen := myStack.Len();
		// print("SLENB=")
		// print(slen)
		// RECEIVE any pointers back to be re-used
		fmt.Println("Get Back Buffers")

	bufloop:
		for {
			select {
			case item := <-toSocket:
				myStack.Push(item)
			default:
				break bufloop
			}

		}
		// print(" SLENA=")
		// slen = myStack.Len();
		// print(slen)
		// print("\n")
	}

	f.Close()
	// print("SOCKET DONE")
}
