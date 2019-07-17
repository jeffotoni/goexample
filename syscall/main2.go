package main

import (
    "os"
    "os/signal"
)

func main() {
    doneCh := make(chan struct{})

    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh, os.Interrupt)

    go receive(signalCh, doneCh)

    <-doneCh
}

func receive(signalCh chan os.Signal, doneCh chan struct{}) {
    for {
        select {
        // Example. Process to receive a message
        // case msg := <-receiveMessage():
        case <-signalCh:
            println("close..")
            doneCh <- struct{}{}
        }
    }
}
