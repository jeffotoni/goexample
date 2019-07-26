package main

import (
    "fmt"
    "runtime"
    "strconv"
    "time"
)

const (
    numElements = 10000000
)

var foo = map[string]int{}

func timeGC() {
    t := time.Now()
    runtime.GC()
    fmt.Printf("gc took: %s\n", time.Since(t))
}

func main() {
    for i := 0; i < numElements; i++ {
        foo[strconv.Itoa(i)] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }
}
