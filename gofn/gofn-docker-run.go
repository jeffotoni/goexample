package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gofn/gofn"
	"github.com/gofn/gofn/provision"
)

// CONCURRENTS
// Number of containers
const CONCURRENTS = 100

func main() {

	wait := &sync.WaitGroup{}
	wait.Add(CONCURRENTS)
	for i := 0; i < CONCURRENTS; i++ {

		go func() {

			buildOpts := &provision.BuildOptions{
				ImageName:               "jeffotoni/testerstress",
				DoNotUsePrefixImageName: true,
			}

			containerOpts := &provision.ContainerOptions{

				//Cmd:   []string{"sh", "-c", "wget https://google.com"},
				//Cmd:   []string{},
				Image: "jeffotoni/testerstress",
			}

			stdout, _, err := gofn.Run(context.Background(), buildOpts, containerOpts)

			wait.Done()

			if err != nil {
				log.Println(err)
			}
			//fmt.Println("Stderr: ", stderr)
			fmt.Printf("%s", stdout)
		}()

	}
	wait.Wait()
}
