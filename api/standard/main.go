package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jeffotoni/goexample/api/standard/crypt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", Healthz)
	mux.HandleFunc("/auth/check", Check)
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	println("Run Server 0.0.0.0:8080")
	log.Fatal(s.ListenAndServe())
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID, err := crypt.RandomID()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"Bad Request RandomID"}`))
		return
	}
	w.Header().Set("ID", ID)
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID, err := crypt.RandomID()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"Bad Request RandomID"}`))
		return
	}
	w.Header().Set("ID", ID)
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"ok", "token":"xxxxxxxxxxxxxxxxxx3993"}`))
}
