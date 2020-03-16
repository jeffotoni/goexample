package main

import "fmt"

func main() {

	//switch

	s := 5

	switch  {

	case s < 5:

		fmt.Println("S é menor que 5")

	case s ==5:

		fmt.Println("S é igual a 5")

		//fallthrough execeuta direto o proximo caso
		fallthrough
	case s > 5:

		fmt.Println("S é maior que 5")

	}

}
