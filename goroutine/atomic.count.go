package main

import (
	"sync/atomic"
	"time"
)

type count32 int32

func (c *count32) Increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) Get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

var c count32

func main() {
	for i := 0; i < 10; i++ {
		go f1()
		go f2()
	}

	time.Sleep(time.Second * 5)
}

func f2() {
	println(c.Get())
}

func f1() {
	c.Increment()
	println("processando f1:", c.Get())
	time.Sleep(time.Second)
}
