package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received person: Name = %s, Age = %d\n", person.Name, person.Age)
}

func main() {
	println("Run Server :8080")
	http.HandleFunc("/v1/user", handlePost)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

