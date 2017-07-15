/*
* Example tollbooth
*
* @package     main
* @author      @jeffotoni
* @size        15/07/2017
*
 */

package main

import (
	"fmt"
	"github.com/didip/tollbooth"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Hello!!"))
}

func Login(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte(`{"msg":"ok"}`))
}

func main() {

	port := "12345"

	limiter := tollbooth.NewLimiter(1, time.Second)

	limiter = tollbooth.NewLimiterExpiringBuckets(1, time.Second, time.Hour, 0)

	limiter.IPLookups = []string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}

	limiter.Methods = []string{"GET", "POST"}

	// limiter.Headers = make(map[string][]string)
	// limiter.Headers["X-Access-Token"] = []string{"abc123", "xyz098"}

	// Create a request limiter per handler.
	//http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(900, time.Millisecond), Hello))

	// Create a request limiter per handler.
	/// tollbooth.LimitFuncHandler(tollbooth.NewLimiter(2, time.Second), Login)
	http.Handle("/login", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(2, time.Second), Login))

	fmt.Println("Start port:", port)
	fmt.Println("Endpoints:")
	fmt.Println("http://localhost:" + port + "/")
	fmt.Println("http://localhost:" + port + "/login")
	fmt.Println("Max bytes:", 1<<20, "bytes")
	//fmt.Println(limiter)

	s := &http.Server{

		Addr: ":" + port,
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
