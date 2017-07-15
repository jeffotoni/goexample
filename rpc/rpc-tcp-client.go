package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Args2 struct {
	A string
}

func main() {

	//client, _ := jsonrpc.Dial("tcp", ":1234")
	client, _ := rpc.Dial("tcp", ":1234")

	// Synchronous call
	args := Args{335, 32}
	var reply int

	// Calling my method
	err := client.Call("Arith.Multiply", args, &reply)
	checkError(err)
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	args2 := Args2{"Rpc client kill program!"} // Defining the arguments
	var stop string

	// Calling my method StopServer Of the type Stop
	err = client.Call("Stop.StopServer", args2, &stop)
	checkError(err)
	fmt.Printf("Stop: %s\n", args2.A)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
