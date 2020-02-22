package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var totalObjects int32

func TotalObjects() int32 {
	return atomic.LoadInt32(&totalObjects)
}

type Object struct {
	p uintptr // C allocated pointer
}

func NewObject() *Object {
	o := &Object{}
	// TODO: perform other initializations
	atomic.AddInt32(&totalObjects, 1)
	runtime.SetFinalizer(o, (*Object).finalizer)
	return o
}

func (o *Object) finalizer() {
	atomic.AddInt32(&totalObjects, -1)
	// TODO: perform finalizations
}

func main() {
	fmt.Println("Total objects:", TotalObjects())
	for i := 0; i < 100; i++ {
		_ = NewObject()
		runtime.GC()
	}
	fmt.Println("Total objects:", TotalObjects())
}
