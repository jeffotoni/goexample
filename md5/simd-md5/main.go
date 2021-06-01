package main

import (
	"fmt"

	"github.com/minio/sha256-simd"
)

func main() {

	fileBlock := []byte("xxxxxxxxxxxxxxxx")
	server := sha256.NewAvx512Server()

	h512 := sha256.NewAvx512(server)
	h512.Write(fileBlock)

	digest := h512.Sum([]byte{})

	fmt.Println(digest)
}
