package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx sync.Mutex
var pending []chan func()
var max = 10
var queued = 0

func release() {
	mtx.Lock()
	defer mtx.Unlock()
	queued--
	if len(pending) == 0 {
		return
	}
	var p chan func()
	p, pending = pending[0], pending[1:]
	p <- release
}

func requestAccess(ch chan func()) {
	mtx.Lock()
	defer mtx.Unlock()
	queued++
	if queued > max {
		pending = append(pending, ch)
		return
	}
	ch <- release
}

func main() {
	// simulate "simultaneous" clients
	for i := 0; i < 16; i++ {
		go func(i int) {
			allowance := make(chan func())
			go requestAccess(allowance)
			releaseFn := <-allowance

			// from this point on client is allowed to proceed,
			// releaseFn signals the end of client execution.
			fmt.Printf("client %d is running\n", i)
			time.Sleep(time.Second)
			releaseFn()
		}(i)
	}

	// unorthodox waitgroup
	time.Sleep(time.Minute)
}
