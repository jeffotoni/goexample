package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	startTCPClient("3000")
}

// Criar um cliente TCP
func startTCPClient(port string) {
	// Criar um endereço
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		panic(err)
	}

	// Criar a conexão
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	// Criar um leitor de mensagens
	reader := bufio.NewReader(conn)

	// Escrever uma mensagem
	fmt.Fprintf(conn, "Olá servidor\n")

	// Lendo a mensagem de volta
	msg, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Imprimir a mensagem
	fmt.Println(msg)
}
