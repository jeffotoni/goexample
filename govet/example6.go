/*
* Example make(chan)
*
* @package     main
* @author      @jeffotoni
* @size        19/02/2018
*
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func f() {}

func main() {
	// Self assignment
	i := 42
	i = i

	// a declared function cannot be nil
	fmt.Println(f == nil)

	// shift too long
	fmt.Println(i >> 32)

	// res used before checking err
	res, err := http.Get("https://www.spreadsheetdb.io/")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
