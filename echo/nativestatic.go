package main

import (
	"net/http"
)

func main() {

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html"))))

	http.ListenAndServe(":8000", nil)
}
