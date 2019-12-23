package handler

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"pong"}`))
}
