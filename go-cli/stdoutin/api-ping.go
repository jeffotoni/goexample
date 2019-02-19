// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ex: ping 127.0.0.1")
		return
	}
	// ip
	Ip := os.Args[1]
	fmt.Println("ip: ", Ip)
	out, _ := exec.Command("ping", Ip, "-c 5", "-i 3", "-w 10").Output()
	fmt.Println(string(out))
	if strings.Contains(string(out), "Destination Host Unreachable") {
		fmt.Println("down")
	} else {
		fmt.Println("alive")
	}
}
