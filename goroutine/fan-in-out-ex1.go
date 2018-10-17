package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// RECEBENDO VARIOS CANAIS
	c := fanIn(boring("jeff"), boring("Mary"))

	for i := 0; i < 10; i++ {

		fmt.Println(i, " :: ", <-c)
	}

	fmt.Println("Isto é chato estou aqui!")
}

// RETORNANDO CANAIS
func boring(msg string) <-chan string {

	c := make(chan string)

	go func() {

		for i := 0; ; i++ {

			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {

	c1 := make(chan string)

	go func() {
		for {
			c1 <- <-input1
		}
	}()

	go func() {
		for {
			c1 <- <-input2
		}
	}()

	return c1
}
