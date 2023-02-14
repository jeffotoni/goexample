package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var (
	path_template = "./aluno.cs"
	path_create   = "./aluno.create.cs"
)

// Prepare some data to insert into the template.
type Table struct {
	Name    string
	Columns []Columns
}

type Columns struct {
	Type string
	Name string
}

type App struct {
	Table Table
}

func main() {

	tmpl := template.Must(template.ParseFiles(path_template))

	var apps = App{Table{"MyClassAluno", []Columns{{"ClassVirtualAlunoOne", "ClassAlunoOne"}, {"ClassVirtualAlunoTwo", "ClassAlunoTwo"}}}}

	//ShowTmpl(tmpl, apps)

	f, err := os.Create(path_create)
	if err != nil {
		log.Println(err)
		return
	}

	err = tmpl.Execute(f, apps)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("criado com sucesso")
}

func ShowTmpl(tmpl *template.Template, apps App) {
	err := tmpl.Execute(os.Stdout, apps)
	if err != nil {
		fmt.Println(err)
		return
	}
}
