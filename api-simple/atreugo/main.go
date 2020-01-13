package main

import (
	"github.com/savsgio/atreugo/v10"
)

func main() {
	config := &atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	// Register a route
	server.Path("GET", "/api/user", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("GET /api/user")
	})

	server.Path("POST", "/api/user", func(ctx *atreugo.RequestCtx) error {

		type t struct {
			Nome  string `json:"nome"`
			Login string `json:login`
		}
		var v = t{"jefferson", "jeffotoni"}
		return ctx.JSONResponse(v)
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
