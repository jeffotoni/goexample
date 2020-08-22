package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gconcat"
)

const PORT = ":8282"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			str := gconcat.Build([]int{1, 2, 3, 4, 5}, " ", []string{"vamos testar nossa concat!!!"})
			w.Write([]byte(str))
		})
	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}
	println("Start Run: ", PORT)
	log.Fatal(server.ListenAndServe())
}
