// nao fazam isto
// isto so utilizei para testar e entender o
// comportamento de context usando handler
package main

import (
	"context"
	"fmt"
	"net/http"
)

type key int

const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

func middleware(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = newContextWithRequestID(ctx, req)
		h.ServeHTTPContext(ctx, rw, req)
	})
}

func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	reqID := requestIDFromContext(ctx)
	fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

func main() {
	h := &ContextAdapter{
		ctx:     context.Background(),
		handler: middleware(ContextHandlerFunc(handler)),
	}
	http.ListenAndServe(":8080", h)
}
