package main

import (
	"log"
	"os"
	"text/template"
)

const letter = `
	apiVersion: {{.Version}}
	kind: ConfigMap
	metadata:
  		name: {{.Release.Name}}-configmap
	data:
  		myvalue: {{.Title}}
  `

type Release struct {
	Name string
}

type App struct {
	Version string
	Release Release
	Title   string
}

func main() {

	var apps = []App{
		{"1.0.0", Release{"metadata one"}, "Lang .Net"},
		{"1.1.0", Release{"metadata two"}, "Lang Go"},
		{"1.2.0", Release{"metadata tree"}, "Lang Rust"},
	}

	tmplFunc(apps)
}

func tmplFunc(apps []App) {
	t := template.Must(
		template.New("letter").Parse(letter),
	)
	for _, r := range apps {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
