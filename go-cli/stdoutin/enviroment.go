package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("./testenv")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	cmd.Env = append(os.Environ(),
		"AWS_REGION=us-test",    // ignored
		"AWS_TEST=xtestawsjeff", // this value is used
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stdout:%v\n", stdout.String())
}
