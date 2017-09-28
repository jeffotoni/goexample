/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	KEY = "AES256Key-32characters1234567891"
)

func EncryptGcm(text string) string {

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(KEY)
	plaintext := []byte("here plain text")

	block, err := aes.NewCipher(key)

	if err != nil {

		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of a repeat.
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	fmt.Println(nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Println(ciphertext)
}

func DecryptGcm(text string) string {

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(KEY)
	ciphertext, _ := hex.DecodeString("2df87baf86b5073ef1f03e3cc738de75b511400f5465bb0ddeacf47ae4dc267d")

	nonce, _ := hex.DecodeString("afb8a7579bf971db9f8ceeed")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
	// Output: exampleplaintext
}

//
// developer
//
func main() {

}
