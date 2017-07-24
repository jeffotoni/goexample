/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"
)

var local = flag.String("local", "", "serve as webserver, example: 0.0.0.0:8000")
var numbPtr = flag.Int("numb", 42, "an int")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Send client
	fmt.Fprintf(w, "Ola sua url é: %q\n", html.EscapeString(r.URL.Path))
}

func main() {

	flag.Parse()

	fmt.Println("local:", *local)
	fmt.Println("local:", *numbPtr)

	// Defining our apis rest
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/insert", viewHandler)

	// Opening a port and listening to it
	http.ListenAndServe(":8080", nil)
}
