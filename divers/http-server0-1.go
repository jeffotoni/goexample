/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
	"html"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Send client
	fmt.Fprintf(w, "Ola sua url é: %q\n", html.EscapeString(r.URL.Path))
}

func main() {

	// Defining our apis rest
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/view", viewHandler)

	// Opening a port and listening to it
	go func() { http.ListenAndServe(":8080", nil) }()
	http.ListenAndServe(":9001", nil)
}
