package main

import "fmt"
import "time"

func main() {
	fmt.Println("vim-go")

	<-time.After(time.Duration(5 * time.Second))

	time.Sleep(time.Second * 2)
}
