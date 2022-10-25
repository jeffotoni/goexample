package main

import (
    "fmt"
    "runtime"
    "time"
)

const (
    numElements = 10000000
)

var foo = map[int]int{}

func timeGC() {
    t := time.Now()
    runtime.GC()
    fmt.Printf("gc took: %s\n", time.Since(t))
}

func main() {
    for i := 0; i < numElements; i++ {
        foo[i] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}
