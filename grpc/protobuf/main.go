package main

//go:generate protoc --go_out=. ./produto/produto.proto

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jeffotoni/goexample/grpc/protobuf/produto"
	"io"
	"os"
	"os/user"
)

var endianness = binary.LittleEndian

const (
	dbPath = "produto.pb"
)

type length uint16

func add(id int64, name string, done bool) (err error) {
	u := &user.User{
		ID:   id,
		Nome: name,
		Done: done,
	}

	b, err := proto.Marshal(u)
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", dbPath, err)
	}

	// add record length to file
	if err = binary.Write(f, endianness, length(len(b))); err != nil {
		return fmt.Errorf("could not encode length of message: %v", err)
	}

	// add rocord to file
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("could not close file %s: %v", dbPath, err)
	}
	return nil
}

func list() (err error) {
	f, err := os.Open(dbPath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", dbPath, err)
	}
	defer func() {
		e := f.Close()
		if e != nil {
			fmt.Println(e)
		}
	}()

	for {
		// load record file
		var l length
		err = binary.Read(f, endianness, &l)
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return fmt.Errorf("could not read file %s: %v", dbPath, err)
		}

		// load record
		bs := make([]byte, l)
		_, err = io.ReadFull(f, bs)
		if err != nil {
			return fmt.Errorf("could not read file %s: %v", dbPath, err)
		}

		// Unmarshal
		var u user.User
		err = proto.Unmarshal(bs, &u)
		if err != nil {
			return fmt.Errorf("could not read user: %v", err)
		}

		// Print
		fmt.Println("id:", u.GetID())
		fmt.Println("name:", u.GetNome())
		fmt.Println("done:", u.GetDone())
		fmt.Println("------------------")
	}

}

func main() {

	err := add(1, "TV LG, 55 polegadas", true)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = list()
	if err != nil {
		fmt.Println(err)
		return
	}
}
