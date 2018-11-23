package main

import "fmt"

type Login struct {
	Email string
	Id    int
}

func NameLogin(r *[]Login) {

	bar := Login{
		Email: "exampel@example.com",
		Id:    123,
	}
	*r = append(*r, bar)

}

func main() {

	r := []Login{}

	NameLogin(&r)
	NameLogin(&r)

	for k, v := range r {
		fmt.Println("dados: ", k, v)
	}

}
