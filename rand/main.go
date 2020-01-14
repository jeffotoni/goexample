package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//var min int = 10
	//var max int = 100
	fmt.Println("Hello, playground")
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(100000))

	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100000))
}
