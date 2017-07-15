/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// client
	fmt.Fprintf(w, "\nWelcome!!\n")
}

func ProductsHandlerForm(w http.ResponseWriter, r *http.Request) {

	// server
	fmt.Println("Name: ", r.PostFormValue("name"))

	// client
	fmt.Fprintf(w, "\nName: %q\n", r.PostFormValue("name"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// server
	fmt.Println("Name: ", vars["name"])

	// client
	fmt.Fprintf(w, "\nName: %q\n", vars["name"])
}

func main() {

	r := mux.NewRouter()
	//r2 := mux.NewRouter()

	//r2.HandleFunc("/postest", HomeHandler)
	r.HandleFunc("/", HomeHandler)

	// curl -X POST localhost:9999/hello -d "name=jefferson"
	// curl -X GET localhost:9999/hello -d "name=jefferson"
	r.HandleFunc("/hello", ProductsHandlerForm)

	// curl -X POST localhost:9999/hello/jeffotoni
	r.HandleFunc("/hello/{name}", ProductsHandler).Methods("POST")

	// curl -X POST localhost:9999
	http.Handle("/", r)

	go http.ListenAndServe(":9999", r)
	go func() { http.ListenAndServe(":8080", nil) }()
	func() { http.ListenAndServe(":8081", nil) }()

}
