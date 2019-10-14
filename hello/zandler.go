package main

import "fmt"

type Filhos []struct {
	Nome, Nascimento string
}

var zandler = Filhos{
	{"Benício", "11/12/2015"},
	{"Théo", "11/12/2015"},
}

func main() {
	fmt.Println(zandler)
}
