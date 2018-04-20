package main

import (
	//"log"
	"os/exec"
)

func main() {

	//cmd := exec.Command("sleep", "1")
	cmd := exec.Command("notify-send", "Test", "Hello World")
	//log.Printf("Running command and waiting for it to finish...")
	cmd.Run()
	//log.Printf("Command finished with error: %v", err)
}
