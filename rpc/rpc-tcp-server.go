package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
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

	arith := new(Arith)
	rpc.Register(arith)

	// Recording the method Stop
	stop := new(Stop)
	rpc.Register(stop)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//jsonrpc.ServeConn(conn)
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
