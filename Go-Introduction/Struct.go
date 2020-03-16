package main

import "fmt"

//Struct
type Pessoa struct {
	nome string
	idade int
}

//struct embutido

type Profissao struct {
	Pessoa
	cargo string
	salario int
}





func main() {
	//Utilizando Struct

	Carlos := Profissao{
		Pessoa:  Pessoa{"User", 20},
		cargo:   "Developer",
		salario: 10000,
	}

	fmt.Println(Carlos)
	fmt.Println(Carlos.idade)

}
