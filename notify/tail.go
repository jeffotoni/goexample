/*
* Golang in action
*
* @package     main
* @author      @jeffotoni
* @size        2019
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var tf = flag.String("file", "", "example: file.txt")
var r *bufio.Reader

func init() {
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.PrintDefaults()
		return
	}
}

func WatchFile() {
	if existsFile(*tf) {
		file, err := os.Open(*tf) // For read access.
		if err != nil {
			log.Fatal(err)
		}
		readFile(file)
	} else {
		fmt.Println("File not exist!")
		flag.PrintDefaults()
	}
}

func tryToOpen() {
	if existsFile(*tf) {
		file, err := os.Open(*tf) // For read access.
		if err != nil {
			log.Fatal(err)
		}
		// r global
		r = bufio.NewReader(file)
	}
}

func readFile(file io.Reader) error {

	// r global
	r = bufio.NewReader(file)

	go func() error {
		for {
			by, err := r.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return err
			}
			fmt.Print(string(by))
			if err == io.EOF {
				time.Sleep(time.Second * 2)
			}
		}
	}()

	// read file
	go func() {
		for {
			time.Sleep(time.Second * 1)
			//println("test file..")
			existgof()
		}
	}()

	for {
	}
	return nil
}

func existgof() {
	if !existsFile(*tf) {
		// not exist
		// try in 10s recursive
		time.Sleep(time.Second * 10)
		tryToOpen()
	}
}

func main() {
	WatchFile()
}

func existsFile(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
