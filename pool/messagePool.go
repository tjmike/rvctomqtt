package pool

import (
	"github.com/golang-collections/collections/stack"
	"rvctomqtt/intf"
)

type Pool struct {
	myStack      *stack.Stack
	factory      *intf.CanThingFactory
	nBuffCreated uint64
}

func NewPool(f *intf.CanThingFactory, initialCapacity int) *Pool {
	var p = Pool{
		myStack:      stack.New(),
		factory:      f,
		nBuffCreated: 0,
	}
	// pre-load the Pool
	{
		var i int = 0
		for {

			if i == initialCapacity {
				break
			}
			var ct *intf.CanThing = (*f).Create()
			p.ReturnToPool(ct)
			p.nBuffCreated++
			i++
		}
	}
	return &p
}
func (p *Pool) getNBuffCreated() uint64 {
	return p.nBuffCreated
}

func (p *Pool) size() int {
	return p.myStack.Len()
}

func (p *Pool) ReturnToPool(thing *intf.CanThing) {
	p.myStack.Push(thing)
}
func (p *Pool) Get() *intf.CanThing {
	var canThingPointer *intf.CanThing

	nBuffers := p.myStack.Len()
	// make sure we have something in the Pool, if not create one
	// for now we can be sure that the stack will not shrink before we pop it
	//
	if nBuffers > 0 {
		tmp := p.myStack.Pop()
		canThingPointer = tmp.(*intf.CanThing)
	} else {
		// Create a new message, don't put on the stack - this message is already "popped"
		// We need some instrumentation - query buffer sizes/etc

		canThingPointer = (*p.factory).Create()
		p.nBuffCreated++
	}
	return canThingPointer

}
