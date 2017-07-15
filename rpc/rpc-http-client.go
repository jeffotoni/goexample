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
)

type Args struct {
	A, B int
}

type Args2 struct {
	A string
}

func main() {

	// type Args2 struct { A string }
	// type Args struct { A, B int }

	// Synchronous call
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	args := Args{335, 32} // Defining the arguments
	var reply int

	// Calling my method Muilt
	client.Call("Matt.Multiply", args, &reply)
	fmt.Printf("Matt: %d*%d=%d\n", args.A, args.B, reply)

	args2 := Args2{"Rpc client kill program!"} // Defining the arguments
	var stop string

	// Calling my method StopServer Of the type Stop
	client.Call("Stop.StopServer", args2, &stop)
	fmt.Printf("Stop: %s\n", args2.A)
}
