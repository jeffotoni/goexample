package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func NovaPessoa(nome string, idade int) *Pessoa {
	return &Pessoa{nome: nome, idade: idade}
}

func main() {
	jeff := NovaPessoa("jeff", 30)
	fmt.Println(jeff.nome, jeff.idade)
}
