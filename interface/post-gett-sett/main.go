package main

import "fmt"

type Apresentavel interface {
	Apresentar() string
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá! Meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type Estudante struct {
	Pessoa
	Curso string
}

func (e Estudante) Apresentar() string {
	return fmt.Sprintf("Meu nome é %s, eu tenho %d anos e estudo %s.", e.Nome, e.Idade, e.Curso)
}

func main() {
	var pessoa1 Apresentavel = Pessoa{"Silver Gama", 28}
	var pessoa2 Apresentavel = Estudante{Pessoa{"jeffotoni", 21}, "Ciência da Computação"}

	fmt.Println(pessoa1.Apresentar())
	fmt.Println(pessoa2.Apresentar())
}
