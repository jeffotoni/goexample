package main

func return80b() [10]int {
	var s [10]int
	return s
}

func return80bPointer() *[10]int {
	var s [10]int
	return &s
}

func return8kb() [1024]int {
	var s [1024]int
	return s
}

func return8kbPointer() *[1024]int {
	var s [1024]int
	return &s
}

func return8mb() [1024 * 1024]int {
	var s [1024 * 1024]int
	return s
}

func return8mbPointer() *[1024 * 1024]int {
	var s [1024 * 1024]int
	return &s
}

func main(){

}