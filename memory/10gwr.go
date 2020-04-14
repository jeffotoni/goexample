package main

import (
	"math"
	"time"
)

func main() {
	ballast := make([]byte, 4000<<20)
	println("byte: ", 4000<<20)
	for i := 0; i < len(ballast)/2; i++ {
		ballast[i] = byte('A')
	}
	<-time.After(time.Duration(math.MaxInt64))
}
