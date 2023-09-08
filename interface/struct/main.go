package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	jeff := Pessoa{}
	jeff.Nome = "jeffotoni"
	jeff.Idade = 35

	fmt.Println("Nome:", jeff.Nome)
	fmt.Println("Idade:", jeff.Idade)
}
