package main

import (
	"fmt"
	"plugin"
)

func main() {
	plug, err := plugin.Open("soma_plugin.so")
	if err != nil {
		fmt.Println("Erro ao carregar o plugin:", err)
		return
	}

	somaFunc, err := plug.Lookup("Soma")
	if err != nil {
		fmt.Println("Erro ao buscar a função 'Soma':", err)
		return
	}

	soma, ok := somaFunc.(func(int, int) int)
	if !ok {
		fmt.Println("Erro ao fazer type assertion para a função 'Soma'")
		return
	}

	resultado := soma(20, 15)
	fmt.Println("Resultado da soma: ", resultado)
}
