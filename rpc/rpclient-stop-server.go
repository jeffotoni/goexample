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
	"math/rand"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"time"
)

type Args2 struct {
	A string
}

// type stop
type Stop string

func (s *Stop) StopClient(args *Args2, replys *string) error {

	// fmt.Println("Stopping
	// the server by rpc!")
	fmt.Println(args.A)

	// it sends an exit to the client
	// to drop and then restart the client
	time.Sleep(50 * time.Millisecond)

	fmt.Println("stop client at now...")
	os.Exit(0)

	return nil
}

func CallRpcServer() {

	// Synchronous call
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	args2 := Args2{"Rpc server bug, kill myself please"} // Defining the arguments
	var stopServer string

	fmt.Println("Send msg rpc server...")

	// Calling my method StopServer Of the type Stop
	client.Call("Stop.StopServer", args2, &stopServer)

}

func simulation() {

	// simulating
	// running
	// program
	func() {
		// loop
		// infinit
		for {

			ii := rand.Intn(10)

			//fmt.Println(time.Millisecond)
			time.Sleep(time.Millisecond * 500)

			if ii == 8 {

				CallRpcServer()
			}
		}
	}()
}

func main() {

	startClient, _ := strconv.Atoi(os.Getenv("START"))

	fmt.Println("Env: ", startClient)

	// start server
	// client
	fmt.Println("Starting client now")
	time.Sleep(time.Second * 1)

	fmt.Println("Start rpc client")
	// Recording the method Stop
	stop := new(Stop)
	rpc.Register(stop)

	// Start handler
	rpc.HandleHTTP()

	if startClient == 0 {

		fmt.Println("call server stop my service a few moments")
		go simulation()

	} else {

		fmt.Println("client online again!")
	}

	fmt.Println("service port: 12345")

	// Opening the port for communication
	err := http.ListenAndServe(":1235", nil)

	if err != nil {

		fmt.Println("Error at now ")
		fmt.Println(err.Error())
	}

	// go func() {
	// }()

}
