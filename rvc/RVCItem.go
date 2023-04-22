package rvc

import (
	"fmt"
	"math"
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

// GetFields - get a list of all the fieds for this item and the data tyoe of that field. This can be used to dynamically
// query the item. Imagine some generic JSON publisher that uses this to publish data to a queue.
// It may turn out that this is NOT needed and all we need to to use reflection to get all the public methods and get
// fields/types based on those method names....
func (i *RvcItem) Fields() *[]dataField {
	return nil
}

// GetFieldUint8 - get the uint8 value of a field. Should only be called for fields that are represented as uint8.
// Undefined if called on the wrong type
func (i *RvcItem) FieldUint8(f dataField) uint8 {
	return NAuint8
}

// GetFieldUint8 - get the uint8\16 value of a field. Should only be called for fields that are represented as uint16
// Undefined if called on the wrong type
func (i *RvcItem) FieldUint16(f dataField) uint16 {
	return NAuint16
}

// GetFieldUint8 - get the uint32 value of a field. Should only be called for fields that are represented as uint32
// Undefined if called on the wrong type
func (i *RvcItem) FieldUint32(f dataField) uint32 {
	return NAuint32
}

// GetFieldUint8 - get the float64 value of a field. Should only be called for fields that are represented as float64
// Undefined if called on the wrong type
func (i *RvcItem) FieldFloat64(f dataField) float64 {
	return NAFloat64
}

// RvcItemIF = Get methods are exported.
type RvcItemIF interface {
	GetName() string
	GetDGN() uint32
	GetTimestamp() time.Time
	GetLastChanged() time.Time
	Fields() *[]dataField
	FieldUint8(f dataField) uint8
	FieldUint16(f dataField) uint16
	FieldUint32(f dataField) uint32
	FieldFloat64(f dataField) float64
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
	fmt.Printf("INIT: RVCITEM: dgn:%d name: %s\n", r.dgn, r.name)
}
