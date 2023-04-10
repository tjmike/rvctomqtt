package can

import "rvctomqtt/intf"

type CanFrameFactory struct {
}

func (f *CanFrameFactory) Create() *intf.CanThing {
	var cf = Frame{}
	var cff intf.CanThing = &cf
	return &cff
}
