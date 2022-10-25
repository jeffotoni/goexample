package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	commands := []string{"echo -n ## deploy ##", "sleep 1", "pwd", "ls -lh"}

	for _, command := range commands {

		cv := strings.Fields(command)
		arg1 := cv[0]
		parts := cv[1:]

		cmd := exec.Command(arg1, parts...)
		cmdReader, _ := cmd.StdoutPipe()

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()
		cmd.Start()
		cmd.Wait()
	}

}
