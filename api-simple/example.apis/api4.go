// Go in action
// @jeffotoni
// 2019-04-05

/// curl localhost:8080/api/hello
//////////////////////////////////////

package main

import (
	//"expvar"
	//_ "expvar"
	"log"
	"net/http"
)

// func init() {

// 	http.HandleFunc("/debug/vars", expvarHandler)

// 	Publish("cmdline", Func(cmdline))

// 	Publish("memstats", Func(memstats))
// }

// var (
// 	exp_points_processed = expvar.NewInt("points_processed")
// )

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, welcome to the world, Go!"))
}

func main() {

	mux := http.NewServeMux()

	//mux.Handle("/api/ping", http.HandlerFunc(Hello))
	mux.Handle("/api/hello", http.HandlerFunc(Hello))
	// mux.Handle("/api/products", http.HandlerFunc(Hello))
	// mux.Handle("/api/products/{id}", http.HandlerFunc(Hello))

	server :=
		&http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Eror while serving metrics: %s", err)
	}
}
