package main

import "fmt"
import "encoding/json"

type Restaurant struct {
	NumberOfCustomers bool `json:"numberofcustomers,omitempty"`
}

func main() {
	d := Restaurant{
		NumberOfCustomers: false,
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))

	var f bool
	fmt.Println(f)
}
