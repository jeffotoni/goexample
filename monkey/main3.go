package main

func a() int { return 1 }

func main() {
	f := a
	f()
}
