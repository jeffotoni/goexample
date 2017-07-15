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

	w.Write([]byte("Hello, Jefferson !"))
}

func main() {

	port := "12345"
	// Create a request limiter per handler.
	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(2, time.Second), Hello))

	fmt.Println("Start port:", port)

	http.ListenAndServe(":"+port, nil)

}
