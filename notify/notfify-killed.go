package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func main() {

	fmt.Println("kill -9 start")

	cmd := exec.Command("/bin/sh", "-c", "watch date > date.txt")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start := time.Now()
	time.AfterFunc(3*time.Second, func() {
		sscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)

}
