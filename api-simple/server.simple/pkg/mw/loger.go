package mw

import (
    "log"
    "net/http"
    "time"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf(
            "\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m\033[0m",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    })
}
