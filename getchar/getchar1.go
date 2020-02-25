package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	fmt.Println("Para sair digite q")
	var total int
	var b []byte = make([]byte, 1)
	for {
		fmt.Printf("Digite um numero:")
		os.Stdin.Read(b)
		//fmt.Println("I got the byte", b, "("+string(b)+")")
		num := string(b)
		if num == "q" {
			break
		}

		n, err := strconv.Atoi(num)
		if err == nil {
			total += n
		}
		fmt.Printf("%d", n)
		fmt.Println("")
	}

	fmt.Println("Soma=", total)
}
