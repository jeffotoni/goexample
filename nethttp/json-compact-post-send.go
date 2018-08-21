package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8189/hello"
	fmt.Println("URL:>", url)

	// values := map[string]string{"username": username, "password": password}
	// jsonValue, _ := json.Marshal(values)
	// resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast.\n"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "Golang Client")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
