package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/heap?seconds=n
// curl http://localhost:6060/debug/pprof/heap --output heap.tar.gz
// go tool pprof heap.tar.gz

// http://localhost:6060/debug/pprof/goroutine
// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/threadcreate
// http://localhost:6060/debug/pprof/block
// http://localhost:6060/debug/pprof/mutex

// http://localhost:6060/debug/pprof/profile
// http://localhost:6060/debug/pprof/trace?seconds=5

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	fmt.Println("hello world")
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(wg)
	wg.Wait()
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
