package main

import (
	"fmt"
)

type Interface interface {
	create() bool
}

type Interface2 interface {
	Interface
}

type User struct {
	Name string
}

var u User

func (u *User) Set(s string) *User {
	return &User{
		Name: s,
	}
}

func (u *User) create() bool {
	fmt.Println("User: ", u.Name)
	return true
}

type Endereco struct {
	Cidade string
}

var e Endereco

func (e *Endereco) Set(s string) *Endereco {
	return &Endereco{
		Cidade: s,
	}
}

func (e *Endereco) create() bool {
	fmt.Println("cidade: ", e.Cidade)
	return true
}

func Create(i Interface2) bool {
	return i.create()
}

func main() {
	// user
	u := &User{"Ketlen"}
	fmt.Println(Create(u))
	// ou
	fmt.Println(Create(u.Set("jeffotoni")))
	// ou
	fmt.Println(Create(&User{"Adicelia"}))

	// endereco
	fmt.Println(Create(e.Set("Belo Horizonte")))
}
