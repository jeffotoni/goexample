package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var cs3 = make(chan bool)
	var csqs = make(chan bool)
	var cwalk = make(chan bool)

	// criando map
	// thread safe
	var sm sync.Map

	go func() {

		for {
			// local
			//sem deadlock

			// map
			sm.Store("login", `{"login":"jeffotoni","id":"123","nome":"jefferson otoni","acess":"12","ativo":true}`)
			x := false
			cs3 <- x
			csqs <- x
			cwalk <- x
		}

		//close(cs3)
		//close(csqs)
		//close(cwalk)
	}()

	go func() {

		for {
			time.Sleep(time.Second * 1)
			x := <-cs3
			fmt.Println("Go1 S3: ", x)

			result, ok := sm.Load("login")

			if ok {
				fmt.Println("Result Go1 S3: ", result)
			} else {
				fmt.Println("erro result Load - Go1 S3", ok)
			}

			if x {

				fmt.Println("Go1 S3 processar...")
			}
		}
	}()

	go func() {

		for {
			time.Sleep(time.Second * 2)
			x := <-csqs
			fmt.Println("Go2 SQS: ", x)

			result, ok := sm.Load("login")

			if ok {
				fmt.Println("Result Go1 SQS: ", result)
			} else {
				fmt.Println("erro result Load - Go1 SQS", ok)
			}

			if x {
				fmt.Println("Go2 SQS processar...")
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			x := <-cwalk
			fmt.Println("Go3 Walk: ", x)

			result, ok := sm.Load("login")

			if ok {
				fmt.Println("Result Go1 Walk: ", result)
			} else {
				fmt.Println("erro result Load - Go1 Walk", ok)
			}

			if x {
				fmt.Println("Go3 Walk processar... ")
			}
		}

	}()

	time.Sleep(time.Hour * 10)
}
