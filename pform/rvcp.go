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
