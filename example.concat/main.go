package main

import (
	"github.com/jeffotoni/gconcat"
)

func main() {
	var ii []interface{}
	ii = append(ii, "jeffotoni")
	ii = append(ii, " ")
	ii = append(ii, "joao")
	ii = append(ii, " ")
	ii = append(ii, 2021)

	var i interface{}
	i = "jeffotoni"

	println(gconcat.Build(ii))
	println(gconcat.Build(i))
	println(gconcat.Build("jeffotoni", " & ", "joao", " ", 20, "/08/", 2020))
	println(gconcat.Build([]string{"2015", " ", "2016", " ", "2017", " ", "2018", " ", "2020"}))
	println(gconcat.Build([]int{12, 0, 11, 0, 10, 11, 12, 23, 3}))
	println(gconcat.Build(10, 9, 10, 20, 30, 40, "x", "&", "."))
	println(gconcat.Build("R$ ", 23456.33, " R$ ", 123.33, " R$ ", 19.11))

}
