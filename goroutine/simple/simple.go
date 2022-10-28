package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {

		fmt.Println("worker", id, "started  job", j)
		// run here
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)

		results <- j * 125
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("worker", a, " results:", <-results)
	}
}
