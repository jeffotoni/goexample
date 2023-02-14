package main

import (
	"fmt"
	"sync"
)

type obj struct {
	msg      string
	receiver int
}

func main() {

	ch := make(chan *obj) // both block or non-block are ok
	var wg sync.WaitGroup
	receiver := 25 // specify receiver count

	sender := func() {
		o := &obj{
			msg:      "hello everyone! test goroutine...",
			receiver: receiver,
		}
		ch <- o
	}

	recv := func(idx int) {
		defer wg.Done()
		o := <-ch
		fmt.Printf("%d received at %d\n", idx, o.receiver)
		o.receiver--
		if o.receiver > 0 {
			ch <- o // forward to others
		} else {
			fmt.Printf("last receiver: %d\n", idx)
		}
	}

	go sender()

	for i := 0; i < receiver; i++ {
		wg.Add(1)
		go recv(i)
	}

	wg.Wait()
}
