// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	// for j := 0; j < len(output); j++ {
	// 	fmt.Printf("%c", output[j])
	// }
}
