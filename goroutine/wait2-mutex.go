package main

import (
	"fmt"
	"sync"
)

func main() {

	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	var m sync.Mutex
	m.Lock()
	go func() {
		v = 1
		m.Unlock()
		wg.Done()
	}()
	go func() {
		m.Lock()
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}
