package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

// curl -i -XGET localhost:8080/v1/user
func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/v1/user":
		if http.MethodGet == r.Method {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Rota /v1/user")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

	case "/v1/user/:codigo":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Minha api Go!")

	default:
		http.NotFound(w, r)
	}
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe("0.0.0.0:8080", mux)
}
