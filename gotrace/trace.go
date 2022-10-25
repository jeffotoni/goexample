package main

import (
	"os"
	"runtime/trace"
	"strconv"
	"time"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	concurrency := 100000

	sem := make(chan bool, concurrency)

	urls := []string{}

	// 1000 urls
	for ii := 0; ii < 100000; ii++ {

		urls = append(urls, "url"+strconv.Itoa(ii))
	}

	for _, url := range urls {

		//println("entrei for nivel 1")

		sem <- true

		go func(url string) {

			//println("entrei func nivel1")

			defer func() {

				//println("entrei func nivel 2")
				time.Sleep(time.Second * 3)
				<-sem
			}() // get the url

		}(url)
	}

	for i := 0; i < cap(sem); i++ {

		//println("for nivel 2")
		sem <- true
	}

	defer trace.Stop()

	// Your program here
}
