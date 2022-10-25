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
	"net/rpc"
	//"os"
)

var startClient = 0

type Args2 struct {
	A string
}

// type stop
type Stop string

func main() {

	// o restart the rpc client
	//os.Setenv("START", "0")

	// Synchronous call
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	args2 := Args2{"Clean Env Server, restart again client..."} // Defining the arguments
	var stop string

	fmt.Println("Allowing you to call the client again")

	client.Call("Stop.CleanEnv", args2, &stop)

}
