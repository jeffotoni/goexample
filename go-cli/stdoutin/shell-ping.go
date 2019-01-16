// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmdName := "ping 127.0.0.1"
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte, 100)
	num := 1
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
		num = num + 1
		if num > 3 {
			//os.Exit(0)
			fmt.Println("[ok]")
		}
	}

	cmd.Wait()
}
