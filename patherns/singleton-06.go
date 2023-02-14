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
	"time"
)

// call somente
// uma unica vez
var once sync.Once

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

	once.Do(func() {

		instance = &DriverPg{conn: "DriverConnectPostgres"}
	})

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

	// 100 goroutine
	for i := 0; i < 100; i++ {

		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", Connect().conn)
		}(i)
	}

	fmt.Scanln()
}
