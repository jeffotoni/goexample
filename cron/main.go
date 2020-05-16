package main

import "fmt"
import "time"

func main() {
    go func() {
        c := time.Tick(1 * time.Second)
        for range c {
            // Note this purposfully runs the function
            // in the same goroutine so we make sure there is
            // only ever one. If it might take a long time and
            // it's safe to have several running just add "go" here.
            RunEverySecond()
        }
    }()

    // Other processing or the rest of your program here.
    time.Sleep(5 * time.Second)

    // Or to block forever:
    //select {}
    // However, if doing that you could just stick the above for loop
    // right here without dropping it into a goroutine.
}

func RunEverySecond() {
    fmt.Println("----")
}
