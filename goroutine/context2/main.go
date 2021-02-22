// nao fazam isto
// isto so utilizei para testar e entender o
// comportamento de context usando handler
package main

import (
    "context"
    "fmt"
    "net/http"
    "sync"
)

var cmap = map[*http.Request]*context.Context{}
var cmapLock sync.Mutex

type key int

const requestIDKey key = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}

// Note that we are returning a pointer to the context, not the
// context itself.
func contextFromRequest(req *http.Request) *context.Context {
    cmapLock.Lock()
    defer cmapLock.Unlock()
    return cmap[req]
}

// Necessary wrapper around all handlers.  Must be the first middleware.
func contextHandler(ctx context.Context, h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctx2 := ctx // make a copy of the root context reference
        cmapLock.Lock()
        cmap[req] = &ctx2
        cmapLock.Unlock()

        h.ServeHTTP(rw, req)

        cmapLock.Lock()
        delete(cmap, req)
        cmapLock.Unlock()
    })
}

func middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctxp := contextFromRequest(req)
        *ctxp = newContextWithRequestID(*ctxp, req)
        h.ServeHTTP(rw, req)
    })
}

func handler(rw http.ResponseWriter, req *http.Request) {
    ctxp := contextFromRequest(req)
    reqID := requestIDFromContext(*ctxp)
    fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
    h := contextHandler(context.Background(), middleware(http.HandlerFunc(handler)))
    http.ListenAndServe(":8080", h)
}
