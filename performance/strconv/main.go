// main.go
package main

import (
	"fmt"
	"strconv"
)

func strconvFmt(a string, b int) string {
	return a + ":" + strconv.Itoa(b)
}

func fmtFmt(a string, b int) string {
	return fmt.Sprintf("%s:%d", a, b)
}

func main() {}
