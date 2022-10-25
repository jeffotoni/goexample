package main

import "fmt"
import "encoding/json"

type PublicKey struct {
	Id  int
	Key string
}

type KeysResponse struct {
	Collection []PublicKey
}

func main() {

	keysBody := []byte(`[{"id": 9090,"key": "383"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)

	keys := make([]PublicKey, 0)
	json.Unmarshal(keysBody, &keys)

	fmt.Printf("%#v\n", keys)

	fmt.Println("value id: ", keys[0].Id)
	fmt.Println("value id: ", keys[0].Key)
}
