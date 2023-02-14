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

func (crw *contextResponseWriter) Flush() {
    if f, ok := crw.ResponseWriter.(http.Flusher); ok {
        f.Flush()
    }
}

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}

type contextResponseWriter struct {
    http.ResponseWriter
    ctx context.Context
}

func contextHandler(ctx context.Context, h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        crw := &contextResponseWriter{rw, ctx}
        h.ServeHTTP(crw, req)
    })
}

func middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        crw := rw.(*contextResponseWriter)
        crw.ctx = newContextWithRequestID(crw.ctx, req)

        h.ServeHTTP(rw, req)
    })
}

func handler(rw http.ResponseWriter, req *http.Request) {
    crw := rw.(*contextResponseWriter)
    reqID := requestIDFromContext(crw.ctx)
    fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
    h := contextHandler(context.Background(), middleware(http.HandlerFunc(handler)))
    http.ListenAndServe(":8080", h)
}
