package main

import (
    "io"
    "os"
    "os/exec"
)

func main() {
    c1 := exec.Command("ls")
    c2 := exec.Command("wc", "-l")

    pr, pw := io.Pipe()
    c1.Stdout = pw
    c2.Stdin = pr
    c2.Stdout = os.Stdout

    c1.Start()
    c2.Start()

    go func() {
        defer pw.Close()

        c1.Wait()
    }()
    c2.Wait()
}
