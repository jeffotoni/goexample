package main

import "runtime"
import "time"

func main() {
	f := make([]int, 21024*8)
	runtime.GC()
	go func() {
		for i := 0; i < 100000; i++ {
			f[i] = i * 2
			time.Sleep(time.Second)
		}
	}()
	//f := make([]int, 1024*8)

	runtime.SetFinalizer(&f, func(f *[]int) {
		println("finalizado")
	})
	f = nil

	println("done")
}
