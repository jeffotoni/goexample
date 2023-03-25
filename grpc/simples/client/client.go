package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jeffotoni/goexample/grpc/simples/proto"
	"google.golang.org/grpc"
)

var ADDR = "localhost:50051"

func main() {
	// Conectar ao servidor GRPC
	conn, err := grpc.Dial(ADDR, grpc.WithInsecure())
	if err != nil {
		log.Printf("Não foi possível conectar ao servidor: %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewClienteServiceClient(conn)

	req := &proto.PostRequest{
		ID:   12345678,
		Nome: "Produto CHAT/GPT-4",
	}

	res, err := client.CreatePost(context.Background(), req)
	if err != nil {
		log.Printf("Não foi possível chamar o serviço: %v", err)
		return
	}

	fmt.Println(res)
}
