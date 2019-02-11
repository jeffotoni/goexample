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
	"flag"
	"fmt"
	"os"
	"strings"
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

	enctext := flag.String("text", "", "example:\n    $ base64 -text mytext normal here")
	flag.Parse()

	if len(os.Args) < 3 {
		flag.PrintDefaults()
		return
	}

	textnew := ""
	//fmt.Println(*enctext)
	// list text
	if len(os.Args) > 3 {
		list := os.Args
		textnew = strings.Join(list[2:], " ")
		fmt.Println(textnew)
	} else {
		textnew = *enctext
	}

	encbase64 := UkkBase64Encode(textnew)

	//decbase64 := UkkBase64Decode(encbase64)

	fmt.Println(encbase64)
	//fmt.Println(decbase64)

}
