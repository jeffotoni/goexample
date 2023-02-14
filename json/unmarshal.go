// Go in action
// @jeffotoni
// 2019-04-06

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	//"log"
	//"strconv"
)

/*
"numero_casas": 10,
"token":"token_do_usuario",
"cifrado": "texto criptografado",
"decifrado": "aqui vai o texto decifrado",
"resumo_criptografico": "aqui vai o resumo
*/

type Cifra struct {
	NumeroCasas         int    `json:"numero_casas"`
	Token               string `json:"token"`
	Cifrado             string `json:"cifrado"`
	Decifrado           string `json:"decifrado"`
	ResumoCriptografico string `json:"resumo_criptografico"`
}

// type Cifra struct {
// 	numero_casas         int
// 	token                string
// 	cifrado              string
// 	decifrado            string
// 	resumo_criptografico string
// }

func getJson(url string, target *Cifra) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&target)
}

func main() {

	url := "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=xxxxxxxxxxxxxxx"

	c := Cifra{}

	err := getJson(url, &c)

	if err != nil {
		os.Exit(0)
	}

	fmt.Println(c)
	os.Exit(0)

}
