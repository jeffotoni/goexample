package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ls", "-lh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("context cancel")
		os.Exit(0)
	}
	fmt.Printf("in all caps: %v\n", out.String())

	//if err := exec.CommandContext(ctx, "sleep", "1").Run(); err != nil {
	//if err := exec.CommandContext(ctx, "ls", "-lh").Run(); err != nil {
	// This will fail after 100 milliseconds. The 5 second sleep
	// will be interrupted.
	//fmt.Println("context cancel")
	// } else {
	// 	fmt.Println("context success:")
	// }
}
