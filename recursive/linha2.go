package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/netcatc/codigos-correio.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // define your buffer size here.

	for {

		n, err := file.Read(buf)

		if n > 0 {

			if string(buf[:n]) == "" {

				fmt.Print("linhe") // your read buffer.
			}

			fmt.Println("value: ", string(buf[:n])) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}

}
