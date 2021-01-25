package main

func main() {
	var s = []int{1, 2, 46, 6}
	doesNotEscape(s)
}

func escapes(input []int) {
	output := make([]int, len(input))
	for _, n := range input {
		output = append(output, n)
	}
	return
}

func doesNotEscape(input []int) {

	output := []int{}
	for _, n := range input {
		output = append(output, n)
	}
	return
}
