package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	proxyStr := "http://usuario:senha@proxy.br:3128"
	proxyURL, _ := url.Parse(proxyStr)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	req, _ := http.NewRequest("GET", "https://minhaurl", nil)
	req.Header.Set("Authorization", "Basic usuario:senha")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}
