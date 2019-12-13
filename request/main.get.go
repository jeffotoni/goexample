// Front-end in Go server
// @jeffotoni

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func ApiGet(urlApi, source string) (string, int) {

	resource := "/products/send"

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(10*time.Second, func() {
		cancel()
	})
	defer afterFuncTimer.Stop()
	u, _ := url.ParseRequestURI(urlApi)
	u.Path = resource
	urlStr := u.String()
	req, err := http.NewRequest("GET", urlStr, nil)
	req = req.WithContext(ctx)

	req.Header.Add("X-User-xxx", "here-user")
	req.Header.Add("X-Api-key", "here-key")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error client.Do: ", err.Error())
		return "", 500
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll: ", err)
	}
	return string(body), resp.StatusCode
}

func main() {

	body, code := ApiGet("https://golang.org", "/pkg")
	fmt.Println(code)
	fmt.Println(body)

}
