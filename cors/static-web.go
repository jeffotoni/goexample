package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var port string = "8282"
var tmpl *template.Template

type Login struct {
	MsgError   string
	IfLabelone string
	Labelone   string
}

func HandlerLoginHtml(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)

	if r.URL.String() == "/login-json" {

		tmpl = template.Must(template.ParseFiles("login-json.html"))

	} else if r.URL.String() == "/login-form" {

		tmpl = template.Must(template.ParseFiles("login-form.html"))

	} else {

		tmpl = template.Must(template.ParseFiles("login-json.html"))
	}

	login := Login{
		MsgError:   "",
		IfLabelone: "",
		Labelone:   "logar!",
	}

	tmpl.Execute(w, login)
}

func main() {

	mux := http.NewServeMux()

	// retorna HTML
	mux.HandleFunc("/login-json", HandlerLoginHtml)

	mux.HandleFunc("/login-form", HandlerLoginHtml)

	// fisico
	fs := http.FileServer(http.Dir("./web/static"))

	// vitual
	mux.Handle("/web/static/",
		http.StripPrefix("/web/static", fs))

	fmt.Println("Server Run: " + port)
	http.ListenAndServe(":"+port, mux)
}
