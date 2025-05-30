package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/rs/cors"
)

var (
	blackList = map[string]string{
		"xx.xx.xx.xx": "xx.xx.xx.xx",
	}
)

//go:embed dist/skote
var contentfs embed.FS

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.FS(contentfs))
	mux.Handle("/", authIp(http.StripPrefix("", fs)))
	handler := cors.Default().Handler(mux) // cors
	log.Println("Executando o servidor na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func authIp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		_, ok := blackList[ip]
		if ok {
			http.Error(w, "Autenticação inválida", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
