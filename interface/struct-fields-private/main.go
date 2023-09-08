package main

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
}

// Getter para nome
func (p *Pessoa) Nome() string {
	return p.nome
}

// Setter para nome
func (p *Pessoa) SetNome(nome string) {
	p.nome = nome
}

// Getter para idade
func (p *Pessoa) Idade() int {
	return p.idade
}

// Setter para idade
func (p *Pessoa) SetIdade(idade int) {
	p.idade = idade
}

func main() {
	jeff := &Pessoa{}
	jeff.SetNome("jeff")
	jeff.SetIdade(30)

	fmt.Println("Nome:", jeff.Nome())
	fmt.Println("Idade:", jeff.Idade())
}
