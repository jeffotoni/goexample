package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func work(n int) {
	f, err := os.Open("/dev/zero")
	if err != nil {
		return
	}
	p := make([]byte, n)
	f.Read(p)
}

func main() {
	fmt.Printf("pid=%d numcpu=%d maxprocs=%d\n", os.Getpid(),
		runtime.NumCPU(), runtime.GOMAXPROCS(0))
	for {
		time.Sleep(time.Second * 10)
		for i := 0; i < 50; i++ {
			go work(i * 1024 * 1024)
		}
	}
}

