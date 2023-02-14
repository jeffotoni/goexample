// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// Funcao que escreve um texto no arquivo e retorna um erro caso tenha algum problema
func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func main() {

	var conteudo []string
	conteudo = append(conteudo, "123")
	conteudo = append(conteudo, "456")
	conteudo = append(conteudo, "789")

	err := escreverTexto(conteudo, "foo.txt")
	if err != nil {
		log.Fatalf("Erro:", err)
	}

	go func() {

		var j int
		var over [][]int

		for {

			j++
			time.Sleep(time.Microsecond)
			fmt.Println("teste 2: ", j)

			// eating a lot of memory
			conteudo = append(conteudo, "789-"+strconv.Itoa(j))
			//a := make([]int, 0, 999999)
			a := make([]int, 0, 99)
			over = append(over, a)

			err := escreverTexto(conteudo, "foo.txt")
			if err != nil {
				log.Fatalf("Erro:", err)
			}

		}
	}()

	var i int
	for {
		i++
		//time.Sleep(time.Second * 3)
		fmt.Println("teste", i)

		d1 := []byte(`hello ` + strconv.Itoa(i) + `\ngo\n`)
		ioutil.WriteFile("foo2.txt", d1, 0644)

	}
}
