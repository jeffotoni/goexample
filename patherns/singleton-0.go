/*
* Example DriverPg Go
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
 */

package main

import (
	"fmt"
	"time"
)

type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

// fazendo uma chamada do
// metodo, antes de qualquer
// chamada
var instanceNew = *Connect()

// funcao retornando
// o ponteiro de nossa
// struct
func Connect() *DriverPg {

	if instance == nil {

		// <--- NOT THREAD SAFE
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func main() {

	// chamada
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println("goroutine 1: ", instanceNew.conn)
	}()

	go func() {

		fmt.Println("goroutine 2: ", *Connect())
	}()

	fmt.Scanln()
}
