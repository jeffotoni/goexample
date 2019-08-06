package main

import "fmt"
import "time"

var fin = make(chan bool)
var stream = make(chan int)

func produce() {
    for i := 0; i < 100; i++ {
        fmt.Println("sending")
        stream <- i
        fmt.Println("sent")
    }

    fmt.Println("No closing channel")
    //close(stream)
    fin <- true
}

func consume() {
    for {
        data := <-stream
        fmt.Println("Consumer: ", data)
        time.Sleep(1500 * time.Millisecond)
    }
}

func main() {
    go produce()
    go consume()
    <-fin
    fmt.Println("After calling DONE")
}
