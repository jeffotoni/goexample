/*
* Example make(chan)
*
* @package     main
* @author      @jeffotoni
* @size        19/02/2018
*
 */

package main

import "fmt"

func main() {

	rate := 42

	// this condition can never be true
	if rate > 60 && rate < 40 {
		fmt.Println("rate %:", rate)
	}
}
