package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var client = &http.Client{}

func main() {
	http.HandleFunc("/v1/client", Get)
	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET] /v1/client")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Engine", "Go")
	w.Header().Set("Location", "/v1/client")
	w.Header().Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, code, err := connect()
	if err != nil {
		log.Println("Error Server:", err, " code:", code)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(``))
		return
	}

	length := strconv.Itoa(len(body))
	w.Header().Set("Content-Length", length)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}

func connect() (body []byte, code int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	var Url string = "http://localhost:3000/v1/customer"
	req, err := http.NewRequestWithContext(ctx, "GET", Url, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	//req.ContentLength = int64(-1)
	//req.TransferEncoding = []string{"identity"}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	code = resp.StatusCode
	return
}
