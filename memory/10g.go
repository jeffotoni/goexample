package main

import (
	"math"
	"time"
)

func main() {
	_ = make([]byte, 4000<<20)
	println("done alocado 10G in memory:", 4000<<20)
	<-time.After(time.Duration(math.MaxInt64))
}
