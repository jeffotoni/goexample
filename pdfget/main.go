package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ModelPdf struct {
	Base64 string `json:"base64"`
}

func GerarPdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	timeout := time.Duration(15) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	url := "https://www.thecampusqdl.com/uploads/files/pdf_sample_2.pdf"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var j ModelPdf
	w.WriteHeader(http.StatusOK)
	sEnc := b64.StdEncoding.EncodeToString(b)
	j.Base64 = string(sEnc)
	b, err = json.Marshal(&j)
	if err != nil {
		log.Println(err)
		return
	}
	println(string(b))
}

func main() {
	log.Println("Run Server port:8005")
	http.HandleFunc("/", GerarPdf)
	log.Fatal(http.ListenAndServe(":8005", nil))
}
