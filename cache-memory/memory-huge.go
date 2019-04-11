//
// simulation large memory
//

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Background GC work generated")

	// Force a GC to set a baseline we can see if we set GODEBUG=gctrace=1
	runtime.GC()

	// A huge allocation to give the GC work to do
	//lotsOf := make([]*int, 15e8)
	lotsOf := make([]int, 15e8) // 1.3G

	// Use up all the CPU doing work that causes allocations that could be cleaned up by the GC.
	numWorkers := runtime.NumCPU()
	fmt.Println(numWorkers)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()

	// Make sure that this memory isn't optimised away
	runtime.KeepAlive(lotsOf)
}

func work() {
	var i int
	for {

		i++
		//work := make([]*int, 1e5) // 1.3G
		//work := make([]int, 1e5) // 1.3G
		if f := factorial(20); f != 2432902008176640000 {
			fmt.Println(f)
			os.Exit(1)
		}
		// work = append(work, i)

		// garante que o objeto não seja liberado e seu
		// finalizador não seja executado antes do ponto
		// no programa em que o KeepAlive é chamado.
		// runtime.KeepAlive(work)
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
