package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type User struct {
	Id    string `json:id`
	Name  string `json:name`
	Email string `json:email`
	Db    string `json:db`
	Msg   string `json:msg`
}

func main() {

	//start := time.Now()

	client := &http.Client{

		Timeout: time.Second * 2, // Maximum of 2 secs

	}

	//apiUrl := "http://192.168.35.251:9010"
	apiUrl := "http://10.221.94.103:9010"

	resource := "/api/user"

	i := 0

	var tempo_resposta time.Duration

	request_erro := 0
	request_success := 0

	go func(i, request_erro, request_success *int, tempo_resposta *time.Duration) {

		for {

			x := *i
			x = x + 1
			*i = x

			data := url.Values{}

			Db64 := lbase64("gofn")

			user := &User{
				Id:    "1",
				Name:  "jeffotoni",
				Email: "jeff.otoni@s3wf.com.bt",
				Db:    Db64,
				Msg:   "ok",
			}

			userJson, err := json.Marshal(user)

			if err != nil {

				log.Println("Error marshal: ", err)
				continue
			}

			u, _ := url.ParseRequestURI(apiUrl)
			u.Path = resource
			urlStr := u.String()
			// fmt.Println(urlStr)
			// r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
			r, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(userJson)) // URL-encoded payload

			if err != nil {
				log.Println("Error newRequest: ", err)
				return
			}
			r.Header.Add("X-key", "1234567890x1234#4")
			r.Header.Add("Authorization", "Bearer ac2168444f4de69c27d6384ea2ccf61a49669be5a2fb037ccc1f")
			r.Header.Add("Content-Type", "application/json")
			r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
			sti := time.Now()
			// tempo de resposta
			// inicio
			resp, err := client.Do(r)
			if err != nil {

				// erro
				// nao respondeu
				e := *request_erro
				e = e + 1
				*request_erro = e

				continue
			}

			body, err := ioutil.ReadAll(resp.Body)
			//defer resp.Body.Close()
			if err != nil {
				continue
			}

			defer resp.Body.Close()

			stf := time.Now()
			stfsub := stf.Sub(sti)
			*tempo_resposta += stfsub

			if resp.StatusCode == 200 {

				// contabiliza... ok
				s := *request_success
				s = s + 1
				*request_success = s

			} else {

				e := *request_erro
				e = e + 1
				*request_erro = e
			}

			fmt.Sprintf("%s", string(body))

			// fmt.Println("status: ", resp.StatusCode)
			// fmt.Println("status: ", http.StatusText(resp.StatusCode))

		}
	}(&i, &request_erro, &request_success, &tempo_resposta)

	//parar em 1 minuto
	time.Sleep(time.Second * 30)
	//fim := time.Now()

	//quantidade...
	fmt.Printf("%d;%d;%d\n", i, request_success, request_erro)

	//fmt.Println("Evios de Requests: ", i)
	//fmt.Println("Duration total: ", fim.Sub(start))
	//fmt.Println("tempo total de requests: ", tempo_resposta)
	//fmt.Println("Requests Success: ", request_success)
	//fmt.Println("Requests Error: ", request_erro)

}

func lbase64(str string) string {

	data := []byte(str)
	str64 := base64.StdEncoding.EncodeToString(data)
	return str64
}
