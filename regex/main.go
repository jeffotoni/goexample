package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(formatPhoneNumber("62983386760"))
}

func formatPhoneNumber(number string) string {
	if len(number) < 11 {
		return "Invalid phone number"
	}

	pattern := `(\d{2})?(\d{5})?(\d{4})`

	repl := `(${1})${2}-${3}`

	phoneRegex := regexp.MustCompile(pattern)

	valueStr := phoneRegex.ReplaceAll([]byte(number), []byte(repl))

	return string(valueStr)
}
