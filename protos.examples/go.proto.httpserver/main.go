package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	protoc "github.com/jeffotoni/go.protobuffer.customer"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/customer/proto", Customer)
	println("Run Server: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Customer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte(`{"msg":"error method!"}`))
		return
	}
	if r.Header.Get("Content-Type") != "application/proto" {
		w.WriteHeader(400)
		w.Write([]byte(`{"msg":"error content-type!"}`))
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		log.Println("Error ioutil:", err)
		return
	}

	var pCustomer = &protoc.Customer{}
	err = proto.Unmarshal(b, pCustomer)
	if err != nil {
		w.WriteHeader(400)
		log.Println("unmarshaling error invalid: ", err)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(`{"id":` + strconv.Itoa(int(pCustomer.Id)) + `,"name":"` + pCustomer.Name + `"}`))
}
