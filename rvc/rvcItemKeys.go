package rvc

type DGNkey struct {
	DGN uint32
}

type DGNInstanceKey struct {
	DGN      uint32
	Instance uint8
}
type ChassisMobilityStatusKey struct {
	// nothing there is only of these
}
type AddressClaimedKey struct {
	sourceAddress uint8
}

type DataPacketKey struct {
	SourceAddress uint8
	PacketNumber  uint8
}

type InformationRequestKey struct {
	DestinationAddress uint8
}

type InitialPacketKey struct {
	SourceAddress uint8
}
