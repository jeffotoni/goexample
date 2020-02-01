package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	messages := make(chan int)
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("start 4")
		time.Sleep(time.Second * 4)
		messages <- 1

	}()
	go func() {
		defer wg.Done()
		fmt.Println("start 2")
		time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		defer wg.Done()
		fmt.Println("start 3")
		time.Sleep(time.Second * 1)
		messages <- 3
	}()

	go func() {
		for i := range messages {
			fmt.Println("done:", i, " goroutine: ", runtime.NumGoroutine())
		}
	}()

	time.Sleep(time.Second * 8)
	wg.Wait()
}
