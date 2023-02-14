package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sendSqs(i int) {
	fmt.Println(i, "user")
	time.Sleep(time.Second)
}

func main() {

	http.HandleFunc("/api/user",
		func(w http.ResponseWriter, r *http.Request) {
			// goroutine
			for i := 0; i < 20000; i++ {
				go sendSqs(i)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})
	log.Println("Server run 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
