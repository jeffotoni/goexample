package main

import (
    "fmt"
    "runtime"
)
func say2(s string){
    for i:=0; i<1000; i ++{
        fmt.Println(s)
    }
}
func say(s string, status int) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func main() {
    //runtime.GOMAXPROCS(2)
    go say2("world")
    say("hello", 0)
    //time.Sleep(time.Second*2)
}
