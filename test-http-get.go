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
	"net/http"
	// "os"
	"time"
)

func main() {

	url := "http://localhost:12345/"

	for {

		time.Sleep(100 * time.Millisecond)

		response, err := http.Get(url)

		if err != nil {

			fmt.Println(err)

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
