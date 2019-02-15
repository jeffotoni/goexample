// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	var proxyPort, port string
	flag.StringVar(&proxyPort, "proxy-port", "5010", "port to proxy of your app")
	flag.StringVar(&port, "port", "1333", "port to listen to on proxy.")
	flag.Parse()

	url, err := url.Parse("http://localhost:" + proxyPort)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
	reverseProxy := httputil.NewSingleHostReverseProxy(url)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	fmt.Printf("Listening on port: %s and forwarding requests to: %s \n", port, proxyPort)
	panic(http.ListenAndServe(":"+port, nil))
}
