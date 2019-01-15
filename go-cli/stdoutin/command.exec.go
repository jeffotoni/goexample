package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-lha")
	//cmd.Stdin = strings.NewReader("echo 'ola'")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %v\n", out.String())
}
