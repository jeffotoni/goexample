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
	"strings"
	//"sync"
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var startClient = 0

type Args2 struct {
	A string
}

// type stop
type Stop string

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func (s *Stop) StopServer(args *Args2, replys *string) error {

	//*replys = args.A + " ok! "
	//fmt.Println("Stopping the server by rpc!")
	fmt.Println(args.A)

	fmt.Println("Stop in 100 millisecond")
	// it sends an exit to the client
	// to drop and then restart the client
	time.Sleep(100 * time.Millisecond)

	CallRpcClient()

	startClient++

	startString := strconv.Itoa(startClient)

	// Create an *exec.Cmd
	cmd := exec.Command("go", "run", "./rpclient-stop-server.go")

	cmd.Env = append(os.Environ(),

		"START="+startString, // ignored
	)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}

	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)

	err := cmd.Run() // will wait for command to return

	printError(err)

	// Only output the commands stdout
	printOutput(cmdOutput.Bytes()) // => go version go1.3 darwin/amd64

	// start application
	// now again
	//cmd := exec.Command("sh", "./start.sh", "&")
	//cmd.Run()
	//cmd.Start()
	//cmd.Wait()

	// cmd.Env = append(os.Environ(),

	// 	"START="+startString, // ignored
	// )

	// if err := cmd.Run(); err != nil {

	// 	fmt.Println(err)
	//}

	return nil
}

func CallRpcClient() {

	// Synchronous call
	client, _ := rpc.DialHTTP("tcp", "localhost:1235")

	args2 := Args2{"Rpc client i'm stopping your service..."} // Defining the arguments
	var stop string

	fmt.Println("Send msg rpc client exit service...")
	// Calling my method StopServer Of the type Stop
	client.Call("Stop.StopClient", args2, &stop)
}

func main() {

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
