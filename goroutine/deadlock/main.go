package main

func main() {
	messages := make(chan string)

	// A go routine that sends to a channel
	go func() {
		//messages <- "I am sending a message."
	}()

	// Main routine has no receiver.
	// No compilation or runtime error.
	// Why doesn't it deadlock?
	//time.Sleep(time.Second)
	<- messages
}