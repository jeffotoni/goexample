package main

import (
	"encoding/json"
	"github.com/rs/cors"
	//"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	_                 = iota
	KB            int = 1 << (10 * iota)
	MB            int = 100 << (10 * iota)
	GB            int = 1000 << (10 * iota)
	MaxHeaderByte     = GB
)

var (
	confServer *http.Server
	ServerPort = "8080"

	CorsAllow          = []string{"http://localhost:8282"}
	CorsAllowedMethods = []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}
	CorsAllowedHeaders = []string{"*"}
)

func main() {

	mux := http.NewServeMux()
	corsx := cors.New(cors.Options{
		AllowedOrigins:   CorsAllow,
		AllowedMethods:   CorsAllowedMethods,
		AllowedHeaders:   CorsAllowedHeaders,
		AllowCredentials: true,
	})

	mux.HandleFunc("/auth", Auth)

	// cors allow
	cors.AllowAll().Handler(mux)

	// cors mux
	handlerCors := corsx.Handler(mux)

	confServer = &http.Server{
		Addr:           ":" + ServerPort,
		Handler:        handlerCors,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: MaxHeaderByte,
	}

	println("\033[0;33mRun Server Cors port: 8080\033[0m")
	if err := confServer.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)

	}
}

type AuthJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(w http.ResponseWriter, r *http.Request) {

	println("[auth access]")
	println(r.Header.Get("Content-Type"))

	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {

		r.ParseForm()
		println("first name: " + r.FormValue("email"))
		println("last name: " + r.FormValue("password"))
		jsonMsg := `{"status":"ok", "msg":"tudo ocorreu bem na Cors..., acessou via form urlencoded"}`
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(jsonMsg))
		return
	}

	if r.Header.Get("Content-Type") == "application/json" {

		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	jsonMsg := `{"status":"error", "msg":"tivemos problemas ao ler o body do request: ` + err.Error() + `"}`
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte(jsonMsg))
		// 	return
		// }

		// println(string(b))

		var auth AuthJson
		// err = json.Unmarshal(b, &auth)
		// if err != nil {
		// 	jsonMsg := `{"status":"error", "msg":"tivemos problemas ao ler seu json: ` + err.Error() + `"}`
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte(jsonMsg))
		// 	return
		// }

		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			jsonMsg := `{"status":"error", "msg":"tivemos problemas ao ler seu json: ` + err.Error() + `"}`
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(jsonMsg))
			return
		}

		println("email: " + auth.Email)
		println("password: " + auth.Password)

		jsonMsg := `{"status":"ok", "msg":"recebido seu json com sucesso"}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonMsg))
		return
	}

}
