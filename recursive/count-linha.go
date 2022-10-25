package main

import (
	"fmt"
	"io"
	"os"
	//"strings"
)

func lineCount(r io.Reader) (n int, err error) {

	buf := make([]byte, 8192)

	i := 0

	for {

		c, err := r.Read(buf)

		if err != nil {

			if err == io.EOF && c == 0 {
				break
			} else {
				return n, err
			}
		}

		for _, b := range buf[:c] {

			if i != 1 && b == '\n' {

				n++
				i = 0
			}

			i++
		}
	}

	if err == io.EOF {
		err = nil
	}

	return n, err
}

func main() {

	file, _ := os.Open("/home/netcatc/codigos-correio.csv")

	n, _ := lineCount(file)

	fmt.Println("numero de linhas: ", n)
}
