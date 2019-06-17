package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func main() {
	trace.Start(os.Stderr)
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
	trace.Stop()
}
