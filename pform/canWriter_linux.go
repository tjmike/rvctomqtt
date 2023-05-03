package pform

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"rvctomqtt/intf"
	"syscall"
)

// TODO: there is  J1939 support in the Linux kernel
// TODO: We need to get a source address and maintain that
func SendCANMessages(message chan *intf.CanFrameIF) {
	var socketInterface = "can0"
	//var socketInterface = "vcan0"
	iface, err := net.InterfaceByName(socketInterface)
	if err != nil {
		fmt.Errorf("Error finding interface: %s", socketInterface)
		return
	}

	// TODO error handler
	s, err := syscall.Socket(syscall.AF_CAN, syscall.SOCK_RAW, unix.CAN_RAW)
	if err != nil {
		fmt.Errorf("Error opening socket on interface: %s", socketInterface)
		return
	}
	addr := &unix.SockaddrCAN{Ifindex: iface.Index}

	// Bind the socket and return intf there's an error
	// TODO better error handling
	// TODO: need source address
	if err := unix.Bind(s, addr); err != nil {
		fmt.Errorf("Error opening socket on interface: %s", socketInterface)
		return
	}

	// This is our file descriptor
	f := os.NewFile(uintptr(s), fmt.Sprintf("fd %d", s))

	// Forever
	for {
		// get the request to send
		var canFrame *intf.CanFrameIF = <-message

		// get the raw byte buffer
		var bytes = (*canFrame).GetMessage()
		// Get the next message into the can raw bytes

		fmt.Printf("Trying to send: %x\n", bytes)
		_, err := f.Write(bytes[0:])
		fmt.Printf("SENT send: %x\n", bytes)

		// TODO Handle Errors
		// We may want to allow a single error or a few errors in a row
		// We could keep error stats too - total errors / recent errors etc
		if err != nil {
			fmt.Println(err)
			break
		}

	}
	f.Close()
}
