package intf

import "time"

type CanThing interface {
	GetMessage() *[16]byte
	SetTimeStamp(t time.Time)
	BuildCanFrameX()
	ToString() string
}
