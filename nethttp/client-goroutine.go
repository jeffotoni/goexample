package main

import (
	"fmt"
	"net/http"
)

func main() {
	requests := 10000
	ch := make(chan string, requests)
	for i := 1; i <= requests; i++ {
		go func(iter int) {
			fmt.Println(iter)
			resp, err := http.Get("http://localhost:8080/api/project")
			if err == nil {
				resp.Body.Close()
			}
			ch <- fmt.Sprint(iter)
		}(i)
	}
	for i := 1; i <= requests; i++ {
		fmt.Println(<-ch)
	}
}
