package pform

// This is the public facing interface. It will call platform specific code.
// In this example doit (lower case is different code on darwin vs linux)

import (
	"fmt"
	//"github.com/tjmike/rvctomqtt/can"
	rvccan "rvctomqtt/can"
)

func init() {
	fmt.Printf("rvcp init\n")
}

func Doit() {
	fmt.Printf("rvcp doit\n")
	doit()
}

func GetRVCMessages(fromSocket, toSocket chan *rvccan.RawCanMessage) {
	GetCANMessages(fromSocket, toSocket)

}
func BuildCanFrame(frame *rvccan.Frame, msg *rvccan.RawCanMessage) {

	// Raspberry PI May be this
	// The test passed the other way on PI - so this means that something may be messed up
	// Leave this comment and keep an eye out. It could be  that the bytes coming from
	// the socket are messed up.
	//frame.ID = binary.LittleEndian.Uint32((*msg).CanMessage[0:])
	// MAC does this
	//frame.ID = binary.BigEndian.Uint32((*msg).CanMessage[0:])
	setFrameID(frame, msg)

	frame.Length = (*msg).CanMessage[4]
	frame.Flags = (*msg).CanMessage[5]
	frame.Res0 = (*msg).CanMessage[6]
	frame.Res1 = (*msg).CanMessage[7]

	// not needed if we set the length and all the bytes
	for i := range frame.Data {
		frame.Data[i] = 0
	}

}
