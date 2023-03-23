package main

import "net/http"

// curl -i -XGET localhost:8080/v1/user
func main() {
	http.HandleFunc("/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Api Go => Rota /v1/user"))
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
