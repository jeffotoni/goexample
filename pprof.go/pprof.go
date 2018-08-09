package main

import (
	"net/http"
	"net/http/pprof"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", hiHandler)

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
