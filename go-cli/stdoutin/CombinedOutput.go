package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "ls", "-lh", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, stdoutStderr)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
