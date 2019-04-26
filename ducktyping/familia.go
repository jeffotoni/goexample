// Golang In Action
// @package     main
// @author      @jeffotoni
// @size        2019

package main

import (
    "encoding/json"
    "fmt"
)

type familia interface {
    dados() string
}

type pai struct {
    Nome  string
    Idade int
    Cpf   string `json:"campo-x"`
}

func (p pai) dados() string {
    return fmt.Sprintf("Nome: %s, Idade: %d", p.Nome, p.Idade)
}

type filho struct {
    pai
    email string
}

func (f filho) dados() string {
    return fmt.Sprintf("Nome: %s, Idade: %d, Email: %s", f.Nome, f.Idade, f.email)
}

func mostraDados(membro familia) {
    fmt.Println(membro.dados())
}
func main() {
    pai := new(pai)
    pai.Nome = "Jeff"
    pai.Idade = 50
    pai.Cpf = "00.xxx.xxx-xx"

    filho := new(filho)
    filho.Nome = "Arthur"
    filho.Idade = 20
    filho.email = "arthur@gmail.com"

    mostraDados(pai)
    mostraDados(filho)

    b, err := json.Marshal(pai)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(b))
}
