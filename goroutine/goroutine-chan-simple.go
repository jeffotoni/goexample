package main

import (
	"fmt"
	"time"
)

// escrevendo no canal
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("escrever:", i, "to ch")
	}
	close(ch)
}

func main() {

	// channel com buffer
	ch := make(chan int, 2)

	// goroutine
	go write(ch)

	//aguarde um pouco
	time.Sleep(1 * time.Second)

	// listando o canal
	for v := range ch {
		fmt.Println("ler", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
