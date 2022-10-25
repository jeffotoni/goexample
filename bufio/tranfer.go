package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// PassThru wraps an existing io.Reader.
//
// It simply forwards the Read() call, while displaying
// the results from individual calls to it.
type PassThru struct {
	io.Reader
	total int64 // Total # of bytes transferred
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.total += int64(n)

	if err == nil {
		fmt.Println("Read", n, "bytes for a total of", pt.total)
		time.Sleep(time.Second * 2)
	}

	return n, err
}

func main() {
	var src io.Reader    // Source file/url/etc
	var dst bytes.Buffer // Destination file/buffer/etc

	// Create some random input data.
	src = bytes.NewBufferString(strings.Repeat("Some random input data", 1000))

	// Wrap it with our custom io.Reader.
	src = &PassThru{Reader: src}

	count, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Transferred", count, "bytes")
}
