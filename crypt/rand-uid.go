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
	//"math/rand"
	crand "crypto/rand"
	"time"
)

//
//
//
func RandUid() string {

	// generate 64 bits timestamp
	unix64bits := uint64(time.Now().UTC().UnixNano())

	buff := make([]byte, 128)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		fmt.Println(err)
	}

	unixUid := fmt.Sprintf("%x", unix64bits)

	//unixUid = GSha1(Blowfish(unixUid))
	return unixUid
}

func main() {

	fmt.Println(RandUid)

}
