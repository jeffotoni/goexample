package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type CabecalhoCliente struct {
	Payload    uint32
	Psecret    uint32
	Step       uint16
	Lastdigits uint16
}

type PacoteA struct {
	Cabecalho CabecalhoCliente
	Mensagem  []byte
	Id        uint32
}

func main() {
	addr := ":12345"
	mensagem := "hello word"
	cab := CabecalhoCliente{
		Payload:    uint32(len(mensagem)),
		Psecret:    0,
		Step:       1,
		Lastdigits: 456,
	}
	pacote := PacoteA{
		Cabecalho: cab,
		Mensagem:  []byte(mensagem),
	}

	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Printf("Error ResolveUDPAddr: %v\n", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Printf("Error DialUDP: %v\n", err)
		return
	}
	encoder := gob.NewEncoder(conn)
	encoder.Encode(pacote)
	conn.Close()
}
