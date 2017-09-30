/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"crypto/sha256"
	"fmt"
)

var (
	HASH_SALT = "salt2888&"
)

//
//
//
func GHash256(password string) string {

	pass := password + HASH_SALT

	h := sha256.New()
	h.Write([]byte(pass))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	password := "1234567890#$"

	hash := GHash256(password)

	fmt.Println(hash)
}
