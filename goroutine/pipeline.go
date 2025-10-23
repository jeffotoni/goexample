package main
// Stage 1: Gera números
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Multiplica por 2
func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// Stage 3: Soma 1
func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

// Encadeando o pipeline
in := gen(1, 2, 3, 4)
stage1 := multiply(in)
stage2 := addOne(stage1)

for result := range stage2 {
fmt.Println(result)
}