package main

import (
	"encoding/json"
	"fmt"
	"github.com/fanliao/go-promise"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	p := promise.NewPromise()

	p.OnSuccess(func(v interface{}) {

		fmt.Println("onsuccess")

	}).OnFailure(func(v interface{}) {

		fmt.Println("onFailure")

	}).OnComplete(func(v interface{}) {

		fmt.Println("onComplete")

	})

	go func() {
		url := "https://api.postmon.com.br/v1/cep/32371090"

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			p.Reject(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			p.Reject(err)
		}
		p.Resolve(body)
	}()

	r, err := p.Get()

	var result map[string]interface{}
	byt := []byte(fmt.Sprintf("%s", r))
	json.Unmarshal(byt, &result)

	fmt.Println("estado:", result["estado"])
	fmt.Println("cidade:", result["cidade"])
	fmt.Println("Bairro:", result["bairro"])
	fmt.Println("logradouro:", result["logradouro"])
	fmt.Println("cep:", result["cep"])

	//info, ok := result["cidade_info"]
	//fmt.Println("info:", info, ok)
	fmt.Println("info:", result["cidade_info"])

	//dumpMap("", result)

	fmt.Println(err)

}

func FuncTestOne() {

	time.Sleep(time.Second * 2)
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {

			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")

		} else {

			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}
