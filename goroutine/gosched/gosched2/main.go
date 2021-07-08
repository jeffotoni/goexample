
package main

import (
	"fmt"
	"time"
)
import "runtime"
import "encoding/json"
import "strings"

const testBytes = `{ "Test": "value" }`

type Message struct {
	Test string
}

func cpuIntensive(p *Message) {
	for i := int64(1); i <= 10000; i++ {
		json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
	}
	//runtime.Gosched()
	fmt.Println("Done intensive thing")
}

func printVar(p *Message) {
	fmt.Printf("print x = %v.\n", *p)
}

func main() {
	runtime.GOMAXPROCS(1)
	runtime.go
	x := Message{}
	cpuIntensive(&x)
	go printVar(&x)
	time.Sleep(1 * time.Nanosecond)
	//time.Sleep(1 * time.Second*1)
}
	