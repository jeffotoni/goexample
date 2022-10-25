/*
* Example tollbooth
*
* @package     main
* @author      @jeffotoni
* @size        16/07/2017
*
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("/path")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString(10) // line defined once separador
	for err != io.EOF {
		fmt.Print(line)              // or any stuff
		line, err = r.ReadString(10) //  line was defined before
	}
}
