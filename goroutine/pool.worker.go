package main

import (
	"sync"
	"time"
)

// Job define job interface
type Job interface {
	Execute()
}

// JobHandler handle consume worker queue
type JobHandler struct {
}

// Execute do the work
func (job *JobHandler) Execute() {
	// do the actual work
}

// Worker defines worker struct
type Worker struct {
	JobQueue chan Job
}

// WorkerQueue queue of worker job queue
var WorkerQueue chan chan Job

// WorkerNum number of workers
var WorkerNum int

func main() {
	wg := new(sync.WaitGroup)
	WorkerNum = 100
	wg.Add(WorkerNum)
	WorkerQueue := make(chan chan Job, WorkerNum)
	for i := 0; i < WorkerNum; i++ {
		// start a worker
		go func() {
			worker := Worker{
				JobQueue: make(chan Job, 1),
			}
			wg.Done()
			for {
				// register channel
				WorkerQueue <- worker.JobQueue
				// wait for job comming
				job := <-worker.JobQueue
				// do work
				job.Execute()
			}
		}()
	}
	wg.Wait()
	timeout := 300 * time.Millisecond
	ticker := time.NewTicker(timeout)
	// consume worker
	select {
	case <-ticker.C:
		panic("timeout reached!")
	case worker := <-WorkerQueue:
		// push job
		job := &JobHandler{}
		worker <- job
	}

}
