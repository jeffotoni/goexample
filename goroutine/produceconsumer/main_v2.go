package main

import "fmt"
import "time"

var stream = make(chan int)
var fin = make(chan bool)

func produce() {
    for i := 0; i < 10; i++ {
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
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go consume()
    go produce()
    <-fin
    fmt.Println("After calling DONE")
}
