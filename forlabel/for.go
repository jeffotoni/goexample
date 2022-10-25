package main

import "fmt"

func main() {

	//OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			fmt.Printf("i=%v, j=%v\n", i, j)

			//if j == 3 {
			break
			//}
		}
	}
	fmt.Println("vim-go")
}
