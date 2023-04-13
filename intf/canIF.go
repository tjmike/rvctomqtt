package intf

import "time"

// CanFrameIF - a simple can frame interface that can be used at a low level to build frames. The lower level code
// need not know if the implementation is really a J1939, RVC or other type of frame.
type CanFrameIF interface {
	GetMessage() *[16]byte    // The raw can message
	SetTimeStamp(t time.Time) // timestamp of the message
	BuildCanFrameX()          // Build an empty frame to be be populated later
	String() string           // this is the standard string interface - we just want all CanFrameIF
}
