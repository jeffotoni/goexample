package main

import "time"

func main() {

	concurrency := 3

	sem := make(chan bool, concurrency)

	urls := []string{"url1", "url2", "url3", "url4", "url5", "url6", "url7"}

	for _, url := range urls {

		println("entrei for nivel 1")

		sem <- true

		go func(url string) {

			println("entrei func nivel1")

			defer func() {

				println("entrei func nivel 2")
				time.Sleep(time.Second * 3)
				<-sem
			}() // get the url

		}(url)
	}

	for i := 0; i < cap(sem); i++ {

		println("for nivel 2")
		sem <- true
	}
}
