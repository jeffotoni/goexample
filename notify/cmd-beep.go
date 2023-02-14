package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("beep", -1)

	log.Println("log: ", cmd.Run())
}
