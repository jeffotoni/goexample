/*
* Example NeWriter
*
* @package     main
* @author      @jeffotoni
* @size        19/02/2019
*
 */

package main

import (
	"bufio"
	"io"
	"os"
)

func WriteJson(out io.Writer) {

	writer := bufio.NewWriter(out)
	writer.WriteString(`{"name":"jeffotoni", "age":"39"}`)
	writer.Flush()
}

func main() {

	myFile, err := os.Create("new.json")
	if err != nil {
		panic(err)
	}
	defer myFile.Close()

	WriteJson(myFile)

}
