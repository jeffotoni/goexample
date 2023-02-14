package main

import (
	"fmt"
	"sync"
)

func sum2(a []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("summing: ", a)
	total := 0
	for _, v := range a {
		total += v
	}
	//fmt.Println("send to c",total)
	c <- total // send total to c
}

func sum(a []int, c chan int) {

	fmt.Println("summing: ", a)

	total := 0
	for _, v := range a {
		total += v
	}
	//fmt.Println("send to c",total)
	c <- total // send total to c
}

func aleatorio() {

	c := make(chan int)

	go sum([]int{1, 2, 3}, c)
	go sum([]int{4, 5, 6}, c)

	x := <-c
	fmt.Println(x)

	x = <-c
	fmt.Println(x)
}

func naOrdem() {

	c := make(chan int)

	go sum([]int{1, 2, 3}, c)

	//use the channel to block until it receives a send
	x := <-c
	fmt.Println(x)

	//then execute the next routine
	go sum([]int{4, 5, 6}, c)

	x = <-c
	fmt.Println(x)
}

func waitCanalNaOrdem() {

	c := make(chan int)
	wg := new(sync.WaitGroup)

	// chama simultaneamente as chamadas simultâneas para soma, permitindo que a execução continue no intervalo do canal
	go func() {
		// incrementa o grupo de espera e passa para a função sum para decrementá-lo quando estiver completo
		wg.Add(1)
		go sum2([]int{1, 2, 3, 4, 5, 6, 6, 7, 87, 8, 8, 9, 9, 9, 9, 0, 0, 0, 0, 0, 34, 4, 4, 4}, c, wg)

		// espera que a chamada acima seja completada
		wg.Wait()

		//and repeat...
		wg.Add(1)
		go sum2([]int{4, 5, 6}, c, wg)
		wg.Wait()
		// todas as chamadas estão completas, feche o canal para permitir que o programa saia de forma limpa
		close(c)

	}()

	//range of the channel
	for theSum := range c {
		x := theSum
		fmt.Println(x)
	}
}

func main() {

	// aguardando..
	waitCanalNaOrdem()

	//aproveitando
	//os bloqueios
	//para determinar
	//naOrdem()

	//aleatorio()
}
