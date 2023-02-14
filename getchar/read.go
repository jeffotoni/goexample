package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Println("Enter an integer value : ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You have entered : ", i)

}
