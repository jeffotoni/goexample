package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sourcegraph/apiproxy"
)

func main() {
	//fmt.Println("vim-go")

	proxy := apiproxy.NewCachingSingleHostReverseProxy("http://localhost:8080/api/v1/ping", httpcache.NewMemoryCache())
	cachingTransport := proxy.Transport.(*httpcache.Transport)
	cachingTransport.Transport = &apiproxy.RevalidationTransport{
		Check: apiproxy.ValidatorFunc(func(url *url.URL, age time.Duration) bool {
			// only revalidate expired cache entries older than 30 minutes
			return age > 30*time.Minute
		}),
	}
	http.Handle("/api/v1/ping", handlers.CombinedLoggingHandler(os.Stdout, proxy))
	http.ListenAndServe(":8080", nil)

}
