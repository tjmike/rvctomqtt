package rvc

import (
	"fmt"
	"math"
	"rvctomqtt/constants"
	"sync"
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

// RvcItem - the general uses case is to take some generic RVC message and turn it into a proper object.
type RvcItem struct {
	name        string    // Once created this does not change
	dgn         uint32    // Once created this does not change
	timestamp   time.Time // System timestamp for this message
	lastChanged time.Time // timestamp of the last change - a sub class is expected to set this

	// have a mutex per parallel item is nice in that multiple items can be queried/manipulated in parallel
	// is there some better approach though. the main case were we mutate the calue is when we're upating some
	// Instance which may happen often - many times per minute

	lock sync.RWMutex
}

func (i *RvcItem) String() string {
	return fmt.Sprintf("name: %s DGN: %x timestamp %s", i.name, i.dgn, i.timestamp)
}

func (i *RvcItem) GetName() string {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.name
}

func (i *RvcItem) GetDGN() uint32 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.dgn
}
func (i *RvcItem) GetTimestamp() time.Time {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.timestamp
}
func (i *RvcItem) GetLastChanged() time.Time {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.lastChanged
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
	// GetName - the name of this DGN
	GetName() string
	GetDGN() uint32
	// GetTimestamp - the timestamp of the data
	GetTimestamp() time.Time
	// GetLastChanged - the timestamp where we last saw a change. Note that some fields are not considered.
	// instance is an example, this is because instance is part of the identifier. It's important that the cached
	// instances of RCVItems works properly for this timestamp to be correct.
	GetLastChanged() time.Time

	// GetInstance - get the instance number/id. If there is no instance then we return DataNotAvailableUint8
	GetInstance() byte

	// Init- init this item from the RVC frame. A "subclass" should give the parent an opportunity to init first.
	Init(f *RvcFrame)
}

// Init - initialize this struct from the provided RVC data frame. It is assume that the parent struct
// will perform any needed synchronization, etc.
// We never want to hold references to the frame we are initializing from. Once we are initialized we expect to
// give the frame back to be reused.
func (r *RvcItem) Init(f *RvcFrame) {
	r.timestamp = f.Timestamp
	r.dgn = uint32(f.DGNHigh()) << 8
	r.dgn = uint32(f.DGNLow()) | r.dgn
	r.name = DGNName(r.dgn)
	//fmt.Printf("INIT: RVCITEM: dgn:%d name: %s\n", r.dgn, r.name)
}
