package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func checkUser(sm *sync.Map) bool {

	sm.Store("login", `{"login":"jeffotoni","id":"123","nome":"jefferson otoni","acess":"12","ativo":true}`)
	return true

}

func SetCount(sm *sync.Map) {

	valor := GetCount(sm)
	valor = valor + 1
	sm.Store("count", valor)
}

func GetCount(sm *sync.Map) int {

	result, _ := sm.Load("count")
	x := fmt.Sprintf("%d", result)
	y, _ := strconv.Atoi(x)
	//fmt.Println("result:", y)
	return y
}

func main() {

	var cs3 = make(chan bool)
	var csqs = make(chan bool)
	var cwalk = make(chan bool)

	// criando map
	// thread safe
	var smx = &sync.Map{} // correct para usar como referencia e alterar em todas funcoes
	// var smx *sync.Map // error

	//checkUser := func(sm sync.Map) bool {

	//sm.Store("login", `{"login":"jeffotoni","id":"123","nome":"jefferson otoni","acess":"12","ativo":true}`)
	//return true
	//}(smx)

	// checando user
	// sincrono uma
	// unica vez
	// ja pre-carrega
	// dados json user
	checkUser(smx)

	// criando um contador
	smx.Store("count", 0)

	// carregando
	// dados
	// de
	// usuario
	// check user
	go func(sm *sync.Map) {

		//dados, ok := sm.Load("login")

		for {

			time.Sleep(time.Millisecond * 700)

			//fmt.Println(GetCount(sm))

			//if ok {

			//encontrou dado usuario e se tiver ativo, executar todas as goroutines

			//} else {

			// nao encontrou valor pausa todas goroutines
			//}
			// ficar sempre
			// checando user
			// ativo ou
			// desativado
			x := checkUser(sm)
			cs3 <- x
			csqs <- x
			cwalk <- x
		}

		//close(cs3)
		//close(csqs)
		//close(cwalk)
	}(smx)

	// Envia os arquivos para S3
	// Consumir channel se true
	// Envia para S3
	// O S3 ira consumir o banco de 10 em 10
	// e setando os enviados com sucesso
	// todos serao enviados para o S3
	go func(sm *sync.Map) {

		for {

			time.Sleep(time.Second * 6)
			x := <-cs3

			if x {

				fmt.Println("Go1 S3: ", x)
				result, ok := sm.Load("login")

				if ok {

					// quebra o json localmente
					fmt.Println("Result Go1 S3: ", result)
					fmt.Println("Go1 S3 processar...")

					// uma forma de capturar valor em mpa
					// do json
					data := []byte(fmt.Sprintf("%s", result))

					// usando json.RawMessage
					//Source := (*json.RawMessage)(&data)
					//var objmap map[string]*json.RawMessage
					//err := json.Unmarshal(*Source, &objmap)
					//////

					// transformando json em objeto
					var objmap map[string]interface{}
					err := json.Unmarshal(data, &objmap)

					if err != nil {

						fmt.Println("Error json", err)

					} else {
						//fmt.Println("Json Processado no S3: ", string(*objmap["login"]))
						fmt.Println("Json Processado no S3: ", objmap["login"])
					}

				} else {

					fmt.Println("erro result Load - Go1 S3", ok)
				}
			}
		}
	}(smx)

	// lendo erros se encontrados
	// dos envios de arquivos
	// os arquivos serao processados
	// depois que os arquivos foram processados
	// precisara quardar em uma fila os erros
	// para que possa ser reenviado novamente
	go func(sm *sync.Map) {

		for {

			time.Sleep(time.Second * 10)
			x := <-csqs

			if x {

				fmt.Println("Go2 Sqs: ", x)
				result, ok := sm.Load("login")

				if ok {

					// quebra o json localmente
					fmt.Println("Result Go2 Sqs: ", result)
					fmt.Println("Go2 sqs processar...")

				} else {

					fmt.Println("erro result Load - Go2 Sqs", ok)
				}
			}
		}
	}(smx)

	// O Walk ler recursivamente
	// a estrutura do diretorio
	// mapeando todos arquivos
	// setando eles para serem
	// enviados
	go func(sm *sync.Map) {

		for {

			//fmt.Println(GetCount(sm))
			if GetCount(sm) == 0 {

				time.Sleep(time.Second * 1)
				SetCount(sm)

			} else {

				time.Sleep(time.Second * 9)
			}

			// primeira vez tempo em segundos
			// segunda vez em hora em hora
			// ou quando der algum erro no
			// notify

			x := <-cwalk

			if x {

				fmt.Println("Go3 walk: ", x)
				result, ok := sm.Load("login")

				if ok {

					// quebra o json localmente
					fmt.Println("Result Go3 Walk: ", result)
					fmt.Println("Go3 Walk processar...")

					// uma forma de capturar valor em mpa
					// do json
					data := []byte(fmt.Sprintf("%s", result))

					// usando json.RawMessage
					//Source := (*json.RawMessage)(&data)
					//var objmap map[string]*json.RawMessage
					//err := json.Unmarshal(*Source, &objmap)
					//////

					// transformando json em objeto
					var objmap map[string]interface{}
					err := json.Unmarshal(data, &objmap)

					if err != nil {

						fmt.Println("Error json Walk:", err)

					} else {
						//fmt.Println("Json Processado no S3: ", string(*objmap["login"]))
						fmt.Println("Json Processado no Walk: ", objmap["login"])
					}

				} else {

					fmt.Println("erro result Load - Go3 Walk", ok)
				}
			}
		}
	}(smx)

	time.Sleep(time.Hour * 10)
}
