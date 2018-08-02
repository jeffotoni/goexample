package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	TotalCount        string `json:"total_count"`
	IncompleteResults bool   `json:"incomplete_results"`
	Name              string `json:"name"`
	//TotaLItem         string `json:"total_item"`
	//TotaLFull         string `json:"total_full"`
	//Repositores       Repositorie `json:"items"`
}

type Repositorie struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
	Stars       string `json:"stargazers_count"`
}

func main() {
	resp, err := http.Get("http://localhost:9002/json")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Body: ", string(body))

	var result Result
	json.Unmarshal(body, &result)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Result: ", result)
	//fmt.Println("Result: ", result.idade)
}
