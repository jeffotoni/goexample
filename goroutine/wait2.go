package main

import (
	"fmt"
	"sync"
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
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}
