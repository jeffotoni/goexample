package main

import (
	"fmt"
	"sync"
)

func main() {

	// canais
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// enviando dados
	// para dois canais
	go send(even, odd)

	// transformando os dois
	// canais com dados diferentes
	// em um unico canal
	go recive(even, odd, fanin)

	// listando o canal unico
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("pronto finalizei")
}

// enviando dados para os canais
// 2 canais diferentes
func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			even <- i * 10
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

// pegando os canais e convertendo em um unico canal
func recive(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {

		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()

	// se nao fechar fanin
	// ira ocorrer all goroutines are asleep - deadlock!
	close(fanin)
}
