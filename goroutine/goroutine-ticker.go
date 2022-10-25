package main

import "time"
import "fmt"

func main() {

	// Simular o time de parada
	// O channel é setado com valor
	// usamos o `range` para listar o conteúdo
	// do channel.
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers podem ser parado como timers.
	// quando parado não receberá mais values
	// em seu channel para depois de 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
