/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page      int
	Framework []string
}

type Response2 struct {
	Page      int      `json:"page"`
	Framework []string `json:"Frameworkorks"`
}

func main() {

	// Create empty slice of
	// struct pointers.
	var resF = []*Response1{}

	// Create struct and
	// append it to the slice
	res0D := new(Response1)

	res0D.Page = 1
	res0D.Framework = []string{"Macaron"}

	resF = append(resF, res0D)

	res1D := &Response1{
		Page:      2,
		Framework: []string{"Gorilla", "Echo", "Gin"},
	}

	resF = append(resF, res1D)

	res3D := new(Response1)

	res3D.Page = 1
	res3D.Framework = []string{"Macaron"}

	resF = append(resF, res3D)

	res1B, _ := json.Marshal(resF)

	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:      1,
		Framework: []string{"Symfony", "Laravel", "Django"}}

	res2B, _ := json.Marshal(res2D)

	fmt.Println("")
	fmt.Println(string(res2B))
}
