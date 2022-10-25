package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Page struct {
	Title    string
	Filename string
	Content  string
}

type Pages []Page

var pages = Pages{

	Page{
		"First Page",
		"page1.txt",
		"This is the 1st Page.",
	},

	Page{
		"Second Page",
		"page2.txt",
		"The 2nd Page is this.",
	},
}

func main() {
	json, err := json.Marshal(pages)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	fmt.Println("  ")
	fmt.Fprintf(os.Stdout, "%s", json)
	fmt.Println("\n")
}
