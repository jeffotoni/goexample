package main

import (
	"fmt"
	//"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {

		json := `{"total_count":"2929","incomplete_results":"true", "totalitem":"22", "name":"jefferson", "idade":"35"}`
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, json)
	})

	log.Fatal(http.ListenAndServe(":9002", nil))

}
