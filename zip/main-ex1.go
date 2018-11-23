package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"sync"
)

func ZipWriter(files chan *os.File) *sync.WaitGroup {
	f, err := os.Create("/tmp/out.zip")
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	zw := zip.NewWriter(f)
	go func() {
		// Note the order (LIFO):
		defer wg.Done() // 2. signal that we're done
		defer f.Close() // 1. close the file
		var err error
		var fw io.Writer
		for f := range files {
			// Loop until channel is closed.
			if fw, err = zw.Create(f.Name()); err != nil {
				panic(err)
			}
			io.Copy(fw, f)
			if err = f.Close(); err != nil {
				panic(err)
			}
		}
		// The zip writer must be closed *before* f.Close() is called!
		if err = zw.Close(); err != nil {
			panic(err)
		}
	}()
	return &wg
}

func main() {

	files := make(chan *os.File)
	wait := ZipWriter(files)

	// Send all files to the zip writer.
	var wg sync.WaitGroup
	wg.Add(len(os.Args) - 1)
	for i, name := range os.Args {
		if i == 0 {
			continue
		}

		log.Println(name)
		// Read each file in parallel:
		go func(name string) {
			defer wg.Done()
			f, err := os.Open(name)
			if err != nil {
				panic(err)
			}
			files <- f
		}(name)
	}

	wg.Wait()
	// Once we're done sending the files, we can close the channel.
	close(files)
	// This will cause ZipWriter to break out of the loop, close the file,
	// and unblock the next mutex:
	wait.Wait()
}
