package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"Event":     "DevOpsFest",
		"Lang":      "Go",
		"instagram": "jeffotoni",
		"langs":     struct{ G string }{G: "gophers"},
		"ints":      []int{1, 2, 3, 4}}
	fmt.Println(m)
}
