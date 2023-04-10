package can

import "rvctomqtt/intf"

type MyCanThingFactory struct {
}

func (f *MyCanThingFactory) Create() *intf.CanFrameIF {
	var cf = Frame{}
	var cff intf.CanFrameIF = &cf
	return &cff
}
