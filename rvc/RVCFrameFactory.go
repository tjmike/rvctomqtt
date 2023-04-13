package rvc

import "rvctomqtt/intf"

type RVCFrameFactory struct {
}

func (f *RVCFrameFactory) Create() *intf.CanFrameIF {
	var cf = RvcFrame{}
	var cff intf.CanFrameIF = &cf
	return &cff
}
