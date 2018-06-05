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
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()
		fmt.Fprintf(w, "\nName: %q\n", r.PostFormValue("name"))

	} else if r.Method == "GET" {

		http.Error(w, "POST only", http.StatusMethodNotAllowed)
	}
}

func main() {

	log.Println("log... teste...")
	//curl -X POST localhost:9999/hello -d "name=jefferson"

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Accept", "Content-Type"})  jefferson
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// corsObj := handlers.AllowedOrigins([]string{"*"})

	http.HandleFunc("/hello", Hello)

	log.Println("Listening...")
	//log.Fatal(http.ListenAndServe(":9999", handlers.CORS(headersOk, methodsOk, corsObj)(router)))

	log.Fatal(http.ListenAndServe(":9999", nil))
}
