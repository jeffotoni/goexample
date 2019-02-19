package main

import (
	"encoding/json"
	"io"
	"os"
)

type Person struct {
	Name string
	Age  uint
}

// writes json representation of Person to Writer
func (p *Person) WriteJson(w io.Writer) error {
	b, err := json.Marshal(*p)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return err
}

func main() {
	p := Person{
		"Sebastian",
		22,
	}

	myFile, err := os.Create("output2.json")
	if err != nil {
		panic(err)
	}
	defer myFile.Close()

	p.WriteJson(myFile)
}
