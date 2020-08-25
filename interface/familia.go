package main

import (
    "fmt"
)

type Familia interface {
    Dados() string
}

type Pai struct {
    Nome  string
    Idade int
    Cpf   string `json:"cpf"`
}

func (p Pai) Dados() string {
    return fmt.Sprintf("Nome: %s, Idade: %d", p.Nome, p.Idade)
}

type Filho struct {
    Pai   Pai
    Idade int
    Nome  string
    Email string
}

func (f Filho) Dados() string {
    return fmt.Sprintf("\nNome: %s, Idade: %d, Email: %s", f.Nome, f.Idade, f.Email)
}

type Filhos []Filho

func (f Filhos) Dados() string {
    concat := ""
    for _, v := range f {
        concat += fmt.Sprintf("\nNome: %s, Idade: %d, Email: %s", v.Nome, v.Idade, v.Email)
    }
    return concat
}

func showDados(membro Familia) {

    fmt.Println(membro.Dados())
}

func showDados2(f []Familia) {
    for _, membro := range f {
        fmt.Println(membro.Dados())
    }
}

func main() {

    //// Pai
    pai := new(Pai)
    pai.Nome = "Jefferson"
    pai.Idade = 38
    pai.Cpf = "00.xxx.xxx-xx"

    var filho Filho
    var filhos Filhos

    //// Filhos
    filho.Pai.Nome = "Pai Nome Here"
    filho.Pai.Idade = 38
    filho.Pai.Cpf = "01.xxx.xxx-xx"

    filho.Nome = "Arthur"
    filho.Email = "arthur@gmail.com"
    filho.Idade = 4
    filhos = append(filhos, filho)

    filho.Nome = "Francisco"
    filho.Email = "francisco@gmail.com"
    filho.Idade = 7
    filhos = append(filhos, filho)

    //// Filha
    filha := new(Filho)
    filha.Nome = "Joyce"
    filha.Idade = 22
    filha.Email = "joycexxx@gmail.com"

    // Show Dados
    showDados(pai)

    showDados(filha)

    showDados(filhos)

    idadeFilhos := []Familia{pai, filha, filhos}
    showDados2(idadeFilhos)

}
