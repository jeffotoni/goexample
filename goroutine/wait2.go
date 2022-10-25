package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		v = 1
		ch <- 1
		wg.Done()
	}()
	go func() {
		<-ch
		fmt.Println("done: ", v)
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(time.Second * 2)
}
