package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// canais
	c1 := make(chan int)
	c2 := make(chan int)

	// enviando dados
	// para dois canais
	go populate(c1)

	// transformando os dois
	// canais com dados diferentes
	// em um unico canal
	go fanOutIn(c1, c2)

	// listando o canal unico
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("pronto finalizei")
}

// enviando dados para os canais
// 2 canais diferentes
func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// pegando os canais e convertendo em um unico canal
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {

			c2 <- timeConsumingWork(v2)
			wg.Done()

		}(v)
	}
	wg.Wait()

	// se nao fechar fanin
	// ira ocorrer all goroutines are asleep - deadlock!
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
