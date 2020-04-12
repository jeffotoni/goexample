package main

import (
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	println("server: hello handler started")
	defer println("server: hello handler ended")

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("hello!"))
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			println(err.Error())
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
