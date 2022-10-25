// Go Api server
// @jeffotoni
// 2019-03-10

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func print(w io.Writer) {
	fmt.Fprintln(w, "output")
}

func main() {
	fmt.Println("print with byes.Buffer:")
	var b bytes.Buffer
	print(&b)
	fmt.Print(b.String())

	fmt.Println("print with os.Stdout:")
	print(os.Stdout)
}
