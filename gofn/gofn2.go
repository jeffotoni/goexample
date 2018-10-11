package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

const CONCURRENTS = 200 // containers

func main() {

	wait := &sync.WaitGroup{}
	wait.Add(CONCURRENTS)

	for i := 0; i < CONCURRENTS; i++ {

		go func(i int) {

			cmdD := "docker"
			cmdR := "run"
			cmdRD := "-d"
			cmdRR := "--rm"
			cmdRN := "--name"
			cmdRS := "stopservice" + fmt.Sprintf("%d", i)
			cmdRName := "jeffotoni/stopservicedocker2:latest"
			cmdRm := "-m"
			cmdRmc := "30M" // definindo memory
			cmdRmcs := "--memory-swap"
			cmdRmcs_v := "-1"

			//cmdStr := "docker run -d --rm --name stopservice jeffotoni/stopservicedocker2:latest -m 30m"
			out, err := exec.Command(cmdD, cmdR, cmdRD, cmdRR, cmdRN, cmdRS, cmdRm, cmdRmc, cmdRmcs, cmdRmcs_v, cmdRName).CombinedOutput()
			wait.Done()

			if err == nil {
				fmt.Printf("%s", out)
				time.Sleep(2 * time.Millisecond)
			} else {

				time.Sleep(200 * time.Millisecond)
				fmt.Println(err)
			}
		}(i)
	}

	wait.Wait()
}
