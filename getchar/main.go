package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	int, _ := strconv.Atoi(in.Text())
	fmt.Println(int)
}
