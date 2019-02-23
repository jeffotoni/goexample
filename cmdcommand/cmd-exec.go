package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	command := "echo -n ################ DEPLOY BUILD LOCAL ###############"
	cv := strings.Fields(command)
	arg1 := cv[0]
	parts := cv[1:]

	cmd := exec.Command(arg1, parts...)
	out, _ := cmd.CombinedOutput()

	fmt.Println(string(out))
}
