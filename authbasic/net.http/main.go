package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

const (
	USERNAME = "admin"
	PASSWORD = "password"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !(pair[0] == USERNAME && pair[1] == PASSWORD) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've been authenticated!")
}

func main() {
	http.HandleFunc("/", BasicAuth(HandleRequest))
	http.ListenAndServe(":8080", nil)
}

