// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func CmdExec(Ip string, outfc chan []byte) {

	var stdout, stderr []byte
	var errStdout, errStderr error

	//pdf := namef
	cmd := exec.Command("ping", Ip, "-c 5", "-i 3", "-w 10")

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr: ", stderr, "\n")
	}
	outfc <- stdout
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Ex: ping 127.0.0.1")
		return
	}

	Ip := os.Args[1]
	cf := make(chan []byte, 1)
	CmdExec(Ip, cf)
	fmt.Println(string(<-cf))
}
