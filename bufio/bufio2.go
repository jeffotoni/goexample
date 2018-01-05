package main

import (
		"bufio"
			"fmt"
				"os"
			)

			func main() {
					
				f, _ := os.Create("/tmp/example.log")
				defer f.Close()
				w := bufio.NewWriter(f)
				//w := bufio.NewWriter(os.Stdout)
				fmt.Fprint(w, "Hello, ")
				fmt.Fprint(w, "world!")
				w.Flush() // Don't forget to flush!
			}
