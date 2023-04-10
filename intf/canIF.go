package intf

import "time"

type CanFrameIF interface {
	GetMessage() *[16]byte
	SetTimeStamp(t time.Time)
	BuildCanFrameX()
	ToString() string
}
