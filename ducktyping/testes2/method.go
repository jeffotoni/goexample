package main

import (
	"fmt"
)

type Login struct {
	Uuid   string
	User   string
	Status int
}

func (a *Login) SetLogin() {
	fmt.Println("Uuid: ", a.Uuid, "User: ", a.User)
}

func main() {
	User := "jeffotoni lima jeff!"
	Uuid := "c7c8acb3-2566-4311-4b23-210ee26a86ed"
	l := Login{Uuid: Uuid, User: User, Status: 10}
	l.SetLogin()
}
