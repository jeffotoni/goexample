// Front-end in Go server
// @jeffotoni
// 2019-01-07

package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {

	for n := 0; n < 10; n++ {
		uuid := uuid.New().String()
		uuid = strings.ToUpper(uuid)
		fmt.Printf("%v", uuid)
	}
}
