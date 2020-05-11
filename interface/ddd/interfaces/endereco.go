package interfaces

import "fmt"

type Endereco struct {
	Cidade string
}

var e Endereco

func EndSet(s string) *Endereco {
	return &Endereco{
		Cidade: s,
	}
}

func (e *Endereco) create() bool {
	fmt.Println("cidade: ", e.Cidade)
	return true
}
