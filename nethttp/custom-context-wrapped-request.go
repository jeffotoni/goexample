// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

type Handler func(*Context)

type RouteStruct struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type App struct {
	Routes       []RouteStruct
	DefaultRoute Handler
}

func Route() *App {
	app := &App{
		DefaultRoute: func(ctx *Context) {
			ctx.Text(http.StatusNotFound, "Not found")
		},
	}

	return app
}

func (a *App) Handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := RouteStruct{Pattern: re, Handler: handler}

	a.Routes = append(a.Routes, route)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{Request: r, ResponseWriter: w}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (c *Context) Text(code int, body string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)

	io.WriteString(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}

func main() {
	app := Route()

	app.Handle(`^/hello$`, func(ctx *Context) {
		ctx.Text(http.StatusOK, "Hello world")
	})

	app.Handle(`/hello/([\w\._-]+)$`, func(ctx *Context) {
		ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s", ctx.Params[0]))
	})

	err := http.ListenAndServe(":9000", app)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
