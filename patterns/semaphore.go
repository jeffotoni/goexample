package main

import (
	"fmt"
	"time"
)

// This example shows how easy it is to implement a semaphore using channels.
// It might be used to control the max number of concurrent requests
// made to a certain destination.

func main() {
	const maxConcurrent = 2

	const totalTasks = 10

	semaphore := make(chan struct{}, maxConcurrent)

	done := make(chan struct{}, 1)

	for i := range make([]int, totalTasks) {
		// blocks until semaphore is released
		semaphore <- struct{}{}

		taskNumber := i
		// executes task async
		go func() {
			fmt.Println("executing task: ", taskNumber)
			time.Sleep(time.Second)

			// release semaphore
			<-semaphore

			// counts finished
			done <- struct{}{}
		}()
	}

	count := 0
	for range done {
		count++
		if count == totalTasks {
			break
		}
	}

	fmt.Println("done")
}

