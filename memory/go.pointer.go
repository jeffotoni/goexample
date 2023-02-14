package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10
	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(&count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

//go:noinline
func increment(inc *int) {

	// Increment the "value of" inc.
	*inc++
	//println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")

}
