package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{}

func main() {
	http.HandleFunc("/v1/client", Get)
	log.Println("Run Server port:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Get(w http.ResponseWriter, r *http.Request) {
	body, code, err := connect()
	if err != nil {
		log.Println("Error Server:", err, " code:", code)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return
}

func connect() (body []byte, code int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5000)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/v1/customer", nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	code = resp.StatusCode
	return
}
