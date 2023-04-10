package pool

import (
	"fmt"
	"rvctomqtt/intf"
	"testing"
	"time"
)

func TestMessagePool(t *testing.T) {
	var ps int = 10
	var m = myMessageFactory{}
	var ctf intf.CanThingFactory = &m
	var pool = NewPool(&ctf, ps)

	var sz = pool.size()

	if !((sz) == ps) {
		t.Errorf("Wrong Pool size. Expected %d got %d", ps, sz)
	}

	for i := 0; i < sz; i++ {
		pool.Get()
	}

	if !(pool.size() == 0) {
		t.Errorf("Wrong Pool size. Expected %d got %d", 0, sz)
	}

	if !(pool.getNBuffCreated() == uint64(ps)) {
		t.Errorf("Wrong Pool created. Expected %d got %d", 0, sz)
	}

	var ct = pool.Get()
	var msg []byte = (*ct).GetMessage()
	fmt.Printf("created = %d\n", pool.getNBuffCreated())
	fmt.Printf("msg = %x\n", msg)

	pool.ReturnToPool(ct)
	var ct1 = pool.Get()
	var ct2 = pool.Get()
	fmt.Printf("created = %d\n", pool.getNBuffCreated())
	pool.ReturnToPool(ct1)
	pool.ReturnToPool(ct2)

	ct1 = pool.Get()
	ct2 = pool.Get()
	fmt.Printf("created = %d\n", pool.getNBuffCreated())

}

type myMessageFactory struct {
	data [32]byte
	ts   time.Time
}

func (f *myMessageFactory) GetMessage() []byte {
	return f.data[0:]
}

func (frame *myMessageFactory) setTimeStamp(t time.Time) {
	frame.ts = t
}

func (f *myMessageFactory) Create() *intf.CanThing {
	var foo intf.CanThing = &myMessageFactory{}
	return &foo
}
