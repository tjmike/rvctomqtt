package rvc

import (
	"fmt"
	"math"
	"rvctomqtt/constants"
	"time"
)

const (
	NABit     uint8   = 0x3
	NAByte    byte    = 0xff
	NAuint8   uint8   = 0xff
	NAuint16  uint16  = 0xffff
	NAuint32  uint32  = 0xffffff
	NAFloat64 float64 = math.MaxFloat64
)

// RvcItem - the general uses case is to take some generic RVC message and turn it into a proper object. All of these
// objects will be based from RVCItem.
type RvcItem struct {
	Name          string    // DGN Name - TODO we could leave this a method and use a map to look it up
	DGN           uint32    // Once created this does not change
	SourceAddress uint8     // Source address (requester/sender?) this just SA OR desired SA for address claim
	Priority      uint8     // 3 bit Priority
	Timestamp     time.Time // System Timestamp for this message
	LastChanged   time.Time // Timestamp of the last change - a sub class is expected to set this
	LastNotified  time.Time // Last time we sent a notification - we receive change data often but don't want to relay unchanged data too often

}

//func (r *RvcItem) GetInstanceKey() interface{} {
//	return DGNInstanceKey{r.DGN, r.GetInstance()}
//}

func (i *RvcItem) Equals(o *RvcItem) bool {
	if i.Name != o.Name {
		return false
	}
	if i.DGN != o.DGN {
		return false
	}
	if i.SourceAddress != o.SourceAddress {
		return false
	}
	if i.Priority != o.Priority {
		return false
	}

	// Don't consider Timestamp / last changed
	return true
}

func (i *RvcItem) String() string {
	return fmt.Sprintf("Name: %s DGN: %x Timestamp %s sa: %x pri: %x", i.Name, i.DGN, i.Timestamp, i.SourceAddress, i.Priority)
}

func (i *RvcItem) GetName() string {
	return i.Name
}

func (i *RvcItem) GetDGN() uint32 {
	return i.DGN
}
func (i *RvcItem) GetTimestamp() time.Time {
	return i.Timestamp
}
func (i *RvcItem) GetLastChanged() time.Time {
	return i.LastChanged
}
func (i *RvcItem) GetLastNotified() time.Time {
	return i.LastNotified
}

// SetLastNotified - set the last notified to the current timestamp. If LastNotified > LastChanged then
// this message has been re-sent even though the data hasn't changed. This is a chatty way for new clients
// to sync up. Another option would be for a client to request current state and then we just fire messages
// for everything.
func (i *RvcItem) SetLastNotified() {
	i.LastNotified = i.Timestamp
}

func (i *RvcItem) GetSourceAddress() uint8 {
	return i.SourceAddress
}

// GetInstance - many items have an instance. It's used enough that this base class supports it an we expect
//
//	"subclass" to override it. This also means that DataNotAvailableUint8 is a special constant that really means this
//	DGN doesn't have instances
func (i *RvcItem) GetInstance() byte {
	return constants.DataNotAvailableUint8
}

// RvcItemIF = Get methods are exported.
type RvcItemIF interface {
	//GetInstanceKey() interface{}

	//MarshalJSON() ([]byte, error)

	// GetName - the Name of this DGN
	GetName() string
	GetDGN() uint32
	// GetTimestamp - the Timestamp of the data
	GetTimestamp() time.Time
	// GetLastChanged - the Timestamp where we last saw a change. Note that some fields are not considered.
	// instance is an example, this is because instance is part of the identifier. It's important that the cached
	// instances of RCVItems works properly for this Timestamp to be correct.
	GetLastChanged() time.Time
	GetLastNotified() time.Time
	SetLastNotified()

	// GetInstance - get the instance number/id. If there is no instance then we return DataNotAvailableUint8
	GetInstance() byte
	GetSourceAddress() byte

	GetPriority() byte
	// Init- init this item from the RVC frame. A "subclass" should give the parent an opportunity to init first.
	Init(f *RvcFrame)
}

func (i *RvcItem) GetPriority() uint8 {
	return i.Priority

}

// Init - initialize this struct from the provided RVC data frame. It is assume that the parent struct
// will perform any needed synchronization, etc.
// We never want to hold references to the frame we are initializing from. Once we are initialized we expect to
// give the frame back to be reused.
func (r *RvcItem) Init(f *RvcFrame) {
	r.Timestamp = f.GetTimeStamp()
	//fmt.Printf("ZZZ RVCItemInit: TS=%s FrameTS=%s\n",
	//	r.Timestamp.Format("01-02-2006 15:04:05.000000"),
	//	f.GetTimeStamp().Format("01-02-2006 15:04:05.000000"),
	//)

	r.DGN = uint32(f.DGNHigh()) << 8
	r.DGN = uint32(f.DGNLow()) | r.DGN
	r.Name = DGNName(r.DGN)
	r.SourceAddress = f.GetSourceAddress()
	r.Priority = f.GetPriority()
	//fmt.Printf("INIT: RVCITEM: DGN:%d Name: %s\n", r.DGN, r.Name)
}
