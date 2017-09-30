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
	//"os"
)

const (
	KEY       = "AES256Key-32Characters1234567890"
	KeySize   = 32
	NonceSize = 12
)

var (
	Nonce []byte
)

func GenerateKey() (*[KeySize]byte, error) {

	key := new([KeySize]byte)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateNonce() (*[NonceSize]byte, error) {

	nonce := new([NonceSize]byte)

	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

func EncryptGcm(text string) string {

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(KEY)

	plaintext := []byte(text)

	block, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of a repeat.
	//Nonce, err := GenerateNonce()

	Nonce = make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, Nonce); err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println("nonce:", string(Nonce))

	if err != nil {
		fmt.Println(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, Nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)
}

func DecryptGcm(text string) string {

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(KEY)
	ciphertext, _ := hex.DecodeString(text)

	nonce, _ := hex.DecodeString(fmt.Sprintf("%x", Nonce))

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		//panic(err.Error())
		fmt.Println(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

//
// developer
//
func main() {

	textCry := EncryptGcm("Let's encrypt our text here.")
	fmt.Println(textCry)

	textDescry := DecryptGcm(textCry)
	fmt.Println(textDescry)
}
