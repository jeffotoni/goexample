package main

import "fmt"

type Speaker interface {
	Say(string)
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Println(p.name+":", message)
}

func SpeakAlphabet(via Speaker) {
	via.Say("abcdefghijlmn")
}

func main() {
	mat := new(Person)
	mat.name = "Jeff"
	SpeakAlphabet(mat)
}
