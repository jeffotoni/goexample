/*
* Example Worker
*
* @package     main
* @author      @jeffotoni
* @sice        01/08/2017
*
 */

// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
	"time"
)

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("👷 Worker %d → started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("✅ Worker %d → finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {

	numJobs := 5
	numWorkers := 3

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, numJobs)
	results := make(chan int, numWorkers)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("🔁 Result received:", result)
	}
}
