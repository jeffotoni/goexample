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

type Driver struct {
	sLock    sync.RWMutex
	instance string
}

type DriverPg struct {
	conn string
	Driver
}

// funcao retornando
// o ponteiro de nossa
// struct
func (D *DriverPg) Connect() *DriverPg {

	// ainda não está perfeita,
	// não é totalmente atomico
	if D.Driver.instance == "" {

		D.Driver.sLock.RLock()

		D.Driver.instance = D.conn // <-- THE TREAD SAFE

		D.Driver.sLock.RUnlock()
	}

	return D
}

func main() {

	// <-- aqui é sicrono
	// Declarando a struct Driver
	var driver = &DriverPg{conn: "DriverConnectPostgres"}
	db := driver.Connect()

	// <-- apartir daqui é assicrono
	// chamada 1
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(db.Driver.instance)
		//fmt.Println(*Connect())
	}()

	// chamada 2
	go func() {

		time.Sleep(time.Millisecond * 1200)
		fmt.Println(db.Driver.instance)
		//fmt.Println(*Connect())
	}()

	// goroutine 2
	go func() {

		for i := 0; i < 100; i++ {

			time.Sleep(time.Millisecond * 700)

			// varias goroutines
			go func(i int) {

				time.Sleep(time.Millisecond * 350)
				fmt.Println("Singleton email :", i, db.Driver.instance)

			}(i)
		}
	}()

	fmt.Scanln()
}
