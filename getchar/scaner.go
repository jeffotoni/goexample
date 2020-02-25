package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("text:")
		fmt.Println(scanner.Text())
	}
}
