package main

import (
	"fmt"
	"time"
)

var workerID int
var publishID int

func main() {

	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)

	time.Sleep(time.Millisecond * 1)
}

func publisher(out chan string) {

	publishID++
	thisID := publishID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publisher %d is data\n", thisID)
		data := fmt.Sprintf("Data de publisher %d Data %d\n", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {

	workerID++
	thisID := workerID

	for {
		fmt.Printf("this is: %d\n", thisID)
		input := <-in
		fmt.Printf("%d input is: %s\n", thisID, input)
	}
}
