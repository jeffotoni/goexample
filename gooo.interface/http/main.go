package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type dataSource interface {
	Get(url string) ([]byte, error)
}

type httpbin struct{}
type httpbin2 struct{}

func (l *httpbin) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("url get error: %s\n", err)
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("body read error: %s\n", err)
		return []byte{}, err
	}

	return body, nil
}

func (l *httpbin2) Get(url string) ([]byte, error) {
	return []byte("jef"), nil
}

func process(n int, ds dataSource) (string, error) {
	url := fmt.Sprintf("http://httpbin.org/links/%d/0", n)

	resp, err := ds.Get(url)
	if err != nil {
		fmt.Printf("data source get error: %s\n", err)
		return "", err
	}

	return string(resp), nil
}

func main() {
	data, err := process(5, &httpbin{})
	if err != nil {
		fmt.Printf("\ndata processing error: %s\n", err)
		return
	}
	fmt.Printf("\nSuccess: %v\n", data)

	data, err = process(6, &httpbin2{})
	if err != nil {
		fmt.Printf("\ndata processing error: %s\n", err)
		return
	}
	fmt.Printf("\nSuccess: %v\n", data)
}
