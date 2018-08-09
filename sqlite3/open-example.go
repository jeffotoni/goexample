package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

var CREATE_TABLE = `
 CREATE TABLE "userinfo" (
       "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
        "username" VARCHAR(64) NULL,
        "departname" VARCHAR(64) NULL,
        "created" DATE NULL
    );
    `
var CREATE_TABLE2 = `CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)`

func Listar(db *sql.DB) {

	// db, err := sql.Open("sqlite3", "./foo.db")
	// checkErr(err)

	// time.Sleep(time.Second * 5)

	var id int
	var firstname string
	var lastname string

	for {
		rows, _ := db.Query("SELECT id, firstname, lastname FROM people")

		for rows.Next() {
			rows.Scan(&id, &firstname, &lastname)
			fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)

			time.Sleep(time.Millisecond * 5)
		}

		fmt.Println("::::: INICIANDO NOVAMENTE LISTAR GO-EXAMPLE:::::")
		time.Sleep(time.Second * 2)
	}

	db.Close()
}

func Add(db *sql.DB) {

	// db, err := sql.Open("sqlite3", "./foo.db")
	// checkErr(err)

	//time.Sleep(time.Second * 6)

	statement, _ := db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")

	i := 10
	for {

		nome := "jefferson " + fmt.Sprintf("%d", i)
		sobrenome := "otoni " + fmt.Sprintf("%d", i)

		statement.Exec(nome, sobrenome)

		fmt.Println("inserido com sucesso: ", nome)
		time.Sleep(time.Millisecond * 50)

		i++
	}

	// nome com i

	db.Close()
}

func createDb(db *sql.DB) {

	//db, err := sql.Open("sqlite3", "./foo.db")
	//checkErr(err)

	statement, _ := db.Prepare(CREATE_TABLE2)
	statement.Exec()

	fmt.Println("criado com sucesso..")
	db.Close()
}

func main() {

	db1, _ := sql.Open("sqlite3", "./foo.db")
	db2, _ := sql.Open("sqlite3", "./foo.db")
	db3, _ := sql.Open("sqlite3", "./foo.db")

	//createDb(db3)

	// Add()
	go Listar(db1)

	go Add(db3)

	Listar(db2)

	time.Sleep(time.Second * 10)

	// db, err := sql.Open("sqlite3", "./foo.db")
	// checkErr(err)

	//statement, _ := db.Prepare(CREATE_TABLE)
	//statement.Exec()

	//statement, _ := db.Prepare(CREATE_TABLE2)
	//statement.Exec()

	// // insert
	// stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)
	// // update
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// // query
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)
	// var uid int
	// var username string
	// var department string
	// var created time.Time

	// for rows.Next() {
	//  err = rows.Scan(&uid, &username, &department, &created)
	//  checkErr(err)
	//  fmt.Println(uid)
	//  fmt.Println(username)
	//  fmt.Println(department)
	//  fmt.Println(created)
	// }

	// rows.Close() //good habit to close

	// // delete
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// db.Close()

}

func checkErr(err error) {

	if err != nil {

		fmt.Println("Error sqlite3", err)
		panic(err)
	}
}
