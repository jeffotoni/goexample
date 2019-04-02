package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	if err := exec.CommandContext(ctx, "sh", "-c", "google-chrome").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		log.Println("exec google-chrome: ", err)
		return
	}
}
