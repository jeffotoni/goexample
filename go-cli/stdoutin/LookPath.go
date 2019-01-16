// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("stdout")
	if err != nil {
		log.Fatal("installing stdout is in your future")
	}
	fmt.Printf("stdout is available at %s\n", path)
}
