// Go in action
// @jeffotoni

/// curl localhost:8080/api/hello
//////////////////////////////////////

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, welcome to the world, Go!"))
	}))
}

// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Hello, welcome to the world, Go!"))
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// func Hello(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hello, welcome to the world, Go!"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/hello", http.HandlerFunc(Hello))
// 	server :=
// 		&http.Server{
// 			Addr:    ":8080",
// 			Handler: mux,
// 		}
// 	server.ListenAndServe()
// }
