// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	fmt.Println("cpus: ", n)
	quit := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(10 * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}
}
