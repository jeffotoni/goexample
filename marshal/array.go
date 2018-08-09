package main

import (
	"encoding/json"
	"fmt"
)

type Addresses []Address

func (a Addresses) MarshalJSON() ([]byte, error) {
	aa := make([]string, len(a))
	for i := range a {
		aa[i] = a[i].Address
	}
	return json.Marshal(aa)
}

/*func (a *Addresses) UnmarshalJSON(content []byte) error {
	var raw []string
	if err := json.Unmarshal(content, &raw); err != nil {
		return err
	}
	for i := range raw {
		*a = append(*a, Address{raw[i]})
	}
	return nil
}*/

type Address struct {
	Address string
}

type Response struct {
	Addresses
}

func main() {

	res := Response{
		Addresses{
			Address{"address 0"},
			Address{"address 1"},
			Address{"address 2"},
			Address{"address 3"},
		},
	}

	j, _ := json.Marshal(res)
	fmt.Println(string(j))

	/* unmarshals
	var res2 Response
	if err := json.Unmarshal(j, &res2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2)
	*/
}
