/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	//crand "crypto/rand"
	"fmt"
	"math/rand"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func HashSeed(n int) string {

	b := make([]rune, n)
	for i := range b {

		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {

	hash := HashSeed(12)

	fmt.Println(hash)
}
