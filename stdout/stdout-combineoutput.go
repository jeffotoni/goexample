package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	// echo stdout; echo 1>&2 stderr
	cmd := exec.Command("sh", "-c", "google-chrome")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
