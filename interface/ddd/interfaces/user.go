package interfaces

import (
	"fmt"
)

type user struct {
	Name string
}

var U user

func UserSet(s string) *user {
	return &user{
		Name: s,
	}
}

func (u *user) create() bool {
	fmt.Println("user: ", u.Name)
	return true
}
