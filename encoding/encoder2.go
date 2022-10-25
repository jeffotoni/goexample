package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	v := make(map[string]string)
	v["key"] = "value with <> symbols"
	if err := enc.Encode(&v); err != nil {
		log.Println(err)
	}

	fmt.Printf("json codec: %v", buf.String())
}
