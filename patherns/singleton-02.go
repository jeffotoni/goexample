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

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

// variavel Global
var instance *DriverPg

// funcao retornando
// o ponteiro de nossa
// struct
func Connect() *DriverPg {

	lock.Lock() // <--- Desnecessario a lock se a instancia já tiver sido criada
	defer lock.Unlock()

	if instance == nil {

		instance = &DriverPg{conn: "DriverConnectPostgres"}
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

	fmt.Scanln()
}
