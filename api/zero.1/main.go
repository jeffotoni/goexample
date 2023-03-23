package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rota /v1/user"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Minha api Go!"))
	})

	http.ListenAndServe("0.0.0.0:8080", mux)
}
