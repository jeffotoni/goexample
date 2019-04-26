// Golang In Action
// @package     main
// @author      @jeffotoni
// @size        2019

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Duck interface {
	Voar()
	//Nadar()
}

type Pato struct {
	Name   string `json:"name"`
	Idade  int    `json:"idade"`
	Id     int    `json:"id"`
	Oculto int    `json:",omitempty"` // se for vazio não mostra
}

type Galinha struct{ Name string }

func (a Pato) Voar() {
	fmt.Println(strings.ToLower(a.Name) + " voa!")
}

func (a Galinha) Voar() {
	fmt.Println(a.Name + " voa!")
}

func ShowAnimal(duck Duck) {
	duck.Voar()
	//duck.Nadar()
}

func main() {

	//a := Pato{"Pato"}
	a := new(Pato)
	a.Id = 1000
	a.Idade = 3
	a.Name = "Pato"
	//a.Oculto = 100

	b := Galinha{"Galinha"}

	ShowAnimal(a)
	ShowAnimal(b)

	// pull a buffer from the pool and pass it along with the struct to Marshal

	buf, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
