package main

import (
	"fmt"
	"time"
)

func main() {

	var response_time time.Duration
	request_erro := 0
	request_success := 0
	i := 0

	// goroutine will run for 30 seconds
	go func(i, request_erro, request_success *int, response_time *time.Duration) {
		for {

			x := *i
			x = x + 1
			*i = x
		}
	}(&i, &request_erro, &request_success, &response_time) // passando parametros

	//stop 5s
	time.Sleep(time.Second * 5)

	//show screen
	fmt.Printf("%d;%d;%d\n", i, request_success, request_erro)
}
