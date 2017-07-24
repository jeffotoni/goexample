/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go run -race race.go
//

func main() {

	start1() // with error
	//start2()
}

func start1() {

	start := time.Now()
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {

		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})

	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {

	return time.Duration(rand.Int63n(1e9))
}

func start2() {

	start := time.Now()
	reset := make(chan bool)

	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {

		fmt.Println(time.Now().Sub(start))
		reset <- true
	})

	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}
