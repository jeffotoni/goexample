package main

import "fmt"

func main() {
	//Criando map
	agenda := map[string] int{

		"uSER" : 777777,
		"USER2" : 888888,

	}
	fmt.Println(agenda)

	//Mostrando apenas um valor

	fmt.Println(agenda["USER10"])

	//adicionando outro valor

	agenda["Go"] = 111111

	fmt.Println(agenda)

	//valor existente ou não

	if sera, ok := agenda["Lang"]; !ok{

		fmt.Println("Não tem")
	} else {
		fmt.Println(sera)
	}


	//Range

	amigos := map [int] string{

		1 : "Wolfs",
		2 : "Evolutie",
	}
	fmt.Println(amigos)

	for key, value := range amigos {

		fmt.Println(key,value)
	}

	//Deletando

	delete(amigos, 2)

	fmt.Println(amigos)

}