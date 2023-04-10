package can

import "rvctomqtt/intf"

type MyCanThingFactory struct {
}

func (f *MyCanThingFactory) Create() *intf.CanThing {
	var cf = Frame{}
	var cff intf.CanThing = &cf
	return &cff
}
