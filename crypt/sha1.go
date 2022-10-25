/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"crypto/sha1"
	"fmt"
)

var (
	SHA1_SALT = "3838x7xx2"
)

//
//
//
func GSha1(key string) string {

	data := []byte(key + SHA1_SALT)
	return (fmt.Sprintf("%x", sha1.Sum(data)))
}

func main() {

	password := "1234567890#$"

	hash := GSha1(password)

	fmt.Println(hash)
}
