package main

import (
	"fmt"
)

func main() {

	channel := make(chan map[string]string, 1)
	m := make(map[string]string)

	m["hello"] = "world"
	channel <- m
	m["hello"] = "data race"

	fmt.Println("teste")
	fmt.Println(<-channel)
}
