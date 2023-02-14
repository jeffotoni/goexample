package main

import (
	"fmt"
	"time"
)

func main() {

	var tempo_resposta time.Duration
	request_erro := 0
	request_success := 0
	i := 0

	// goroutine será executada por 30 segundos
	go func(i, request_erro, request_success *int, tempo_resposta *time.Duration) {
		for {

			x := *i
			x = x + 1
			*i = x
		}
	}(&i, &request_erro, &request_success, &tempo_resposta) // passando parametros

	//parar em 5 segundos
	time.Sleep(time.Second * 5)

	//show
	fmt.Printf("%d;%d;%d\n", i, request_success, request_erro)
}
