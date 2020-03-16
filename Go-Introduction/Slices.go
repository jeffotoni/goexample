package main

import "fmt"

func main() {
	//Cria uma slice
	nome := []int {1,2,3}

	fmt.Println(nome)



	//Range slice
	slice := []string{"A", "B", "C", "D","E"}
	slice2 := []string{"Carro", "Moto", "Aviao"}
	for indice, valor := range slice {

		fmt.Println("No indice ", indice ,"temos o valor ", valor)
	}

	//Fatiando Slice
	fatia := slice[0:4]

	//Anexando elementos com Append

	slice2 = append(slice2, "FORD","Ferrari")

	fmt.Println(slice2)

	//Anexando Slices

	slice2 = append(slice2, slice...)

	fmt.Println(slice2)

	//Deletando

	fmt.Println(fatia)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)


	//Make Slice

	slice3 := make([]int, 2,2)

	//Recebe posições

	slice3[0], slice3[1] = 1,2

	fmt.Println(slice3)

	//A capacidade cap dobra ao adicionar com append

	slice3= append(slice3, 3)

	fmt.Println(slice3, len(slice3), cap(slice3))

	//Slice multidimensional

	ss := [][]int{


		[]int{0,1,2}, //Indice 0
		[]int{2,3,4}, //Indice 1
		[]int{5,6,7}, //Indice 2

	}
	fmt.Println(ss)

}