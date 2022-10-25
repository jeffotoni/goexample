package main

import (
	"bufio"
	"embed"
	"log"
	"time"
)

//go:embed funcionario.json
var f embed.FS

func main() {
	start := time.Now()

	// exemplo usando ReadFile
	// data, _ := f.ReadFile("funcionario.json")
	// print(string(data))

	file, err := f.Open("funcionario.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}

	end := time.Now()
	println("time:", end.Sub(start).String())
}
