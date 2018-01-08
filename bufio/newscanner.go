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
	"os"
)

func readLine(path string) {

	inFile, _ := os.Open(path)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}
}

func main() {

	// read line by line
	readLine("/path")

}
