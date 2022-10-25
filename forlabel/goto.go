package main

import (
	"fmt"
)

func main() {
	i := 0
Start:
	fmt.Println(i)
	if i > 2 {
		goto End
	} else {
		i += 1
		goto Start
	}
End:
}
