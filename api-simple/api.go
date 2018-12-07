package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const LayoutDateLog = "2006-01-02 15:04:05"
const LayoutDate = "2006-01-02"
const LayoutHour = "15:04:05"

type StatusMsg struct {
	Msg  string `json:msg`
	Nome string `json:nome`
}

var (
	err         error
	PORT_SERVER = "5000"
)

func Now() string {

	Now := time.Now().Format(LayoutDateLog)
	Now = fmt.Sprintf("%v", Now)

	return Now
}

func DateHora() string {

	return fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
}

func Date() string {

	return fmt.Sprintf("%s", time.Now().Format("2006-01-02"))
}

func goApiHello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Header.Get("Content-Type") != "application/json" {

		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, "Error content-type")
		return
	}

	if r.Header.Get("X-key") != "123456" {

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error de chave em seu json!")
		return
	}

	if strings.ToUpper(r.Method) == "POST" {

		b, err := ioutil.ReadAll(r.Body)

		defer r.Body.Close()

		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var v = &StatusMsg{}

		err = json.Unmarshal(b, &v)

		if err != nil {

			w.WriteHeader(http.StatusUnprocessableEntity)
			io.WriteString(w, "Error ao fazer merge em struct, confira seu json!")
			return
		}

		// removendo aspas..
		// v.IdUserDb = strings.Replace(v.IdUserDb, `"`, "", -1)

		if strings.ToUpper(v.Msg) == "OK" {

			jsonstr := `{"status":"ok","msg":"Hello ` + v.Nome + `, seja bem vindo, vamos testar os acentos: ~á é í ó ú ç tão lindão!"}`

			w.WriteHeader(http.StatusOK)
			io.WriteString(w, jsonstr) // send data to client side

		} else {

			jsonstr := `{"status":"error","msg":"Hello, error, sua mensagem tem que ter OK! ~á é í ó ú ç tão error...!"}`

			w.WriteHeader(http.StatusOK)
			io.WriteString(w, jsonstr)
		}

	} else {

		jsonstr := `{"status":"error","msg":"Method tem que ser POST.."}`
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, jsonstr) // send data to client side
	}
}

func maxClientsFunc(h http.Handler, n int) http.Handler {

	sema := make(chan struct{}, n)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sema <- struct{}{}

		defer func() { <-sema }()

		h.ServeHTTP(w, r)
	})
}

// decode base64
func decbase64(str64 string) string {

	data, err := base64.StdEncoding.DecodeString(str64)
	if err != nil {
		log.Println("error base64:", err, " Dado: ", str64)
		return str64
	}

	return fmt.Sprintf("%q", data)
}

func init() {

	////////// inicio
	port_tmp := os.Getenv("PORT_SERVER")

	if port_tmp != "" {

		PORT_SERVER = port_tmp

	} else {

		//if for argumentos OK
		if len(os.Args) == 2 && os.Args[1] != "" {

			PORT_SERVER = os.Args[1]

		}
	}

	log.Println("server run port: " + PORT_SERVER) // print

	////////////////////////////////
	///
}

// api-server
// @jeffotoni
func main() {

	const maxClients = 10000

	handlerApiUp := http.HandlerFunc(goApiHello)

	// fazendo o controle de conexoes
	http.Handle("/v1/api/user", maxClientsFunc(handlerApiUp, maxClients))

	srv := &http.Server{

		Addr: ":" + PORT_SERVER,
		// timeout para o nosso http
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe() // set listen port
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
