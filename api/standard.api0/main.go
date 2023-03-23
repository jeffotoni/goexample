package main

import "net/http"

// curl -i -XGET localhost:8080
func main() {
	http.ListenAndServe("0.0.0.0:8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Minha api Go!"))
		}))
}
