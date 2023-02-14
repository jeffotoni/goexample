package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

var (
    output_path = filepath.Join("./output")
    bash_script = filepath.Join("_script.sh")
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}
func exe_cmd(cmds []string) {
    os.RemoveAll(output_path)
    err := os.MkdirAll(output_path, os.ModePerm|os.ModeDir)
    checkError(err)
    file, err := os.Create(filepath.Join(output_path, bash_script))
    checkError(err)
    defer file.Close()
    file.WriteString("#!/bin/sh\n")
    file.WriteString(strings.Join(cmds, "\n"))
    err = os.Chdir(output_path)
    checkError(err)
    out, err := exec.Command("sh", bash_script).Output()
    checkError(err)
    fmt.Println(string(out))
}

func main() {
    commands := []string{
        "echo newline >> foo.o",
        "echo newline >> f1.o",
        "echo newline >> f2.o",
    }
    exe_cmd(commands)
}
