// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

type S string

var (
	String = "@jeffotoni"
)

func main() {
	var text string
	var str S

	mypicture := "@Photograph-jeffotoni"
	str = "@workshop-devOpsBh"
	text = "@jeffotoni-golang"

	fmt.Println(str)
	fmt.Println(String)
	fmt.Println(text)
	fmt.Println(mypicture)

	// example string
	s := "日本語"
	fmt.Printf("Glyph:             %q\n", s)
	fmt.Printf("UTF-8:             [% x]\n", []byte(s))
	fmt.Printf("Unicode codepoint: %U\n", []rune(s))
}
