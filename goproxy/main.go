package main

import (
	"fmt"
	//"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyTarget := "http://127.0.0.1:8080"
	proxySource := "localhost:8888"
	proxySourceIP := "34.70.223.100"

	proxy := NewReverseProxy(proxyTarget)

	server := &http.Server{
		Addr:    proxySource,
		Handler: ProxyMiddleware(proxy, proxySourceIP),
	}

	fmt.Printf("Servidor proxy escutando em %s...\n", proxySource)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func NewReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func ProxyMiddleware(next http.Handler, sourceIP string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(sourceIP)
		r.Header.Set("X-Real-IP", sourceIP)
		next.ServeHTTP(w, r)
	})
}

