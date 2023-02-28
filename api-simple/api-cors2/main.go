package main

import (
	"io"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Cors struct {
	handler *cors.Cors
}

func New() *Cors {
	return &Cors{
		handler: cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		}),
	}
}

func (c *Cors) Handler(next http.Handler) http.Handler {
	return c.handler.Handler(next)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/v1/user", http.HandlerFunc(User))

	// corsmux := cors.Default().Handler(mux)
	corsmux := New().Handler(mux)

	server :=
		&http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: corsmux,
		}
	log.Println("Server Run port: 0.0.0.0:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error server: %s", err)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error:` + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
