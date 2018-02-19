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
	var i int

	// always true
	fmt.Println(i != 0 || i != 1)

	// always false
	fmt.Println(i == 0 && i == 1)

	// redundant check
	fmt.Println(i == 0 && i == 0)
}
