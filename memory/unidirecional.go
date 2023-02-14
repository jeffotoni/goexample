package main

// unidirecional
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// unidirecional
	sendch := make(chan<- int)
	go sendData(sendch)

	//fmt.Println(<-sendch)
}
