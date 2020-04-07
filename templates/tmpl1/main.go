package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
	apiVersion: {{.Version}}
	kind: ConfigMap
	metadata:
  		name: {{.Release.Name}}-configmap
	data:
  		myvalue: {{.Title}}
  `

	// Prepare some data to insert into the template.
	type Release struct {
		Name string
	}

	type App struct {
		Version string
		Release Release
		Title   string
	}
	var apps = []App{
		{"1.0.0", Release{"seu nome do metadata one"}, "Programando .Net"},
		{"1.1.0", Release{"seu nome do metadata two"}, "Programando Go"},
		{"1.2.0", Release{"seu nome do metadata tree"}, "Programando Rust"},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range apps {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
