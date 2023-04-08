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

func GetRVCMessages(fromSocket, toSocket chan *rvccan.Frame) {
	GetCANMessages(fromSocket, toSocket)

}
func BuildCanFrame(frame *rvccan.Frame) {

	// Raspberry PI May be this
	// The test passed the other way on PI - so this means that something may be messed up
	// Leave this comment and keep an eye out. It could be  that the bytes coming from
	// the socket are messed up.
	//frame.ID = binary.LittleEndian.Uint32((*msg).CanMessage[0:])
	// MAC does this
	//frame.ID = binary.BigEndian.Uint32((*msg).CanMessage[0:])
	//frame.Timestamp = (*msg).Timestamp
	setFrameID(frame)
	frame.Length = (*frame).MessageBytes[4]
	frame.Flags = (*frame).MessageBytes[5]
	frame.Res0 = (*frame).MessageBytes[6]
	frame.Res1 = (*frame).MessageBytes[7]

	// not needed if we set the length and all the bytes
	for i := range frame.Data {
		frame.Data[i] = 0xff
	}
	{
		var fl int = int(frame.Length)
		for i, v := range (*frame).MessageBytes[8:15] {
			if i < fl {
				frame.Data[i] = v
			} else {
				frame.Data[i] = 0xff
			}
		}
	}

}
