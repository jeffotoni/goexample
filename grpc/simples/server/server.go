package main

//go:generate protoc --go_out=plugins=grpc:proto -I=proto produto.proto

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pro "github.com/jeffotoni/goexample/grpc/simples/proto"
)

var ADDR = "0.0.0.0:50051"

type postServiceServer struct{}

func (s *postServiceServer) CreatePost(ctx context.Context, in *pro.PostRequest) (*pro.PostResponse, error) {

	fmt.Println("chamada realizada por um client, e estamos enviando:", in)
	return &pro.PostResponse{Status: 201, Msg: fmt.Sprintf("Criando Produto: %d", in.Id)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pro.RegisterPostServiceServer(s, &postServiceServer{})
	reflection.Register(s)

	fmt.Println("Run Server port:", ADDR)
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}
