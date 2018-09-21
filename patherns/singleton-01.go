/*
* Example Pattern Singleton
*
* @package     main
* @author      @jeffotoni
* @size        09/09/2018
*
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")

	DB_PORT  = "5432"
	DB_SSL   = "disable"
	DB_SORCE = "postgres"
)

// Struct singleton
type Singleton struct {
	Pgdb *sql.DB
}

var (
	err    error
	PostDb Singleton
)

func PgConnect() {

	if PostDb.Pgdb == nil {

		//return PostDb.Pgdb
		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

		// Struct Singleton
		PostDb.Pgdb, err = sql.Open(DB_SORCE, DBINFO) // <--- NOT SAFE

		if err != nil {
			log.Println("Error PgConnect::", err)
		}
	}
}

func getEmail(logi_id int) (email string) {

	err := PostDb.Pgdb.QueryRow("select logi_email FROM ad_login where logi_id=$1", logi_id).Scan(&email)

	if err != nil {

		log.Println("error ao buscar email!")
	}

	return
}

// usando
// init
// para subir
// struct
// singleton
func init() {

	// connect
	// ele é
	// chamado
	// depois
	// da variavel
	// instanciada
	// e antes
	// do func
	// main
	PgConnect()
}

func main() {

	fmt.Println("Singleton::", PostDb.Pgdb.Ping())

	// goroutine 1
	go func() {

		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Println("Singleton email  3:", getEmail(3))
		}
	}()

	// goroutine 2
	go func() {

		for i := 0; i < 100; i++ {

			time.Sleep(time.Millisecond * 50)

			// varias goroutines
			go func() {
				time.Sleep(time.Millisecond * 300)
				fmt.Println("Singleton email 16:", getEmail(16))

			}()
		}
	}()

	// goroutine 3
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Singleton email  1:", getEmail(1))
	}()

	// presione
	// uma tecla
	// para sair
	fmt.Scanln()
}
