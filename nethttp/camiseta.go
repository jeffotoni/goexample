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
)

func Hello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()
		fmt.Fprintf(w, "\nName: %q\n", r.PostFormValue("name"))

	} else if r.Method == "GET" {

		http.Error(w, "POST only", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8189", nil)
}
