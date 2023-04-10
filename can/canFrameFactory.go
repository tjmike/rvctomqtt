package can

import "rvctomqtt/intf"

type CanFrameFactory struct {
}

func (f *CanFrameFactory) Create() *intf.CanFrameIF {
	var cf = Frame{}
	var cff intf.CanFrameIF = &cf
	return &cff
}
