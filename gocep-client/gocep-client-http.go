package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jeffotoni/gocep/pkg/cep"
)

var (
	Port = ":8080"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cep/", HandlerCep)
	mux.HandleFunc("/cep", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Println("port", Port)
	log.Fatal(server.ListenAndServe())
}

func HandlerCep(w http.ResponseWriter, r *http.Request) {

	cepstr := strings.Split(r.URL.Path[1:], "/")[1]
	if len(cepstr) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := cep.Search(cepstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}
