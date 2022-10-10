package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

// This program has the objective show people how easy is using linux commands in golang
func runLs() error {

	ls := exec.Command("ls")

	var outb bytes.Buffer
	var outErr bytes.Buffer

	ls.Stdout = &outb
	ls.Stderr = &outErr

	if err := ls.Run(); err != nil {
		return err
	}

	if outErr.String() != "" {
		return errors.New(outErr.String())
	}

	log.Printf("\nls output:\n%s\n", outb.String())
	return nil
}
