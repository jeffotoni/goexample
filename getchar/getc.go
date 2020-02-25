package main

import (
	"fmt"
	"os"
)

type pseudoStdin struct {
	buf []byte
	i   int
}

var stdin *pseudoStdin

func newStdin() *pseudoStdin {
	s := &pseudoStdin{}
	s.buf = make([]byte, 1024*1024)
	os.Stdin.Read(s.buf)
	return s
}
func getc(stdin *pseudoStdin) (byte, error) {
	b := stdin.buf[stdin.i]
	if b == byte(0) {
		return b, fmt.Errorf("EOL")
	}
	stdin.i++
	return b, nil
}
func ungetc(c byte, stdin *pseudoStdin) {
	stdin.i--
	return
}

func main() {
	stdin = newStdin()
	c, err := getc(stdin)
	if err != nil {
		fmt.Printf("EOF")
		return
	}
	fmt.Printf("first char  = %c\n", c)
	c, err = getc(stdin)
	if err != nil {
		fmt.Printf("EOF")
		return
	}
	fmt.Printf("second char = %c\n", c)
	// go back 1 char
	ungetc(c, stdin)
	c, err = getc(stdin)
	if err != nil {
		fmt.Printf("EOF")
		return
	}
	fmt.Printf("second char = %c\n", c)
	c, err = getc(stdin)
	if err != nil {
		fmt.Printf("EOF")
		return
	}
}
