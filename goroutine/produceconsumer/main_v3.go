package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "sync"
)

func produce(in chan string) {
    defer close(in)
    producedPassword := strconv.Itoa(rand.Intn(100))
    in <- producedPassword
}

func consume(in, out chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    for s := range in {
        //currentHash := crypt(s)
        // success condition
        hashToCrack := "x93939"
        currentHash := "x93939"
        if currentHash == hashToCrack {
            fmt.Println(currentHash, "\n")
            out <- s
        }
    }
}

func stop(out chan string, wg *sync.WaitGroup) {
    wg.Wait()
    close(out)
}

func main() {

    in, out := make(chan string), make(chan string)
    wg := &sync.WaitGroup{}
    go produce(in)

    //for i:=0; i<runtime.NumCPU();i++ {
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go consume(in, out, wg)
    }
    go stop(out, wg)
    fmt.Println(<-out)
}
