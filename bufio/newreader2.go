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

	f, err := os.Open("/home/netcatc/Documentos/postgresql.log.2017-11-30-14")

	if err != nil {

		fmt.Println("error opening file ", err)
		os.Exit(1)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for {

		path, err := r.ReadString(10) // 0x0A separator = newline

		// fmt.Println(path)

		if err == io.EOF {

			fmt.Println("FIM")

			// do something here
			break

		} else if err != nil {

			fmt.Println(err)

			// return err // if you return error
		}
	}
}
