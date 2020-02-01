package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var myWaitGroup sync.WaitGroup

	myWaitGroup.Add(2)

	go func() {
		fmt.Println("Start gorouting 1")
		time.Sleep(time.Second * 2)
		fmt.Println("End gorouting 1")
		myWaitGroup.Done()
	}()

	go func() {
		fmt.Println("Start gorouting 2")
		time.Sleep(time.Second * 4)
		fmt.Println("End gorouting 2")
		myWaitGroup.Done()
	}()

	fmt.Println("Waiting for all goroutines to exit")
	myWaitGroup.Wait()
	fmt.Println("Waited for all goroutines to exit")
}
