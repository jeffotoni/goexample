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
	"io/ioutil"
	// "io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {

	vetUrl := make(map[int]string)

	vetUrl[0] = "http://localhost:12345/"
	vetUrl[1] = "http://localhost:12345"
	//vetUrl[2] = "http://localhost:12345/login"
	//vetUrl[3] = "http://localhost:12345/login/"

	min := 0
	max := 2

	for {

		seed := Seed(min, max)

		time.Sleep(4 * time.Millisecond)

		// fmt.Println("")
		// fmt.Println("::::: http get:::::")
		// fmt.Println(vetUrl[seed])
		response, err := http.Get(vetUrl[seed])

		if err != nil {

			fmt.Println(err)
			os.Exit(1)

		} else {

			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)

			//_, err := io.Copy(os.Stdout, response.Body)

			if err != nil {

				fmt.Println(err)
			}

			fmt.Println(string(contents))
		}
	}
}

func Seed(min int, max int) int {

	return (min + rand.Intn(max-min))
}
