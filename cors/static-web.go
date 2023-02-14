package main

import (
	"net/http"
	"text/template"
)

var port string = "9090"
var tmpl *template.Template

type Login struct {
	MsgError      string
	IfLabelone    string
	Labelone      string
	IfLabeloneSub string
}

func HandlerLoginHtml(w http.ResponseWriter, r *http.Request) {

	println(r.URL.String())

	if r.URL.String() == "/login-json" {

		tmpl = template.Must(template.ParseFiles("login-json.html"))

	} else if r.URL.String() == "/login-form" {

		tmpl = template.Must(template.ParseFiles("login-form.html"))

	} else {

		tmpl = template.Must(template.ParseFiles("login-json.html"))
	}

	login := Login{
		MsgError:      "",
		IfLabelone:    "",
		Labelone:      "logar!",
		IfLabeloneSub: "logar submit",
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

	println("\033[0;36mRun Server Cors port: " + port + "\033[0m")
	http.ListenAndServe(":"+port, mux)
}
