// Golang In Action
// @package     main
// @author      @jeffotoni
// @size        2019

package main

import (
  "fmt"
  "strings"
)

type Duck interface {
  Voar()
  //Nadar()
}

type Pato struct{ Name string }
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
  a.Name = "Pato"
  b := Galinha{"Galinha"}

  ShowAnimal(a)
  ShowAnimal(b)
}
