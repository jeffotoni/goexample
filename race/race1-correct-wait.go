package main

import (
	"fmt"
	"sync"
	"time"
)

var Lock sync.Mutex

var i = 10

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {

	// Initialize a waitgroup variable
	var wg sync.WaitGroup

	// `Add(1) signifies that there is 1 task that we need to wait for
	wg.Add(1)
	go func() {

		Lock.Lock()
		i++
		i++
		Lock.Unlock()

		// Calling `wg.Done` indicates that we are done with the task we are waiting fo
		defer wg.Done()

	}()

	wg.Add(1)
	go func() {

		Lock.Lock()

		i++
		time.Sleep(1 * time.Nanosecond)

		Lock.Unlock()

		// Calling `wg.Done` indicates that we are done with the task we are waiting fo
		defer wg.Done()

	}()

	// `wg.Wait` blocks until `wg.Done` is called the same number of times
	// as the amount of tasks we have (in this case, 1 time)
	wg.Wait()
	return i
}
