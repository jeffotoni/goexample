package main

//go:generate protoc --go_out=. ./proto/poduto.proto

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pro "github.com/jeffotoni/goexample/grpc/simples/proto"
)

type postServiceServer struct{}

func (s *postServiceServer) CreatePost(ctx context.Context, in *pro.PostRequest) (*pro.PostResponse, error) {

	fmt.Println("Capturando dados:", in)
	return &pro.PostResponse{Status: 201, Msg: fmt.Sprintf("Criando Produto: %s", in.Id)}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pro.RegisterPostServiceServer(s, &postServiceServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}
