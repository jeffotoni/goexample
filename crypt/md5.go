/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//
//
//
func Md5(text string) string {

	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}

func main() {

	password := "1234567890#$"

	hash := Md5(password)

	fmt.Println(hash)
}
