package main

import "fmt"

type Login struct {
	Email string
	Id    int
}

type Loginx []*Login

//func NameLogin(B []*Login) {
func NameLogin(B Loginx) Loginx {

	//var B []Login
	//B = []*Login{}
	bar := new(Login)
	bar.Email = "jefferson"
	bar.Id = 200
	B = append(B, bar)

	return B
}

func AltInt(foo *[]int) {
	bar := []int{1, 2, 3}
	*foo = append(*foo, bar...)
}

func main() {

	fmt.Println("vim-go")
	b1 := Loginx{}
	//b1 := []*Login{}
	n := NameLogin(b1)
	for k, v := range n {
		fmt.Println("dados: ", k, v)
	}
}
