// Golang In Action
// @package     main
// @author      @jeffotoni
// @size        2019

package main

import (
	"fmt"
	"github.com/bet365/jingo"
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
	Oculto int    // anything we don't annotate doesn't get emitted.
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

// Create an encoder, once, letting it know which type of struct we're going to be encoding.
var enc = jingo.NewStructEncoder(Pato{})

func main() {

	//a := Pato{"Pato"}
	a := new(Pato)
	a.Id = 1000
	a.Idade = 3
	a.Name = "Pato"
	a.Oculto = 100

	b := Galinha{"Galinha"}

	ShowAnimal(a)
	ShowAnimal(b)

	// pull a buffer from the pool and pass it along with the struct to Marshal
	buf := jingo.NewBufferFromPool()
	enc.Marshal(a, buf)

	fmt.Println(buf)
}
