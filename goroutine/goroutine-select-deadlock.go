package main

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	}
}
