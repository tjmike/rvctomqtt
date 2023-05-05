package pool

import (
	"fmt"
	"rvctomqtt/intf"
	"sync"
	"sync/atomic"
)

// Pool Simple pool for can frames. It's supposed to be thread safe so the pool can be used among multiple
// go routines. The use case is gorutine 1 pulls a buffer from the pool, populates it, and sends it to some other
// goroutine(2). Goroutine 2 utilizes the buffer and then when done returns it to the pool. Never retain a reference to
// an item buffer if it's going back into the pool.

type Pool struct {
	pool         sync.Pool
	factory      *intf.CanFrameFactory
	nBuffCreated uint64
}

func (p *Pool) Create() *intf.CanFrameIF {
	atomic.AddUint64(&p.nBuffCreated, 1)
	fmt.Printf("CREATED %d:\n", atomic.LoadUint64(&p.nBuffCreated))
	return (*p.factory).Create()
}

// NewPool - crate a new pool and populate with initialCapacity items. Use the CanFrameFactory to crete new items
// if needed.
func NewPool(f *intf.CanFrameFactory) *Pool {
	var p = Pool{
		factory:      f,
		nBuffCreated: 0,
	}

	p.pool = sync.Pool{
		New: func() interface{} {
			return p.Create()
		},
	}
	return &p
}

func (p *Pool) getNBuffCreated() uint64 {
	return atomic.LoadUint64(&p.nBuffCreated)
}

func (p *Pool) ReturnToPool(canFrame *intf.CanFrameIF) {
	p.pool.Put(canFrame)
}

func (p *Pool) Get() *intf.CanFrameIF {
	var canThingPointer = p.pool.Get().(*intf.CanFrameIF)
	return canThingPointer
}
