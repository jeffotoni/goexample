/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	b64 "encoding/base64"
	"fmt"
)

//
func UkkBase64Encode(textString string) string {

	text := []byte(textString)

	sEnc := b64.StdEncoding.EncodeToString(text)

	return sEnc
}

//
func UkkBase64Decode(textString string) string {

	//text := []byte(textString)

	sDec, _ := b64.StdEncoding.DecodeString(textString)

	return string(sDec)
}

func main() {

	encbase64 := UkkBase64Encode("lets text now here...")
	decbase64 := UkkBase64Decode(encbase64)

	fmt.Println(encbase64)
	fmt.Println(decbase64)

}
