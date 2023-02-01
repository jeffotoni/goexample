package main

import (
	"bufio"
	"fmt"
	"net"
	"log"
)

func main() {
	log.Println("Run Client :3000")
	startTCPClient(":3000")
}

// Criar um cliente TCP
func startTCPClient(port string) {
	// Criar um endereço
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Println("Error resolver TCP:",err)
		return 
	}

	// Criar a conexão
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("Error server tcp not exist:",err)
                return
	}

	// Criar um leitor de mensagens
	reader := bufio.NewReader(conn)

	// Escrever uma mensagem
	fmt.Fprintf(conn, "Olá servidor\n")

	// Lendo a mensagem de volta
	msg, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error ReadString:",err)
                return
	}

	// Imprimir a mensagem
	fmt.Println(msg)
}
