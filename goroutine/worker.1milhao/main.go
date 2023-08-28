package main

import (
    "os"
    "strconv"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        println("worker", id, "started  job", j)
        time.Sleep(time.Second * 10)
        println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    numRoutines := 1000000
    if len(os.Args) > 1 {
        n, err := strconv.Atoi(os.Args[1])
        if err == nil {
            numRoutines = n
        }
    }

    jobs := make(chan int, numRoutines+1)
    results := make(chan int, numRoutines+1)

    for w := 1; w <= 20000; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numRoutines; j++ {
        jobs <- j
    }

    close(jobs)

    for a := 1; a <= numRoutines; a++ {
        <-results
    }
    println("done")
}
