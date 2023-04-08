package pform

import (
	"encoding/binary"
	"fmt"
	"rvctomqtt/can"
)

// we get a by buffer that's really a c ctruct with an int(unsigned?). We need to convert that
// int into a go uint32. Since that's a platform specific operation we need to provide a per CPU
// architecture implementation
func setFrameID(frame *can.Frame) {
	frame.ID = binary.LittleEndian.Uint32((*frame).MessageBytes[0:])
}

func init() {
	fmt.Printf("darwin init ")
}

func doit() {
	fmt.Printf("darwin doit")
}
