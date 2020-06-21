package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	var ch = make(chan bool)
	for i := 0; i <= 1000; i++ {
		go SendPing()
	}

	<-ch
}

func SendPing() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	endpoint := "http://35.199.113.208/ping"
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return
	}

	req.Header.Set("X-GO-APP-Key", "*%&$*$($(xxx")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error ClientDo:%s", err.Error())
		return
	}

	if response.StatusCode != 200 {
		log.Printf("Error Status:%s", response.Status)
		return
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error Readfile:%s", err.Error())
		return
	}

	log.Printf("...Success %s\n...body\n%s", response.Status, string(b))
	return
}
