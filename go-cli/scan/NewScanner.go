// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"bufio"
	"os"
)

func main() {
	println("Type something, to close enter [exit]")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "exit" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
