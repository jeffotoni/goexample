package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err, out)
	}
	fmt.Printf("The date is %s\n", out)
}
