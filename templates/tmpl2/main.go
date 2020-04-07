package main

import (
	"log"
	"os"
	"text/template"
)

var (
	path_template = "./template1.cs"
	path_create   = "./template1.create.cs"
)

// Prepare some data to insert into the template.
type Release struct {
	Name string
}

type App struct {
	Version string
	Release Release
	Title   string
	Kind    string
}

func main() {

	//tmpl, err := template.ParseFiles(path_template)
	// or
	// tmpl := template.Must(template.ParseFiles(path_template))

	tmpl := template.Must(template.ParseFiles(path_template))

	var apps = App{"1.0.0", Release{"seu nome do metadata one"}, "Programando .Net", "ConfigMap"}

	f, err := os.Create(path_create)
	if err != nil {
		log.Println(err)
		return
	}

	err = tmpl.Execute(f, &apps)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("criado com sucesso")
}
