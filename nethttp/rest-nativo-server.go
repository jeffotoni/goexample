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

	//curl -X POST localhost:9999/hello -d "name=jefferson"

	http.HandleFunc("/hello", Hello)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
