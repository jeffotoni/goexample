package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Crawl(url string, depth int, fetcher int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	see := randInt(1, 10)
	time.Sleep(time.Second * time.Duration(see))
	return fmt.Sprintf(" - secound %s", strconv.Itoa(see))
}

func main() {

	var depth, fetcher int
	var urls = []string{"http://google.com", "http://s3wf.com.br", "http://s3apis.com"}
	done := make(chan string)
	for _, u := range urls {
		fmt.Printf("-> Crawling %s\n", u)
		go func(url string) {
			second := Crawl(url, depth-1, fetcher)
			done <- url + second
		}(u)
	}
	defer close(done)

	for range urls {
		fmt.Printf("<- Waiting for next child\n")
		u := <-done
		fmt.Printf("  Done... %s\n", u)
	}

	time.Sleep(time.Second * 1)
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("pronto")
}
