package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// não impede a criação
	// de mais threads.
	// Apenas impede que ele
	// execute em
	// múltiplos processadore
	runtime.GOMAXPROCS(-1)

	// criando 100 goroutines
	for i := 0; i < 100; i++ {

		go func() {
			for {
				s := "DevOps Belo Horizonte 2018"
				fmt.Printf("%s\n", s)
				time.Sleep(time.Millisecond * 300)
			}
		}()
	}

	// criando 200 goroutines
	for i := 0; i < 200; i++ {
		go func() {
			for {
				s := "Goroutine 200 vezes."
				fmt.Printf("%s\n", s)
				time.Sleep(time.Millisecond * 200)
			}
		}()
	}

	// stop em 2 segundos
	time.Sleep(time.Second * 2)

	// P = processadores lógicos
	fmt.Println("gmx: ", runtime.GOMAXPROCS(0))

	// numeros de cores
	fmt.Println("cpuf: ", runtime.NumCPU())

	// numero de goroutines
	fmt.Println("goroutine: ", runtime.NumGoroutine())

	println("\n######################################")
	// P = processadores lógicos
	fmt.Println("gmx: ", runtime.GOMAXPROCS(0))

	// numeros de cores
	fmt.Println("cpuf: ", runtime.NumCPU())

	// numero de goroutines
	fmt.Println("goroutine: ", runtime.NumGoroutine())
	println("######################################\n")

}
