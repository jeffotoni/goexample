package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

// Função construtora para criar uma nova Pessoa
func New(nome string, idade int) *pessoa {
	return &pessoa{nome: nome, idade: idade}
}

// Getter para nome
func (p *pessoa) Nome() string {
	return p.nome
}

// Setter para nome
func (p *pessoa) SetNome(nome string) {
	p.nome = nome
}

// Getter para idade
func (p *pessoa) Idade() int {
	return p.idade
}

// Setter para idade
func (p *pessoa) SetIdade(idade int) {
	p.idade = idade
}

func main() {
	jeff := New("jeffotoni", 35)
	fmt.Println("Nome:", jeff.Nome())
	fmt.Println("Idade:", jeff.Idade())
}
