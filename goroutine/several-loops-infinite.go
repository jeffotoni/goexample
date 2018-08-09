package main

import "fmt"
import "time"

//import "sync"

func LoopOne() {

	// criando func watcher
	doneChan := make(chan bool)
	fmt.Println("Loop One.")

	for {

		fmt.Println("Estou dentro do Loop ##One")
		go func(doneChan chan bool) {

			defer func() {
				doneChan <- true
			}()

			fmt.Println("Estou dentro da funcao do Loop [*One].")

			time.Sleep(1 * time.Second)

		}(doneChan)

		<-doneChan
	}
}

func Loop(str string) {

	//time.Sleep(1 * time.Second)
	fmt.Println("Loop " + str)

	doneChan := make(chan bool)

	for {

		fmt.Println("Estou dentro do Loop **" + str + "**.")

		go func(doneChan chan bool) {

			defer func() {
				doneChan <- true
			}()

			fmt.Println("Estou dentro da funcao do Loop **" + str + "**.")
			time.Sleep(2 * time.Second)

		}(doneChan)

		<-doneChan
		// listar todo banco
		//gbolt.ListAllKeys()

		// <-doneChan

		//}(doneChan)
	}
}

func LoopYes(str string) {

	fmt.Println("Estou dentro da funcao do LoopYes **")

	for {

		fmt.Println("Estou dentro da funcao do Loop **" + str + "**.")
		time.Sleep(1 * time.Second)
	}
}

func LoopYesTwo(str string) {

	fmt.Println("Estou dentro da funcao LoopTwo**")

	for {

		fmt.Println("Estou dentro da funcao do Loop **" + str + "**.")
		time.Sleep(1 * time.Second)
	}
}

func main() {

	//vamos fazer varios loops infinitos e chama-los
	//
	//

	// loop
	// inicial

	// one
	// goroutine
	// go LoopOne()

	//
	// go Loop("2")

	// go Loop("3")

	//go LoopYes("4")

	//LoopYesTwo("5")
	i := 0
	go func(i int) {

		fmt.Println("Estou dentro da funcao da func local **")

		for {
			fmt.Println("dentro do loop ", i)
			<-time.After(time.Duration(3 * time.Second))
			i++
		}
	}(i)

	for {

		fmt.Println("Estou dentro da funcao do Loop...")
		time.Sleep(1 * time.Second)
	}
}
