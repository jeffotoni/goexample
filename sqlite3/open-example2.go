package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func Listar(db *sql.DB, chanStruct chan<- struct{}) {

	// db, err := sql.Open("sqlite3", "./foo.db")
	// checkErr(err)

	time.Sleep(time.Second * 1)

	var id int
	var firstname string
	var lastname string

	//go func() {

	chanStruct <- struct{}{}

	for {

		//time.Sleep(time.Millisecond * 300)

		rows, _ := db.Query("SELECT id, firstname, lastname FROM people")

		for rows.Next() {

			rows.Scan(&id, &firstname, &lastname)
			fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)

			time.Sleep(time.Millisecond * 200)
		}

		fmt.Println("::::: INICIANDO NOVAMENTE LISTAR GO-EXAMPLE2:::::")
		time.Sleep(time.Second * 1)

	}
	//}()
	//
	////<-chanStruct

	db.Close()
}

func Add(db *sql.DB) {

	time.Sleep(time.Second * 2)
	// checkErr(err)

	//time.Sleep(time.Second * 6)

	statement, _ := db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")

	i := 3000
	for {

		nome := "jefferson " + fmt.Sprintf("%d", i)
		sobrenome := "otoni " + fmt.Sprintf("%d", i)

		statement.Exec(nome, sobrenome)

		fmt.Println("inserido com sucesso: ", nome)
		time.Sleep(time.Millisecond * 1)

		i++
	}

	// nome com i

	db.Close()
}

func main() {

	//db1, _ := sql.Open("sqlite3", "./foo.db")
	db2, _ := sql.Open("sqlite3", "./foo.db")
	db3, _ := sql.Open("sqlite3", "./foo.db")

	chanStruct := make(chan struct{})

	go Listar(db3, chanStruct)

	<-chanStruct

	// createDb()

	chanStruct2 := make(chan struct{})

	Listar(db2, chanStruct2)

	<-chanStruct2

	//Add(db1)

	time.Sleep(time.Second * 5)

}

func checkErr(err error) {

	if err != nil {

		time.Sleep(time.Second * 30)
		fmt.Println("Error sqlite3", err)
		panic(err)
	}
}
