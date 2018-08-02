package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/netcatc/codigos-correio.csv")
	i := 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator

		if scanner.Text() != "" {
			i++
		}
	}

	fmt.Println("quantidade de linhas: ", i)
}
