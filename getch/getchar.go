package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// desabilitar buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	// na apresentar no display
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	var b []byte = make([]byte, 1)

	for {
		os.Stdin.Read(b)
		fmt.Println("opa, o usuário tinha apertado o botão: ", b, "("+string(b)+")")
	}
}
