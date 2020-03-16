package main

import "fmt"
//função Basica

func Soma (x, y int) int {

	return x+y
}

//defer -> fechar por ultimo

func main() {

	teste := Soma(5,5)
	fmt.Println(teste)

}
