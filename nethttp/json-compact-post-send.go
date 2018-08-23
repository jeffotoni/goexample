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

	var jsonStr = []byte(`{
"id": 0,
"jsonrpc": "2.0",
"result": {
"address": "47VDcqrjmfuBv5Z45uAGCjiG2Y8uMmyewPWJzciszdN64NLigTUNjh6DJLkTC2TnU2JvhFqyins7qJYrN98Wxm7MHsZJPUb",
"addresses": [{
"address": "47VDcqrjmfuBv5Z45uAGCjiG2Y8uMmyewPWJzciszdN64NLigTUNjh6DJLkTC2TnU2JvhFqyins7qJYrN98Wxm7MHsZJPUb",
"address_index": 0,
"label": "Primary account",
"used": false
},{
"address": "862mLCYYkh4AU2e2RYGx51STwE16aNLsFf7jhkggLYzD7ZuHhycdGtYArYbW2nFtYG7HwsLxcQ7LsGka3YHKHpTbBgmEhm2",
"address_index": 1,
"label": "teste",
"used": false
}]
}
}`)
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
