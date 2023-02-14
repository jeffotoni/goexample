package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	qntFiles := 30
	Goroutine := 200

	//strings.
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, qntFiles)
	results := make(chan string, qntFiles)

	for w := 1; w <= Goroutine; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= qntFiles; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= qntFiles; a++ {
		qntFiles := <-results
		fmt.Println(qntFiles)
	}

}

func worker(id int, jobs <-chan int, results chan<- string) {

	for j := range jobs {

		//time.Sleep(time.Second)
		results <- GenerateFilename(j, id)
	}
}

func GenerateFilename(j, id int) string {
	//start := time.Now()
	//dateTime := start.Format("20060102150405")
	return fmt.Sprintf("job_%d_id_%d.wav", j, id)
	//return fmt.Sprintf("%s_%03d_00000_x%d_y%d_%s_%s.wav", "1", 1, j, id, dateTime, "1")
}
