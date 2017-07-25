/*
* Golang config with gcfg
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	good := []byte("Vamos deixar acentuação certinha no código, áéíóú ÁÉÍÓÚ Çç ôê ãõ ")
	newS := ""

	for len(good) > 0 {

		r, size := utf8.DecodeRune(good)

		fmt.Printf("%c %v\n", r, size)

		newS = newS + fmt.Sprintf("%c", r)
		good = good[size:]
	}

	fmt.Println(newS)
}
