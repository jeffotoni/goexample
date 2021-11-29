package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

// go run .
func main() {

	// echo "testando tcp server" | netcat localhost 9000
	go tcpConnect()

	// nc -u localhost 3000
	go updConnect()

	// curl -i localhost:8080/api/v1/grpc
	go grpcConnect()

	// curl -i localhost:8080
	httpConnect()
}

func httpConnect() {
	println("Run Server HTTP:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/v1/grpc" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("call grpc client"))
				clientGrpc()
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Let's Go to DevOpsFest..."))

		})))
}

func tcpConnect() {
	tcp, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to listen tcp:", err)
	}
	println("Run Server TCP:9000")
	for {
		conn, _ := tcp.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(string(message))
		if message == "" {
			conn, _ = tcp.Accept()
		} else {
			fmt.Println("Client sent:", string(message))
			conn.Write([]byte("Welcome at the DevOpsFest!" + "\n"))
			conn.Close()
		}
	}
}

func updConnect() {
	addr := ":3000"
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatal("Error ResolveUDPAddr:\n", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("Error ListenUDP:\n", err)
	}
	defer conn.Close()

	println("Run Server UPD:3000")
	for {
		message := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}
		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("received: %s from %s\n", data, remote)

		dataCli := []byte("UDP message server for you..\n")
		_, err = conn.WriteToUDP(dataCli, remote)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func grpcConnect() {
	lis, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}
	println("Run Server GRPC:9001")
	grpcServer := grpc.NewServer()
	RegisterMainServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func clientGrpc() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("0.0.0.0:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := NewMainServiceClient(conn)

	response, err := c.SayHello(context.Background(), &Message{Body: "Hello Client GRPC..."})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	response, err = c.BroadcastMessage(context.Background(), &Message{Body: "Message to Broadcast grpc tests!"})
	if err != nil {
		log.Fatalf("Error when calling Broadcast Message: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
