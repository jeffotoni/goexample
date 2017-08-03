/*
* Example sync.Mutex
*
* @package     main
* @author      @jeffotoni
* @size        01/08/2017
*
 */

package main

import "fmt"
import "time"
import "math/rand"

func f(n int) {
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {

		// time.Sleep(time.Millisecond * 100)
		fmt.Println(n, ":", i)

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func f2(n int) {

	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {

	go f(0)

	// go f2(1)

	for i := 0; i < 10; i++ {

		go f(i)
	}

	var input string

	fmt.Scanln(&input)
}
