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

func readFile(file io.Reader) error {
	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Print(string(by))
		if err == io.EOF {
			time.Sleep(time.Second)
		}
	}
}

func main() {
	tf := flag.String("file", "", "example: file.txt")
	flag.Parse()
	if len(os.Args) < 3 {
		flag.PrintDefaults()
		return
	}

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

func existsFile(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
