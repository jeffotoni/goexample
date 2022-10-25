// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-lhs")
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run: %v", err)
	}
	//log.Printf("Waiting for command to finish...")
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Printf("Command finished with error: %v", err)
	// }
}
