package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/user",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Evento Go BH ❤️!"))
		})

	http.ListenAndServe("0.0.0.0:8080", mux)
}
