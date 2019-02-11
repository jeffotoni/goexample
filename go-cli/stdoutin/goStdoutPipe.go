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

func CmdExec(namef string, outfc chan []byte) {

	var stdout, stderr []byte
	var errStdout, errStderr error

	pdf := namef
	cmd := exec.Command("sh", "docker-pdf-totext.sh")
	newEnv := append(os.Environ(), "PATH_PDF="+pdf)
	cmd.Env = newEnv
	//out, err := cmd.CombinedOutput()

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
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
	//outStr, errStr := string(stdout), string(stderr)
	//fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	outfc <- stdout

	//return stdout
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

	cf := make(chan []byte, 1)
	CmdExec("/tmp/golang.pdf", cf)
	fmt.Println(string(<-cf))
}
