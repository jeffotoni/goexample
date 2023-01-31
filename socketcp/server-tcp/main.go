package main

// import (
// 	"net"
// )

// Criar um servidor TCP
// func startTCPServer(port string) {
// 	// Criar um endereço
// 	addr, err := net.ResolveTCPAddr("tcp", port)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Criar o servidor
// 	listener, err := net.ListenTCP("tcp", addr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Aceitar e lidar com solicitações
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Iniciar uma goroutine para lidar com a conexão
// 		go handleConnection(conn)
// 	}
// }

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	startTCPServer("3000")
}

// Criar um servidor TCP
func startTCPServer(port string) {
	// Criar um endereço
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		panic(err)
	}

	// Criar o servidor
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	// Aceitar e lidar com solicitações
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// Iniciar uma goroutine para lidar com a conexão
		go handleConnection(conn)
	}
}

// Função para lidar com conexões
func handleConnection(conn net.Conn) {
	// Criar um scanner de leitura
	scanner := bufio.NewScanner(conn)

	// Lendo as mensagens
	for scanner.Scan() {
		// Imprimir a mensagem
		fmt.Println(scanner.Text())

		// Enviar a mensagem de volta
		fmt.Fprintf(conn, "Recebido: %s\n", scanner.Text())
	}
}
