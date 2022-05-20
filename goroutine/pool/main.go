package main

import (
    "fmt"

    "sync"
)

func main() {
    myPool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Creating new instance here.")
            return struct{}{}
        },
    }

    myPool.Get()
    instance := myPool.Get()
    fmt.Println("Here instance:", instance)

    myPool.Put(instance)
    myPool.Get()

    var numCalcsCreated int
    calcPool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("new calc pool here")

            numCalcsCreated += 1
            mem := make([]byte, 1024)

            return &mem
        },
    }

    fmt.Println("calcPool.New:", calcPool.New())

    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())

    calcPool.Get()

    const numWorkers = 1024 * 1024
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    for i := numWorkers; i > 0; i-- {
        go func() {
            defer wg.Done()

            mem := calcPool.Get().(*[]byte)
            defer calcPool.Put(mem)
        }()
    }

    wg.Wait()
    fmt.Printf("%d calculators were created.", numCalcsCreated)
}
