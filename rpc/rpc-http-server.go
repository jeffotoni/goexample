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
	"net/http"
	"net/rpc"
	"os"
	"time"
)

// Method Multiply arguments
type Args struct {
	A, B int
}

// Kind for my method
type Matt int

// My method Multiply
func (t *Matt) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Method StopServer arguments
type Args2 struct {
	A string
}

// type stop
type Stop string

// My method StopServer
func (s *Stop) StopServer(args *Args2, replys *string) error {

	*replys = args.A + " ok! "
	fmt.Println("Stopping the server by rpc!")

	var count = 5
	for i := 0; i < count; i++ {

		fmt.Println("service[", i, "]", "stop")
		time.Sleep(2 * time.Second)
	}

	os.Exit(1)
	return nil
}

func main() {

	// Recording the method Matt
	matt := new(Matt)
	rpc.Register(matt)

	// Recording the method Stop
	stop := new(Stop)
	rpc.Register(stop)

	// Start handler
	rpc.HandleHTTP()

	// Opening the port for communication
	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
