/*
* Example make(chan)
*
* @package     main
* @author      @jeffotoni
* @size        19/02/2018
*
 */

package main

import "fmt"

func main() {
	lang := []string{"Golang", "PHP", "Python"}

	for _, word := range lang {
		go func() {
			fmt.Println(word)
		}()
	}
}
