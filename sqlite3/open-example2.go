package main

import (
	"database/sql"
	"fmt"
	//"github.com/google/gops/agent"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	_ "net/http/pprof"

	"flag"
	"os"
	"runtime"
	"runtime/pprof"
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

		fmt.Println(rows)

		if rows != nil {
			for rows.Next() {

				rows.Scan(&id, &firstname, &lastname)
				fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)

				time.Sleep(time.Millisecond * 200)
			}

			fmt.Println("::::: INICIANDO NOVAMENTE LISTAR GO-EXAMPLE2:::::")

		} else {

			fmt.Println("::::: NAO ENCONTROU NADA NO BANCO LISTAR GO-EXAMPLE2:::::")
		}

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

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()

	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create CPU profile: ", err)
	// 	}
	// 	if err := pprof.StartCPUProfile(f); err != nil {
	// 		log.Fatal("could not start CPU profile: ", err)
	// 	}

	// 	//f, _ := os.Create("/tmp/profile.pb.gz")
	// 	defer f.Close()
	// 	runtime.GC()
	// 	pprof.WriteHeapProfile(f)

	// 	defer pprof.StopCPUProfile()
	// }

	// // ... rest of the program ...

	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create memory profile: ", err)
	// 	}
	// 	runtime.GC() // get up-to-date statistics
	// 	if err := pprof.WriteHeapProfile(f); err != nil {
	// 		log.Fatal("could not write memory profile: ", err)
	// 	}
	// 	f.Close()
	// }

	//Add(db1)

	// log.Println(http.ListenAndServe("localhost:6060", nil))

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	//db1, _ := sql.Open("sqlite3", "./foo.db")
	db2, _ := sql.Open("sqlite3", "./foo.db")
	db3, _ := sql.Open("sqlite3", "./foo.db")

	chanStruct := make(chan struct{})

	go Listar(db3, chanStruct)

	<-chanStruct

	// createDb()

	chanStruct2 := make(chan struct{})

	go Listar(db2, chanStruct2)

	<-chanStruct2

	// we need a webserver to get the pprof webserver
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}

		//f, _ := os.Create("/tmp/profile.pb.gz")
		defer f.Close()
		runtime.GC()
		pprof.WriteHeapProfile(f)

		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	time.Sleep(time.Hour)

	// f, _ := os.Create("./profile.pb.gz")

	// defer f.Close()

	// runtime.GC()

	// pprof.WriteHeapProfile(f)

	// time.Sleep(10 * time.Second)

}

func checkErr(err error) {

	if err != nil {

		time.Sleep(time.Second * 30)
		fmt.Println("Error sqlite3", err)
		panic(err)
	}
}
