package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	idx   int
	value int
}

func main() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	out := make([]int, len(data))

	jobs := make(chan job)

	go func() {
		for i, val := range data {
			jobs <- job{i, val}
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for job := range jobs {
				time.Sleep(time.Millisecond * time.Duration(randGen.Intn(10)))

				if job.idx == 0 {
					time.Sleep(time.Second * 2)
				}

				fmt.Println("job", job)
				out[job.idx] = job.value
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(out)
}
