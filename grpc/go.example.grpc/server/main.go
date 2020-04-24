package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// CepSearch implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	//log.Printf("Handling CepSearch request [%v] with context %v", in, ctx)
	return &pb.HelloReply{Message: "Vamos testar..: " + in.Name}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Println("failed to listen: %v", err)
		return
	}

	creds, err := credentials.NewServerTLSFromFile("../certs/server.crt", "../certs/server.key")
	if err != nil {
		log.Println("failed to TlsFromFile: %v", err)
		return
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
	)

	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
