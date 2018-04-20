package main

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

	println("\n" + a.Name)
	Sum(5, 5)
}
