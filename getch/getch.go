package main

import (
	"bytes"
	"fmt"

	"github.com/pkg/term"
)

func getch() []byte {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)

	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)
	t.Restore()
	t.Close()
	if err != nil {
		return nil
	}

	return bytes[0:numRead]
}

func main() {
	for {
		c := getch()
		switch {
		case bytes.Equal(c, []byte{3}):
			return
		case bytes.Equal(c, []byte{27, 91, 68}): // left
			fmt.Println("LEFT pressed")
		default:
			fmt.Println("Unknown pressed", c, string(c))
		}
	}
	return
}
