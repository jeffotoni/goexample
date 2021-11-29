package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

// go run main.go chat.go chat.pb.go
func main() {

	// echo "testando tcp server" | netcat localhost 9000
	go tcpConnect()

	// nc -u localhost 3000
	go updConnect()

	// go run chat.pb.go chat.go client.grpc.go
	go grpcConnect()

	// curl -i localhost:8080
	httpConnect()
}

func httpConnect() {
	println("Run Server HTTP:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Let's Go to DevOpsFest ..."))
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
