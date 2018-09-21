/*
* Example DriverPg Go
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// manter o estado
var atomicinz uint64

// lock mutex
var lock = &sync.Mutex{}

// driver
type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

// funcao retornando
// o ponteiro de nossa
// struct
func Connect() *DriverPg {

	// garantindo que já entrou
	if atomic.LoadUint64(&atomicinz) == 1 {

		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	// entra somente uma
	// únic vez
	if atomicinz == 0 {

		instance = &DriverPg{conn: "DriverConnectPostgres"}
		atomic.StoreUint64(&atomicinz, 1)
	}

	return instance
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	go func() {

		fmt.Println(*Connect())
	}()

	// 50 goroutine
	for i := 0; i < 50; i++ {
		go func() {
			for {
				time.Sleep(time.Millisecond * 60)
				fmt.Println(Connect().conn, " - ", i)
			}
		}()
	}

	fmt.Scanln()
}
