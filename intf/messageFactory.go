package intf

// CanFrameFactory - implementations should be able to create a new empty CanFrame at will.
// We use an interface in hopes that an implementation that creates J1939 or RVC frames will
// simply plug in here.
type CanFrameFactory interface {
	Create() *CanFrameIF
}
