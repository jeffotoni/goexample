package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jeffotoni/goexample/grpc/simples/server/proto"
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

	client := proto.NewPostServiceClient(conn)

	req := &proto.PostRequest{
		Id:   12345678,
		Nome: "Produto CHAT/GPT-4",
	}

	res, err := client.CreatePost(context.Background(), req)
	if err != nil {
		log.Printf("Não foi possível chamar o serviço: %v", err)
		return
	}

	fmt.Println("Objeto:\n", res)
	b, err := json.Marshal(&res)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Json:\n", string(b))
}
