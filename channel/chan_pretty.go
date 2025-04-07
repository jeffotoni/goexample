package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	// read the channel
	go func() {
		for msg := range ch {
			fmt.Println("Received:", msg)
		}
		fmt.Println("Dated channel. Enclosed reading.")
	}()

	// Simulates 3 workers sending data to the channel
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ch <- fmt.Sprintf("worker's message %d", id)
		}(i)
	}

	// Wait for everyone to finish and here's the date or channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Wait (just to prevent main from finalizing before this example)
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		wg.Wait()
		done.Done()
	}()
	done.Wait()
}
