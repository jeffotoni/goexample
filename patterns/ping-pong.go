package main

import "time"

func main() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(3 * time.Second)
	println(<-table)
}

func player(table chan int) {
	println("player")
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
