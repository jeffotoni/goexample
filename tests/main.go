package main

import (
	"log"
)

type App struct {
	Name string
}

func Sum(x int, y int) int {
	return x + y
}

var a App

func main() {
	a = App{}
	a.Name = "jefferson"

	log.Println("teste..")
	println("\n" + a.Name)
	Sum(5, 5)
}
