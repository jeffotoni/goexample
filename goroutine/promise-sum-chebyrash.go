package main

import (
	"errors"
	"fmt"
	"github.com/chebyrash/promise"
)

func main() {

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {

		// Faça algo de forma assíncrona
		const sum = 2 + 2

		// Se o seu trabalho foi bem sucedido, chame resolve () passando o resultado.
		if sum == 4 {
			resolve(sum)
			return
		}

		// Se você encontrou um erro, chame reject () passando o erro.
		if sum != 4 {
			reject(errors.New("2 + 2 não é igual 4"))
			return
		}

		// Se você esqueceu de verificar se há erros e sua função entra em pânico, a promise
		// automatically reject.
		// panic() == reject()
	})

	// Uma promise é um objeto retornado ao qual você anexa retornos de chamada
	p.Then(func(data interface{}) interface{} {
		fmt.Println("Resultado 1:", data)
		// soma o resultado
		return data.(int) + 3
	})

	// Uma promise é um objeto retornado ao qual você anexa retornos de chamada
	p.Then(func(data interface{}) interface{} {
		fmt.Println("Resultado 2:", data)
		// soma o resultado
		return data.(int) + 10
	})

	// Os retornos de chamada podem ser adicionados mesmo após o sucesso ou a falha da operação assíncrona.
	// Vários retornos de chamada podem ser adicionados chamando. Then ou .Catch várias vezes,
	// para ser executado independentemente em ordem de inserção.
	p.
		Then(func(data interface{}) interface{} {
			fmt.Println("Resultado novo:", data)
			return nil
		}).
		Catch(func(error error) error {
			fmt.Println("Error during execution:", error.Error())
			return nil
		})

	// Como as chamadas de retorno são executadas de forma assíncrona, você pode esperar por elas.
	p.Await()
}
