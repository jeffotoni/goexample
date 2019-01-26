// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"os"
	"time"
)

type Promise struct {
	Result chan string
	Error  chan error
}

var (
	ch1  = make(chan *Promise)  // received a pointer from the structure
	ch2  = make(chan string, 1) // allows only 1 channels
	ch3  = make(chan int, 2)    // allows only 2 channels
	ch4  = make(chan float64)   // has not been set can freely receive
	ch5  = make(chan []byte)    // by default the capacity is 0
	ch6  = make(chan bool, 1)   // non-zero capacity
	ch7  = make(chan time.Time, 2)
	ch8  = make(chan struct{}, 2)
	ch9  = make(chan struct{})
	ch10 = make(map[string](chan int)) // map channel
	ch11 = make(chan error)
	ch12 = make(chan error, 2)
	// receives a zero struct
	ch14 <-chan struct{}
	ch15 = make(<-chan bool)          // can only read from
	ch16 = make(chan<- []os.FileInfo) // // can only write to
	// holds another channel as its value
	ch17 = make(chan<- chan bool) // // can read and write to
)

// Parameters of Func
// (jobs <-chan int, results chan<- int)

// Receives Value, only read
// jobs <-chan int //receives the value

// Receives Channel, only write
// results chan<- int // receive channel
// or
// results chan int // receive channel

// Receives Channel variadic
// results ...<-chan int

func main() {

	ch2 <- "okay"
	defer close(ch2)
	fmt.Println(ch2, &ch2, <-ch2)

	ch7 <- time.Now()
	ch7 <- time.Now()
	fmt.Println(ch7, &ch7, <-ch7)
	fmt.Println(ch7, &ch7, <-ch7)
	defer close(ch7)

	ch3 <- 1 // okay
	ch3 <- 2 // okay
	// deadlock
	// ch3 <- 3 // does not accept any more values, if you do it will error : deadlock
	defer close(ch3)
	fmt.Println(ch3, &ch3, <-ch3)
	fmt.Println(ch3, &ch3, <-ch3)

	ch10["lambda"] = make(chan int, 2)
	ch10["lambda"] <- 100
	defer close(ch10["lambda"])
	fmt.Println(<-ch10["lambda"])
}
