package mw

import (
	"net/http"
)

type AuthenticationMiddleware struct {
	TokenUsers map[string]string
}

func (amw *AuthenticationMiddleware) Populate() {
	amw.TokenUsers["your-key-here"] = "user-here"
}

func (amw *AuthenticationMiddleware) MiddlewareToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Key-Token")
		if _, found := amw.TokenUsers[token]; found {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
