package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sourcegraph/apiproxy"
	"github.com/sourcegraph/httpcache"
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

// func NewProxy(spec *serviceSpec.ServiceSpec, director func(*http.Request),
// respDirector func(*http.Response) error, dialTimeout, dialKAlive, transTLSHTimeout,
// transRHTimeout time.Duration) *MultiReverseProxy {
// 	return &MultiReverseProxy{
// 		proxy: &httputil.ReverseProxy{
// 			Director:       director, //Request director function
// 			ModifyResponse: respDirector,
// 			Transport: &http.Transport{
// 				Dial: (&net.Dialer{
// 					Timeout:   dialTimeout, //limits the time spent establishing a TCP connection (if a new one is needed).
// 					KeepAlive: dialKAlive,  //limits idle keep a live connection.
// 				}).Dial,
// 				TLSHandshakeTimeout:   transTLSHTimeout, //limits the time spent performing the TLS handshake.
// 				ResponseHeaderTimeout: transRHTimeout,   //limits the time spent reading the headers of the response.
// 			},
// 		},
