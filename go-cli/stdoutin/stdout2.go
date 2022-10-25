// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("only -> out:\n%s\nerr:\n%s\n", outStr, errStr)
}
