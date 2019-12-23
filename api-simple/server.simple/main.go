package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pkg/handler"
	"pkg/mw"
)

var amw = mw.AuthenticationMiddleware{make(map[string]string)}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/ping", handler.Ping).Methods("POST")
	r.HandleFunc("/my/endpoint/one/{id}", handler.MetodoOne).Methods(http.MethodGet, http.MethodOptions)
	//r.HandleFunc("/my/endpoint/two/{id}", handler.MetodoTwo).Methods("GET")
	//r.HandleFunc("/my/endpoint3/three", handler.MetodoThree).Methods("POST")
	//r.HandleFunc("/my/endpoint4/four", handler.MetodoFour).Methods("PUT")
	amw.Populate()
	r.Use(amw.MiddlewareToken)
	r.Use(mw.Logger)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("Run Service port:8080")
	http.ListenAndServe(":8080", r)
}
