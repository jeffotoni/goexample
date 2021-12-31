package main

import (
	"embed"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var (
	HTTP_PORT = "0.0.0.0:8080"
)

//go:embed static/* static/css/* static/fonts/* static/images/* static/js/*
//go:embed static/index.html
var contentfs embed.FS

type LoginPage struct {
	IfLabelone string
	Labelone   string
	MsgError   string
}

func LoginHtml(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFS(contentfs, "static/index.html")
	if err != nil {
		log.Println("error:", err)
		return
	}
	data := LoginPage{
		Labelone: "Sign in",
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/pedido", LoginHtml)

	fs := http.FileServer(http.FS(contentfs))
	mux.Handle("/", http.StripPrefix("/", fs))
	mux.HandleFunc("/order", Order)

	log.Println("Run Server:", HTTP_PORT)
	http.ListenAndServe(HTTP_PORT, mux)
}

func Order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(`{"msg":"error decode json"}`))
		return
	}
	log.Println(string(b))
	w.WriteHeader(200)
	codeVet := strings.Split(uuid.NewString(), "-")
	w.Write([]byte(`{"idpedido":"` + codeVet[0] + `"}`))
}
