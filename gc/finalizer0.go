package main

import (
	"fmt"
	"runtime"
	"time"
)

type Test struct {
	A int
}

func test() {
	// create pointer
	a := &Test{}
	// add finalizer which just prints
	runtime.SetFinalizer(a, func(a *Test) { fmt.Println("I AM DEAD") })
}

func main() {
	test()
	// run garbage collection
	runtime.GC()
	// sleep to switch to finalizer goroutine
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("done")
}
