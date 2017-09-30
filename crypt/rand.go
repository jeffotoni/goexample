/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
//
//
func Random(min, max int) int { rand.Seed(time.Now().Unix()); return rand.Intn(max-min) + min }

func main() {

	fmt.Println(Random(10, 100))

}
