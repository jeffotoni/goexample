package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type CabecalhoServidor struct {
	Payload    uint32
	Psecret    uint32
	Step       uint16
	Lastdigits uint16
}

type RecebeA struct {
	Cabecalho CabecalhoServidor
	Mensagem  []byte
}

func ClienteA(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	rA := &RecebeA{}
	dec.Decode(&rA)
	if len(string(rA.Mensagem)) > 0 {
		fmt.Println("Mensagem: ", string(rA.Mensagem))
	}
	conn.Close()
}

func main() {
	var err error
	addr := ":12345"
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Printf("Error ResolveUDPAddr: %v\n", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("Error ListenUDP: %v\n", err)
		return
	}

	log.Println("Listen:", addr)
	for {
		conn, err = net.ListenUDP("udp", udpAddr)
		if err != nil {
			continue
		}
		ClienteA(conn)
	}
}
